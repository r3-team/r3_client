package file

import (
	"fmt"
	"os"
	"path/filepath"
	"r3_client/config"
	"r3_client/file/open"
	"r3_client/log"
	"r3_client/tools"
	"r3_client/tray"
	"r3_client/types"
	"sort"
	"sync"

	"github.com/gofrs/uuid"
)

var (
	files_mx        sync.Mutex
	files           = make(map[uuid.UUID]types.File) // key: file ID
	logContext      = "fileManager"
	tempDir         = ""    // from OS
	tempDirAttempts = 1000  // how many attempts to create file temp directory
	tempDirPrefix   = "r3_" // directory prefix
)

func init() {
	tempDir = os.TempDir()
}

func Open(attributeId uuid.UUID, fileId uuid.UUID,
	fileHash string, fileName string, chooseApp bool) error {

	var err error
	defer updateTouched(fileId)

	files_mx.Lock()
	f, exists := files[fileId]
	files_mx.Unlock()

	if exists {
		// file reference exists, check for file as well
		exists, err = tools.Exists(filepath.Join(tempDir, f.DirName, f.FileName))
		if err != nil {
			log.Error(logContext, "failed to check file", err)
			return err
		}
	}

	// file or file reference does not exist, create
	if !exists {
		dirName := ""
		for tries := 0; tries < tempDirAttempts; tries++ {

			dirName = fmt.Sprintf("%s%d", tempDirPrefix, tools.RandNumber(100000, 499999))
			dirExists, err := tools.Exists(filepath.Join(tempDir, dirName))
			if err != nil {
				log.Error(logContext, "failed to check temporary directory", err)
				return err
			}
			if !dirExists {
				if err := os.Mkdir(filepath.Join(tempDir, dirName), 0750); err != nil {
					log.Error(logContext, "failed to create temporary directory", err)
					return err
				}
				break
			}
		}

		if dirName == "" {
			return fmt.Errorf("failed to create temporary directory after %d attempts", tempDirAttempts)
		}

		f.AttributeId = attributeId
		f.DirName = dirName
		f.FileHash = fileHash
		f.FileName = fileName
	}
	filePath := filepath.Join(tempDir, f.DirName, f.FileName)

	if exists {
		if f.FileHash == fileHash {

			// correct file version is already available, just open it
			log.Info(logContext, "already has the correct file version available, opens it")
			return open.WithLocalSystem(filePath, chooseApp)
		} else {
			// file exists but is outdated, remove it
			if err := os.Remove(filePath); err != nil {
				// cannot remove temporary file, still being accessed, nothing to do
				log.Warning(logContext, "failed to delete older file version", err)
				return err
			}
		}
	}

	// download file
	scheme := "https"
	if !config.File.Ssl {
		scheme = "http"
	}
	fileUrl := fmt.Sprintf("%s://%s:%d/data/download/%s?attribute_id=%s&file_id=%s&token=%s",
		scheme, config.File.HostName, config.File.HostPort, fileName,
		attributeId, fileId, config.GetAuthToken())

	log.Info(logContext, fmt.Sprintf("downloading file from '%s'", fileUrl))
	if err := download(fileUrl, filePath); err != nil {
		return err
	}

	// TEMP
	// todo: check downloaded hash against expected file hash

	// register file
	files_mx.Lock()
	files[fileId] = f
	files_mx.Unlock()

	if err := watcherAdd(filepath.Join(tempDir, f.DirName)); err != nil {
		return err
	}
	if err := CacheStore(); err != nil {
		return err
	}
	return open.WithLocalSystem(filePath, chooseApp)
}

func updateTouched(fileId uuid.UUID) {
	files_mx.Lock()
	if f, exists := files[fileId]; exists {
		f.Touched = tools.GetTimeUnix()
		files[fileId] = f
	}
	files_mx.Unlock()
	updateTray()
}

func updateTray() {
	files_mx.Lock()
	defer files_mx.Unlock()

	// show latest 3 touched files in systray
	keys := make([]uuid.UUID, 0, len(files))
	for k := range files {
		keys = append(keys, k)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return files[keys[i]].Touched > files[keys[j]].Touched
	})

	filesShow := make([]types.File, 0)
	for i := 0; i < 3 && i < len(keys); i++ {
		filesShow = append(filesShow, files[keys[i]])
	}
	tray.SetFiles(filesShow)
}
