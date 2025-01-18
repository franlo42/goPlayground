package main

import (
	"fmt"
	"time"
)

func main() {
	total := 50 // Tamaño de la barra
	for i := 0; i <= total; i++ {
		// Calcula el progreso en porcentaje
		percent := float64(i) / float64(total) * 100

		// Genera la barra
		bar := fmt.Sprintf("[%s%s]", string(repeat('#', i)), string(repeat('-', total-i)))

		// Imprime la barra con el progreso
		fmt.Printf("\r%s %.2f%%", bar, percent)
		time.Sleep(100 * time.Millisecond) // Simula el progreso
	}
	fmt.Println("\n¡Descarga completada!")
}

// repeat es una función auxiliar para repetir un carácter n veces.
func repeat(char rune, count int) []rune {
	result := make([]rune, count)
	for i := 0; i < count; i++ {
		result[i] = char
	}
	return result
}
