//go:build linux

package open

// no good solution yet
// 'mimeopen -d' exists but must be installed and runs via command line
// it seems, we´d need a solution for every file manager (nautilus, dolphin, etc.)
func chooseApp(filePath string) error {
	return nil
}
