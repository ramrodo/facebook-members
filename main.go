package main

import (
	"fmt"
	"log"

	"github.com/ramrodo/facebook-members/model"
	"github.com/ramrodo/facebook-members/repository"
)

type Key struct {
	Id string
}

type Value struct {
	Counter int
	Member  model.Member
}

func main() {
	allMembers, err := repository.GetAll()

	if err != nil {
		log.Fatalln(err)
	}

	//fmt.Printf("allMembers: %d\n\n", len(allMembers))

	membersMap := make(map[Key]Value)

	for _, member := range allMembers {
		key := Key{Id: member.ProfileId}

		if v, exists := membersMap[key]; exists {
			membersMap[key] = Value{Counter: v.Counter + 1, Member: member}
		} else {
			membersMap[key] = Value{Counter: 1, Member: member}
		}
	}

	for _, v := range membersMap {
		if v.Counter >= 5 {
			fmt.Println(v.Member.Fullname, v.Member.ProfileLink)
		}
	}
}
