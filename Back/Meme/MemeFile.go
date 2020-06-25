package Meme

import (
	"Back/Error"
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"math/rand"
	"os"
	"strconv"
	"time"
)

type MemeFile struct {
	FileName         string
	FileData         []byte
	FileDataInBase64 string
}

func (file *MemeFile) GetAndPrepareMemeFileFromRequest(context *gin.Context) (errorMessage string) {
	fileFromRequest, err := context.FormFile("file")
	if err != nil {
		return Error.HandleCommonError(err)
	}

	file.generateFileName(fileFromRequest.Filename)

	err = context.SaveUploadedFile(fileFromRequest, file.FileName)
	if err != nil {
		return Error.HandleCommonError(err)
	}

	err = file.readMemeFileData()
	if err != nil {
		return Error.HandleCommonError(err)
	}

	file.prepareMemeData()

	err = file.deleteMemeFile()
	if err != nil {
		return Error.HandleCommonError(err)
	}
	return
}

func (file *MemeFile) generateFileName(filename string) {
	uploadMemeTime := time.Now()
	file.FileName = "memeTemp" +
		strconv.Itoa(uploadMemeTime.Day()) +
		strconv.Itoa(uploadMemeTime.Hour()) +
		strconv.Itoa(uploadMemeTime.Minute()) +
		strconv.Itoa(uploadMemeTime.Second()) +
		strconv.Itoa(rand.Int()) + filename
}

func (file *MemeFile) readMemeFileData() (err error) {
	file.FileData, err = ioutil.ReadFile(file.FileName)
	return
}

func (file *MemeFile) deleteMemeFile() (err error) {
	err = os.Remove(file.FileName)
	return
}

func (file *MemeFile) prepareMemeData() {
	file.FileDataInBase64 = base64.StdEncoding.EncodeToString(file.FileData)
}
