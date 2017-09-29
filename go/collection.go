package main 

import "strings"
import "fmt"



func Index(vs []string, t string) int{
    for i, v := range vs {
        if v == t {
            return i
        }
    }
    return -1
}

func Include(vs []string, t string) bool {
    return Index(vs, t) >= 0
}

func Any (vs []string, f func(string) bool) bool{
    for _, v := range vs {
        if f(v){
            return true
        }
    }

    return false;
}

func All(vs []string, f func(string) bool) bool{
    
    for _, v := range vs {
        if !f(v){
            return false
        }
    }
    return true
}

func Filter(vs []string, f func(string) bool) []string {
    vsf := make([]string, 0)
    for _,v := range vs{
        if f(v){
            vsf = append(vsf, v)
        }
    }
    return vsf
}

func Map(vs []string, f func(string) string) []string{
    vsm := make([]string, len(vs))
    for i, v := range vs {
        vsm[i] = f(v)
    }
    return vsm
}

func main() {
    var strs = []string{"apple", "pear", "plum" , "peach"}
    fmt.Println(Index(strs, "peach"))

    fmt.Println(Include(strs, "peach"))


    fmt.Println(Any(strs, func(v string) bool {
        return strings.HasPrefix(v, "p")
    }))

    fmt.Println(All(strs, func(v string) bool{
        return strings.HasPrefix(v, "p")
    }))

    fmt.Println(Filter(strs, func (v string) bool {
        return strings.HasPrefix(v, "p")
    }))

    fmt.Println(Map(Filter(strs, func (v string) bool {
        return strings.HasPrefix(v, "p")
    }),func (v string) string {
        return strings.ToUpper("-" + v  + "-")
    }))
}











