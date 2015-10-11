package bktree

import (
    "unicode/utf8"
    "runtime"
    "fmt"
)

func DEBUG_PRINT(format string, v ...interface{}) {
    _, fname, line, _ := runtime.Caller(1)
    fmt.Printf("(%s,%d) ", fname, line); fmt.Printf(format, v...)
}

type BKTreeNode struct {
    dist  int     // distance from its parent
    str   string
    child []*BKTreeNode
}

func NewBKTreeNode(s string, d int) *BKTreeNode {
    return &BKTreeNode{ str: s, dist: d, child: nil, }
}

func (this *BKTreeNode) display(level int) {
    for i := 0; i < level; i++ {
        fmt.Print("    ") // indentation
    }

    if this.child != nil && len(this.child) == 0 {
        fmt.Printf("str='%s',dist=%d\n", this.str, this.dist)
        return
    }


    fmt.Printf("str='%s',dist=%d\n", this.str, this.dist)

    if this.child != nil {
        for i, _ := range this.child {
            this.child[i].display(level+1)
        }
    }
}

func (this *BKTreeNode) Size() int {
    if this.str == "" { // Root is "", None is added yet
        return 0
    }

    if this.child != nil && len(this.child) == 0 {
        return 1
    }

    sum := 1
    if this.child != nil {
        for i, _ := range this.child {
            sum += this.child[i].Size()
        }
    }

    return sum
}

func (this *BKTreeNode) Display() {
    this.display(0)
}

func (this *BKTreeNode) Insert(s string) bool {
    if this.str == "" { // Root node
        this.str = s
        this.dist = 0
        return true
    }

    if s == this.str {
        return false // don't insert, duplicated
    }

    d := Levenshtein(this.str, s)

    if this.child == nil {
        this.child = make([]*BKTreeNode, 0)
    }

    for _, c := range this.child {
        if c.dist == d {
            if c.str == s {
                return false // don't insert, duplicated
            }
            return c.Insert(s)
        }
    }

    this.child = append(this.child, NewBKTreeNode(s, d))
    return true
}

func (this *BKTreeNode) Find(s string, k int) (ret []string) {
    d := Levenshtein(this.str, s)
    minD := d - k
    maxD := d + k

    if d <= k {
        ret = append(ret, this.str)
    }

    for _, c := range this.child {
        if minD <= c.dist && c.dist <= maxD {
            ret = append(ret, c.Find(s, k)...)
        }
    }

    return ret
}

// calculate Edit Distance of two utf-8 strings
func Levenshtein(s1, s2 string) int {
    m, n := utf8.RuneCountInString(s1), utf8.RuneCountInString(s2)
    runes1, runes2 := make([]rune, m), make([]rune, n)

    // copy runes, array of utf-8 characters, not bytes
    i, j := 0, 0
    for _, v := range s1 { // _ points the byte position within utf-8 string
        runes1[i] = v; i++
    }

    for _, v := range s2 { // _ points the byte position within utf-8 string
        runes2[j] = v; j++
    }

    d := make([][]int, 2)  // roll array
    d[0] = make([]int, n+1)
    d[1] = make([]int, n+1)

    turn, pre := 0, 0
    for i = 0; i <= n; i++ {
        d[turn][i] = i
    }
    for i = 1; i <= m; i++ {
        pre = turn
        turn = (turn + 1) % 2
        d[turn][0] = i

        for j = 1; j <= n; j++ {
            if runes1[i-1] == runes2[j-1] {
                d[turn][j] = d[pre][j-1]
            } else {
                d[turn][j] = min(min(d[pre][j]+1,d[turn][j-1]+1),d[pre][j-1]+1)
            }
        }
    }

    return d[turn][n]
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
