package ocr

import (
	"TFThExternal/constants"
	"fmt"
	"github.com/go-vgo/robotgo"
	"strconv"
	"strings"
)

//Fix by not using ocr but instead using pixel search with photos for each champ text
func GetShopChampions() []constants.Champ {
	ChampPosLen := len(constants.ChampShopPos)
	Champs := make([]constants.Champ, ChampPosLen)
	for i := 0; i < ChampPosLen; i++ {
		cap := robotgo.CaptureScreen(constants.ChampShopPos[i], constants.ChampShopPosEdges[0],constants.ChampShopPosEdges[1],constants.ChampShopPosEdges[2])
		imgPath := fmt.Sprintf("./tmp/ch%v.png", i)
		robotgo.SaveBitmap(cap, imgPath)
		text := strings.ToLower(RemoveStringGarbage(OCR(imgPath)))
		Champs[i]=constants.Champs[text]
	}
	return Champs
}

func GetLvl() int {
	cap := robotgo.CaptureScreen(constants.LvlPos...)
	robotgo.SaveBitmap(cap,"./tmp/lvl.png")
	text := OCR("./tmp/lvl.png")
	r,_ := strconv.Atoi(KeepOnlyNumbers(text))
	return r
}

func GetGold() int {
	cap := robotgo.CaptureScreen(constants.GoldPos...)
	robotgo.SaveBitmap(cap,"./tmp/gold.png")
	text := OCR("./tmp/gold.png")
	r,_ := strconv.Atoi(KeepOnlyNumbers(text))
	return r
}

func KeepOnlyNumbers(text string) string {
	builder := strings.Builder{}
	for i:=0; i<len(text); i++ {
		//if text[i] ascii code is >= 48 which is one and <=57 which is 9
		asciiCode:=int(text[i])
		if asciiCode>47 && asciiCode<58 {
			builder.WriteByte(text[i])
		}
	}
	return builder.String()
}

func RemoveStringGarbage(text string) string {
	builder := strings.Builder{}
	for i := 0; i < len(text); i++ {
		asciiCode := int(text[i])
		//from A to Z and from a to z
		if (asciiCode > 64 && asciiCode < 91) || (asciiCode > 96 && asciiCode < 123) {
			builder.WriteByte(text[i])
		}
	}
	return builder.String()
}
