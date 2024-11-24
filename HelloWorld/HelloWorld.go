package main

import (
	"fmt"
)

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

	if d == "HelloWorld"{
		sayhi();
	}else{
		fmt.Println("else");
	}

	switch names[1]{

	case "sina":
		fmt.Println("sina")

	case "mahomad":
		fmt.Println("mahomad")

	case "ali","dara":
		fmt.Println("Baddies!!")

	default:
		fmt.Println("._.");
	}

	for i:=0 ; i<10 ; i++ {
		fmt.Println(i)
		if i>5{
			break
		}
	}

	for _,val := range names{
		fmt.Println(val)
	}

	for idx,val := range names{
		fmt.Println(idx,val)
	}

	response  := sayWordWithAllah("ya")

	fmt.Println(response)

	a,b := return10bomb()

	println(a,b)

    var person1 person 
	person1.Name="ali"
	person1.Age=19

	var mapsample = make(map[int]string)
	
	mapsample[1]="ali"
	mapsample[2]="lia"
	
	delete(mapsample,2)

	val,exist := mapsample[1]

	println(exist,val)
}

type person struct{
	Age int
	Name string
}

func sayhi(){
	fmt.Println("hiiiiiiiiiiiiii")
}

func sayWordWithAllah(preallah string)string{
	fmt.Println(preallah+" allah")
	return "ok"
}

func return10bomb()(res int,txt string){
	res=10
	txt="bomb"
	return
}
//Command to Run: "go run ./HelloWorld.go"