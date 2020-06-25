package API

import (
	"Back/Objects"
)

var APIRouter Objects.APIRouter

func init() {
	APIRouter.InitializeRouter()
}

//func findMeme(context *gin.Context) {
//	memeToFind := getAndPrepareMemeData(context)
//
//	memeDataInJSON, err := memeToFind.ConvertTagsToJSON()
//	if err != nil {
//		returnErrorToClient(context, err)
//		return
//	}
//
//	memeImage, err := makeFindMemeRequestToDBServer(memeDataInJSON)
//	if err != nil {
//		returnErrorToClient(context, err)
//		return
//	}
//
//	err = encodeImageToBase64(context, memeImage)
//	if err != nil {
//		returnErrorToClient(context, err)
//	}
//	context.Status(http.StatusOK)
//}
//
//func makeFindMemeRequestToDBServer(memeDataInJSON []byte) (memeImage []byte, err error){
//	requestToDBServer, err := prepareHttpRequestToDBServer(memeDataInJSON, "GET")
//	if err != nil {
//		return
//	}
//
//	httpClient := &http.Client{}
//	response, err := httpClient.Do(requestToDBServer)
//	if err != nil {
//		return
//	}
//
//	if response.StatusCode == 500 {
//		if response.Header.Get("err") == "Nothing was found!" {
//			err = fmt.Errorf(MemeNotFound)
//		} else {
//			err = fmt.Errorf(DefaultError)
//		}
//		return
//	}
//
//	defer response.Body.Close()
//
//	memeImage, _ = ioutil.ReadAll(response.Body)
//	return
//}
//
//func prepareHttpRequestToDBServer(memeDataInJSON []byte, requestType string) (requestToDBServer *http.Request, err error) {
//	requestToDBServer, err = http.NewRequest(requestType, DBServerAddress+"meme", bytes.NewBuffer(memeDataInJSON))
//	if err != nil {
//		return
//	}
//
//	requestToDBServer.Header.Set("Content-Type", "application/json")
//	return
//}
//
//func encodeImageToBase64(context *gin.Context, image []byte) (err error) {
//	encoder := base64.NewEncoder(base64.StdEncoding, context.Writer)
//	defer encoder.Close()
//
//	_, err = encoder.Write(image)
//	return
//}
//
//func addMeme(context *gin.Context) {
//	file, err := context.FormFile("file")
//	if err != nil {
//		returnErrorToClient(context, err)
//		return
//	}
//
//	memeFile, err := generateMemeImageFile(file, context)
//	if err != nil {
//		returnErrorToClient(context, err)
//		return
//	}
//
//	newMeme := Structures.Meme{
//		MemeFileBase64: base64.StdEncoding.EncodeToString(memeFile.FileData),
//		MemeTags:       getAndPrepareMemeData(context),
//	}
//
//	err = memeFile.DeleteMemeFile()
//	if err != nil {
//		returnErrorToClient(context, err)
//		return
//	}
//
//	memeDataInJSON, err := newMeme.ConvertTagsToJSON()
//	if err != nil {
//		returnErrorToClient(context, err)
//		return
//	}
//
//	err = makeAddMemeRequestToDBServer(memeDataInJSON)
//	if err != nil {
//		returnErrorToClient(context, err)
//		return
//	}
//
//	context.Status(http.StatusOK)
//}
//
//func generateMemeImageFile(file *multipart.FileHeader, context *gin.Context) (memeImage Structures.MemeFile, err error){
//	memeImage.GenerateFileName(file.Filename)
//
//	err = context.SaveUploadedFile(file, memeImage.FileName)
//	if err != nil {
//		return
//	}
//
//	err = memeImage.ReadMemeFileData()
//	return
//}
//
//func getAndPrepareMemeData(context *gin.Context) (memeToFind Structures.MemeTags) {
//	if context.Request.Method == "GET" {
//		memeToFind = getMemeDataFromGETRequest(context, memeToFind)
//	} else {
//		memeToFind = getMemeDataFromPOSTRequest(context, memeToFind)
//	}
//
//	memeToFind.PrepareTagsForWork()
//
//	return memeToFind
//}
//
//func getMemeDataFromGETRequest(context *gin.Context, memeToFind Structures.MemeTags) Structures.MemeTags {
//	memeToFind = Structures.MemeTags{
//		MainTags:        strings.Split(context.Query("mainTags"), ","),
//		AssociationTags: strings.Split(context.Query("associationTags"), ","),
//	}
//	return memeToFind
//}
//
//func getMemeDataFromPOSTRequest(context *gin.Context, memeToFind Structures.MemeTags) Structures.MemeTags {
//	memeToFind = Structures.MemeTags{
//		MainTags:        strings.Split(context.PostForm("mainTags"), ","),
//		AssociationTags: strings.Split(context.PostForm("associationTags"), ","),
//	}
//	return memeToFind
//}
//
//func makeAddMemeRequestToDBServer(memeDataInJSON []byte) (err error){
//	requestToDBServer, err := prepareHttpRequestToDBServer(memeDataInJSON, "POST")
//	if err != nil {
//		return
//	}
//
//	httpClient := &http.Client{}
//	response, err := httpClient.Do(requestToDBServer)
//	if err != nil {
//		return
//	}
//
//	if response.StatusCode != 200 {
//		err = fmt.Errorf(DefaultError)
//	}
//	return
//}
//
//func returnAuthErrorToClient(context *gin.Context, err error) {
//	context.String(http.StatusInternalServerError, err.Error())
//}
//
//func returnErrorToClient(context *gin.Context, err error) {
//	log.Print(err)
//	context.String(http.StatusInternalServerError, err.Error())
//}
