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
	memeToFind := GetMemeData(context)

	memeImage, err := MakeFindRequestToDBServer(memeToFind)
	if err != nil {
		returnErrorToClient(context, err)
		return
	}

	err = EncodeImageToBase64(context, memeImage)

	if err != nil {
		returnErrorToClient(context, err)
	}
}

func GetMemeData(context *gin.Context) (memeToFind Structures.MemeTags) {
	memeToFind = Structures.MemeTags{
		MainTags:        strings.Split(context.Query("mainTags"), ","),
		AssociationTags: strings.Split(context.Query("associationTags"), ","),
	}

	memeToFind.PrepareTagsForWork()

	return memeToFind
}

func MakeFindRequestToDBServer(memeToFind Structures.MemeTags) (memeImage []byte, err error){
	requestToDBServer, err := PrepareHttpRequestToDBServer(memeToFind)
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

func PrepareHttpRequestToDBServer(memeToFind Structures.MemeTags) (requestToDBServer *http.Request, err error) {
	memeDataInJSON, err := memeToFind.ConvertTagsToJSON()
	if err != nil {
		return
	}

	requestToDBServer, err = http.NewRequest("GET", DBServerAddress+"meme", bytes.NewBuffer(memeDataInJSON))
	if err != nil {
		return
	}

	requestToDBServer.Header.Set("Content-Type", "application/json")
	return
}

func EncodeImageToBase64(context *gin.Context, image []byte) (err error) {
	encoder := base64.NewEncoder(base64.StdEncoding, context.Writer)
	defer encoder.Close()

	_, err = encoder.Write(image)
	return
}

func addMeme(context *gin.Context) {
	//I NEED THIS CODE COMMENTED!!!
	//file, err := context.FormFile("file")
	//if err != nil {
	//	log.Println(err)
	//	return
	//}

	//newMeme := Structures.Meme{
	//	MemeFile: file,
	//	MemeTags: GetMemeData(context),
	//}
}

func returnErrorToClient(context *gin.Context, err error) {
	log.Print(err)
	context.Status(http.StatusInternalServerError)
}
