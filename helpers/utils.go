package helpers

import (
	"io/fs"
	"log"
	"os"
	"path/filepath"
)


func HandlesError(err error) {
	log.Fatal(err)
}

func Chdir(path string)  {
	err := os.Chdir(path)
	if err != nil {
		HandlesError(err)
	}
}
func GetCurrentDir() string {
	success, err := os.Getwd()
	if err != nil {
		HandlesError(err)
	}
	return success
}

func CreateDir(nameDir string, permission os.FileMode) bool {
	err := os.Mkdir(nameDir, permission)

	if err != nil && !os.IsExist(err) {
		HandlesError(err)
		return false
	} else {
		return true
	}
}

func ReadFile(name string) []byte {
	file, err := os.ReadFile(name)
	if err != nil && !os.IsExist(err) {
		HandlesError(err)
	}
	return file
}

func WriteFile(name string, dataToWrite []byte, permission os.FileMode) bool {
	err := os.WriteFile(name, dataToWrite, permission)
	if err != nil {
		HandlesError(err)
		return false
	}
	return true
}

// func RenameFile is the same as moving

func RenameFile(oldPath, newPath string) bool {
	err := os.Rename(oldPath, newPath)
	if err != nil && !os.IsNotExist(err) {
		HandlesError(err)
		return false
	}
	return true
}

func UserHomeDirectory() string {
	dir, err := os.UserHomeDir()
	if err != nil && !os.IsNotExist(err) {
		HandlesError(err)
	}
	return dir
}

func GetDirs(workingDirectory string) []os.DirEntry {
	dirs, err := os.ReadDir(workingDirectory)
	if err != nil {
		HandlesError(err)
	}
	return dirs
}

func GetFileSize(filepath string) (int64, error) {
	fi, err := os.Stat(filepath)
	if err != nil {
		return 0, err
	}
	// get the size
	return fi.Size(), nil
}

func JoinPath(elem ...string) string {
	return filepath.Join(elem...)
}

func DirSize(path string) (int64, error) {
	var size int64
	err := filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			size += info.Size()
		}
		return err
	})
	return size, err
}

// This is more efficient 
func DirSizes(path string) (int64, error) {
	var size int64
	err := filepath.WalkDir(path, func(p string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		info, erro := d.Info()
		if erro != nil {
			return erro
		}
		if !info.IsDir() {
			size += info.Size()
		}
		return err

	})
	return size, err
}

