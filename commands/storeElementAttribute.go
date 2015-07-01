package commands

import (
	"fmt"
	"unidriver/errors"
)

func init() {
	CommandList["storeElementAttribute"] = storeElementAttribute
}

func storeElementAttribute(_url interface{}) {

	url := _url.(string)

	fmt.Print("[storeElementAttribute]: " + url + " ")
	err := WD.Get(url)
	errors.Fatal(err)
	fmt.Print("SUCCESS")
	fmt.Println("")

}
