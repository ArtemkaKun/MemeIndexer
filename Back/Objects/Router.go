package Objects

import (
	"github.com/gin-gonic/gin"
	"log"
)

type APIRouter struct {
	router *gin.Engine
}

func (apiRouter *APIRouter) InitializeRouter() {
	initializeRouter(apiRouter)
	setRoutingPaths(apiRouter)
}

func initializeRouter(apiRouter *APIRouter) {
	gin.SetMode(gin.ReleaseMode)

	apiRouter.router = gin.Default()

	const PathToMainPage string = "../Front/index.html"
	const PathToFrontResources = "../Front/Resources"

	apiRouter.router.Static("/Resources", PathToFrontResources)
	apiRouter.router.LoadHTMLGlob(PathToMainPage)
}

func setRoutingPaths(apiRouter *APIRouter) {
	setGetRequestEndpoints(apiRouter.router)
	setPostRequestEndpoints(apiRouter.router)
}

func setGetRequestEndpoints(router *gin.Engine) {
	router.GET("/", LoadMainPage)
	router.GET("/userAuth", AuthenticateUser)
	//router.GET("/meme", findMeme)
}

func setPostRequestEndpoints(router *gin.Engine) {
	//return router.POST("/meme", addMeme)
}

func (apiRouter *APIRouter) RunServer (port string) {
	log.Panic(apiRouter.router.Run(port))
}