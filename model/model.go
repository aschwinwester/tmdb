package model

import (
	"strings"
)

type Movie struct {
	Id string
	Title string
	Actors []Actor
}

func(movie Movie) HasActors() bool {
	return movie.Actors != nil && len(movie.Actors) > 0
}

func(movie Movie) SearchActorsByFirstName(firstName string) []Actor {
	actors := make([]Actor, 0)
	for _, actor := range movie.Actors {
		
		names :=strings.Split(actor.Name, " ")
		if (len(names) > 0 && names[0] == firstName) {
			actors = append(actors, actor )
		}
	
	}
	return actors;
}

func(movie Movie) SearchCharacter(character string) []Actor {
	actors := make([]Actor, 0)
	for _, actor := range movie.Actors {
		
		isCharacter := strings.Contains(actor.Character, character)
		if (isCharacter) {
			actors = append(actors, actor )
		}
	}
	return actors;
}

func(movie Movie) IsActorInMovie(actor Actor) bool {
	for _, someactor := range movie.Actors {
		if (someactor.Name == actor.Name) {
			return true;
		}
	}
	return false;
}


type Actor struct {
	Name string `json:"name"`
	Id int  `json:"cast_id"`
	Character string `json:"character"`
}
