package main

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"log"
	"os"

	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
)

const (
    width = 80
    height = 80
)
func main() {
    numberOfImages := 10
    for i := 0; i < numberOfImages; i++ {
        fmt.Printf("%d\n", i)
        m := image.NewRGBA(image.Rect(0, 0, width, height))
        colorImage(m, color.RGBA{100, 200, 200, 0xff})
        addText(m, 10, 10, fmt.Sprintf("test: %d", i))

        f, err := os.Create(fmt.Sprintf("img%d.jpg", i))
        if err != nil {
            panic(err)
        }
        defer f.Close()

        if err = jpeg.Encode(f, m, nil); err != nil {
            log.Printf("failed to encode: %v", err)
        }
    }

}


func colorImage(m *image.RGBA, color color.RGBA) {
    for x := 0; x < width; x++ {
        for y := 0; y < height; y++ {
            m.Set(x, y, color)
        }
    }
}

func addText(m *image.RGBA, xPos, yPos int, text string) {
    textColor := color.RGBA{0, 0, 0, 255}
    coord := fixed.Point26_6{X: fixed.I(xPos), Y: fixed.I(yPos)}

    d := &font.Drawer{
        Dst: m,
        Src: image.NewUniform(textColor),
        Face: basicfont.Face7x13,
        Dot: coord,
    }
    d.DrawString(text)
}
