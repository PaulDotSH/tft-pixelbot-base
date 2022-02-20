package ocr

import (
	"TFThExternal/constants"
	"github.com/go-vgo/robotgo"
	"log"
	"os/exec"
	"strconv"
)

func OCR(path string) string {
	body, err := exec.Command("gocr", "-i", path).Output()
	if err != nil {
		log.Fatalln(err)
	}
	return string(body)
}

func GetPlayerHP() int {
	capture := robotgo.CaptureScreen(constants.PlayersHpBarPos...)
	//robotgo.SaveBitmap(capture,"./tmp/testcap.png")
	x, y := robotgo.FindBitmap(constants.PlayerSymbolBmp, capture, 0.05)
	if x != -1 {
		capture2 := robotgo.CaptureScreen(constants.PlayersHpBarPos[0]+x+10,constants.PlayersHpBarPos[1]+y,40,30)
		robotgo.SaveBitmap(capture2,"./tmp/hp.png")
		r,_ := strconv.Atoi(KeepOnlyNumbers(OCR("./tmp/hp.png")))
		return r
	}
	return -100
}
