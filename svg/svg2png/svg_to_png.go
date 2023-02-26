package main

import (
	"github.com/srwiley/oksvg"
	"github.com/srwiley/rasterx"
	"image"
	"image/png"
	"log"
	"os"
)

func main() {

	w, h := 512, 512

	in, err := os.Open("aquarius.svg")
	if err != nil {
		log.Println(err)
	}
	defer in.Close()

	icon, _ := oksvg.ReadIconStream(in)
	icon.SetTarget(0, 0, float64(w), float64(h))
	rgba := image.NewRGBA(image.Rect(0, 0, w, h))
	icon.Draw(rasterx.NewDasher(w, h, rasterx.NewScannerGV(w, h, rgba, rgba.Bounds())), 1)

	out, err := os.Create("aquarius.png")
	if err != nil {
		log.Println(err)
	}
	defer out.Close()
	//Encode 不是image.NRGBA 的图像可能会有损编码
	err = png.Encode(out, rgba)
	if err != nil {
		log.Println(err)
	}

}
