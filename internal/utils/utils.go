package utils

import (
    "os"
    "bufio"
    "errors"
    "strings"
    "fmt"
)


type FileInfo struct {

    NumLines int
    NumCols int
    WordSpace int

}


//------------------------------------------------------------------
func GetLines( file *os.File) (int, error) {

    numLines := 0
    var err error

    if file == nil {
        return 0, errors.New("GetLines: no file")
    }
        
    //Reset pointer to the beginning of the file
    _, err = file.Seek(0, 0)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {

        numLines++

    }
    return numLines, nil

}

//------------------------------------------------------------------
func SplitByCommas( data []byte, atEOF bool) (advanced int,token []byte, err error) {

    if atEOF && len(data) == 0 {
        return 0, nil, nil
    }

    for i, b := range data {
        if b == ',' {

            return i + 1, data[0:i], nil
        }
    }

    if atEOF {
        return len(data), data, nil
    }
    return

}
//------------------------------------------------------------------

func GetCols (file *os.File) (int, error) {

    numCols := 0
    var err error

    if file == nil {
        return 0, errors.New("GetCols: no file")
    }

    //Reset pointer to the beginning of the file
    _, err = file.Seek(0, 0)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

    scanner := bufio.NewScanner(file)
    if scanner.Scan() {
        
        for range strings.Split(scanner.Text(),",") {

            numCols++

        }

    }
    return numCols, nil


}


//------------------------------------------------------------------
func GetWordSpace (file *os.File) (int, error) {

    wordSpace := 0
    var err error

    if file == nil {
        return 0, errors.New("GetWordSpace: no file")
    }

    //Reset pointer to the beginning of the file
    _, err = file.Seek(0, 0)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

    scanner := bufio.NewScanner(file)
    scanner.Split(SplitByCommas)

    for scanner.Scan(){
    
        fmt.Println(scanner.Text())
        if len(scanner.Text()) > wordSpace {
            fmt.Println(wordSpace)
            wordSpace = len(scanner.Text())
        }
    }
    return wordSpace, nil

}

//------------------------------------------------------------------
//func GetMatrix (file *os.File, numLines int, numCols int) [][]string {}

//------------------------------------------------------------------
func GetInfo (file *os.File) (*FileInfo, error) {

    //info *FileInfo
    info := new(FileInfo)
    var err error

    if file == nil {
        return nil, errors.New("GetInfo: no file")
    }
     
    info.NumLines, err = GetLines(file)   
    if (err != nil){
        return nil, err
    }

    info.NumCols, err = GetCols(file)  
    if (err != nil){
        return nil, err
    }

    info.WordSpace, err = GetWordSpace(file)
    if (err != nil){
        return nil, err
    }
    return info, nil

 }
//------------------------------------------------------------------




