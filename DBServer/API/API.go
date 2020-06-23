package API

import (
	"DBServer/DB"
	"DBServer/MemeSeeker"
	"DBServer/Structures"
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

var Router *gin.Engine
const MemeImageFolderPath string = "./MemeImages/"

func init() {
	initRouter()
	setRouterPoints()
}

func initRouter() {
	gin.SetMode(gin.ReleaseMode)
	Router = gin.Default()
}

func setRouterPoints() {
	Router.GET("/userAuth", authenticateUser)

	Router.GET("/meme", findMeme)
	Router.POST("/meme", addMeme)
}

func authenticateUser(context *gin.Context) {
	receivedUserAuthData := getUserAuthDataForRequest(context)

	userId, err := DB.FindUser(receivedUserAuthData)
	if err != nil {
		sendUnauthorizedError(context, err)
	} else {
		context.String(http.StatusOK, "userId", userId)
	}
}

func getUserAuthDataForRequest(context *gin.Context) (receivedUserAuthData Structures.UserAuthData) {
	receivedUserAuthData = Structures.UserAuthData{
		Login:    context.Query("login"),
		Password: context.Query("pass"),
	}
	return
}

func sendUnauthorizedError(context *gin.Context, err error) {
	log.Println(err)
	context.Status(http.StatusUnauthorized)
}

func findMeme(context *gin.Context) {
	receivedMemeTags, err := getMemeTagsFromRequest(context)
	if err != nil {
		sendInternalServerError(context, err)
		return
	}

	memeFilePath, err := MemeSeeker.FindMeme(receivedMemeTags)
	if err != nil {
		sendInternalServerError(context, err)
		return
	}

	context.File(MemeImageFolderPath + memeFilePath)
}

func getMemeTagsFromRequest(context *gin.Context) (receivedMemeTags Structures.MemeTags, err error) {
	err = context.BindJSON(&receivedMemeTags)
	return
}

func addMeme(context *gin.Context) {
	meme, err := getMemeData(context)
	if err != nil {
		sendInternalServerError(context, err)
		return
	}

	fileName, err := writeMemeImageToStorage(meme)
	if err != nil {
		sendInternalServerError(context, err)
		return
	}

	meme.MemeFilePath = fileName

	err = DB.AddNewMeme(meme)
	if err != nil {
		sendInternalServerError(context, err)
	} else {
		context.Status(http.StatusOK)
	}
}

func getMemeData(context *gin.Context) (receivedMeme Structures.Meme, err error) {
	err = context.BindJSON(&receivedMeme)
	return
}

func writeMemeImageToStorage(meme Structures.Meme) (fileName string, err error) {
	data, err := base64.StdEncoding.DecodeString(meme.MemeFilePath)
	if err != nil {
		return
	}

	uploadMemeTime := time.Now()
	fileName = "meme" +
		strconv.Itoa(uploadMemeTime.Day()) +
		strconv.Itoa(uploadMemeTime.Hour()) +
		strconv.Itoa(uploadMemeTime.Minute()) +
		strconv.Itoa(uploadMemeTime.Second()) + ".png"

	err = ioutil.WriteFile(MemeImageFolderPath + fileName, data, 0644)
	return
}

func sendInternalServerError(context *gin.Context, err error) {
	log.Println(err)
	context.Header("err", err.Error())
	context.Status(http.StatusInternalServerError)
}
