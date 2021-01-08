package main

import "fmt"



func array_func(){

    var arr1 [10]int
    arr2 :=[...]float64{1,2,3,4,5}
    fmt.Printf("%d",len(arr1))
    fmt.Printf("%d",len(arr2))
}


func main() {
    var first_name = "rocky"
    last_name :="chen"
    first_string := first_name[:3]
    string_list :=[]byte(first_name)
    fmt.Printf("First name : %s, Last name : %s\n",first_name,last_name)
    fmt.Printf("first chareter %s\n",first_string)
    fmt.Printf("%s\n",string_list)
    fmt.Printf("End\n")

}
