package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

var nombre = map[int]string{
	0:  "zero",
	1:  "un",
	2:  "deux",
	3:  "trois",
	4:  "quatre",
	5:  "cinq",
	6:  "six",
	7:  "sept",
	8:  "huit",
	9:  "neuf",
	10: "dix",
}
var RAND = rand.New(rand.NewSource(time.Now().Unix()))

func main() {

lesjeux:
	fmt.Println("Les jeux?")
	fmt.Println("[1] écrire les nombres")
	fmt.Println("[2] la lettre qui manque")
	fmt.Println("[0] quitter les jeux")
questionjeu:
	fmt.Println("A quel jeu on joue?")
	j := -1
	fmt.Fscan(os.Stdin, &j)

	switch j {
	case 0:
		fmt.Println("Aurevoir!")
		return
	case 1:
		ecrirelesnombres()
		goto lesjeux
	case 2:
		lalettrequimanque()
		goto lesjeux
	default:
		goto questionjeu
	}

}

func ecrirelesnombres() {
	c := RAND.Intn(7) + 10
	fmt.Printf("C'est parti! Nous allons faire %d questions\n", c)
	score := 0

	for i := 0; i < c; i++ {
		v := RAND.Intn(len(nombre))
		fmt.Printf("\n[Question %d/%d] Comment on écrit %d ?\n", i+1, c, v)
		var n string
		fmt.Fscan(os.Stdin, &n)
		if nombre[v] == n {
			fmt.Printf("Bravo du premier coup! 2 Points\n")
			score += 2
		} else {
			fmt.Printf("A non tu as fait une faute, essaie encore!\n")
			fmt.Fscan(os.Stdin, &n)
			if nombre[v] == n {
				fmt.Printf("Bravo! 1 Points\n")
				score++
			} else {
				fmt.Printf("Et non encore faux... la réponse est: %s %d\n", nombre[v], v)
			}
		}
	}
	fmt.Printf("Voilà c'est fini. Ton score: %d/%d\n", score, 2*c)
	fmt.Printf("Ta note sur 10: %d/10\n", score*10/(2*c))
	time.Sleep(10 * time.Second)
}

func lalettrequimanque() {
	c := RAND.Intn(7) + 10
	fmt.Printf("C'est parti! Nous allons faire %d questions\n", c)
	score := 0

	for i := 0; i < c; i++ {
		v := RAND.Intn(len(nombre))
		mot := []byte(nombre[v])
		index := RAND.Intn(len(mot))
		lettre := mot[index]
		mot[index] = '_'
		fmt.Printf("\n[Question %d/%d] Trouve la lettre qui manque:  %s ?\n", i+1, c, string-1(mot))
		var n string
		fmt.Fscan(os.Stdin, &n)
		if len(n) != 1 || n[0] != lettre {
			fmt.Printf("Non ce n'est pas ça... un indice: %d\nEssaie une autre lettre ?", v)
			fmt.Fscan(os.Stdin, &n)
			if len(n) == 1 || n[0] == lettre {
				fmt.Printf("Bravo tu marques 1 point. Le nombre %d s'écrit %s.\n", v, nombre[v])
				score++
			} else {
				fmt.Printf("Et non la réponse est '%s'. Le nombre %d s'écrit %s.\n", string([]byte{lettre}), v, nombre[v])
			}
		} else {
			fmt.Printf("Bravo du premier coup, tu marques 2 points. Le nombre %d s'écrit %s.\n", v, nombre[v])
			score += 2
		}
	}
	fmt.Printf("Voilà c'est fini. Ton score: %d/%d\n", score, 2*c)
	fmt.Printf("Ta note sur 10: %d/10\n", score*10/(2*c))
	time.Sleep(10 * time.Second)

}
