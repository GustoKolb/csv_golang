package main

import (

    "fmt"
    "os"
    "log"
    "bufio"
    "github.com/GustoKolb/csv/internal/utils"
)

func main() {

    pessoa := utils.Pessoa{
            Nome: "augusto",
            Idade: 20,
        }

    pessoa.Apresentar()

    args := os.Args

    if len(args) == 1 {
        fmt.Println("No File Selected")
        return 
        //Requisitar ao usuário que insira algum arquivo
    }
        
    file_path := args[1]
    fmt.Println(file_path)

    file, err := os.Open(file_path)

    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("File Opened")
    defer file.Close()

    scanner := bufio.NewScanner(file)

    for scanner.Scan() {

        fmt.Println(scanner.Text())
    }

    //Conseguir o número de linhas e o número de colunas
    //Alocar memória para uma matriz






}
