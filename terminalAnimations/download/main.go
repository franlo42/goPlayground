package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	total := 1024 * 1024 // 1 MB en bytes
	downloaded := 0

	fmt.Println("Descargando archivo...")
	for downloaded < total {
		// Simula bytes descargados
		step := rand.Intn(10000)
		if downloaded+step > total {
			step = total - downloaded
		}
		downloaded += step

		// Imprime el progreso
		percent := float64(downloaded) / float64(total) * 100
		fmt.Printf("\rProgreso: %d/%d bytes (%.2f%%)", downloaded, total, percent)
		time.Sleep(100 * time.Millisecond) // Simula el tiempo entre descargas
	}

	fmt.Println("\n¡Descarga completada!")
}
