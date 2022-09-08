package config

import (
	"crypto/tls"
	"fmt"
	"r3_client/types"

	"github.com/gofrs/uuid"
)

func GetAppVersion() string {
	access_mx.Lock()
	defer access_mx.Unlock()

	return appVersion
}
func GetIsAuthenticated(instanceId uuid.UUID) bool {
	access_mx.Lock()
	defer access_mx.Unlock()

	token, exists := instanceIdMapToken[instanceId]
	if !exists {
		return false
	}
	return token != ""
}
func GetAuthToken(instanceId uuid.UUID) string {
	access_mx.Lock()
	defer access_mx.Unlock()

	token, exists := instanceIdMapToken[instanceId]
	if !exists {
		return ""
	}
	return token
}
func GetAutoStart() bool {
	access_mx.Lock()
	defer access_mx.Unlock()

	return file.AutoStart
}
func GetDarkIcon() bool {
	access_mx.Lock()
	defer access_mx.Unlock()
	return file.DarkIcon
}
func GetDebug() bool {
	access_mx.Lock()
	defer access_mx.Unlock()
	return file.Debug
}
func GetFileName() string {
	return fileName
}
func GetFileNameCache() string {
	return fileNameCache
}
func GetFileNameLog() string {
	return fileNameLog
}
func GetInstance(instanceId uuid.UUID) (types.Instance, error) {
	access_mx.Lock()
	defer access_mx.Unlock()

	inst, exists := file.Instances[instanceId]
	if !exists {
		return inst, fmt.Errorf("unknown instance '%s'", instanceId)
	}
	return inst, nil
}
func GetInstances() map[uuid.UUID]types.Instance {
	access_mx.Lock()
	defer access_mx.Unlock()
	return file.Instances
}
func GetKeepFilesSec() int64 {
	access_mx.Lock()
	defer access_mx.Unlock()

	return file.KeepFilesSec
}
func GetLanguageCode() string {
	access_mx.Lock()
	defer access_mx.Unlock()

	return file.LanguageCode
}
func GetPathApp() string {
	access_mx.Lock()
	defer access_mx.Unlock()

	return pathApp
}
func GetPathUser() string {
	access_mx.Lock()
	defer access_mx.Unlock()

	return pathUser
}
func GetSsl() bool {
	access_mx.Lock()
	defer access_mx.Unlock()

	return file.Ssl
}
func GetSslVerify() bool {
	access_mx.Lock()
	defer access_mx.Unlock()

	return file.SslVerify
}
func GetTlsConfig() tls.Config {
	access_mx.Lock()
	defer access_mx.Unlock()

	tlsConfig := tls.Config{
		PreferServerCipherSuites: true,
	}
	if !file.SslVerify {
		tlsConfig.InsecureSkipVerify = true
	}
	return tlsConfig
}
