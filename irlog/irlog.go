package irlog

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"sync"
	"time"
)

// SynchronousLog asd
const SynchronousLog LogModeType = 0

// AsynchronousLog asdasd
const AsynchronousLog LogModeType = 1

// LogModeType asdfq
type LogModeType int

// LogType struckt
type LogType struct {
	FilePath          string
	File              *os.File
	MaxFileSize       int
	CurrentFileSize   int
	LogLevel          int
	ArchiveFolderPath string
	Mutex             *sync.Mutex
	ModeType          LogModeType
	Buffer            *bufio.Writer
	MaxBufferSize     int
}

// InitializeLogger sevdiani
func (l *LogType) InitializeLogger(path string, maxFileSize int, level int, archiveFolderPath string, mode LogModeType, maxBufferSize int) {

	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		file = createNewFile(path)
	}
	l.File = file
	stat, err := file.Stat()
	if err != nil {
		panic(err)
	}
	l.ModeType = mode
	l.FilePath = path
	l.MaxFileSize = maxFileSize
	l.MaxBufferSize = maxBufferSize
	l.CurrentFileSize = int(stat.Size())
	l.LogLevel = level
	l.ArchiveFolderPath = archiveFolderPath
	l.Mutex = &sync.Mutex{}
	l.Buffer = bufio.NewWriterSize(l.File, maxBufferSize)
}

// ShutDown asdasd
func (l *LogType) ShutDown() {
	if l.Buffer.Size() > 0 {
		l.Buffer.Flush()
	}
}

// Log Data
func (l *LogType) Log(logText string, level int) {
	_, file, no, ok := runtime.Caller(1)
	if ok {
		fmt.Printf("called from %s#%d\n", file, no)
		logText = time.Now().Format("2006-01-02 15:04:05.000") + " " + file + " " + strconv.Itoa(no) + ": " + logText
	}
	l.Mutex.Lock()
	defer l.Mutex.Unlock()

	if level <= l.LogLevel == false {
		return
	}

	if l.ModeType == SynchronousLog {
		_, err := l.File.WriteString(logText + "\n")
		if err != nil {
			panic(err)
		}
		l.CurrentFileSize += len(logText) + 1
		l.checkFileSize()

	} else if l.ModeType == AsynchronousLog {
		_, err := l.Buffer.WriteString(logText + "\n")
		if err != nil {
			panic(err)
		}
		l.CurrentFileSize += len(logText) + 1
		l.checkFileSizeProd()

	} else {
		panic("Wrong Environment")
	}
}

// check file size buffer
func (l *LogType) checkFileSizeProd() {
	if l.CurrentFileSize > l.MaxFileSize {
		fmt.Println("created new file!!!")
		l.Buffer.Flush()
		l.File.Close()
		moveFileToArchive(l.FilePath, l.ArchiveFolderPath)
		l.File = createNewFile(l.FilePath)
		l.Buffer = bufio.NewWriterSize(l.File, l.MaxBufferSize)
		l.CurrentFileSize = 0
	}
}

// check file size
func (l *LogType) checkFileSize() {
	if l.CurrentFileSize > l.MaxFileSize {
		l.File.Close()
		moveFileToArchive(l.FilePath, l.ArchiveFolderPath)
		l.File = createNewFile(l.FilePath)
		l.CurrentFileSize = 0
	}
}

// create new file
func createNewFile(path string) *os.File {
	fl, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	return fl
}

// move full file to archive folder
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
