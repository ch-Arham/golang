package main

import "fmt"

type User struct {
	name    string
	email   string
	age     uint
	comment Comment // nested struct
	avatar          // embedded struct
}

type Comment struct {
	likes uint
	cmt   string
}

type avatar struct {
	platform string
	image    string
}

func main() {
	newUser := User{
		name:  "Arham",
		email: "email@.cm",
		age:   27,
		comment: Comment{
			likes: 0,
			cmt:   "asdasdasda",
		},
		avatar: avatar{
			platform: "firebase",
			image:    "https...",
		},
	}

	userAttributes := struct { // Anon struct
		height uint
		color  string
	}{
		height: 10,
		color:  "white",
	}

	fmt.Println("newUser:", newUser.comment.likes)
	fmt.Println("newUser:", newUser.platform)

	fmt.Println("userAttribute:", userAttributes.color)
}
