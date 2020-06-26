package API

import (
	"Back/Meme"
	"Back/User"
	"github.com/gin-gonic/gin"
	"net/http"
)

var Router APIRouter

func init() {
	Router.InitializeRouter()
}

func loadMainPage(context *gin.Context) {
	context.HTML(http.StatusOK, "index.html", gin.H{})
}

func authenticateUser(context *gin.Context) {
	var user User.User
	user.AuthenticateUser(context)
}

func findMeme(context *gin.Context) {
	var meme Meme.Meme
	meme.FindMemeInDB(context)
}

func addMeme(context *gin.Context) {
	var meme Meme.Meme
	meme.InsertMemeInDB(context)
}
