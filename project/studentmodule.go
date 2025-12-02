package project

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

var (
	Errinvalidage   = errors.New("Invalid age")
	Errinvalidmarks = errors.New("Invalid marks")
	ErrDuplicate    = errors.New("Duplicate student")
)

type student struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Marks int    `json:"marks"`
}

func Validstudents(s student) error {
	if s.Age < 0 || s.Age > 100 {
		return fmt.Errorf("age %d out of range %w", s.Age, Errinvalidage)
	}
	if s.Marks < 0 || s.Marks > 100 {
		return fmt.Errorf("marks %d out of range %w", s.Marks, Errinvalidmarks)
	}
	if s.Name == "" {
		return fmt.Errorf("Invalid name")
	}
	return nil
}

type DB struct {
	Store map[string]student
}

func NewDB() *DB {
	return &DB{Store: make(map[string]student)}
}
func (db *DB) Addstudent(s student) error {
	if err := Validstudents(s); err != nil {
		return fmt.Errorf("Validation fail for %q: %w", s.Name, err)
	}
	if _, ok := db.Store[s.Name]; ok {
		return fmt.Errorf("add %q :%w", s.Name, ErrDuplicate)
	}
	db.Store[s.Name] = s
	return nil
}
func (db *DB) Averagemarks() float64 {
	if len(db.Store) == 0 {
		return 0
	}
	sum := 0
	for _, s := range db.Store {
		sum += s.Marks
	}
	return float64(sum) / float64(len(db.Store))
}
func (db *DB) Savetofile(path string) error {
	list := make([]student, 0, len(db.Store))
	for _, s := range db.Store {
		list = append(list, s)
	}
	data, err := json.MarshalIndent(list, "", " ")
	if err != nil {
		return fmt.Errorf("Marshal students: %w", err)
	}
	if err := os.WriteFile(path, data, 0644); err != nil {
		return fmt.Errorf("Write file %s : %W", path, err)
	}
	return nil
}

func (db *DB) Loadfromfile(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("read file%s :%w", path, err)
	}
	var list []student
	if err := json.Unmarshal(data, &list); err != nil {
		return fmt.Errorf("Unmarshal student : %w", err)
	}
	newstore := make(map[string]student)
	for _, s := range list {
		newstore[s.Name] = s
	}
	db.Store = newstore
	return nil
}
