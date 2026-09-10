package main

import "fmt"

func main() {

	var a map[string]int
	fmt.Println(a)

	//puting values in a map

	// m["ali"] = 4

	// causig error because a nill map cannot get values in a way
	//we have to use make for inserting values

	fmt.Println(a == nil) //because map is nill not empty

	a = make(map[string]int) // intialize the map
	fmt.Println(a == nil)    // not nill becuse its empty but not null

	//now putting the values
	a["irfan"] = 22
	a["ali"] = 21

	fmt.Println(a)

	//others ways of delcaring the maps

	b := map[string]int{
		"ali":   1,
		"usman": 2,
	}
	fmt.Println("___________map=b___________")
	fmt.Println(b)

	//adding more values
	b["yaseen"] = 14
	b["hunza"] = 12

	fmt.Println("__after putting new values__")
	fmt.Println(b)

	//updating values
	b["hunza"] = -999
	fmt.Println("__ updating hunza values__")
	fmt.Println(b)

	//accesting map with keys
	fmt.Println("_____accesting map with keys _____")
	fmt.Println(b["hunza"])

	//3rd way of map declaraing

	c := make(map[string]int)

	fmt.Println("_______is nill or empty________")
	fmt.Println(c)
	fmt.Print("____is empty:")
	fmt.Println(c == nil)

	//putting values
	c["yousaf"] = 2
	fmt.Println(c)
	fmt.Println(c["yousaf"] == 2)

	c["ali"] = 2
	fmt.Println(c["ali"] == c["yousaf"])

	var d = map[int]string{

		1: "ali", 2: "usman",
		3: "yaqoob"}

	fmt.Println(d)
	fmt.Println("____deleting the key (2)______")
	delete(d, 2)
	fmt.Println(d)

	kasii, ok := d[1]

	if ok {
		fmt.Println("kasii gets:", kasii)
	}

	//maps with mising keys
	fmt.Println("__________________handling msiing keys___________")
	fmt.Println(d[0])
	nader, exist := d[0]
	if exist {
		fmt.Print(nader)

	} else {
		fmt.Print("value not exist")
	}
}
