package main

import (
    . "g2d"
    "math/rand"
    "fmt"
)

var image = LoadImage("ball.png")
var x, y, dx = 10, rand.Intn(360), 5
var screen = Size{480, 360}
var count = 0

func tick() {
    if x + dx > screen.W + 100 { 
        x = -100
    }
    if x +dx < -100 {
        x = screen.W + 100
    }
    ClearCanvas()
    DrawImage(image, Point{x, y})
    //x += dx
    if KeyPressed("Enter") { 
        for i := 0; i < 5; i++ {
            x += dx
        }
        count += 1
        fmt.Println("Enter è stato premuto ", count, " volte")
    }
    if KeyPressed("+") {
        dx = -dx
    }
}

func main() {
    InitCanvas(screen)
    MainLoop(tick)
    
}
