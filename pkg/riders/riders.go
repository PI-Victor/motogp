package riders

import (
	"github.com/PI-Victor/motogp/pkg/teams"
)

// Rider holds information about the rider and the teams he belongs to.
type Rider struct {
	Name string
	Rank Rank
	Team teams.Team
}

// Rank holds the rank based on what weather the track is experiencing.
type Rank struct {
	Rain  int
	Hot   int
	Cold  int
	Clear int
}

func (r *Rider) GetTeam() *teams.Team {
	return nil
}
