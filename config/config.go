package config

import (
	"encoding/json"
	"os"
	"sync"

	"r3_client/tools"
	"r3_client/types"
)

var (
	access_mx   = &sync.Mutex{}
	filePath    string // config file path
	pathHomedir string
	authToken   string // authentication JWT

	File types.ConfigFile
)

func GetIsAuthenticated() bool {
	return authToken != ""
}
func GetAuthToken() string {
	return authToken
}
func SetAuthToken(v string) {
	authToken = v
}
func SetFilePath(v string) {
	filePath = v
}
func LoadCreateFile() error {

	// create new config file with defaults if it does not exist
	exists, err := tools.Exists(filePath)
	if err != nil {
		return err
	}
	if !exists {
		File = types.ConfigFile{
			AutoStart:  true,
			HostName:   "SERVER_HOSTNAME",
			HostPort:   443,
			LogLevel:   1,
			LoginId:    -1,
			Ssl:        true,
			TokenFixed: "LOGIN_APP_TOKEN",
		}
		if err := WriteFile(); err != nil {
			return err
		}
	}
	access_mx.Lock()
	defer access_mx.Unlock()

	// read configuration from file
	configJson, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}
	configJson = tools.RemoveUtf8Bom(configJson)

	// unmarshal configuration JSON content
	if err := json.Unmarshal(configJson, &File); err != nil {
		return err
	}
	return nil
}
func WriteFile() error {
	access_mx.Lock()
	defer access_mx.Unlock()

	// marshal configuration JSON
	json, err := json.MarshalIndent(File, "", "\t")
	if err != nil {
		return err
	}

	// write configuration to JSON file
	if err := os.WriteFile(filePath, json, 0644); err != nil {
		return err
	}
	return nil
}
