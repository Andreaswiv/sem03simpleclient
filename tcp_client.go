package main

import (
    "log"
    "net"
    "os"
    "strings"

    "github.com/Andreaswiv/is105sem03/mycrypt"
)

func main() {
    conn, err := net.Dial("tcp", "172.17.0.4:8888")
    if err != nil {
        log.Fatal(err)
    }

    log.Println("os.Args[1] = ", os.Args[1])

    var response string

    if strings.ToLower(os.Args[1]) == "ping" {
        _, err = conn.Write([]byte("ping"))
        if err != nil {
            log.Fatal(err)
        }

        buf := make([]byte, 1024)
        n, err := conn.Read(buf)
        if err != nil {
            log.Fatal(err)
        }

        response = string(buf[:n])
    } else {
        kryptertMelding := mycrypt.Krypter([]rune(os.Args[1]), mycrypt.ALF_SEM03, 4)
        log.Println("Kryptert melding: ", string(kryptertMelding))

        _, err = conn.Write([]byte(string(kryptertMelding)))
        if err != nil {
            log.Fatal(err)
        }

        buf := make([]byte, 1024)
        n, err := conn.Read(buf)
        if err != nil {
            log.Fatal(err)
        }

        kryptertSvar := string(buf[:n])
        log.Println("Dekryptert svar: ", kryptertSvar)

        dekryptertSvar := mycrypt.Krypter([]rune(kryptertSvar), mycrypt.ALF_SEM03, len(mycrypt.ALF_SEM03)-4)
        log.Println("Dekryptert svar: ", string(dekryptertSvar))

        response = string(dekryptertSvar)
    }

    log.Println("Svar fra serveren: ", response)
}
