## bktree (BK Tree)

A BK Tree implementation (a tree structure for spelling checker - fuzzy search of a string)

~~~ go
package main

import (
	"fmt"
	"github.com/nomota/bktree"
	"log"
)

func main() {

    if bktree.Levenshtein("hello", "Aelo") != 2 {
        log.Fatal()
    }

    if bktree.Levenshtein("我爱你", "你爱我") != 2 {
        log.Fatal()
    }

    bk := bktree.NewBKTreeNode("", 0) // <--- Initialize the tree

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

    bk.Display() // <--- display indentated tree structure

    fmt.Printf("bk.Size(): %d\n", bk.Size())  // <-- how many words are in the tree

    ret := bk.Find("AABB", 2)
    fmt.Println(ret)

    ret = bk.Find("joan", 1)
    fmt.Println(ret)
}

~~~

## LICENSE

MIT
