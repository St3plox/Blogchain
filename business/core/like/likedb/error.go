package likedb

import (
	"fmt"
)

// Custom errors for Like operations
var (
	ErrInvalidIDFormat       = fmt.Errorf("invalid ID format")
	ErrLikeCreationFailed    = fmt.Errorf("like creation failed")
	ErrLikeNotFound          = fmt.Errorf("like not found")
	ErrLikeQueryFailed       = fmt.Errorf("failed to query Like by ID")
	ErrLikeUpdateFailed      = fmt.Errorf("like update failed")
	ErrLikeDeletionFailed    = fmt.Errorf("like deletion failed")
	ErrInvalidUserIDFormat   = fmt.Errorf("invalid User ID format")
)

// IsInvalidIDFormat checks if an error is due to invalid ID format.
func IsInvalidIDFormat(err error) bool {
	return err == ErrInvalidIDFormat
}

// IsLikeNotFound checks if an error is due to a like not being found.
func IsLikeNotFound(err error) bool {
	return err == ErrLikeNotFound
}

// IsLikeCreationFailed checks if an error is due to like creation failure.
func IsLikeCreationFailed(err error) bool {
	return err == ErrLikeCreationFailed
}

// IsLikeUpdateFailed checks if an error is due to like update failure.
func IsLikeUpdateFailed(err error) bool {
	return err == ErrLikeUpdateFailed
}

// IsLikeDeletionFailed checks if an error is due to like deletion failure.
func IsLikeDeletionFailed(err error) bool {
	return err == ErrLikeDeletionFailed
}

// IsInvalidUserIDFormat checks if an error is due to invalid User ID format.
func IsInvalidUserIDFormat(err error) bool {
	return err == ErrInvalidUserIDFormat
}
