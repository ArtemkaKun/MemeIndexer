package Meme

import (
	"Back/Error"
	"Back/Request"
	"encoding/base64"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Meme struct {
	MemeFile MemeFile `json:"memeFile,omitempty"`
	MemeTags MemeTags `json:",omitempty"`
}

func (meme *Meme) FindMemeInDB(context *gin.Context) {
	meme.MemeTags.GetAndPrepareTagsFromRequest(context)
	tagsInJSON, errorMessage := meme.MemeTags.ConvertTagsToJSON()
	if errorMessage != "" {
		context.String(http.StatusInternalServerError, errorMessage)
		return
	}

	request, errorMessage := Request.PrepareJSONRequest("GET", "meme", tagsInJSON)
	if errorMessage != "" {
		context.String(http.StatusInternalServerError, errorMessage)
		return
	}

	memeImageFromDB, errorMessage := Request.MakeRequestToDBServer(request)
	if errorMessage != "" {
		context.String(http.StatusInternalServerError, errorMessage)
		return
	}

	errorMessage = convertImageToBase64AndSendToClient(context, memeImageFromDB)
	if errorMessage != "" {
		context.String(http.StatusInternalServerError, errorMessage)
		return
	}

	context.Status(http.StatusOK)
}

func convertImageToBase64AndSendToClient(context *gin.Context, image []byte) (errorMessage string) {
	encoder := base64.NewEncoder(base64.StdEncoding, context.Writer)
	defer encoder.Close()

	_, err := encoder.Write(image)
	if err != nil {
		return Error.HandleCommonError(err)
	}

	return
}

func (meme *Meme) InsertMemeInDB(context *gin.Context) {
	errorMessage := meme.MemeFile.GetAndPrepareMemeFileFromRequest(context)
	if errorMessage != "" {
		context.String(http.StatusInternalServerError, errorMessage)
		return
	}

	meme.MemeTags.GetAndPrepareTagsFromRequest(context)
	memeInJSON, errorMessage := meme.ConvertMemeToJSON()
	if errorMessage != "" {
		context.String(http.StatusInternalServerError, errorMessage)
		return
	}

	request, errorMessage := Request.PrepareJSONRequest("POST", "meme", memeInJSON)
	if errorMessage != "" {
		context.String(http.StatusInternalServerError, errorMessage)
		return
	}

	_, errorMessage = Request.MakeRequestToDBServer(request)
	if errorMessage != "" {
		context.String(http.StatusInternalServerError, errorMessage)
		return
	}

	context.Status(http.StatusOK)
}

func (meme *Meme) ConvertMemeToJSON() (jsonMemeData []byte, errorMessage string) {
	memeDataContainer := MemeData{
		MemeFile: meme.MemeFile.FileDataInBase64,
		MemeTags: meme.MemeTags,
	}

	jsonMemeData, err := json.Marshal(memeDataContainer)
	if err != nil {
		return nil, Error.HandleCommonError(err)
	}
	return
}
