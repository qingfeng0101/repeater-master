package logging

import (
	"errors"
	"io"
	"strconv"

	"os"
	"path"
	"time"
)

type logfile struct {
	file        *os.File
	isOpen      bool
	rolte       <-chan time.Time
	preFileName string
}

func (lf *logfile) Write(b []byte) (int, error) {
	select {
	case <-lf.rolte:

		err := lf.SetFile(lf.preFileName + time.Now().Format("2006-01-02") + ".log")
		if err != nil {
			return 0, err
		}

	default:
		if !lf.isOpen {
			return 0, errors.New("logfile is not config")
		}
	}

	return lf.file.Write(b)
}

func newLogFile(preFileName string) *logfile {
	var lf logfile
	lf.preFileName = preFileName
	err := lf.SetFile(lf.preFileName + time.Now().Format("2006-01-02") + ".log")
	if err != nil {
		panic(err)
	}

	go func() {
		//用于第一次凌晨的日志滚动
		t := time.Now()
		nx := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location()).AddDate(0, 0, 1)
		<-time.After(nx.Sub(t))
		lf.SetFile(lf.preFileName + time.Now().Format("2006-01-02") + ".log")
		//后续24小时滚动一次
		lf.rolte = time.NewTicker(time.Hour * 24).C
	}()

	return &lf
}

func (lf *logfile) SetFile(filename string) error {
	var filepath string
	wd, err := os.Getwd()
	if err == nil {
		filepath = path.Join(wd, "logs/"+filename)
	} else {
		panic(err)
	}

	if pathIsExist(filepath) {
		err := os.Rename(filepath, filepath+strconv.FormatInt(time.Now().Unix(), 10))
		if err != nil {
			return err
		}
	}
	if lf.isOpen {
		lf.Close()
	}
	file, err := os.OpenFile(filepath, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0664)
	if err != nil {
		return err
	}
	lf.file = file
	lf.isOpen = true
	return nil
}

func (lf *logfile) Close() {
	lf.file.Close()
}

var fli io.Writer

func init() {
	if !pathIsExist("logs") {
		err := os.Mkdir("logs", os.ModePerm)
		if err != nil {
			panic(err)
		}
	}
}

func LoggerWriter() io.Writer {
	if fli == nil {
		fli = newLogFile("repeater")
	}
	return fli
}

func pathIsExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}
