package irlog

import (
	"errors"
	"fmt"
	"os"
	"sync"
	"syscall"
	"time"
)

// var LogTypes = [5]string{"Error", "Warning", "Mesame", "Meotxe", "Mexute"}

//LogData structure
type LogData struct {
	Path            string
	File            *os.File
	ArchivePath     string
	LogLevel        int
	MaxFileSize     int
	CurrentFileSize int
	mutex           sync.Mutex
}

//InitializeLogger function
func InitializeLogger(path string, archivePath string, level int, maxFileSize int) (*LogData, error) {

	logger := &LogData{
		Path:            path,
		File:            nil,
		ArchivePath:     archivePath,
		LogLevel:        level,
		MaxFileSize:     maxFileSize,
		CurrentFileSize: 0,
		mutex:           sync.Mutex{},
	}

	file, err := os.OpenFile(logger.Path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	logger.File = file

	if f, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		logger.createNewFile(path)
	} else {
		logger.CurrentFileSize = int(f.Size())
	}

	// fmt.Println("aqamde")

	return logger, nil
}

//Log Main Function
func (logger *LogData) Log(logType int, logText string) error {
	// fmt.Println("lock aqamde")
	logger.mutex.Lock()
	defer logger.mutex.Unlock()

	if logger.LogLevel < logType {
		return nil
	}

	if logger.CurrentFileSize+len(logText) >= logger.MaxFileSize {

		logger.File.Close()
		moveFullFile(logger.Path, logger.ArchivePath)
		logger.createNewFile(logger.Path)
	}

	// lock
	syscall.Flock(int(logger.File.Fd()), syscall.LOCK_EX)

	logger.File.Write([]byte(logText + "\n"))

	syscall.Flock(int(logger.File.Fd()), syscall.LOCK_UN)

	logger.CurrentFileSize += len(logText) + 1

	return nil
}

func (logger *LogData) createNewFile(path string) error {
	f, err := os.Create(path)
	if err != nil {
		fmt.Println("Can't create file on " + path + " directory")
		return err
	}
	fmt.Println("File Created Successfully!")
	logger.CurrentFileSize = 0
	logger.File = f
	return nil
}

func moveFullFile(oldPath string, newPath string) error {
	newPath += time.Now().Format("2006-01-02 15:04:05.000")
	err := os.Rename(oldPath, newPath)
	return err
}
