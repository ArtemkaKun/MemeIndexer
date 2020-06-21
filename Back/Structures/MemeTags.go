package Structures

import (
	"encoding/json"
	"github.com/mxmCherry/translit/ruicao"
	"golang.org/x/text/transform"
	"strings"
)

type MemeTags struct {
	MainTags        []string `json:"mainTags"`
	AssociationTags []string `json:"associationTags"`
}

func (tags MemeTags) PrepareTagsForWork() {
	prepareTagsText(tags.MainTags)
	prepareTagsText(tags.AssociationTags)
}

func prepareTagsText(tagsArray []string) {
	for i, tag := range tagsArray {
		if len(tag) == 0 {
			continue
		}

		tagsArray[i] = strings.ToLower(tag)
		tagsArray[i] = strings.Trim(tag, " ")
		tagsArray[i] = transliterateTags(&tag)
	}
}

func transliterateTags(tag *string) (transliteratedTagText string) {
	transliteratedTagText, _, _ = transform.String(ruicao.ToLatin().Transformer(), *tag)
	return
}

func (tags MemeTags) ConvertTagsToJSON() (jsonMemeData []byte, err error) {
	jsonMemeData, err = json.Marshal(tags)
	return
}