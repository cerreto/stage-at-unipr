package main

import . "g2d"

func main() {
    radius := ToInt(Prompt("Forniscimi un valore del raggio tra 0 e 200"))
    if radius > 0 && radius <= 200 {
        InitCanvas(Size{640, 480})
        SetColor(Color{255, 255, 0})
        FillCircle(Point{320, 240}, radius)
        MainLoop()
    } else {
        Alert("Pensaci bene ragionando e leggendo")
    }
}
