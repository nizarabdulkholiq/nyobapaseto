package main

import (
	"fmt"

	"github.com/o1egl/paseto"
)

func main() {
    // create token
    token := paseto.JSONToken{
        Issuer:   "my_app",
        Subject:  "1234",
        Audience: "my_client_app",
    }
    key := []byte("my_secret_key")
    footer := "my_footer"
    encryptedToken, err := paseto.Encrypt(key, token, footer)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Printf("Encrypted Token: %s\n", encryptedToken)

    // verify token
    var payload paseto.JSONToken
    _, err = paseto.Decrypt(encryptedToken, key, &payload, &footer)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Printf("Payload: %v\n", payload)
}