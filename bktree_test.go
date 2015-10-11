package bktree

import (
    "fmt"
    "./bktree"
    "log"
    "strings"
)

func main() {
    if bktree.Levenshtein("hello", "Aelo") != 2 {
        log.Fatal()
    }

    if bktree.Levenshtein("我爱你", "你爱我") != 2 {
        log.Fatal()
    }

    bk := bktree.NewBKTreeNode("", 0)

    bk.Insert("ABCD")
    bk.Insert("ACED")
    bk.Insert("SBDE")
    bk.Insert("가나다")
    bk.Insert("가나다라")
    bk.Insert("AABC")
    bk.Insert("abcd")
    bk.Insert("acd")
    bk.Insert("go")
    bk.Insert("went")
    bk.Insert("love")
    bk.Insert("bool")
    bk.Insert("fact")
    bk.Insert("join")
    bk.Insert("eleven")
    bk.Insert("seventeen")
    bk.Insert("element")
    bk.Insert("test")
    
    bk.Display()  // <----- Display in tree structured data

    fmt.Printf("bk.Size(): %d\n", bk.Size())  // <-- how many word is in this tree

    ret := bk.Find("AABB", 2)
    fmt.Println(ret)

    ret = bk.Find("joan", 1)
    fmt.Println(ret)
}
