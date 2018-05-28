package main

import (
	"fmt"
	"time"

	"github.com/nickysemenza/hyperion/backend/api"
	"github.com/nickysemenza/hyperion/backend/cue"
	"github.com/nickysemenza/hyperion/backend/light"
)

func main() {
	fmt.Println("Hello!")

	light.ReadLightConfigFromFile("./light/testconfig.json")

	cue.CM = cue.Master{}
	CueMaster := &cue.CM

	mainCueStack := cue.Stack{Priority: 2, Name: "main"}
	for x := 1; x <= 2; x++ {
		a := CueMaster.New([]cue.Frame{
			CueMaster.NewFrame([]cue.FrameAction{
				CueMaster.NewFrameAction(time.Millisecond*1500, light.RGBColor{R: 255}, "hue1"),
				CueMaster.NewFrameAction(0, light.RGBColor{R: 255}, "hue2"),
			}),
			CueMaster.NewFrame([]cue.FrameAction{
				CueMaster.NewFrameAction(time.Second*time.Duration(x), light.RGBColor{G: 255}, "hue1"),
				CueMaster.NewFrameAction(0, light.RGBColor{B: 255}, "hue2"),
			}),
			CueMaster.NewFrame([]cue.FrameAction{
				CueMaster.NewFrameAction(0, light.RGBColor{B: 255}, "hue1"),
				CueMaster.NewFrameAction(0, light.RGBColor{R: 255}, "hue2"),
			}),
			CueMaster.NewFrame([]cue.FrameAction{
				CueMaster.NewFrameAction(time.Second*2, light.RGBColor{B: 255}, "hue1"),
				CueMaster.NewFrameAction(0, light.RGBColor{B: 255}, "hue2"),
			}),
		}, fmt.Sprintf("Cue #%d", x))
		mainCueStack.Cues = append(mainCueStack.Cues, a)
	}
	CueMaster.CueStacks = append(CueMaster.CueStacks, mainCueStack)
	// spew.Dump(CueMaster)

	//go api.ServeRPC(8888)
	go api.ServeHTTP()
	// CueMaster.ProcessForever()
	fmt.Println("faaa")

	for {
		time.Sleep(1 * time.Second)
	}
}