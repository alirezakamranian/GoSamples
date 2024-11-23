package main
import ("fmt")

var vari = 12 

var (
	a int
	b int =2
	c string
)

var names = [3]string{"ali","sina","mehdi"}
var names2 = [...]string{"ali","sina","mehdi"}

func main() {
	d := "HelloWorld"

	//nums1 := [5]int{4,5,6,7,8}
	//nums2 := [...]int{4,5,6,7,8}

	fmt.Println(d)
	sayhi();
}

func sayhi(){
	fmt.Println("hiiiiiiiiiiiiii")
}
//Command to Run: "go run ./HelloWorld.go"