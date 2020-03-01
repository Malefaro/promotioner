package main

import (
	"fmt"
	"promotioner/cmd/promotionerAPI/models/enums"
)

func test() {
	e := enums.PromotionTypeEnum("")
	err := e.Scan([]byte("asdasd"))
	fmt.Println(err, e)
}
func main() {
	// todo: connection to authService
	// todo: run API server
	test()
}
