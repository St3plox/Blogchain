package contract

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// AddressConfig represents a mapping of names to their respective addresses
type AddressConfig map[string]struct {
	Address string `json:"address"`
}

// Config holds the dynamic mapping
var config AddressConfig

// LoadConfig loads the JSON configuration from a file
func LoadConfig(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("could not open config file: %v", err)
	}
	defer file.Close()

	byteValue, err := ioutil.ReadAll(file)
	if err != nil {
		return fmt.Errorf("could not read config file: %v", err)
	}

	if err := json.Unmarshal(byteValue, &config); err != nil {
		return fmt.Errorf("could not unmarshal config JSON: %v", err)
	}

	return nil
}

// GetAddress returns the address for the given name
func GetAddress(name string) (string, error) {
	entry, exists := config[name]
	if !exists {
		return "", fmt.Errorf("unknown name: %s", name)
	}
	return entry.Address, nil
}
