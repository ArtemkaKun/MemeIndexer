package Structures

type User struct {
	Id uint `json:"id" bson:"_id"`
	UserAuthData
}

type UserAuthData struct {
	Login    string `bson:"login"`
	Password string `bson:"pass"`
}

type Meme struct {
	MemeFilePath string `json:"memeFilePath" bson:"memeFilePath"`
	MainTags []string `json:"mainTags" bson:"mainTags"`
	AssociationTags []string `json:"associationTags" bson:"associationTags"`
}

type MemeTags struct {
	MainTags []string `json:"mainTags" bson:"mainTags"`
	AssociationTags []string `json:"associationTags" bson:"associationTags"`
}