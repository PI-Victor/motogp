package tracks

type Track struct {
	Name     string
	Location string
	Laps     int
}

func (t *Track) GetTrackName() string {
	return t.Name
}

func (t *Track) GetTrackLaps() int {
	return t.Laps
}
