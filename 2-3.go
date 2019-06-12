package main

import . "g2d"
import "math/rand"

var image = LoadImage("ball.png")
var x, y, dx = 10, rand.Intn(360), 5
var screen = Size{480, 360}

func tick() {
    if KeyPressed("Enter") { 
        dx = -dx
    }
    if x + dx > screen.W + 100 { 
        x = -100
    }
    if x +dx < -100 {
        x = screen.W + 100
    }
    ClearCanvas()
    DrawImage(image, Point{x, y})
    x += dx
}

func main() {
    InitCanvas(screen)
    MainLoop(tick)
}
