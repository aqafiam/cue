package commands

import (
	"fmt"
	"unidriver/errors"
)

func init() {
	CommandList["get"] = get
}

func get(_url interface{}) {

	url := _url.(string)
	fmt.Print("[get]: " + url + " ")
	err := WD.Get(url)
	errors.AssertResult(err)
}
