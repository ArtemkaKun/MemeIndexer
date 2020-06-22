package Structures

type User struct {
	Id uint `json:"id,omitempty" bson:"_id,omitempty"`
	UserAuthData `json:",omitempty" bson:",omitempty"`
}

type UserAuthData struct {
	Login    string `bson:"login,omitempty"`
	Password string `bson:"pass,omitempty"`
}