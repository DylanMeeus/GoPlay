package main

import (
    "github.com/hajimehoshi/ebiten"
    "github.com/hajimehoshi/ebiten/ebitenutil"
)

var (
    evilImg *ebiten.Image
)

func update(screen *ebiten.Image) error {
    if ebiten.IsDrawingSkipped() {
        return nil
    }
    screen.DrawImage(evilImg,nil)
    return nil
}

func init() {
    var err error
    evilImg, _, err = ebitenutil.NewImageFromFile("evil.png", ebiten.FilterDefault)
    if err != nil {
        panic(err)
    }
}

func main() {
    if err := ebiten.Run(update, 640, 480, 1, "hello, world"); err != nil {
        panic(err)
    }
}
