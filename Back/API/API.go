package API

import (
	"Back/Structures"
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mxmCherry/translit/ruicao"
	"golang.org/x/text/transform"
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

	requestToDBServer, err := PrepareHttpRequestToDBServer(memeToFind)
	if err != nil {
		returnErrorToClient(context, err)
		return
	}

	httpClient := &http.Client{}
	response, err := httpClient.Do(requestToDBServer)
	if err != nil {
		returnErrorToClient(context, err)
		return
	}
	defer response.Body.Close()

	memeImage, _ := ioutil.ReadAll(response.Body)
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

	prepareTagsText(memeToFind.MainTags, memeToFind.AssociationTags)

	return memeToFind
}

func prepareTagsText(tagsArrays... []string) {
	for _, tagsArray := range tagsArrays {
		for i, tag := range tagsArray {
			if len(tag) == 0 {
				continue
			}

			tagsArray[i] = strings.ToLower(tag)
			tagsArray[i] = strings.TrimSpace(tag)
			tagsArray[i] = transliterateTags(&tag)
		}
	}
}

func transliterateTags(tag *string) (transliteratedTagText string) {
	transliteratedTagText, _, _ = transform.String(ruicao.ToLatin().Transformer(), *tag)
	return
}

func PrepareHttpRequestToDBServer(memeToFind Structures.MemeTags) (requestToDBServer *http.Request, err error) {
	jsonMemeData, err := ConvertDataToJSON(memeToFind)
	if err != nil {
		return
	}

	requestToDBServer, err = http.NewRequest("GET", DBServerAddress+"meme", bytes.NewBuffer(jsonMemeData))
	if err != nil {
		return
	}

	requestToDBServer.Header.Set("Content-Type", "application/json")
	return
}

func ConvertDataToJSON(dataToConvert interface{}) (jsonMemeData []byte, err error) {
	jsonMemeData, err = json.Marshal(dataToConvert)
	return
}

func EncodeImageToBase64(context *gin.Context, image []byte) (err error) {
	encoder := base64.NewEncoder(base64.StdEncoding, context.Writer)
	defer encoder.Close()

	_, err = encoder.Write(image)
	return
}

func addMeme(context *gin.Context) {
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
