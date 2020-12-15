package main

import (
	"fmt"

	"github.com/xopenapi/uim-api-go2"
)

func main() {
	api := uim.New("YOUR_TOKEN_HERE")
	user, err := api.GetUserInfo("U023BECGF")
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	fmt.Printf("ID: %s, Fullname: %s, Email: %s\n", user.ID, user.Profile.RealName, user.Profile.Email)
}
