package Structures

import (
	"mime/multipart"
)

type Meme struct {
	MemeFile *multipart.FileHeader
	MemeTags
}
