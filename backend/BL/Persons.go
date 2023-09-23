package BL

import (
	"encoding/json"
	"fmt"
	"log/slog"
)

type Persons []Person

func (p *Persons) ToJSON() string {
	json, err := json.Marshal(p)
	if err != nil {
		slog.Warn("Unable to create json", "json error", err)
		return "nil"
	}

	return string(json)
}
func (p *Persons) FromJSON(s string) error {
	fmt.Println(s)
	err := json.Unmarshal([]byte(s), p)
	if err != nil {
		slog.Warn("SHIT!", "parse error", err)
		return fmt.Errorf("failed to parse person from json: %w", err)
	}

	return nil
}
