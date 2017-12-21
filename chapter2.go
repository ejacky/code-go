package main

import (
    "fmt"
    "math"
    "runtime"
)

type person struct {
    name string
    age int
}

func Older(p1, p2 person) (person, int) {
    if p1.age > p2.age {
        return p1, p1.age - p2.age
    }
    return p2, p2.age - p1.age
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func add(a int, b *int) (int, int) {
    a = a + 1
    *b = *b + 1
    return a, *b
}

func sumAndMulti(A, B int) (int, int) {
    return A + B, A * B
}

func getSum() {
    sum := 0
    for index := 0; index < 10; index++ {
        sum += index
    }

	fmt.Println("sum is equal to ", sum)
}

type testInt func(int) bool

func isOdd(integer int) bool {
    if integer % 2 == 0 {
        return false
    }
    return true
}

func isEven(integer int) bool {
    if integer % 2 == 0 {
        return true
    }
    return false
}

func filter(slice [] int, f testInt) []int {
    var result []int
    for _, value := range slice {
        if f(value) {
            result = append(result, value)
        }
    }
    return result
}

type Rectangle struct {
    width, height float64
}

type Circle struct {
    radius float64
}

func (r Rectangle) area() float64 {
    return r.width * r.height
}

func (c Circle) area() float64 {
    return c.radius * c.radius * math.Pi
}

const (
    WHITE = iota
    BLACK
    BLUE
    RED
    YELLOW
)

type Color byte

type Box struct {
    width, height, depth float64
    color Color
}

type BoxList []Box

func (b Box) Volume() float64 {
    return b.width * b.height * b.depth
}

func (b *Box) SetColor(c Color) {
    b.color = c
}

func (b1 BoxList) BiggestColor() Color {
    v := 0.00
    k := Color(WHITE)
    for _, b := range b1 {
        if b.Volume() > v {
            v = b.Volume()
            k = b.color
        }
    }
    return k
}

func (b1 BoxList) PaintItBlack() {
    for i, _ := range b1 {
        b1[i].SetColor(BLACK)
    }
}

func (c Color) String() string {
    strings := []string {"WHITE", "BLACK", "BLUE", "RED", "YELLOW"}
    return strings[c]
}

type Human struct {
    name string
    age int
    phone string
}

type Student struct {
    Human 
    school string
}

type Employee struct {
    Human 
    company string
}

func (h *Human) SayHi() {
    fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

func (e *Employee) SayHi() {
    fmt.Printf("Hi, I am %s , I work at %s, you can call me on %s\n", e.name, e.company, e.phone)
}

func say(s string) {
    for i := 0; i < 5; i++ {
        runtime.Gosched()
        fmt.Println(s)
    }
}
func main() {
    go say("world")
    say("hello")





/**
* 2.5.1
*
    r1 := Rectangle{12, 2}
    r2 := Rectangle{9, 4}
    c1 := Circle{10}
    c2 := Circle{25}

    fmt.Println("Area of r1 is : ", r1.area())
    fmt.Println("Area of r2 is : ", r2.area())
    fmt.Println("Area of c1 is : ", c1.area())
    fmt.Println("Area of c2 is : ", c2.area())
*/

/**
* 2.5.2
*
boxes := BoxList {
    Box{4, 4, 4, RED},
    Box{4, 4, 4, RED},
    Box{4, 4, 4, RED},
    Box{4, 4, 4, RED},
    Box{4, 4, 4, RED},
    Box{4, 4, 4, RED},
}

fmt.Printf("We have %d boxes in our set\n", len(boxes))
fmt.Println("The volume of the first one is", boxes[0].Volume(), "cm3")
fmt.Println("The color of the last one is", boxes[len(boxes)-1].color.String())
fmt.Println("The biggest one if ", boxes.BiggestColor().String())
fmt.Println("Let's paint them all black")
boxes.PaintItBlack()
fmt.Println("The color of the second one is", boxes[1].color.String())
fmt.Println("Obviously, now, the biggest one is", boxes.BiggestColor().String())
*/

/**
* 2.5.2
    mark := Student{Human{"Mark", 25, "222-222-YYYY"}, "MIT"}
    sam := Employee{Human{"Sam", 45, "111-888-XXX"}, "Golang Inc"}

    mark.SayHi()
    sam.SayHi()
*/

        //getSum()


/**
*  
    x := 3
    y := 4
    plus, times := sumAndMulti(x, y)

    fmt.Printf("%d + %d = %d\n", x, y, plus)
    fmt.Printf("%d * %d = %d\n", x, y, times)
*/

/**
*
    x := 3
    y := 4
    x1, y1 := add(x, &y)

    fmt.Println("x+1=", x1)
    fmt.Println("x=", x)

    fmt.Println("y+1=", y1)
    fmt.Println("y=", y)
*/

/**
 *
    slice := []int {1, 2, 3, 4, 5, 7}
    fmt.Println("slice=", slice)
    odd := filter(slice, isOdd)
    fmt.Println("Odd element of slice are:", odd)
    even := filter(slice, isEven)
    fmt.Println("Even elements of slice are:", even)
*/

/**
* 2.4
*
    var tom,paul,bob person

    tom.name, tom.age = "Tom", 18
    paul.name, paul.age = "Paul",43
    bob.name, bob.age = "Bob", 25
    // why the two other cann't 
    // paul := person("Paul", 43)
    // bob := person("Bob", 25)
    //bob := person(age:25, name:"Bob")

    tb_Older, tb_diff := Older(tom, bob)
    tp_Older, tp_diff := Older(tom, paul)
    bp_Older, bp_diff := Older(bob, paul)

    fmt.Printf("Of %s and %s, %s is older by %d years\n",
        tom.name, bob.name, tb_Older.name, tb_diff)
    fmt.Printf("Of %s and %s, %s is older by %d years\n",
        tom.name, paul.name, tp_Older.name, tp_diff)
    fmt.Printf("Of %s and %s, %s is older by %d years\n",
        bob.name, paul.name, bp_Older.name, bp_diff)

    fmt.Println("hello world!")
*/
}
