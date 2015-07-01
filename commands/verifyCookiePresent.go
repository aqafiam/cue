package commands

import (
	"fmt"
	"unidriver/errors"
)

func init() {
	CommandList["verifyCookiePresent"] = verifyCookiePresent
}

func verifyCookiePresent(_url interface{}) {

	url := _url.(string)

	fmt.Print("[verifyCookiePresent]: " + url + " ")
	err := WD.Get(url)
	errors.Fatal(err)
	fmt.Print("SUCCESS")
	fmt.Println("")

}
