package main

import (
	"fmt"
	"time"
)

func spinner(delay time.Duration) {
	// Array con los caracteres del spinner.
	chars := []rune{'|', '/', '-', '\\'}
	for {
		for _, char := range chars {
			fmt.Printf("\r%c Cargando...", char) // \r regresa el cursor al inicio de la línea
			time.Sleep(delay)                    // Pausa entre cada frame
		}
	}
}

func main() {
	go spinner(100 * time.Millisecond) // Lanza el spinner en una gorutina
	time.Sleep(5 * time.Second)        // Simula una tarea que tarda 5 segundos
	fmt.Println("\r¡Completado!     ") // Limpia la línea al terminar
}
