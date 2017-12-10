package main

import "fmt"

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

func main() {
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
}
