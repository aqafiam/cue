package commands

import (
	"fmt"
	"unidriver/errors"
)

func init() {
	CommandList["waitForTitle"] = waitForTitle
}

func waitForTitle(_url interface{}) {

	url := _url.(string)

	fmt.Print("[waitForTitle]: " + url + " ")
	err := WD.Get(url)
	errors.Fatal(err)
	fmt.Print("SUCCESS")
	fmt.Println("")

}
