package API

import (
	"Back/Objects"
)

var APIRouter Objects.APIRouter

func init() {
	APIRouter.InitializeRouter()
}

