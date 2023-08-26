package main

import (
	"fmt"
	"os"
)

type Friend struct {
	id         uint
	name       string
	address    string
	occupation string
	reason     string
}

var friends []Friend = []Friend{
	{
		id:         1,
		name:       "Andry",
		address:    "Kuningan",
		occupation: "Software Engineer",
		reason:     "Karena golang laku dipasaran",
	},
	{
		id:         2,
		name:       "Yuni",
		address:    "Cipete",
		occupation: "QA",
		reason:     "Karena golang populer",
	},
	{
		id:         3,
		name:       "Willy",
		address:    "Ciledug",
		occupation: "UI/UX Developer",
		reason:     "Karena golang lebih cepat dan efisien",
	},
	{
		id:         4,
		name:       "Zee",
		address:    "Tigaraksa",
		occupation: "IT Consultant",
		reason:     "Karena golang memiliki fitur garbage collection",
	},
}

func main() {
	var arg = os.Args

	if len(arg) != 2 {
		fmt.Println("Anda belum memasukkan nama")
	} else {
		findMyFriend(arg[1])
	}
}

func findMyFriend(name string) {
	for _, eachFriend := range friends {
		if name == eachFriend.name {
			fmt.Println("id:", eachFriend.id)
			fmt.Println("nama:", eachFriend.name)
			fmt.Println("alamat:", eachFriend.address)
			fmt.Println("pekerjaan:", eachFriend.occupation)
			fmt.Println("alasan mempelajari golang:", eachFriend.reason)

		}
	}
}
