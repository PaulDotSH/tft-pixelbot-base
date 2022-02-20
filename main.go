package main

import (
	"TFThExternal/ocr"
	"TFThExternal/settings"
	"fmt"
	"github.com/robotn/gohook"
)

func main() {
	add()
}

func add() {
	fmt.Println("hook add...")
	s := hook.Start()
	defer hook.End()

	for {
		i := <-s
		if i.Kind == hook.KeyDown {
			switch i.Rawcode {
			case settings.SemiColon:

				//fmt.Println(ocr.GetShopChampions())
				//fmt.Println(ocr.GetGold())
				//fmt.Println(ocr.GetGameState())
				//fmt.Println(ocr.GetLvl())
				//fmt.Println(ocr.IsLoseScreen())
				fmt.Println(ocr.GetPlayerHP())

				//case settings.SemiColon:
				//	fmt.Println("Exiting")
				//	os.Exit(0)
			}
		}

		if i.Kind != hook.KeyDown {
			continue
		}

		//fmt.Println(i.Kind, i.When, i.Keycode, i.Keychar)

	}
}
