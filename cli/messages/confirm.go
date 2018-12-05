package messages

import (
	"fmt"
	"log"
	"strings"
)

func AskForConfirmation(message string) bool {
	fmt.Println(message)
	var response string
	_, err := fmt.Scanln(&response)
	if err != nil {
		log.Fatal(err)
	}
	return confirmationResponse(strings.ToLower(response))
}

func confirmationResponse(response string) bool {
	okResponses := [2]string{"y", "yes"}
	for _, ok := range okResponses {
		if response == ok {
			return true
		}
	}
	return false
}