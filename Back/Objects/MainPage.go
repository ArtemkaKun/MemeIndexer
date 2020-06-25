package Objects

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func LoadMainPage(context *gin.Context) {
	context.HTML(http.StatusOK, "index.html", gin.H{})
}

func AuthenticateUser(context *gin.Context) {
	var user User
	user.AuthenticateUser(context)
}

func FindMeme(context *gin.Context) {
	var meme Meme
	meme.FindMemeInDB(context)
}

func AddMeme(context *gin.Context) {
	var meme Meme
	meme.InsertMemeInDB(context)
}