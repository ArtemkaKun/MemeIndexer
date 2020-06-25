package Objects

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/mxmCherry/translit/ruicao"
	"golang.org/x/text/transform"
	"strings"
)

type MemeTags struct {
	MainTags        []string `json:"mainTags"`
	AssociationTags []string `json:"associationTags"`
}

func (tags *MemeTags) GetAndPrepareTagsFromRequest(context *gin.Context) {
	if context.Request.Method == "GET" {
		tags.MainTags = strings.Split(context.Query("mainTags"), ",")
		tags.AssociationTags = strings.Split(context.Query("associationTags"), ",")
	} else {
		tags.MainTags = strings.Split(context.PostForm("mainTags"), ",")
		tags.AssociationTags = strings.Split(context.PostForm("associationTags"), ",")
	}
	tags.prepareTagsForWork()
}

func (tags *MemeTags) prepareTagsForWork() {
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

func (tags MemeTags) ConvertTagsToJSON() (jsonMemeData []byte, errorMessage string) {
	jsonMemeData, err := json.Marshal(tags)
	if err != nil {
		return nil, HandleCommonError(err)
	}
	return
}
