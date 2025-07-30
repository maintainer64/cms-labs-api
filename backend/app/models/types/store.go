package types

import (
	"database/sql/driver"
	"errors"

	"github.com/goccy/go-json"
)

type UserStore map[string]interface{} // Произвольное JSON-хранилище

// Value реализует интерфейс driver.Valuer для сериализации в JSON
func (us UserStore) Value() (driver.Value, error) {
	return json.Marshal(us)
}

// Scan реализует интерфейс sql.Scanner для десериализации из JSON
func (us *UserStore) Scan(value interface{}) error {
	if value == nil {
		*us = UserStore{}
		return nil
	}
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	err := json.Unmarshal(b, us)
	return err
}
