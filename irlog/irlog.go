package irlog

import (
	"fmt"
	"os"
	"sync"
	"time"
)

// LogType struckt
type LogType struct {
	FilePath          string
	File              *os.File
	MaxFileSize       int
	CurrentFileSize   int
	LogLevel          int
	ArchiveFolderPath string
	Mutex             *sync.Mutex
	Junk              int
}

// InitializeLogger sevdiani
func (l *LogType) InitializeLogger(path string, maxSize int, level int, archiveFolderPath string) {
	l.FilePath = path
	l.MaxFileSize = maxSize
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		file = createNewFile(path)
	}
	l.File = file
	stat, err := file.Stat()
	if err != nil {
		panic(err)
	}
	l.CurrentFileSize = int(stat.Size())
	l.LogLevel = level
	l.ArchiveFolderPath = archiveFolderPath
	l.Mutex = &sync.Mutex{}
	fmt.Println("Came here")
}

// Log Data
func (l *LogType) Log(logText string, level int) {
	// l.Junk++
	l.Mutex.Lock()

	if level > l.LogLevel {
		fmt.Println(level, l.LogLevel)
		return
	}

	if l.CurrentFileSize+len(logText) > l.MaxFileSize {
		l.File.Close()
		moveFileToArchive(l.FilePath, l.ArchiveFolderPath)
		l.File = createNewFile(l.FilePath)
		l.CurrentFileSize = 0
	}

	_, err := l.File.Write([]byte(logText + "\n"))
	if err != nil {
		panic(err)
	}
	l.CurrentFileSize += len(logText) + 1
	l.Mutex.Unlock()

}

func createNewFile(path string) *os.File {
	fl, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	return fl
}

func moveFileToArchive(oldPath string, newPath string) {
	newFilePath := newPath + time.Now().Format("2006-01-02 15:04:05.000")
	if _, err := os.Stat(newFilePath); err == nil {
		newFilePath += "Damtxveva"
	}
	err := os.Rename(oldPath, newFilePath)
	if err != nil {
		panic(err)
	}
}
