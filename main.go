package main

import (
    "fmt"
    "os"
)

type authService struct{}

func (auth *authService) Exit() {
    fmt.Println("GoodBye")
    os.Exit(0)
}

func main() {
    auth := &authService{}

    for {
        fmt.Println("--- Welcome to System ---")
        fmt.Println("1. Register")
        fmt.Println("2. Login")
        fmt.Println("3. Forgot Password")
        fmt.Println("0. Exit")

        var menu int
        fmt.Print("Choose menu : ")
        fmt.Scan(&menu)

        switch menu {
        case 0:
            auth.Exit()
        default:
            fmt.Println("Feature under development")
        }
    }
}