package main 

import (
    "bytes"
    "fmt"
    "io"
    "net"
    "os"
)


func main(){
    service := ":80"
    conn, err := net.Dial("tcp", service)
    checkError(err)

    _, err = conn.Write([] byte("HEAD / HTTP/1.0\r\n\r\n"))
    checkError(err)

    result, err := readFull(conn)
    checkError(err)

    fmt.Println(string(result))
    os.Exit(0)
}

func readFull(conn net.Conn)([]byte, error){
    defer conn.Close()

    result := bytes.NewBuffer(nil)
    var buf [512] byte
    for {
        n, err := conn.Read(buf[0:])
        result.Write(buf[0:n])
        if err != nil{
            if err == io.EOF{
                break
            }
            return nil, err
        }
    }
    return result.Bytes(), nil
}
func checkError(err error) { 
    if err != nil {
        fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
        os.Exit(1) 
    }
}