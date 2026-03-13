# Go-клиент для сервиса «WB кошелек»

Минимальный Go-клиент для API WB Pay.

## Установка

```bash
go get github.com/alewon/wb-wallet-go-client.git
```

Минимальная версия Go: `1.18`.

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

## Что умеет клиент

- реализует все методы из публичной документации WB Pay
- содержит все request/response структуры
- автоматически подставляет `Authorization`
- автоматически генерирует `X-Signature` через `ED25519`
- автоматически подставляет `X-Request-Country`, `X-Request-Region` и `X-Wbpay-Id`, когда это возможно

## Лицензия

MIT. См. [LICENSE](LICENSE).
