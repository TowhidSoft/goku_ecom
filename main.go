package main

import "ecommerce/cmd"

func main() {
	cmd.Serve()

	// jwt, err := util.CreateJwt("mysecret", util.Payload{
	// 	Sub:         45,
	// 	FirstName:   "John",
	// 	LastName:    "Doe",
	// 	Email:       "john.doe@example.com",
	// 	IsShopOwner: true,
	// })
	// if err != nil {
	// 	panic(err)
	// }
	// println(jwt)
}
