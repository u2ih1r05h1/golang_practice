package main

// comments

/* comments */

/*
    comments
*/

import (
    "fmt"
)

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
}

