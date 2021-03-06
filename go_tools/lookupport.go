package main 

import (
    "net"
    "os"
    "fmt"
)
//go run lookupport.go tcp http
func main (){
    if len(os.Args) != 3{
        fmt.Fprintf(os.Stderr, "Usage: %s net work-type service\n", os.Args[0])
        os.Exit(1)
    }

    netWorkType := os.Args[1]
    service := os.Args[2]

    port, err := net.LookupPort(netWorkType, service)
    if err != nil {
        fmt.Println("Error: ", err.Error())
        os.Exit(2)
    }

    fmt.Println("service port:", port)
    os.Exit(0)
}