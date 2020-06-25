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

func LoadMainPage(context *gin.Context) {
	context.HTML(http.StatusOK, "index.html", gin.H{})
}

func AuthenticateUser(context *gin.Context) {
	var user User.User
	user.AuthenticateUser(context)
}

func FindMeme(context *gin.Context) {
	var meme Meme.Meme
	meme.FindMemeInDB(context)
}

func AddMeme(context *gin.Context) {
	var meme Meme.Meme
	meme.InsertMemeInDB(context)
}
