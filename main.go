package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Analysis struct {
	Nombre    string `json:"nombre"`
	Condicion string `json:"condicion"`
	Tipo      string `json:"tipo"`
	Promesa   string `json:"promesa"`
	Precio    string `json:"precio"`
}

func printAnalysis(analysis Analysis) {
	fmt.Println("-------------------------------------------------")
	fmt.Printf("+Nombre: %s\n", analysis.Nombre)
	fmt.Printf("+Condicion: %s\n", analysis.Condicion)
	fmt.Printf("+Tipo: %s\n", analysis.Tipo)
	fmt.Printf("+Promesa: %s\n", analysis.Promesa)
	fmt.Printf("+Precio: %s\n", analysis.Precio)
	fmt.Println("-------------------------------------------------")
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	file, err := os.ReadFile("./data.json")
	if err != nil {
		fmt.Print(err.Error())
		panic(1)
	}

	analysisList := []Analysis{}
	err = json.Unmarshal(file, &analysisList)
	if err != nil {
		fmt.Print(err.Error())
		panic(1)
	}

	for {
		fmt.Println("")
		fmt.Printf("Search: ")
		analysisName, err := reader.ReadBytes('\n')
		if err != nil {
			fmt.Print(err.Error())
			panic(1)
		}

		name := strings.Trim(string(analysisName), "\n")

		for _, analysis := range analysisList {
			found := strings.Contains(strings.ToLower(analysis.Nombre), strings.ToLower(name))
			if found {
				printAnalysis(analysis)
			}
		}

	}
}
