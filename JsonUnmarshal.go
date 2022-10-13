package main

import (
	"encoding/json"
	"fmt"
)

type SeaNettles struct {
	Species      string
	SpeciesLatin string
}

func main() {
	var jsonBlob = []byte(`[
		{"Species": "Black sea nettle", "SpeciesLatin": "Chrysaora achlyos"},
		{"Species": "Atlantic sea nettle", "SpeciesLatin": "Chrysaora quinquecirrha"},
		{"Species": "Pacific sea nettle", "SpeciesLatin": "Chrysaora fuscescens"}
	]`)

	var seaNettles []SeaNettles

	err := json.Unmarshal(jsonBlob, &seaNettles)
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Printf("%+v", seaNettles)
}
