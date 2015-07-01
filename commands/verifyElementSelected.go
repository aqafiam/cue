package commands

import (
	"fmt"
	"unidriver/errors"
)

func init() {
	CommandList["verifyElementSelected"] = verifyElementSelected
}

func verifyElementSelected(_url interface{}) {

	url := _url.(string)

	fmt.Print("[verifyElementSelected]: " + url + " ")
	err := WD.Get(url)
	errors.Fatal(err)
	fmt.Print("SUCCESS")
	fmt.Println("")

}
