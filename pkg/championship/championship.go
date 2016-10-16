package championship

import (
	"errors"

	"github.com/PI-Victor/motogp/pkg/riders"
	"github.com/PI-Victor/motogp/pkg/tracks"
)

var (
	errTrackDoesNotExist = errors.New("Track does not exist")
)

// Championship is the struct that holds the order of the tracks in which
// the championship takes place.
type Championship struct {
	Year     int
	Riders   []riders.Rider
	Winner   riders.Rider
	Order    TrackOrder
	Standing Standing
}

// Standing holds the current championship standing.
type Standing struct {
	Position int
	Rider    riders.Rider
}

func newChampionShip() *Championship {
	return &Championship{}
}

// TrackOrder is a type that holds the mapping of the order of the tracks in a
// championship.
type TrackOrder map[int]*tracks.Track

// GetNextTrack returns the track succeeding the one that's been passed as a
// parameter.
func (c *Championship) GetNextTrack(trackIndex int) (*tracks.Track, error) {
	successor, exists := c.Order[trackIndex+1]
	if !exists {
		return nil, errTrackDoesNotExist
	}
	return successor, nil
}

// GetPreviousTrack returns the track preceeding the one that's been passed as
// a parameter.
func (c *Championship) GetPreviousTrack(trackIndex int) (*tracks.Track, error) {
	predecessor, exists := c.Order[trackIndex-1]
	if !exists {
		return nil, errTrackDoesNotExist
	}
	return predecessor, nil
}
