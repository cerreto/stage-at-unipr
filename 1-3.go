package main

import . "g2d"

func main() {
    n := ToInt(Prompt("Inserisci il numero di cerchi desiderato"))
    InitCanvas(Size{640, 480})
    x := 0

    for i := 0; i < n-1; i++ {
        SetColor(Color{255-(i*35), 0, 0})
        FillCircle(Point{320, 240}, 200-(i*25))
        x = 200-(i*25)
    }
    SetColor(Color{0, 0, 0})
    FillCircle(Point{320, 240}, x)
    
    MainLoop()      
}
