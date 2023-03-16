package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

type name struct {
    fname string
    lname string
}

func main() {
    var fileInput string

    fmt.Print("Enter the name of the text file: ")
    fmt.Scan(&fileInput)

    file, err := os.Open(fileInput)
    if err != nil {
        fmt.Println(err)
        return
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    var names []name

    for scanner.Scan() {
        line := scanner.Text()
        nameParts := strings.Split(line, " ")

        n := name{
            fname: nameParts[0],
            lname: nameParts[1],
        }

        names = append(names, n)
    }

    if err := scanner.Err(); err != nil {
        fmt.Println(err)
        return
    }

    for _, n := range names {
        fmt.Printf("First name: %s, Last name: %s\n", n.fname, n.lname)
    }
}

