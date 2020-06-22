package Structures

type Meme struct {
	MemeFilePath string `json:"memeFilePath,omitempty" bson:"memeFilePath,omitempty"`
	MainTags []string `json:"mainTags,omitempty" bson:"mainTags,omitempty"`
	AssociationTags []string `json:"associationTags,omitempty" bson:"associationTags,omitempty"`
}

type MemeTags struct {
	MainTags []string `json:"mainTags,omitempty" bson:"mainTags,omitempty"`
	AssociationTags []string `json:"associationTags,omitempty" bson:"associationTags,omitempty"`
}