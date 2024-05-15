package main

import "fmt"

func main() {
	// this works, but if someone looks at this code, we just see URLs. Adding a label could make sense
	// and if you needed to retrieve a url, it'd need to be by index
	// maos can solve this
	//websites := []string{"https://www.google.com", "https://www.facebook.com", "https://www.twitter.com"}
	// empty map initialized
	// map[key]value{}
	websites := map[string]string{
		"google":              "https://www.google.com",
		"amazon web services": "https://www.aws.com",
	}

	fmt.Println(websites)

}
