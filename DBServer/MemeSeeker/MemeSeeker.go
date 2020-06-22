package MemeSeeker

import (
	"DBServer/DB"
	"DBServer/Structures"
	"github.com/juliangruber/go-intersect"
)

func FindMeme(memeTagsFromRequest Structures.MemeTags) (memeImagePath string, err error) {
	foundedMemes, err := FindMemesWithMainTagsMatches(memeTagsFromRequest.MainTags)
	if err != nil {
		return
	}

	if needToSearchByAssociationTags(memeTagsFromRequest, foundedMemes) {
		foundedMemes, err = findMemesWithTheMostTagsMatch(foundedMemes, memeTagsFromRequest.AssociationTags, false)
		if err != nil {
			return
		}
	}

	return foundedMemes[0].MemeFilePath, nil
}

func FindMemesWithMainTagsMatches(mainTagsFromRequest []string) (foundedMemes []Structures.Meme, err error) {
	memesWithAtLeastOneTagMatch, err := DB.FindMemesByMainTagsWithAtLeastOneTag(mainTagsFromRequest)
	if err != nil {
		return
	}

	foundedMemes, err = findMemesWithTheMostTagsMatch(memesWithAtLeastOneTagMatch, mainTagsFromRequest, true)
	return
}

func findMemesWithTheMostTagsMatch(memes []Structures.Meme, tagsFromRequest []string, matchByMainTags bool) (memesWithTheMostMatches []Structures.Meme, err error) {
	var theHighestMatchesCount int

	for _, meme := range memes {
		var countOfMatches int

		if matchByMainTags {
			countOfMatches = findCountOfMatches(meme.MainTags, tagsFromRequest)
		} else {
			countOfMatches = findCountOfMatches(meme.AssociationTags, tagsFromRequest)
		}

		if theHighestMatchesCount < countOfMatches {
			theHighestMatchesCount = countOfMatches
			memesWithTheMostMatches = []Structures.Meme{meme}
		} else if theHighestMatchesCount == countOfMatches {
			memesWithTheMostMatches = append(memesWithTheMostMatches, meme)
		}
	}
	return
}

func findCountOfMatches(actualMemeTags, tagsFromRequest []string) (lengthOfIntersection int) {
	tagsMatches := intersect.Simple(actualMemeTags, tagsFromRequest)

	lengthOfIntersection = len(tagsMatches.([]interface{}))
	return
}

func needToSearchByAssociationTags(memeTagsFromRequest Structures.MemeTags, foundedMemes []Structures.Meme) bool {
	return len(foundedMemes) != 1 && len(memeTagsFromRequest.AssociationTags) != 0
}
