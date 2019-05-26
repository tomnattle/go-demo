package main

import(
    "fmt"
    "golang.org/x/crypto/ssh"
    "log"
    "time"
    "bytes"
)

type loginInfo struct{
    Ip string
    Port int
    User  string
    Password string
}

func main(){
    for i := 6; i <= 6; i++ {
    chs := make([] chan int, 255)
    validip := []string{}

    for j := 0; j <= 254; j++ {

        chs[j] = make(chan int)
        go func(d int,i int, ch chan int) {
            status, ip := login(loginInfo{
                Ip : fmt.Sprintf("192.168.%d.%d", d, i),
                Port : 22,
                User  : "tommy",
                Password : "zz123456",
            })
            if status {
                validip = append(validip, ip)
                log.Printf("success:%s", ip)
            }
            ch <- j
        }(i, j, chs[j])
    }

    for _, ch := range(chs) {
        <-ch
    }

    log.Println("success ids:", validip)
    }
}

func login(info loginInfo) (bool, string){
    log.Print(fmt.Sprintf("%s:%d\n", info.Ip, info.Port))

    client, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", info.Ip, info.Port), &ssh.ClientConfig{
        User: info.User,
        Auth: []ssh.AuthMethod{ssh.Password(info.Password)},
        Timeout: time.Second * 2,
        HostKeyCallback: ssh.InsecureIgnoreHostKey(),
    })

    if err != nil {
        log.Printf("%s:%s error: %v", info.Ip, "", err)
        return false, info.Ip
    }

    session, err := client.NewSession()
    if err != nil {
        //log.Printf("%s:%s warning: %v", info.Ip, "", "")
        return false, info.Ip
    }

    defer session.Close()
    var b bytes.Buffer
    session.Stdout = &b
    if err := session.Run("ls"); err != nil {
        return false, info.Ip
    }
    fmt.Println(b.String())

    //log.Println("success: ", info.Ip)
    return true, info.Ip
}
