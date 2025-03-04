package utils

import {
    "fmt"
    "os"
    "bufio"
}


type fileInfo struct {

    numLines int
    numCol int
    wordSpace int

}



func GetLines( file *os.File) (int, error) {

    numLines := 0

    if file == nil
        return 0, errors.New("GetLines: no file")
        
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {

        numLines++

    }
    return numLines, nil

}

func SplitByCommas( data []byte, atEOF bool) (advanced int, 
token []byte, err error) {

    if atEOF && len(data) == 0 {
        return 0, nil, nil
    }

    for i, b := range data {
        if b == ',' {

            return i + 1, data[0:i], nil
        }
    }


    }
    if atEOF {
        return len(data), data, nil
    }
    return

}

func GetCols (file *os.File) int {}

func GetMatrix (file *os.File, numLines int, numCols int) [][]string {}

func getInfo (file *os.File) (fileInfo, error) {

    info *fileInfo

    if file == nil {
        return 0, errors.New("GetInfo: no file")
    }
     
    numLines := 0   
    numCols := 0

    //pegar o numero de linhas
    //adicionar um reader para pegar a primeira linha com o numero de colunas
    // pegar palavra por palavra e por num slice 2d

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {



    scanner := bufio.NewScanner(file)
    for scanner.Scan() {

        numLines++
        wordScanner := bufio.NewScanner(scanner.Text)
        scanner.Split(SplitByCommas)

        

    }




}




