package BL

import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
)

type Person struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Age       int    `json:"age"`
}

func (Person) TableName() string {
	return "persons"
}

func (p *Person) ToJSON() string {
	json, err := json.Marshal(p)
	if err != nil {
		slog.Warn("Unable to create json", "json error", err)
		return "nil"
	}

	return string(json)
}

func (p *Person) FromJSON(s string) error {
	err := json.Unmarshal([]byte(s), p)

	if err != nil {
		slog.Warn("SHIT!", "parse error", err)
		return fmt.Errorf("failed to parse person from json: %w", err)
	}

	return nil
}

func (p *Person) FromJSONReder(r io.Reader) error {
	// fmt.Println(s)
	err := json.NewDecoder(r).Decode(p)
	if err != nil {
		return fmt.Errorf("failed to parse person from json: %w", err)
	}
	return nil
}
