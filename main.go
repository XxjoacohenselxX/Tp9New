package main

import "fmt"

func bubbleSort(arr []int) {
    n := len(arr)
    for i := 0; i < n-1; i++ {
        // En cada iteración, el elemento más grande "burbujea" hasta la posición correcta
        changed := false
        for j := 0; j < n-i-1; j++ {
            if arr[j] > arr[j+1] {
                // Intercambiar elementos si están en orden incorrecto
                arr[j], arr[j+1] = arr[j+1], arr[j]
                changed = true
            }
        }

        // Confirmación para saber si se ordeno depués del arrays
        if !changed {
            break
        }
    }
}

func main() {
    // Ejemplo de array desordenado
    arr := []int{64, 34, 25, 12, 22, 11, 90}

    fmt.Println("Array original:", arr)
    bubbleSort(arr)
    fmt.Println("Array ordenado:", arr)
}
