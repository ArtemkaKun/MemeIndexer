package Structures

import (
	"encoding/json"
	"mime/multipart"
)

type Meme struct {
	MemeFile *multipart.FileHeader `json:"memeImage,omitempty"`
	MemeTags	`json:",omitempty"`
}

func (meme Meme) ConvertTagsToJSON() (jsonMemeData []byte, err error) {
	jsonMemeData, err = json.Marshal(meme)
	return
}