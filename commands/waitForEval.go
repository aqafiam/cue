package commands

import (
	"fmt"
	"unidriver/errors"
)

func init() {
	CommandList["waitForEval"] = waitForEval
}

func waitForEval(_url interface{}) {

	url := _url.(string)

	fmt.Print("[waitForEval]: " + url + " ")
	err := WD.Get(url)
	errors.Fatal(err)
	fmt.Print("SUCCESS")
	fmt.Println("")

}
