package types

import "github.com/gofrs/uuid"

// handled files
type File struct {
	AttributeId uuid.UUID `json:"attributeId"` // file attribute
	DirName     string    `json:"dirName"`     // file directory name, inside temp directory
	FileHash    string    `json:"fileHash"`    // known file hash (to check for changes)
	FileName    string    `json:"fileName"`    // file name
	Version     int       `json:"version"`     // file version
}
type FilesSaved struct {
	Files map[uuid.UUID]File `json:"files"`
}

// configuration file
type ConfigFile struct {
	AutoStart  bool   `json:"autoStart"`
	HostName   string `json:"hostName"`
	HostPort   int    `json:"hostPort"`
	LogLevel   int    `json:"logLevel"`
	LoginId    int64  `json:"loginId"`
	Ssl        bool   `json:"ssl"`
	TokenFixed string `json:"tokenFixed"`
}
