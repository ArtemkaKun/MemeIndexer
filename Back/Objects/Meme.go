package Objects

import (
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Meme struct {
	MemeFileContent string `json:"memeFile,omitempty"`
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