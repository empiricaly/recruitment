package model

// Datum is a single piece of custom data set by a player or the Game.
type Datum struct {
	Base

	// key identifies the unique key of the Datum.
	Key string `json:"key"`

	// val is the value of the Datum. It can be any JSON encodable value.
	// Passing null will delete the Datum.
	Val *string `json:"val"`
}

// IsNode is needed by the GraphQL schema Node interface
func (Datum) IsNode() {}

// Previous returns previous values for the Datum (they will all have the same ID)
func (d *Datum) Previous() ([]*Datum, error) {
	return nil, nil
}
