//go:build linux || darwin

package open

// no good solution yet
// on linux 'mimeopen -d' exists but must be installed and runs via command line
func chooseApp(filePath string) error {
	return nil
}
