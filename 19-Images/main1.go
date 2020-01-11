package main

import (
	"math/rand"
	"os"

	svg "github.com/ajstarks/svgo"
)

func rn(n int) int { return rand.Intn(n) }

func main() {
	/*
		f, err := os.OpenFile("demo.svg", os.O_RDWR|os.O_CREATE, 0755)
		if err != nil {
			panic(err)
		}
		defer f.Close()
	*/

	canvas := svg.New(os.Stdout)
	data := []struct {
		Month string
		Usage int
	}{
		{"Jan", 171},
		{"Feb", 180},
		{"March", 100},
		{"April", 87},
		{"May", 66},
		{"Jun", 40},
		{"July", 32},
		{"Augest", 55},
	}

	width := len(data)*60 + 10
	height := 300
	threshold := 160
	max := 0
	for _, item := range data {
		if item.Usage > max {
			max = item.Usage
		}
	}
	canvas.Start(width, height)
	for i, val := range data {
		percent := val.Usage * (height - 50) / max
		canvas.Rect(i*60+10, (height-50)-percent, 50, percent, "fill:rgb(77,200,232)")
		canvas.Text(i*60+35, height-24, val.Month, "font-size:14pt;fill:rgb(150,150,150);text-anchor:middle")
	}
	threshPercent := threshold * (height - 50) / max
	canvas.Line(0, height-threshPercent, width, height-threshPercent, "stroke: rgb(255,100,100); opacity:0.8; stroke-width: 2px")
	canvas.Rect(0, 0, width, height-threshPercent, "fill:rgb(255,100,100); opacity: 0.4")
	canvas.Line(0, height-50, width, height-50, "stroke: rgb(150,150,150); stroke-width:2")
	/*
		rand.Seed(time.Now().Unix())
		canvas.Start(width, height)
		canvas.Rect(0, 0, width, height)
		for i := 0; i < nstars; i++ {
			canvas.Circle(rn(width), rn(height), rn(3), "fill:white")
		}
		canvas.Circle(width/2, height, width/2, "fill:rgb(77, 117, 232)")
		canvas.Text(width/2, height*4/5, "hello, world", style)
	*/
	canvas.End()

}

//rsvg-convert -w 1000 demo.svg > demo.png
