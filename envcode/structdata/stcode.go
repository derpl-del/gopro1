package structdata

import "github.com/derpl-del/gopro1/envcode/structcode"

// Combined for result.html
type Combined struct {
	RsData
	ListArticles []structcode.Article
}

//EnvData for data
type EnvData struct {
}

//HpStruct a
type HpStruct struct {
	ListArticles []structcode.Article
}

//TypeStruct a
type TypeStruct struct {
	ListType []structcode.TypePokemon
}

// RsData for result.html
type RsData struct {
	Id           string
	Name         string
	FrontDefault string
	BackDefault  string
	FrontShiny   string
	BackShiny    string
	DataBefore   int
	DataAfter    int
	Type1        string
	Type2        string
	structcode.Stat
}
