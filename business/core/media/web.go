package media

import (
	"errors"
	"net/http"

	v1 "github.com/St3plox/Blogchain/business/web/v1"
	"github.com/gabriel-vasile/mimetype"
)

// ParseMedia function is utility function that hepls hendlers to parse media from http request
func (c *Core) ParseMedia(r *http.Request) (NewMedia, error) {
	// Parse the multipart form, allowing for a maximum upload MaxFileSizeMb
	err := r.ParseMultipartForm(c.MaxFileSizeMb << 20)
	if err != nil {
		return NewMedia{}, v1.NewRequestError(errors.New("failed to parse multipart form: "+err.Error()), http.StatusBadRequest)
	}

	file, handler, err := r.FormFile("file")
	if err != nil {
		return NewMedia{}, v1.NewRequestError(errors.New("read file error: "+err.Error()), http.StatusBadRequest)
	}
	defer file.Close()

	buf := make([]byte, handler.Size)
	_, err = file.Read(buf)
	if err != nil {
		return NewMedia{}, v1.NewRequestError(errors.New("error reading file: "+err.Error()), http.StatusInternalServerError)
	}

	// Check if the file is an image
	mime := mimetype.Detect(buf)
	if !mime.Is("image/jpeg") && !mime.Is("image/png") && !mime.Is("image/gif") && !mime.Is("image/bmp") {
		return NewMedia{}, v1.NewRequestError(errors.New("file is not a valid image type"), http.StatusBadRequest)
	}

	return NewMedia{
		Filename:  handler.Filename,
		Length:    handler.Size,
		FileBytes: buf,
	}, nil
}

func (c *Core) ParseMultipleMedia(r *http.Request) ([]NewMedia, error) {

	// Parse the multipart form, allowing for a maximum upload MaxFileSizeMb
	err := r.ParseMultipartForm(c.MaxFileSizeMb << 20)
	if err != nil {
		return nil, v1.NewRequestError(errors.New("failed to parse multipart form: "+err.Error()), http.StatusBadRequest)
	}

	files := r.MultipartForm.File["files"]
	if len(files) == 0 {
		return nil, v1.NewRequestError(errors.New("no files provided"), http.StatusBadRequest)
	}

	media := make([]NewMedia, len(files))

	for _, handler := range files {
		file, err := handler.Open()
		if err != nil {
			return nil, v1.NewRequestError(errors.New("read file error: "+err.Error()), http.StatusBadRequest)
		}
		defer file.Close()

		buf := make([]byte, handler.Size)
		_, err = file.Read(buf)
		if err != nil {
			return nil, v1.NewRequestError(errors.New("error reading file: "+err.Error()), http.StatusInternalServerError)
		}

		// Check if the file is an image
		mime := mimetype.Detect(buf)
		if !mime.Is("image/jpeg") && !mime.Is("image/png") && !mime.Is("image/gif") && !mime.Is("image/bmp") {
			return nil, v1.NewRequestError(errors.New("file is not a valid image type"), http.StatusBadRequest)
		}

		newMedia := NewMedia{
			Filename:  handler.Filename,
			Length:    handler.Size,
			FileBytes: buf,
		}
		media = append(media, newMedia)
	}

	return media[len(media)/2:], nil
}
