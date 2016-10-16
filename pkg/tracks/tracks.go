package tracks

import (
	"fmt"
)

// Track holds all the information about a track.
type Track struct {
	Name     string
	Location string
	Laps     int
	Weather  Weather
}

var (
	// WeatherType holds the different weather types that can exist and the
	// handicap percentage that the riders have, based on the weather experienced
	// at the track.
	WeatherType = map[string]int{
		"Clear": 0,  // 0 percent handicap.
		"Rain":  10, // 10 percent handicap, because rain affects riders the most.
		"Hot":   4,
		"Cold":  7,
	}
)

// Weather holds the type of weather that will affect performance at tracks.
type Weather struct {
}

// GetTrackName returns a user friendly track name and location.
func (t *Track) GetTrackName() string {
	return fmt.Sprintf("%s - %s", t.Name, t.Location)
}

// GetTrackLaps returns the number of laps that the track has defined.
func (t *Track) GetTrackLaps() int {
	return t.Laps
}

// SetTrackWeather sets the weather for a certain track based on climate.
func (t *Track) SetTrackWeather() {

}
