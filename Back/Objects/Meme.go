package Objects

import (
	"encoding/base64"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Meme struct {
	MemeFile MemeFile `json:"memeFile,omitempty"`
	MemeTags MemeTags       `json:",omitempty"`
}

func (meme *Meme) FindMemeInDB(context *gin.Context) {
	meme.MemeTags.GetAndPrepareTagsFromRequest(context)
	tagsInJSON, errorMessage := meme.MemeTags.ConvertTagsToJSON()
	if errorMessage != "" {
		context.String(http.StatusInternalServerError, errorMessage)
		return
	}

	request, errorMessage := PrepareJSONRequest("GET", "meme", tagsInJSON)
	if errorMessage != "" {
		context.String(http.StatusInternalServerError, errorMessage)
		return
	}

	memeImageFromDB, errorMessage := MakeRequestToDBServer(request)
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
		return HandleCommonError(err)
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
	memeInJSON, errorMessage := meme.ConvertTagsToJSON()
	if errorMessage != "" {
		context.String(http.StatusInternalServerError, errorMessage)
		return
	}

	request, errorMessage := PrepareJSONRequest("POST", "meme", memeInJSON)
	if errorMessage != "" {
		context.String(http.StatusInternalServerError, errorMessage)
		return
	}

	_, errorMessage = MakeRequestToDBServer(request)
	if errorMessage != "" {
		context.String(http.StatusInternalServerError, errorMessage)
		return
	}

	context.Status(http.StatusOK)
}

func (meme *Meme) ConvertTagsToJSON() (jsonMemeData []byte, errorMessage string) {
	memeDataContainer := MemeData{
		MemeFile: meme.MemeFile.FileDataInBase64,
		MemeTags: meme.MemeTags,
	}

	jsonMemeData, err := json.Marshal(memeDataContainer)
	if err != nil {
		return nil, HandleCommonError(err)
	}
	return
}