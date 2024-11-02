package oj

import (
	"bytes"
	"encoding/json"
)

// OrderedJSON represents a JSON object that maintains field order
type OrderedJSON struct {
	Order []string
	Data  map[string]interface{}
}

// UnmarshalJSON implements the json.Unmarshaler interface
func (o *OrderedJSON) UnmarshalJSON(data []byte) error {
	dec := json.NewDecoder(bytes.NewReader(data))

	_, err := dec.Token()
	if err != nil {
		return err
	}

	o.Order = make([]string, 0)
	o.Data = make(map[string]interface{})

	for dec.More() {
		token, err := dec.Token()
		if err != nil {
			return err
		}

		key := token.(string)
		o.Order = append(o.Order, key)

		var value interface{}
		if err := dec.Decode(&value); err != nil {
			return err
		}
		o.Data[key] = value
	}

	_, err = dec.Token()

	return err
}

// MarshalJSON implements the json.Marshaler interface
func (o OrderedJSON) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	enc.SetIndent("", "  ")

	buf.WriteString("{\n")

	for i, key := range o.Order {
		keyJSON, err := json.Marshal(key)
		if err != nil {
			return nil, err
		}

		buf.Write([]byte("  "))
		buf.Write(keyJSON)
		buf.WriteString(": ")

		valueJSON, err := json.Marshal(o.Data[key])
		if err != nil {
			return nil, err
		}
		buf.Write(valueJSON)

		if i < len(o.Order)-1 {
			buf.WriteString(",\n")
		} else {
			buf.WriteString("\n")
		}
	}

	buf.WriteString("}")

	return buf.Bytes(), nil
}
