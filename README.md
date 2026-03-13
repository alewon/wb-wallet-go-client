# Go-клиент для сервиса VB-кошелек

## Установка

```bash
go get github.com/alewon/wb-wallet-go-client.git
```

## Использование

```go
package main

import (
	"context"
	"fmt"

	wbwallet "github.com/alewon/wb-wallet-go-client.git"
)

func main() {
	privateKey, err := wbwallet.LoadEd25519PrivateKeyFromPEMFile("private.pem")
	if err != nil {
		panic(err)
	}

	client := wbwallet.NewClientWithCredentials(
		"https://api.wbpay.ru",
		nil,
		"<token>",
		privateKey,
		"ru",
		"ru-mow",
	)

	result, err := client.GeneratePayerToken(context.Background(), &wbwallet.GeneratePayerTokenRequest{
		Body: wbwallet.GeneratePayerTokenRequestBody{
			TerminalID:  "terminal-id",
			PhoneNumber: "79991234567",
			CreatedAt:   1752057656,
			ClientID:    "client-id",
		},
	})
	if err != nil {
		panic(err)
	}

	if result.OK != nil {
		fmt.Println(result.OK.Data.RegistrationID)
	}
}
```
