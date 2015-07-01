package commands

import (
	"fmt"
	"unidriver/errors"
)

func init() {
	CommandList["waitForPageSource"] = waitForPageSource
}

func waitForPageSource(_url interface{}) {

	url := _url.(string)

	fmt.Print("[waitForPageSource]: " + url + " ")
	err := WD.Get(url)
	errors.Fatal(err)
	fmt.Print("SUCCESS")
	fmt.Println("")

}
