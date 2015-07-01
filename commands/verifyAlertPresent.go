package commands

import (
	"fmt"
	"unidriver/errors"
)

func init() {
	CommandList["verifyAlertPresent"] = verifyAlertPresent
}

func verifyAlertPresent(_url interface{}) {

	url := _url.(string)

	fmt.Print("[verifyAlertPresent]: " + url + " ")
	err := WD.Get(url)
	errors.Fatal(err)
	fmt.Print("SUCCESS")
	fmt.Println("")

}
