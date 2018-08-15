package main

import (
	"fmt"

	"github.com/jyellick/examples/protos/common"
	"github.com/jyellick/examples/protos/dependant"
)

func main() {
	fmt.Println(&dependant.Other{
		Reply: &common.Reply{
			Status:  common.Status_SUCCESS,
			Message: "some-message",
		},
		Payload: []byte("some-payload"),
	})
}
