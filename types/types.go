package types

import "github.com/gofrs/uuid"

// handled files
type File struct {
	AttributeId uuid.UUID `json:"attributeId"` // file attribute
	DirName     string    `json:"dirName"`     // file directory name, inside temp directory
	FileHash    string    `json:"fileHash"`    // known file hash (to check for changes)
	FileName    string    `json:"fileName"`    // file name
	Touched     int64     `json:"touched"`     // unix timestamp of last time file was touched
}
type FilesSaved struct {
	Files map[uuid.UUID]File `json:"files"`
}

// configuration file
type ConfigFile struct {
	AutoStart    bool   `json:"autoStart"`
	Debug        bool   `json:"debug"`
	DeviceName   string `json:"deviceName"`
	HostName     string `json:"hostName"`
	HostPort     int    `json:"hostPort"`
	KeepFilesSec int64  `json:"keepFilesSec`
	LanguageCode string `json:"languageCode"`
	LoginId      int64  `json:"loginId"`
	Ssl          bool   `json:"ssl"`
	SslVerify    bool   `json:"sslVerify"`
	TokenFixed   string `json:"tokenFixed"`
}
