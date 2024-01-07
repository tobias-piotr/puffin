package database

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type JSON map[string]any

func (c *JSON) Scan(value any) error {
	v, ok := value.([]byte)
	if !ok {
		panic(fmt.Sprintf("unsupported type: %T", v))
	}
	return json.Unmarshal(v, &c)
}

func (c JSON) Value() (driver.Value, error) {
	return json.Marshal(c)
}
