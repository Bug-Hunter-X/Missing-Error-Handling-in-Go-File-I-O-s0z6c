package main

import (


        "fmt"

        "os"


)

func main() {

        data, err := os.ReadFile("my_file.txt")

        if err != nil {

                if os.IsNotExist(err) {

                        fmt.Println("Error: File not found.")

                } else {

                        fmt.Printf("Error reading file: %v", err)

                }

                return

        }

        fmt.Println(string(data))

}