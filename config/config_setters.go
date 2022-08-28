package config

import "github.com/gofrs/uuid"

func SetAppVersion(v string) {
	access_mx.Lock()
	defer access_mx.Unlock()

	appVersion = v
}
func SetAutoStart(v bool) {
	access_mx.Lock()
	defer access_mx.Unlock()

	file.AutoStart = v
}
func SetDebug(v bool) {
	access_mx.Lock()
	defer access_mx.Unlock()

	file.Debug = v
}
func SetSsl(v bool) {
	access_mx.Lock()
	defer access_mx.Unlock()

	file.Ssl = v
}
func SetSslVerify(v bool) {
	access_mx.Lock()
	defer access_mx.Unlock()

	file.SslVerify = v
}
func SetPathApp(v string) {
	access_mx.Lock()
	defer access_mx.Unlock()

	pathApp = v
}
func SetPathUser(v string) {
	access_mx.Lock()
	defer access_mx.Unlock()

	pathUser = v
}
func SetInstanceToken(instanceId uuid.UUID, token string) {
	access_mx.Lock()
	defer access_mx.Unlock()

	instanceIdMapToken[instanceId] = token
}
