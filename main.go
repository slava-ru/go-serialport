package main

import (
    "fmt"
    "github.com/tarm/goserial"
    //"time"
)

func main() {
    c := &serial.Config{Name: "/dev/ttyACM0", Baud: 115200}
    s, err := serial.OpenPort(c)

    if err != nil {
            fmt.Println(err)
    }

    _, err = s.Write([]byte("\x0a"))

    if err != nil {
            fmt.Println(err)
    }

    //time.Sleep(time.Second/2)

    buf := make([]byte, 40)
    n, err := s.Read(buf)

    if err != nil {
            fmt.Println(err)
    }

    fmt.Print(string(buf[:n]))

    s.Close()

    main()

}