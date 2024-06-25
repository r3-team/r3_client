package types

import (
	"github.com/gofrs/uuid"
)

// handled files
type File struct {
	AttributeId uuid.UUID `json:"attributeId"` // file attribute
	DirName     string    `json:"dirName"`     // file directory name, inside temp directory
	FileHash    string    `json:"fileHash"`    // known file hash (to check for changes)
	FileName    string    `json:"fileName"`    // file name
	InstanceId  uuid.UUID `json:"instanceId"`  // ID of R3 instance that file belongs to
	Touched     int64     `json:"touched"`     // unix timestamp of last time file was touched
}
type FilesSaved struct {
	Files map[uuid.UUID]File `json:"files"`
}

// configuration file
type Event struct {
	Action          string `json:"action"`          // callJsFunction
	Event           string `json:"event"`           // onHotkey, onConnect, onDisconnect
	HotkeyModifier1 string `json:"hotkeyModifier1"` // CTRL, ALT, SHIFT
	HotkeyModifier2 string `json:"hotkeyModifier2"` // CTRL, ALT, SHIFT
	HotkeyChar      string `json:"hotkeyChar"`      // a-Z, 0-9

	// for action: callJsFunction
	JsFunctionArgs []string      `json:"jsFunctionArgs"` // arguments to deliver to callback function, in order (clipboard, hostname, username, windowTitle)
	JsFunctionId   uuid.NullUUID `json:"jsFunctionId"`   // JS function to call inside the browser session
}
type Instance struct {
	DeviceName string  `json:"deviceName"`
	Events     []Event `json:"events"`
	HostName   string  `json:"hostName"`
	HostPort   int     `json:"hostPort"`
	LoginId    int64   `json:"loginId"`
	TokenFixed string  `json:"tokenFixed"`
}
type ConfigFile struct {
	AutoStart    bool                   `json:"autoStart"`
	DarkIcon     bool                   `json:"darkIcon"`
	Debug        bool                   `json:"debug"`
	Instances    map[uuid.UUID]Instance `json:"instances"`
	KeepFilesSec int64                  `json:"keepFilesSec"`
	LanguageCode string                 `json:"languageCode"`
	Ssl          bool                   `json:"ssl"`
	SslVerify    bool                   `json:"sslVerify"`
}
