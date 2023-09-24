package DA

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/lapayka/rsoi/BL"
)

type DB struct {
	db *gorm.DB
}

func New(db_name string) (*DB, error) {
	dsn := fmt.Sprintf("host=localhost user=postgres password=1234 dbname=%s port=5433 sslmode=disable", db_name)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("unable to connect database", err)
	}

	return &DB{db: db}, nil
}

func (d *DB) GetPersons() (BL.Persons, error) {
	persons := BL.Persons{}
	d.db.Find(&persons)

	if len(persons) == 0 {
		return nil, nil
	}

	return persons, nil
}

func (d *DB) GetPersonById(Id int) (BL.Person, error) {
	person := BL.Person{}
	d.db.First(&person, Id)

	return person, nil
}

func (d *DB) CreatePerson(person *BL.Person) error {
	error := d.db.Create(&person).Error

	if error != nil {
		return fmt.Errorf("Cannot create person: %s", person.ToJSON())
	}
	return nil
}

func (d *DB) UpdatePerson(id int, person BL.Person) error {
	tx := d.db.Begin()

	p := BL.Person{ID: id}
	if err := tx.First(&p).Error; err != nil {
		tx.Rollback()

		return err
	}

	if person.FirstName != "" {
		p.FirstName = person.FirstName
	}
	if person.LastName != "" {
		p.LastName = person.LastName
	}
	if person.Age != 0 {
		p.Age = person.Age
	}

	if err := tx.Save(&p).Error; err != nil {
		tx.Rollback()

		return err
	}

	tx.Commit()

	return nil
}

func (d *DB) DeletePerson(id int) error {
	tx := d.db.Begin()

	person := BL.Person{}
	tx.First(&person, id)

	err_eq := BL.Person{}
	if person == err_eq {
		tx.Rollback()
		return fmt.Errorf("No such person with id=%d, can't delete it", id)
	}

	err := tx.Delete(&BL.Person{}, id).Error

	if err != nil {
		tx.Rollback()
		return fmt.Errorf("Some error while deleting %s", err)
	}

	tx.Commit()

	return nil
}
