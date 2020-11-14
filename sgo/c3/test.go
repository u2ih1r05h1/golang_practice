package main

// comments

/* comments */

/*
    comments
*/

import (
    "fmt"
)

var n0 = 100

func main() {

    var i int // comments
    i = 1
    fmt.Println(i)

    fmt.Println("Hello, World!")
    fmt.Println("Hello, World!");

    a := [3]string{
        "Yamada Taro",
        "Sato Hanako",
        "Suzuki Kenji",
    }
    fmt.Println(a)

    n := 1
    fmt.Println(n)

    n1 := 1 + 2 + 3 // n == 6

    fmt.Println(n1)

    fmt.Println("Hello, World!") // => "Hello, World!"

    fmt.Println("Hello, Golang!")

    fmt.Println("My", "name", "is", "Taro")

    fmt.Printf("数値=%d\n", 5)

    fmt.Printf("10進数=%d 2進数=%B 8進数=%o 16進数=%x\n", 17, 17, 17, 17)

    fmt.Printf("%d年%d月%d日\n", 2015, 12)

    fmt.Printf("%d年%d月%d日\n", 2015, 12, 25, 17)

    fmt.Printf("数値=%v 文字列=%v 配列=%v\n", 5, "Golang", [...]int{1, 2, 3})

    fmt.Printf("数値=%#v 文字列=%#v 配列=%#v\n", 5, "Golang", [...]int{1, 2, 3})

    fmt.Printf("数値=%T 文字列=%T 配列=%T\n", 5, "Golang", [...]int{1, 2, 3})

    print("Hello, World!")
    println("Hello, World!")

    println(1, 2, 3)

    var n2 int
    n2 = 1
    fmt.Println(n2)

    var x, y, z int
    x, y, z = 1, 2, 3
    fmt.Println(x, y, z)

    var (
        x1, y1 int
        name string
    )
    x1, y1 = 4, 5
    name = "gotaro"
    fmt.Println(x1, y1, name)

    var n3 int
    n3 = 5
    fmt.Println(n3)

    var x2, y2 int
    x2, y2 = 1, 2
    fmt.Println(x2, y2)

    i1 := 1
    fmt.Println(i1)

    b := true

    i2 := 1

    f := 3.14

    s := "abc"

    fmt.Println(b, i2, f, s)

    n4 := one()
    fmt.Println(n4)

    var a1 int
    a1 = 1
    a1 = 2

    b1 := 1
    b1 = 2

    fmt.Println(a1, b1)

    var a2 = 1
    fmt.Println(a2)

    var (
        n5 = 1
        s1 = "string"
        b2 = true
    )
    fmt.Println(n5, s1, b2)

    n0 = n0 + 1
    fmt.Printf("n0=%d\n", n0)
}

func one() int {
    return 1
}
