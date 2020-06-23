package Structures

import (
	"io/ioutil"
	"os"
	"strconv"
	"time"
)

type MemeFile struct {
	FileName string
	FileData []byte
}

func (file *MemeFile) GenerateFileName(filename string) {
	uploadMemeTime := time.Now()
	file.FileName = "memeTemp" + filename +
		strconv.Itoa(uploadMemeTime.Day()) +
		strconv.Itoa(uploadMemeTime.Hour()) +
		strconv.Itoa(uploadMemeTime.Minute()) +
		strconv.Itoa(uploadMemeTime.Second())
}

func (file *MemeFile) ReadMemeFileData() (err error) {
	file.FileData, err = ioutil.ReadFile(file.FileName)
	return
}

func (file *MemeFile) DeleteMemeFile() (err error) {
	err = os.Remove(file.FileName)
	return
}
