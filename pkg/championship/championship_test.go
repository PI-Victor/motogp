package championship

import (
	"reflect"
	"testing"

	"github.com/PI-Victor/motogp/pkg/riders"
	"github.com/PI-Victor/motogp/pkg/teams"
	"github.com/PI-Victor/motogp/pkg/tracks"
)

var (
	yamahaMotoGPTeam = teams.Team{
		Name: "Yamaha MotoGP Team",
	}
	hondaRacingCorporation = teams.Team{
		Name: "Honda Racing Corporation",
	}
	ducatiMotoGPTeam = teams.Team{
		Name: "Ducati MotoGP Team",
	}

	losailTrack = &tracks.Track{
		Name:     "Losail International Circuit",
		Location: "Qatar",
	}
	argentinaTrack = &tracks.Track{
		Name:     "Termas de Río Hondo",
		Location: "Argentina",
	}
	americasTrack = &tracks.Track{
		Name:     "Circuit of the Americas",
		Location: "Americas",
	}

	lorenzoRider = riders.Rider{
		Name: "Jorge Lorenzo",
		Rank: 9.2,
		Team: yamahaMotoGPTeam,
	}
	marquezRider = riders.Rider{
		Name: "Marc Marquez",
		Rank: 9.4,
		Team: hondaRacingCorporation,
	}
	doviziosoRider = riders.Rider{
		Name: "Andrea Dovizioso",
		Rank: 8.9,
		Team: ducatiMotoGPTeam,
	}
)

func TestNewChampionShip(t *testing.T) {
	if championship := newChampionShip(); championship == nil {
		t.Fatalf("Could not create a new championship instance!")
	}
}

func mockNewTracks() TrackOrder {
	return TrackOrder{1: losailTrack, 2: argentinaTrack, 3: americasTrack}
}

func mockNewRiders() []riders.Rider {
	return []riders.Rider{lorenzoRider, marquezRider, doviziosoRider}
}

func mockNewChampionship() *Championship {
	riders := mockNewRiders()
	tracks := mockNewTracks()

	return &Championship{Riders: riders, Order: tracks}
}

func TestGetNextTrack(t *testing.T) {
	c := mockNewChampionship()
	nextTrack, err := c.GetNextTrack(2)
	if err != nil && err != errTrackDoesNotExist {
		t.Fatalf("An error occured while getting next track: %#v", err)
	}
	if !reflect.DeepEqual(nextTrack, americasTrack) {
		t.Errorf("Expected track to be %#v, got %#v", nextTrack, nextTrack)
	}
}

func TestGetNextTrackInvalid(t *testing.T) {
	c := mockNewChampionship()
	nextTrack, err := c.GetNextTrack(10)
	if nextTrack != nil {
		t.Errorf("Expected nexTrack to be nil, got %#v", nextTrack)
	}
	if err != errTrackDoesNotExist {
		t.Errorf("Expected error to be %#v, got %#v", errTrackDoesNotExist, err)
	}
}

func TestGetPreviousTrack(t *testing.T) {
	c := mockNewChampionship()
	previousTrack, err := c.GetPreviousTrack(2)
	if err != nil && err != errTrackDoesNotExist {
		t.Fatalf("An error occured while getting the previous track: %#v", err)
	}
	if !reflect.DeepEqual(previousTrack, losailTrack) {
		t.Errorf("Expected track to be %#v, got %#v", losailTrack, previousTrack)
	}
}

func TestGetPreviousTrackInvalid(t *testing.T) {
	c := mockNewChampionship()
	previousTrack, err := c.GetNextTrack(10)
	if previousTrack != nil {
		t.Errorf("Expected previousTrack to be nil, got %#v", previousTrack)
	}
	if err != errTrackDoesNotExist {
		t.Errorf("Expected error to be %#v, got %#v", errTrackDoesNotExist, err)
	}
}
