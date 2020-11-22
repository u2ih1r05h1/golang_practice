package main

// comments

/* comments */

/*
    comments
*/

import (
    "fmt"
    "math"
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

    fmt.Printf("10進数=%d 2進数=%b 8進数=%o 16進数=%x\n", 17, 17, 17, 17)

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

    var b3 bool
    b3 = true
    fmt.Println(b3)

    b4 := false
    fmt.Println(b4)

    n6 := 9223372036854775807
    fmt.Println(n6)

    n7 := 123
    n8 := 0755
    n9 := 0x0719BEEF
    fmt.Println(n7, n8, n9)

    n10 := uint(17)
    fmt.Println(n10)

    n11 := 1
    b5  := byte(n11)
    i64 := int64(n11)
    u32 := uint32(n11)
    fmt.Println(n11, b5, i64, u32)

    n12 := 256
    b6  := byte(n12)
    fmt.Println(n12, b6)

    b7 := byte(255)
    b7 = b7 + 1
    fmt.Println(b7)

    n13 := -1
    b8  := byte(n13)
    fmt.Println(n13, b8)

    b9  := byte(255)
    b10 := b9 + byte(255)
    fmt.Println(b9, b10)

    ui_1 := uint32(400000000)
    ui_2 := uint32(4000000000)
    sum := ui_1 + ui_2
    fmt.Printf("%d + %d = %d\n", ui_1, ui_2, sum)

    fmt.Printf("uint32 max value = %d\n", math.MaxUint32)

    fmt.Println(doSomething(1, 2))

    i3  := 1
    f64 := 1.0
    f32 := float32(1.0)
    fmt.Println(i3, f64, f32)

    zero :=  0.0
    pinf :=  1.0  / zero
    ninf := -1.0  / zero
    nan  :=  zero / zero
    fmt.Println(zero, pinf, ninf, nan)

    fmt.Println(1.0e2)
    fmt.Println(1.0e+2)
    fmt.Println(1.0e-2)

    fmt.Printf("value = %v\n", 1.0000000000000000)
    fmt.Printf("value = %v\n", 1.0000000000000001)
    fmt.Printf("value = %v\n", 1.0000000000000002)
    fmt.Printf("value = %v\n", 1.0000000000000003)
    fmt.Printf("value = %v\n", 1.0000000000000004)
    fmt.Printf("value = %v\n", 1.0000000000000005)
    fmt.Printf("value = %v\n", 1.0000000000000006)
    fmt.Printf("value = %v\n", 1.0000000000000007)
    fmt.Printf("value = %v\n", 1.0000000000000008)
    fmt.Printf("value = %v\n", 1.0000000000000009)

    fmt.Printf("value = %v\n", float32(1.0000000000000000))
    fmt.Printf("value = %v\n", float32(1.0000000000000001))
    fmt.Printf("value = %v\n", float32(1.0000000000000002))
    fmt.Printf("value = %v\n", float32(1.0000000000000003))
    fmt.Printf("value = %v\n", float32(1.0000000000000004))
    fmt.Printf("value = %v\n", float32(1.0000000000000005))
    fmt.Printf("value = %v\n", float32(1.0000000000000006))
    fmt.Printf("value = %v\n", float32(1.0000000000000007))
    fmt.Printf("value = %v\n", float32(1.0000000000000008))
    fmt.Printf("value = %v\n", float32(1.0000000000000009))

    fmt.Println(float32(1.0) / float32(3.0))
    fmt.Println(float64(1.0) / float64(3.0))

    f1   := 3.14
    n14 := int(f1)
    fmt.Println(f1, n14)

    f2 := -3.14
    n15 := int(f2)
    fmt.Println(f2, n15)

    c := 1.0 + 3i
    fmt.Println(c)

    c1 := complex(1.0, 3)
    fmt.Println(c1)

    c2 := 1.3 + 4.2i

    fmt.Println(real(c2))
    fmt.Println(imag(c2))

    r := '松'
    fmt.Printf("%v\n", r)

    s2 := "Goの文字列"
    fmt.Printf("%v\n", s2)

    s3 := `
    Goの
    RAW文字列リテラルによる
    複数行に渡る
    文字列
    `
    fmt.Printf("%v\n", s3)

    s4 := `abc`
    s5 := `\n
    \n`
    fmt.Println(s4, s5)
}

func one() int {
    return 1
}

func doSomething(a, b uint32) bool {
    if (math.MaxUint32 - a) < b {
        return false
    } else {
        sum := a + b
        fmt.Println(sum)
        return true
    }
}
