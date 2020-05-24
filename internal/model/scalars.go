package model

import (
	"encoding/json"
	"fmt"
	"io"
	"strconv"

	"github.com/99designs/gqlgen/graphql"
)

// MarshalUint32 helps marshalling a uint32 for the Uint32 Graphql scalar
func MarshalUint32(i uint32) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		_, _ = io.WriteString(w, strconv.FormatInt(int64(i), 10))
	})
}

// UnmarshalUint32 helps unmarshal a uint32 for the Uint32 Graphql scalar
func UnmarshalUint32(v interface{}) (uint32, error) {
	switch v := v.(type) {
	case string:
		iv, err := strconv.ParseUint(v, 10, 32)
		if err != nil {
			return 0, err
		}
		return uint32(iv), nil
	case int:
		return uint32(v), nil
	case int64:
		return uint32(v), nil
	case json.Number:
		iv, err := strconv.ParseUint(string(v), 10, 32)
		if err != nil {
			return 0, err
		}
		return uint32(iv), nil
	default:
		return 0, fmt.Errorf("%T is not an int", v)
	}
}
