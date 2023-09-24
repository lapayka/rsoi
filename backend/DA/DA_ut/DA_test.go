package main

import (
	"testing"

	"github.com/lapayka/rsoi/BL"
	"github.com/lapayka/rsoi/DA"
)

func TestCreate(t *testing.T) {
	db, _ := DA.New("rsoi_persons_test")

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
	db, _ := DA.New("rsoi_persons_test")

	p, err := db.GetPersonById(1)

	if err != nil {
		t.Errorf("Error when call func %s", err)
	}

	if p.ID == 0 {
		t.Errorf("Person Id must be new, current=%d", p.ID)
	}
}

func TestGetAll(t *testing.T) {
	db, _ := DA.New("rsoi_persons_test")

	p, err := db.GetPersons()

	if err != nil {
		t.Errorf("Error when call func %s", err)
	}

	if len(p) == 0 {
		t.Errorf("Empty list of persons")
	}
}

func TestUpdatePerson(t *testing.T) {
	db, _ := DA.New("rsoi_persons_test")

	p := BL.Person{FirstName: "abc", LastName: "acd"}
	err := db.UpdatePerson(1, p)

	if err != nil {
		t.Errorf("Error when call func %s", err)
	}
}

func TestDeletePerson(t *testing.T) {
	db, _ := DA.New("rsoi_persons_test")

	err := db.DeletePerson(1)

	if err != nil {
		t.Errorf("Error when call func %s", err)
	}
}
