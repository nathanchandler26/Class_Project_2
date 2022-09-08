package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("Project2_Names.csv")
	if err != nil {
		log.Fatal(err)
		// fmt.Println("Hello world")
	}

	//defer f.Close()

	reader := csv.NewReader(f)
	records, _ := reader.ReadAll()

	fmt.Println(records)

}
