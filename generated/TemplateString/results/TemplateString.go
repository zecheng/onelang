package main

import "fmt"
import "strconv"

type TestClass struct {
    
}

func NewTestClass() *TestClass {
    this := new(TestClass)
    
    return this
}

func (this *TestClass) TestMethod()  {
    str_val := "str"
    num := 1337
    b := true
    result := fmt.Sprintf("before %v, num: %v, true: %v after", str_val, num, b)
    fmt.Println(result)
    fmt.Println(fmt.Sprintf("before %v, num: %v, true: %v after", str_val, num, b))
    
    result2 := "before " + str_val + ", num: " + strconv.Itoa(num) + ", true: " + strconv.FormatBool(b) + " after"
    fmt.Println(result2)
}

func main() {
    c := (TestClass{})
    c.TestMethod();
}