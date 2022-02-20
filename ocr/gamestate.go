package ocr

import (
	"TFThExternal/constants"
	"errors"
	"github.com/go-vgo/robotgo"
)

type Gamestate int

const (
	PVP Gamestate = iota
	Carrousel
	Augment
	Minions
	PVE
)

const (
	maxOffset = 0.15
)

//Should be able to make this faster somehow idk
func GetGameState() (Gamestate, error) {
	capture := robotgo.CaptureScreen(constants.GamestatePos...)
	x, _ := robotgo.FindBitmap(constants.PvpBmp, capture, maxOffset)
	if x != -1 {
		return PVP, nil
	}
	x, _ = robotgo.FindBitmap(constants.AugmentBmp, capture, maxOffset)
	if x != -1 {
		return Augment, nil
	}
	x, _ = robotgo.FindBitmap(constants.CarrouselBmp, capture, maxOffset)
	if x != -1 {
		return Carrousel, nil
	}
	x, _ = robotgo.FindBitmap(constants.MinionsBmp, capture, maxOffset)
	if x != -1 {
		return Minions, nil
	}
	x, _ = robotgo.FindBitmap(constants.DrakeBmp, capture, maxOffset)
	if x != -1 {
		return PVE, nil
	}
	x, _ = robotgo.FindBitmap(constants.RaptorsBmp, capture, maxOffset)
	if x != -1 {
		return PVE, nil
	}
	x, _ = robotgo.FindBitmap(constants.WolvesBmp, capture, maxOffset)
	if x != -1 {
		return PVE, nil
	}
	return -1, errors.New("no gamestate found")
}

func IsWinScreen() bool {
	capture := robotgo.CaptureScreen(constants.CheckWinPos...)
	x, _ := robotgo.FindBitmap(constants.WinBmp, capture)
	if x != -1 {
		return true
	}
	return false
}

func IsLoseScreen() bool {
	capture := robotgo.CaptureScreen(constants.CheckLosePos...)
	x, _ := robotgo.FindBitmap(constants.ExitBmp, capture)
	if x != -1 {
		return true
	}
	return false
}


