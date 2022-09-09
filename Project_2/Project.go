package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	// Open the csv file
	f, err := os.Open("Project2_Names.csv")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

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

	// Get the number of teams and number of extra people
	var numTeams int = len(records) / numPeople
	var extra int = len(records) % numPeople

	// Initializing teams
	teams := [][]string{}

	// Adding people to teams
	for i := 0; i < numTeams; i++ {
		var temp []string
		for j := 0; j < numPeople; j++ {
			temp = append(temp, records[i*numPeople+j][0])
		}
		teams = append(teams, temp)
	}

	// Add extra people to teams
	for i := 0; i < extra; i++ {
		teams[i] = append(teams[i], records[len(records)-i-1][0])
	}

	// Print teams
	for i := 0; i < len(teams); i++ {
		fmt.Println("Team " + strconv.Itoa(i+1) + ":" + strings.Join(teams[i], ", "))
	}
}
