package Structures

import "mime/multipart"

type Meme struct {
	MemeFile *multipart.FileHeader
	MemeTags
}

type MemeTags struct {
	MainTags []string `json:"mainTags"`
	AssociationTags []string `json:"associationTags"`
}