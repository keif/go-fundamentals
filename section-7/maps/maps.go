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

	fmt.Println("websites: ", websites)
	// the key won't error if mispelled, it just won't print out anything.
	fmt.Println("ex: ", websites["amazon web services"])
	websites["LinkedIn"] = "https://www.linkedin.com"
	fmt.Println("websites: ", websites)

	delete(websites, "google")
	fmt.Println("websites: ", websites)
}
