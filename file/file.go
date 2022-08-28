package file

import (
	"fmt"
	"os"
	"path/filepath"
	"r3_client/config"
	"r3_client/file/open"
	"r3_client/log"
	"r3_client/tools"
	"r3_client/types"
	"sync"

	"github.com/gofrs/uuid"
)

var (
	files           = make(map[uuid.UUID]types.File) // key: file ID
	files_mx        sync.Mutex
	logContext      = "fileManager"
	tempDir         = ""    // from OS
	tempDirAttempts = 1000  // how many attempts to create file temp directory
	tempDirPrefix   = "r3_" // directory prefix
	trayFileCnt     = 5     // number of last accessed files, shown in system tray
)

func init() {
	tempDir = os.TempDir()
}

func GetDirPath(dirName string) string {
	return filepath.Join(tempDir, dirName)
}
func GetFilePath(dirName string, fileName string) string {
	return filepath.Join(tempDir, dirName, fileName)
}
func setFile(id uuid.UUID, f types.File, trayShouldUpdate bool) {
	files_mx.Lock()
	files[id] = f
	files_mx.Unlock()

	CacheStore()

	if trayShouldUpdate {
		updateTray()
	}
}
func setFileHash(id uuid.UUID, hash string) {
	files_mx.Lock()
	f, exists := files[id]
	files_mx.Unlock()

	if exists {
		f.FileHash = hash
		setFile(id, f, false)
	}
}
func setFileTouchedToNow(id uuid.UUID) {
	files_mx.Lock()
	f, exists := files[id]
	files_mx.Unlock()

	if exists {
		f.Touched = tools.GetTimeUnix()
		setFile(id, f, true)
	}
}

func Open(instanceId uuid.UUID, attributeId uuid.UUID, fileId uuid.UUID,
	fileHash string, fileName string, chooseApp bool) error {

	var err error
	files_mx.Lock()
	f, exists := files[fileId]
	files_mx.Unlock()

	if exists {
		// file reference exists, check for file as well
		exists, err = tools.Exists(GetFilePath(f.DirName, f.FileName))
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
			return fmt.Errorf("failed to create temporary directory after %d attempts",
				tempDirAttempts)
		}

		f.AttributeId = attributeId
		f.DirName = dirName
		f.FileHash = fileHash
		f.FileName = fileName
		f.InstanceId = instanceId
		f.Touched = tools.GetTimeUnix()
	}
	filePath := GetFilePath(f.DirName, f.FileName)

	if exists {
		if f.FileHash == fileHash {

			// correct file version is already available, just open it
			log.Info(logContext, "already has the correct file version available, opens it")
			setFileTouchedToNow(fileId)
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
	inst, err := config.GetInstance(instanceId)
	if err != nil {
		return err
	}

	scheme := "https"
	if !config.GetSsl() {
		scheme = "http"
	}
	fileUrl := fmt.Sprintf("%s://%s:%d/data/download/%s?attribute_id=%s&file_id=%s&token=%s",
		scheme, inst.HostName, inst.HostPort, fileName,
		attributeId, fileId, config.GetAuthToken(instanceId))

	log.Info(logContext, fmt.Sprintf("downloading file from '%s'", fileUrl))
	if err := download(fileUrl, filePath); err != nil {
		return err
	}

	// register file
	setFile(fileId, f, true)
	if err := watcherAdd(filepath.Join(tempDir, f.DirName)); err != nil {
		return err
	}
	return open.WithLocalSystem(filePath, chooseApp)
}
