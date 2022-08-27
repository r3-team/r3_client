package config

import (
	"crypto/tls"
	"encoding/json"
	"os"
	"path/filepath"
	"sync"

	"r3_client/tools"
	"r3_client/types"
)

var (
	access_mx   = &sync.Mutex{}
	appVersion  string
	authToken   string // authentication JWT
	fileName    = "r3_client.conf"
	pathApp     string // application path
	pathUser    string // user home path
	pathHomedir string

	File types.ConfigFile
)

func GetTlsConfig() tls.Config {
	tlsConfig := tls.Config{
		PreferServerCipherSuites: true,
	}
	if !File.SslVerify {
		tlsConfig.InsecureSkipVerify = true
	}
	return tlsConfig
}
func GetAppVersion() string {
	return appVersion
}
func GetIsAuthenticated() bool {
	return authToken != ""
}
func GetAuthToken() string {
	return authToken
}
func GetFileName() string {
	return fileName
}
func GetPathApp() string {
	return pathApp
}
func GetPathUser() string {
	return pathUser
}
func SetAppVersion(v string) {
	appVersion = v
}
func SetAuthToken(v string) {
	authToken = v
}
func SetPathApp(v string) {
	pathApp = v
}
func SetPathUser(v string) {
	pathUser = v
}
func LoadCreateFile() error {

	// create new config file with defaults if it does not exist
	filePath := filepath.Join(pathApp, fileName)
	exists, err := tools.Exists(filePath)
	if err != nil {
		return err
	}
	if !exists {
		File = types.ConfigFile{
			AutoStart:    true,
			Debug:        false,
			DeviceName:   "DEVICE_NAME",
			HostName:     "SERVER_HOSTNAME",
			HostPort:     443,
			LanguageCode: "en_us",
			LoginId:      -1,
			Ssl:          true,
			SslVerify:    true,
			TokenFixed:   "LOGIN_APP_TOKEN",
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
	return json.Unmarshal(configJson, &File)
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
	return os.WriteFile(filepath.Join(pathApp, fileName), json, 0644)
}
