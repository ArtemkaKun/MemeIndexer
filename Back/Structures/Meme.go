package Structures

import (
	"encoding/json"
)

type Meme struct {
	MemeFileBase64 string `json:"memeFile,omitempty"`
	MemeTags       `json:",omitempty"`
}

func (meme Meme) ConvertTagsToJSON() (jsonMemeData []byte, err error) {
	jsonMemeData, err = json.Marshal(meme)
	return
}