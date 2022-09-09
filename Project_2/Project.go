package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

func main() {
	f, err := os.Open("Project2_Names.csv") // Open the csv file
	if err != nil {
		log.Fatal(err)
	}

	//defer f.Close()

	// Get user input and save it to an int for number of people on each team
	fmt.Println("Enter the number of people on each team:")
	var numPeople int
	fmt.Scanln(&numPeople)

	// Read the csv file
	reader := csv.NewReader(f)
	records, _ := reader.ReadAll()

	// Shuffle the records to randomize
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(records), func(i, j int) { records[i], records[j] = records[j], records[i] })
	fmt.Println(records)

}
