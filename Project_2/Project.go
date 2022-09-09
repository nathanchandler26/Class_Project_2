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
	f, err := os.Open("Project2_Names.csv")
	if err != nil {
		log.Fatal(err)
		// fmt.Println("Hello world")
	}

	//defer f.Close()

	reader := csv.NewReader(f)
	records, _ := reader.ReadAll()

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(records), func(i, j int) { records[i], records[j] = records[j], records[i] })
	fmt.Println(records)
}
