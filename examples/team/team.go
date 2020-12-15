package main

import (
	"fmt"

	"github.com/xopenapi/uim-api-go2"
)

func main() {
	api := uim.New("YOUR_TOKEN_HERE")
	//Example for single user
	billingActive, err := api.GetBillableInfo("U023BECGF")
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	fmt.Printf("ID: U023BECGF, BillingActive: %v\n\n\n", billingActive["U023BECGF"])

	//Example for team
	billingActiveForTeam, _ := api.GetBillableInfoForTeam()
	for id, value := range billingActiveForTeam {
		fmt.Printf("ID: %v, BillingActive: %v\n", id, value)
	}

}
