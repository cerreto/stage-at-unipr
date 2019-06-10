package main

import . "g2d"
import "math/rand"

func main() {
    n := ToInt(Prompt("Inserisci il numero di quadrati desiderato"))
    InitCanvas(Size{640, 480})

    for i := 0; i < n; i++ {
        SetColor(Color{rand.Intn(255), rand.Intn(255), rand.Intn(255)})
        x := rand.Intn(200)
        FillRect(Rect{rand.Intn(640-x), rand.Intn(480-x), x, x})
    }
    
    MainLoop()      
}
