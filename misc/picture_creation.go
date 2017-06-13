package main

import "image"
import "image/color"
import "image/png"
import "os"
import "time"

func main() {

  img := image.NewRGBA(image.Rect(500, 500, 0, 0))
  for x := 0; x < 100; x++ {
    for y := 0; y < 100; y++ {
      // img := image.NewRGBA(image.Rect(0, 0, 0, 0))
      img.Set(x, y, color.RGBA{255, 0, 0, 255})

      f, _ := os.OpenFile("out.png", os.O_WRONLY|os.O_CREATE,0600)
      defer f.Close()
      png.Encode(f, img)
      time.Sleep(10 * time.Millisecond)
    }
  }
}
