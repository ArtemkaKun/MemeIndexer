package DB

import (
	"DBServer/Structures"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var DatabaseConnections = Structures.DBConnections{}

func init() {
	DatabaseConnections.InitializeDBConnections()
}

func FindUser(loginData Structures.UserAuthData) (userId uint, err error) {
	foundedUser := new(Structures.User)

	err = DatabaseConnections.UsersCollection.FindOne(context.Background(), loginData).Decode(&foundedUser)
	if err != nil {
		return
	}

	if foundedUser.Id == 0 {
		return 0, fmt.Errorf("Log in data incorrect!")
	}
	return foundedUser.Id, nil
}

func FindMemesByMainTagsWithAtLeastOneTag(mainTagsFromRequest []string) (memesWithAtLeastOneTag []Structures.Meme, err error) {
	searchFilter := bson.M{"mainTags": bson.M{"$in": mainTagsFromRequest}}

	foundedDocuments, err := DatabaseConnections.MemesCollection.Find(context.Background(), searchFilter)
	if err != nil {
		return
	}
	defer foundedDocuments.Close(context.Background())

	memesWithAtLeastOneTag, err = storeFoundedMemesInSlice(foundedDocuments)
	if len(memesWithAtLeastOneTag) == 0 && err == nil {
		return nil, fmt.Errorf("Nothing was found!")
	}

	return
}

func storeFoundedMemesInSlice(foundedDocuments *mongo.Cursor) (memesSlice []Structures.Meme, err error) {
	for foundedDocuments.Next(context.Background()) {
		var meme Structures.Meme

		err = foundedDocuments.Decode(&meme)
		if err != nil {
			return
		}

		memesSlice = append(memesSlice, meme)
	}

	err = foundedDocuments.Err()
	return
}
