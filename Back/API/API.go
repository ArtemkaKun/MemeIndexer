package API

import (
	"Back/Structures"
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
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
	memeToFind := GetAndPrepareMemeData(context)

	memeDataInJSON, err := memeToFind.ConvertTagsToJSON()
	if err != nil {
		returnErrorToClient(context, err)
		return
	}

	memeImage, err := MakeFindMemeRequestToDBServer(memeDataInJSON)
	if err != nil {
		returnErrorToClient(context, err)
		return
	}

	err = EncodeImageToBase64(context, memeImage)

	if err != nil {
		returnErrorToClient(context, err)
	}
}

func MakeFindMemeRequestToDBServer(memeDataInJSON []byte) (memeImage []byte, err error){
	requestToDBServer, err := PrepareHttpRequestToDBServer(memeDataInJSON, "GET")
	if err != nil {
		return
	}

	httpClient := &http.Client{}
	response, err := httpClient.Do(requestToDBServer)
	if err != nil {
		return
	}
	defer response.Body.Close()

	memeImage, _ = ioutil.ReadAll(response.Body)
	return
}

func PrepareHttpRequestToDBServer(memeDataInJSON []byte, requestType string) (requestToDBServer *http.Request, err error) {
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

	newMeme := Structures.Meme{
		MemeFile: file,
		MemeTags: GetAndPrepareMemeData(context),
	}

	memeDataInJSON, err := newMeme.ConvertTagsToJSON()
	if err != nil {
		returnErrorToClient(context, err)
		return
	}

	responseCode, err := MakeAddMemeRequestToDBServer(memeDataInJSON)

	if responseCode == "200 OK" {
		context.Status(http.StatusOK)
	} else {
		context.Status(http.StatusUnauthorized)
	}
}

func MakeAddMemeRequestToDBServer(memeDataInJSON []byte) (responseCode string, err error){
	requestToDBServer, err := PrepareHttpRequestToDBServer(memeDataInJSON, "POST")
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

func GetAndPrepareMemeData(context *gin.Context) (memeToFind Structures.MemeTags) {
	if context.Request.Method == "GET" {
		memeToFind = GetMemeDataFromGETRequest(context, memeToFind)
	} else {
		memeToFind = GetMemeDataFromPOSTRequest(context, memeToFind)
	}

	memeToFind.PrepareTagsForWork()

	return memeToFind
}

func GetMemeDataFromGETRequest(context *gin.Context, memeToFind Structures.MemeTags) Structures.MemeTags {
	memeToFind = Structures.MemeTags{
		MainTags:        strings.Split(context.Query("mainTags"), ","),
		AssociationTags: strings.Split(context.Query("associationTags"), ","),
	}
	return memeToFind
}

func GetMemeDataFromPOSTRequest(context *gin.Context, memeToFind Structures.MemeTags) Structures.MemeTags {
	memeToFind = Structures.MemeTags{
		MainTags:        strings.Split(context.PostForm("mainTags"), ","),
		AssociationTags: strings.Split(context.PostForm("associationTags"), ","),
	}
	return memeToFind
}

func EncodeImageToBase64(context *gin.Context, image []byte) (err error) {
	encoder := base64.NewEncoder(base64.StdEncoding, context.Writer)
	defer encoder.Close()

	_, err = encoder.Write(image)
	return
}

func returnErrorToClient(context *gin.Context, err error) {
	log.Print(err)
	context.Status(http.StatusInternalServerError)
}
