package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	type whatIsIt string
	secret := "aWFuZ25vVzpOQU06RU5JTDp0YTpzdTpuaW9K"
	sd, err := base64.StdEncoding.DecodeString(secret)
	fmt.Println("whatIsIt = ", string(sd))
	if err != nil {
		fmt.Printf("Error decoding Base64 encoded data %v", err)
	}

}
