package main

import "fmt"
import "github.com/bit101/go-ansi"
import gk "github.com/gookit/color"

const BgBlueFgLightGrey string = "\033[38:5:254m\033[48:5:33m"
const Reset string = "\033[0m"

const BgBlueFgWhite string = "\033[38;2;255;255;255m\033[48;2;30;30;140m"

func main() {
	fmt.Printf("%vI love CBT Nuggets!%v\n", BgBlueFgLightGrey, Reset)
	fmt.Printf("%vI love learning!!%v\n", BgBlueFgWhite, Reset)

	ansi.SetFg(ansi.BoldBlue)
	ansi.SetBg(ansi.Black)
	ansi.SetUnderline(true)
	ansi.SetReversed(true)
	fmt.Println("My favorite CBT Nuggets trainer is John McGovern!")
	ansi.ResetAll()

	GookitTesting()
}

func GookitTesting() {
	style01 := gk.Style{gk.FgCyan, gk.BgDarkGray}
	style01.Println("Hello from Gookit!")

	fgColor := gk.RGB(30, 190, 40)
	bgColor := gk.RGB(25, 25, 25)
	stylergb01 := gk.RGBStyle{Name: "Trevor01"}
	stylergb01.SetFg(fgColor)
	stylergb01.SetBg(bgColor)

	stylergb01.AddOpts(gk.Bold, gk.OpItalic, gk.OpUnderscore)
	stylergb01.Println("Hello from CBT Nuggets!")
}
