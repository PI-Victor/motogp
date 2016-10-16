package riders

import (
	"github.com/PI-Victor/motogp/pkg/teams"
)

// Rider holds information about the rider and the teams he belongs to.
type Rider struct {
	Name string
	Rank float32
	Team teams.Team
}
