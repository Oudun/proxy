package main

import (
    "fmt"
    "os"
    "net"
)

const (
    connHost = "localhost"
    connPort = "8001"
    connType = "tcp"
)

type address struct {
    length []byte
    address []byte
    port [2]byte
}

func method_0(conn net.Conn) {
    req := []byte{5, 0}
    conn.Write(req)
    res := make([]byte, 4)
    conn.Read(res)
    fmt.Println("Response", res)
    if (res[3]==3) {
        len_buf := make([]byte, 1)
        conn.Read(len_buf)
        fmt.Println("Length", len_buf)
        addr_buf := make([]byte, int(len_buf[0]))
        conn.Read(addr_buf)
        fmt.Println("Address", addr_buf)
        port_buf := make([]byte, 2)
        conn.Read(port_buf)
        fmt.Println("Port", port_buf)
    }
    
    res1 := make([]byte, 100)
    i, _ := conn.Read(res1)
    fmt.Println("Response", string(res1[0:i-2]))    
}

func method_2(conn net.Conn) {
    fmt.Printf("Method 2 not implemented")
}

func main() {
    fmt.Println("Starting " + connType + " server on " + connHost + ":" + connPort)
    l, err := net.Listen(connType, connHost+":"+connPort)
    if err != nil {
        fmt.Println("Error listening:", err.Error())
        os.Exit(1)
    }
    defer l.Close()
    for {
        c, e := l.Accept()
        if e != nil {
            fmt.Println("Error accepting", e.Error())
            os.Exit(1)
        }
        buf := make([]byte, 257)
        _, e1 := c.Read(buf)
        if e1 != nil {
            fmt.Println("Error reading request", e.Error())
        }
        fmt.Println("Num methods", int(buf[1]))        
        for i := 0; i < int(buf[1]); i++ {
            fmt.Printf("Method: %d\n", buf[2+i])
            if buf[2+i] == 0 {
                method_0(c)
            }
            if buf[2+i] == 2 {
                method_2(c)
            }
        }
        fmt.Println("End")
    }

}