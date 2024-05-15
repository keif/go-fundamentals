package main

import "fmt"

type Product struct {
	title string
	id    string
	price float64
}

func main() {
	hobbies := [...]string{"hobby1", "hobby2", "hobby3"}
	fmt.Println("First example:")
	fmt.Println(hobbies)
	fmt.Println("Second example:")
	fmt.Println(hobbies[0])
	fmt.Println(hobbies[1:])

	fmt.Println("Third example:")
	//mainHobbies := hobbies[0:2]
	mainHobbies := hobbies[:2]
	fmt.Println(mainHobbies)

	fmt.Println("Fourth example:")
	fmt.Println(cap(mainHobbies))
	mainHobbies = mainHobbies[1:3]
	fmt.Println(mainHobbies)

	fmt.Println("Fifth example:")
	courseGoals := []string{"Learn Go", "Learn basics"}
	fmt.Println(courseGoals)

	fmt.Println("Sixth example:")
	courseGoals[1] = "Learn Details"
	courseGoals = append(courseGoals, "Learn Stuff")
	fmt.Println(courseGoals)

	fmt.Println("SeventhR example:")
	products := []Product{
		Product{
			title: "Title One",
			id:    "0",
			price: 1.99,
		},
		Product{
			title: "Title Two",
			id:    "1",
			price: 2.99,
		},
	}
	fmt.Println(products)
	newProduct := Product{"Title Three", "3", 3.99}
	products = append(products, newProduct)
	fmt.Println(products)
}

// Time to practice what you learned!

// 1) Create a new array (!) that contains three hobbies you have
// 		Output (print) that array in the command line.
// 2) Also output more data about that array:
//		- The first element (standalone)
//		- The second and third element combined as a new list
// 3) Create a slice based on the first element that contains
//		the first and second elements.
//		Create that slice in two different ways (i.e. create two slices in the end)
// 4) Re-slice the slice from (3) and change it to contain the second
//		and last element of the original array.
// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
// 7) Bonus: Create a "Product" struct with title, id, price and create a
//		dynamic list of products (at least 2 products).
//		Then add a third product to the existing list of products.
