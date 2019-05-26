package main
 
import (
    "os/exec"
)
 
func main() {
    exec.Command("sh", "-c", "pkill -SIGINT test1").Output()
}
