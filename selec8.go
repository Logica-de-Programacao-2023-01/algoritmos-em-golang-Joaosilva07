package main

import "fmt"

func main() {
    var num int
    fmt.Println("Digite um número inteiro entre 1 e 7:")
    fmt.Scanln(&num)
    switch num {
    case 1:
        fmt.Println("Domingo")
    case 2:
        fmt.Println("Segunda-feira")
    case 3:
        fmt.Println("Terça-feira")
    case 4:
        fmt.Println("Quarta-feira")
    case 5:
        fmt.Println("Quinta-feira")
    case 6:
        fmt.Println("Sexta-feira")
    case 7:
        fmt.Println("Sábado")
    default:
        fmt.Println("Número inválido")
    }
}
