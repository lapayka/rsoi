package main

import (
	"testing"

	"github.com/lapayka/rsoi/BL"
	"github.com/lapayka/rsoi/DA"
)

func TestCreate(t *testing.T) {
	db, _ := DA.New("postgres", "postgres", "rsoi_persons_test", "1234")

	p := BL.Person{FirstName: "Denis", LastName: "Ivanov", Age: 25}
	err := db.CreatePerson(&p)

	if err != nil {
		t.Errorf("Error when call func %s", err)
	}

	if p.ID == 0 {
		t.Errorf("Person Id must be new, current=%d", p.ID)
	}
}

func TestGetById(t *testing.T) {
	db, _ := DA.New("postgres", "postgres", "rsoi_persons_test", "1234")

	p, err := db.GetPersonById(1)

	if err != nil {
		t.Errorf("Error when call func %s", err)
	}

	if p.ID == 0 {
		t.Errorf("Person Id must be new, current=%d", p.ID)
	}
}

func TestGetAll(t *testing.T) {
	db, _ := DA.New("postgres", "postgres", "rsoi_persons_test", "1234")

	p, err := db.GetPersons()

	if err != nil {
		t.Errorf("Error when call func %s", err)
	}

	if len(p) == 0 {
		t.Errorf("Empty list of persons")
	}
}

func TestUpdatePerson(t *testing.T) {
	db, _ := DA.New("postgres", "postgres", "rsoi_persons_test", "1234")

	p := BL.Person{FirstName: "abc", LastName: "acd"}
	err := db.UpdatePerson(1, p)

	if err != nil {
		t.Errorf("Error when call func %s", err)
	}
}

func TestDeletePerson(t *testing.T) {
	db, _ := DA.New("postgres", "postgres", "rsoi_persons_test", "1234")

	err := db.DeletePerson(1)

	if err != nil {
		t.Errorf("Error when call func %s", err)
	}
}
