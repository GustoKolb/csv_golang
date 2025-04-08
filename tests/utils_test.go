package utils

import (

    "testing"
    "os"
    "fmt"
    "github.com/GustoKolb/csv/internal/utils"
)

func TestGetInfo(t *testing.T) {

    file, err := os.Open("../cmd/main/texto.csv")
    if err != nil {
        fmt.Println("No file")
    }

    info, err := utils.GetInfo(file)
    fmt.Printf("linhas: %d, colunas: %d, wordspace: %d", info.NumLines, info.NumCols, info.WordSpace)



}
