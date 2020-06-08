package structcode

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// A Response struct to map the Entire Response
type Response struct {
	Name    string    `json:"name"`
	Pokemon []Pokemon `json:"pokemon_entries"`
}

// A Pokemon Struct to map every pokemon to.
type Pokemon struct {
	EntryNo int            `json:"entry_number"`
	Species PokemonSpecies `json:"pokemon_species"`
}

// PokemonSpecies A struct to map our Pokemon's Species which includes it's name
type PokemonSpecies struct {
	Name string `json:"name"`
}

//Response2 for pokedata
type Response2 struct {
	Abilities []struct {
		Ability struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"ability"`
		IsHidden bool `json:"is_hidden"`
		Slot     int  `json:"slot"`
	} `json:"abilities"`
	BaseExperience int `json:"base_experience"`
	Forms          []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"forms"`
	GameIndices []struct {
		GameIndex int `json:"game_index"`
		Version   struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"version"`
	} `json:"game_indices"`
	Height                 int           `json:"height"`
	HeldItems              []interface{} `json:"held_items"`
	ID                     int           `json:"id"`
	IsDefault              bool          `json:"is_default"`
	LocationAreaEncounters string        `json:"location_area_encounters"`
	Moves                  []struct {
		Move struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"move"`
		VersionGroupDetails []struct {
			LevelLearnedAt  int `json:"level_learned_at"`
			MoveLearnMethod struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"move_learn_method"`
			VersionGroup struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version_group"`
		} `json:"version_group_details"`
	} `json:"moves"`
	Name    string `json:"name"`
	Order   int    `json:"order"`
	Species struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"species"`
	Sprites struct {
		BackDefault      string      `json:"back_default"`
		BackFemale       interface{} `json:"back_female"`
		BackShiny        string      `json:"back_shiny"`
		BackShinyFemale  interface{} `json:"back_shiny_female"`
		FrontDefault     string      `json:"front_default"`
		FrontFemale      interface{} `json:"front_female"`
		FrontShiny       string      `json:"front_shiny"`
		FrontShinyFemale interface{} `json:"front_shiny_female"`
	} `json:"sprites"`
	Stats []struct {
		BaseStat int `json:"base_stat"`
		Effort   int `json:"effort"`
		Stat     struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"type"`
	} `json:"types"`
	Weight int `json:"weight"`
}

//A Article to map every pokemon to.
type Article struct {
	EntryNo int
	Species string
}

//Articles array
var Articles []Article

//A Stat to map every pokemon to.
type Stat struct {
	HP    int
	ATK   int
	DEF   int
	SPATK int
	SPDEF int
	SPD   int
}

//A TypePokemon to map every pokemon to.
type TypePokemon struct {
	No   int
	Type string
}

//ListType array
var ListType []TypePokemon

//GetValue Func
func GetValue() *Response {
	response, err := http.Get("http://pokeapi.co/api/v2/pokedex/kanto/")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseObject Response
	json.Unmarshal(responseData, &responseObject)
	return &responseObject
}

//GetPokeData funct
func GetPokeData(input string) *Response2 {
	pathfix := "https://pokeapi.co/api/v2/pokemon/" + input
	response, err := http.Get(pathfix)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseObject Response2
	json.Unmarshal(responseData, &responseObject)
	//fmt.Println(responseObject.Types)
	return &responseObject
}

//SearchFunc Func
func SearchFunc() {
	Articles = []Article{}
	data := GetValue()
	for i := 0; i < len(data.Pokemon); i++ {
		newdata := Article{EntryNo: data.Pokemon[i].EntryNo, Species: data.Pokemon[i].Species.Name}
		Articles = append(Articles, newdata)
	}
}

//GetType Func
func GetType(input *Response2) {
	ListType = []TypePokemon{}
	for i := 0; i < len(input.Types); i++ {
		newdata := TypePokemon{input.Types[i].Slot, input.Types[i].Type.Name}
		ListType = append(ListType, newdata)
	}
}

//GetStat Func
func GetStat(users *Response2) Stat {
	var HP, ATK, DEF, SPATK, SPDEF, SPD int
	for i := 0; i < len(users.Stats); i++ {
		if users.Stats[i].Stat.Name == "hp" {
			HP = users.Stats[i].BaseStat
		} else if users.Stats[i].Stat.Name == "attack" {
			ATK = users.Stats[i].BaseStat
		} else if users.Stats[i].Stat.Name == "defense" {
			DEF = users.Stats[i].BaseStat
		} else if users.Stats[i].Stat.Name == "special-attack" {
			SPATK = users.Stats[i].BaseStat
		} else if users.Stats[i].Stat.Name == "special-defense" {
			SPDEF = users.Stats[i].BaseStat
		}
		SPD = users.Stats[i].BaseStat
	}
	ListStat := Stat{HP, ATK, DEF, SPATK, SPDEF, SPD}
	return ListStat
}
