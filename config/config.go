package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"r3_client/tools"
	"r3_client/types"

	"github.com/gofrs/uuid"
)

var (
	access_mx = &sync.Mutex{}

	appVersion  string
	file        types.ConfigFile // config file
	fileName    = "r3_client.conf"
	pathApp     string // application path
	pathUser    string // user home path
	pathHomedir string

	instanceIdMapToken = make(map[uuid.UUID]string) // map of instance JWTs, key: instance ID
)

func ReadFile() error {

	// create new config file with defaults if it does not exist
	filePath := filepath.Join(pathApp, fileName)
	exists, err := tools.Exists(filePath)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("config file does not exist at '%s'", filePath)
	}

	// read configuration from file
	configJson, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}
	configJson = tools.RemoveUtf8Bom(configJson)

	// unmarshal configuration JSON content
	access_mx.Lock()
	defer access_mx.Unlock()
	return json.Unmarshal(configJson, &file)
}
func WriteFile() error {
	access_mx.Lock()
	defer access_mx.Unlock()

	// marshal configuration JSON
	json, err := json.MarshalIndent(file, "", "\t")
	if err != nil {
		return err
	}

	// write configuration to JSON file
	return os.WriteFile(filepath.Join(pathApp, fileName), json, 0644)
}
