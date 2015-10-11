## bktree (BK Tree)

A BKTree implementation

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

	bk := bktree.New()
	bk.SetLevenshteinLimit(50)

	bk.Insert("ABCD")
	bk.Insert("ACED")
	bk.Insert("SBDE")

	ret := bk.Find("AABB", 3)
	fmt.Println(ret)
}

~~~

## LICENSE

MIT
