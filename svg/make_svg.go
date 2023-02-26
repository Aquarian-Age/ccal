package main

import (
	"os"

	svg "github.com/ajstarks/svgo"
)

// go run main . > abc.svg
func main() {
	w, h := 256, 256
	name := `奇`
	s := svg.New(os.Stdout)
	s.Start(w, h)
	//s.Circle(w/2, h/2, w/4, "fill:none;stroke:black")
	s.Text(w/2, 180, name, "text-anchor:middle;font-size:150px;fill:red")
	s.End()
}
