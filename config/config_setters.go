package config

import (
	"regexp"
	"strconv"

	"github.com/gofrs/uuid"
)

func SetAppVersion(versionFull string) error {
	access_mx.Lock()
	defer access_mx.Unlock()

	build, err := strconv.Atoi(regexp.MustCompile(`^\d+\.\d+\.\d+\.`).ReplaceAllString(versionFull, ""))
	if err != nil {
		return err
	}

	appVersionFull = versionFull
	appVersionBuild = build
	return nil
}
func SetAutoStart(v bool) {
	access_mx.Lock()
	defer access_mx.Unlock()

	file.AutoStart = v
}
func SetDarkIcon(v bool) {
	access_mx.Lock()
	defer access_mx.Unlock()

	file.DarkIcon = v
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
