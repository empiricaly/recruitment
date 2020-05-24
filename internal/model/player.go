package model

// Player is a participant of a Game
type Player struct {
	Base
	Sorted

	// playerID is the public unique identifier of the player. This must be given by
	// the player to participate in a Game. It can also be automatically given by a
	// system managing players before they enter a Game.
	PlayerID string `json:"playerID"`
}

// IsNode is needed by the GraphQL schema Node interface
func (Player) IsNode() {}

// Data returns all custom data that has been set on the Player.
func (p *Player) Data() ([]*Datum, error) {
	return nil, nil
}
