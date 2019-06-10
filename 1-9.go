package main

import . "g2d"
import "math/rand"

func main() {
    player := [2]int { rand.Intn(5), rand.Intn(5) }
    enemy := [2]int { rand.Intn(5), rand.Intn(5) }
    victory := [2]int { rand.Intn(5), rand.Intn(5) }
    
    Alert("Howdy!")
    for {
        d := Prompt("Where do you want to move? - u for up and d for down, l for left and r for right")
        if d == "u" {
            if player[1] != 5 {
                player[1] = player[1] + 1
            }
        }
        if d == "" {
            if player[1] != 0 {
                player[1] = player[1] - 1
            }
        }
        if d == "l" {
            if player[0] != 0 {
                player[0] = player[0] - 1
            }
        }
        if d == "r" {
            if player[0] != 5 {
                player[0] = player[0] + 1
            }
        }
        if player == enemy {
            Alert("You have lost")
            break
            }
        if player == victory {
            Alert("You have won")
            break
            }
    }
}
