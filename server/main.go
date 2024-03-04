package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type problem struct {
	Text   string
	Answer string
}

func (p problem) String() string {
	return fmt.Sprintf("[%s | %s]", p.Text, p.Answer)
}

var problems map[string]*problem

func loadProblems(file string) error {
	problemsFile, err := os.Open(file)
	if err != nil {
		return err
	}
	defer problemsFile.Close()

	problemsData, err := io.ReadAll(problemsFile)
	if err != nil {
		return err
	}

	err = json.Unmarshal(problemsData, &problems)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	err := loadProblems("mathematics.json")
	if err != nil {
		fmt.Println("Խնդիրները բեռնելը չհաջողվեց։")
		fmt.Println(err)
	}

	for id, pr := range problems {
		fmt.Println(id, pr)
	}
	// Բեռնել արդեն լուծված խնդիրների ցուցակը
	// Խնդիրներից առանձնացնել դեռ չլուծվածները
	// ցուցադրել խնդիր
	//
}
