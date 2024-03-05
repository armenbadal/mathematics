package mdb

import (
	"encoding/json"
	"io"
	"os"
)

var mdb struct {
	UsersFile    string `json:"users-file"`
	ProblemsFile string `json:"problems-file"`
	AnswersFile  string `json:"answers-file"`
}

type User struct {
	Id       int
	Name     string
	Password string
}

var users []*User

type Problem struct {
	Id       int
	Text     string
	Solution string
	Answer   string
}

var problems []*Problem

type Answer struct {
	Id        int
	UserId    int
	ProblemId int
	Solution  string
	Result    string
}

var answers []*Answer
var answersChanged bool

func LoadDatabase() error {
	if err := loadJsonData("db.json", &mdb); err != nil {
		return err
	}

	if err := loadUsers(mdb.UsersFile); err != nil {
		return err
	}

	if err := loadProblems(mdb.ProblemsFile); err != nil {
		return err
	}

	if err := loadAnswers(mdb.AnswersFile); err != nil {
		return err
	}

	return nil
}

func loadJsonData(file string, obj any) error {
	dbf, err := os.Open(file)
	if err != nil {
		return err
	}
	defer dbf.Close()

	dbData, err := io.ReadAll(dbf)
	if err != nil {
		return err
	}

	return json.Unmarshal(dbData, &obj)
}

func loadUsers(file string) error {
	return loadJsonData(file, &users)
}

func loadProblems(file string) error {
	return loadJsonData(file, &users)
}

func loadAnswers(file string) error {
	return loadJsonData(file, &answers)
}
