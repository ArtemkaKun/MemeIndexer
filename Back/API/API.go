package API

import (
	"Back/Objects"
)

var APIRouter Objects.APIRouter

func init() {
	APIRouter.InitializeRouter()
}

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
//
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
