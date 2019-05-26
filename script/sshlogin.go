package main
import "fmt"
import "bytes"
import "golang.org/x/crypto/ssh"

func main(){
    ip := "192.168.3.192:22"
    user := "root"
    password := "zz@123456"
    if login(ip, user, password){
        fmt.Println(ip + " ok")
    }
}

func login(ip string, user string, password string) bool {
    config := &ssh.ClientConfig{
        User: user,
        Auth: []ssh.AuthMethod{
            ssh.Password(password),
        },
        HostKeyCallback: ssh.InsecureIgnoreHostKey(),
    }
    client, err := ssh.Dial("tcp", ip, config)
    if err != nil {
        panic("Failed to dial: " + err.Error())
    }

    session, err := client.NewSession()
    if err != nil {
        panic("Failed to create session: " + err.Error())
    }
    defer session.Close()

    var b bytes.Buffer
    session.Stdout = &b
    if err := session.Run("ls /opt/"); err != nil {
        panic("Failed to run: " + err.Error())
    }
    fmt.Println(b.String())
    return true
    
}
