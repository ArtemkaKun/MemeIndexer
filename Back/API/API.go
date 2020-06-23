package API

import (
	"Back/Structures"
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"strings"
)

var Router *gin.Engine

const (
	DBServerAddress string = "http://localhost:8001/"
)

func init() {
	initializeRouter()

	setRoutingPaths()
}

func initializeRouter() {
	gin.SetMode(gin.ReleaseMode)
	Router = gin.Default()

	Router.Static("/Resources", "../Front/Resources")
	Router.LoadHTMLGlob("../Front/index.html")
}

func setRoutingPaths() {
	Router.GET("/", loadMainPage)
	Router.GET("/userAuth", authenticateUser)

	Router.GET("/meme", findMeme)
	Router.POST("/meme", addMeme)
}

func loadMainPage(context *gin.Context) {
	context.HTML(http.StatusOK, "index.html", gin.H{})
}

func authenticateUser(context *gin.Context) {
	receivedLogin := context.Query("login")
	receivedPass := context.Query("pass")

	userAuthAddress := fmt.Sprintf("userAuth?login=%v&pass=%v", receivedLogin, receivedPass)
	response, err := http.Get(DBServerAddress + userAuthAddress)
	if err != nil {
		returnErrorToClient(context, err)
		return
	}

	if response.Status == "200 OK" {
		context.Status(http.StatusOK)
	} else {
		context.Status(http.StatusUnauthorized)
	}
}

func findMeme(context *gin.Context) {
	memeToFind := getAndPrepareMemeData(context)

	memeDataInJSON, err := memeToFind.ConvertTagsToJSON()
	if err != nil {
		returnErrorToClient(context, err)
		return
	}

	memeImage, err := makeFindMemeRequestToDBServer(memeDataInJSON)
	if err != nil {
		returnErrorToClient(context, err)
		return
	}

	err = encodeImageToBase64(context, memeImage)

	if err != nil {
		returnErrorToClient(context, err)
	}
}

func makeFindMemeRequestToDBServer(memeDataInJSON []byte) (memeImage []byte, err error){
	requestToDBServer, err := prepareHttpRequestToDBServer(memeDataInJSON, "GET")
	if err != nil {
		return
	}

	httpClient := &http.Client{}
	response, err := httpClient.Do(requestToDBServer)
	if err != nil {
		return
	}

	if response.StatusCode == 500 {
		err = fmt.Errorf(response.Header.Get("err"))
		return
	}

	defer response.Body.Close()

	memeImage, _ = ioutil.ReadAll(response.Body)
	return
}

func prepareHttpRequestToDBServer(memeDataInJSON []byte, requestType string) (requestToDBServer *http.Request, err error) {
	requestToDBServer, err = http.NewRequest(requestType, DBServerAddress+"meme", bytes.NewBuffer(memeDataInJSON))
	if err != nil {
		return
	}

	requestToDBServer.Header.Set("Content-Type", "application/json")
	return
}

func addMeme(context *gin.Context) {
	file, err := context.FormFile("file")
	if err != nil {
		returnErrorToClient(context, err)
		return
	}

	memeFile, err := generateMemeImageFile(file, context)
	if err != nil {
		returnErrorToClient(context, err)
		return
	}

	newMeme := Structures.Meme{
		MemeFileBase64: base64.StdEncoding.EncodeToString(memeFile.FileData),
		MemeTags:       getAndPrepareMemeData(context),
	}

	err = memeFile.DeleteMemeFile()
	if err != nil {
		log.Panic(err)
	}

	memeDataInJSON, err := newMeme.ConvertTagsToJSON()
	if err != nil {
		returnErrorToClient(context, err)
		return
	}

	responseCode, err := makeAddMemeRequestToDBServer(memeDataInJSON)

	if responseCode == "200 OK" {
		context.Status(http.StatusOK)
	} else {
		context.Status(http.StatusUnauthorized)
	}
}

func generateMemeImageFile(file *multipart.FileHeader, context *gin.Context) (memeImage Structures.MemeFile, err error){
	memeImage.GenerateFileName(file.Filename)

	err = context.SaveUploadedFile(file, memeImage.FileName)
	if err != nil {
		return
	}

	err = memeImage.ReadMemeFileData()
	return
}

func makeAddMemeRequestToDBServer(memeDataInJSON []byte) (responseCode string, err error){
	requestToDBServer, err := prepareHttpRequestToDBServer(memeDataInJSON, "POST")
	if err != nil {
		return
	}

	httpClient := &http.Client{}
	response, err := httpClient.Do(requestToDBServer)
	if err != nil {
		return
	}

	return response.Status, nil
}

func getAndPrepareMemeData(context *gin.Context) (memeToFind Structures.MemeTags) {
	if context.Request.Method == "GET" {
		memeToFind = getMemeDataFromGETRequest(context, memeToFind)
	} else {
		memeToFind = getMemeDataFromPOSTRequest(context, memeToFind)
	}

	memeToFind.PrepareTagsForWork()

	return memeToFind
}

func getMemeDataFromGETRequest(context *gin.Context, memeToFind Structures.MemeTags) Structures.MemeTags {
	memeToFind = Structures.MemeTags{
		MainTags:        strings.Split(context.Query("mainTags"), ","),
		AssociationTags: strings.Split(context.Query("associationTags"), ","),
	}
	return memeToFind
}

func getMemeDataFromPOSTRequest(context *gin.Context, memeToFind Structures.MemeTags) Structures.MemeTags {
	memeToFind = Structures.MemeTags{
		MainTags:        strings.Split(context.PostForm("mainTags"), ","),
		AssociationTags: strings.Split(context.PostForm("associationTags"), ","),
	}
	return memeToFind
}

func encodeImageToBase64(context *gin.Context, image []byte) (err error) {
	encoder := base64.NewEncoder(base64.StdEncoding, context.Writer)
	defer encoder.Close()

	_, err = encoder.Write(image)
	return
}

func returnErrorToClient(context *gin.Context, err error) {
	log.Print(err)
	context.Header("err", err.Error())
	context.Status(http.StatusInternalServerError)
}
