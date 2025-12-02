package project

import (
	"errors"
	"os"
	"path/filepath"
	"testing"
)

func TestValidatestudent(t *testing.T) {
	cases := []struct {
		name    string
		input   student
		wantErr error
	}{
		{"valid", student{"Akash", 20, 80}, nil},
		{"age negative", student{"Bunty", -5, 81}, Errinvalidage},
		{"age too large", student{"Kushal", 150, 68}, Errinvalidage},
		{"marks negative", student{"Tanmay", 21, -7}, Errinvalidmarks},
		{"marks too large", student{"Sujan", 21, 132}, Errinvalidmarks},
		{"Empty name", student{"", 21, 87}, nil},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			err := Validstudents(tc.input)
			if tc.wantErr == nil {
				if tc.input.Name == "" {
					if err == nil {
						t.Fatalf("Expected error from empty name, got nil")
					}
					return
				}
				if err != nil {
					t.Fatalf("Unexpected error")
				}
				return
			}
			if !errors.Is(err, tc.wantErr) {
				t.Fatalf("Expected error containing %v, got %v", tc.wantErr, err)
			}
		})
	}

}
func TestAddstudentsAndDuplicate(t *testing.T) {
	db := NewDB()

	s1 := student{"Akash", 20, 90}
	if err := db.Addstudent(s1); err != nil {
		t.Fatalf("unexpected add error: %v", err)
	}
	if err := db.Addstudent(s1); err == nil {
		t.Fatalf("Expected duplicate error , got nil")
	} else if !errors.Is(err, ErrDuplicate) {
		t.Fatalf("expected Errduplicate, got :%v", err)
	}
}
func TestAveragemarks(t *testing.T) {
	db := NewDB()

	if got := db.Averagemarks(); got != 0 {
		t.Fatalf("Expected 0 from emplt db,got %v", got)
	}
	students := []student{
		{"A", 20, 80},
		{"B", 22, 70},
		{"C", 19, 90},
	}
	for _, s := range students {
		if err := db.Addstudent(s); err != nil {
			t.Fatalf("Add failed : %v", err)
		}
	}
	Want := 80.0
	got := db.Averagemarks()
	if got != Want {
		t.Fatalf("Expected average %v, got %v", Want, got)
	}
}
func TestSaveLoadjson(t *testing.T) {
	db := NewDB()
	students := []student{
		{"X", 18, 65},
		{"Y", 21, 75},
	}
	for _, s := range students {
		if err := db.Addstudent(s); err != nil {
			t.Fatalf("add failed: %v", err)
		}
	}
	tmpDir := t.TempDir()
	path := filepath.Join(tmpDir, "students.json")

	if err := db.Savetofile(path); err != nil {
		t.Fatalf("saved failed :%v", err)
	}
	db2 := NewDB()
	if err := db2.Loadfromfile(path); err != nil {
		t.Fatalf("load failed: %v", err)
	}
	if len(db2.Store) != len(db.Store) {
		t.Fatalf("expected %d students after load, got %d", len(db.Store), len(db2.Store))
	}
	for name, s1 := range db.Store {
		s2, ok := db2.Store[name]
		if !ok {
			t.Fatalf("students %q missing after load", name)
		}
		if s1 != s2 {
			t.Fatalf("student mismatched for %q :got %+v want %+v", name, s2, s1)
		}
	}
	_ = os.Remove(path)

}
