# wb-wallet-go-client

Минимальный Go-клиент для API WB Pay.

В репозитории:
- `client.go` — HTTP-клиент и все методы API
- `types.go` — все request/response структуры, статусы и коды ошибок
- `doc.md` — зафиксированная сводка по документации WB Pay

## Возможности

- покрывает все 24 метода из опубликованной документации API WB Pay
- содержит отдельные структуры для каждого метода без унификации
- не выполняет валидацию входных данных
- возвращает разбор ответов `200`, `400` и `403`

## Установка

```bash
go get wb-wallet-go-client
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

	result, err := client.DoOnlinePaymentByPhone(context.Background(), &wbwallet.DoOnlinePaymentByPhoneRequest{
		Body: wbwallet.DoOnlinePaymentByPhoneRequestBody{
			OrderID: "order-id",
		},
	})
	if err != nil {
		panic(err)
	}

	if result.OK != nil {
		fmt.Println(result.OK.Data.DeepLink)
	}
}
```

Если в `Client` заданы токен, ключ и региональные заголовки, клиент автоматически:
- добавляет `Authorization: Bearer ...`
- добавляет `Content-Type: application/json`
- добавляет `X-Request-Country` и `X-Request-Region`
- генерирует `X-Signature` для методов с body-подписью
- подставляет `X-Wbpay-Id`, если его можно вывести из request

## Статус проекта

Проект генерировался по публичной документации WB Pay и задуман как прозрачная тонкая обёртка над HTTP API.

Если документация WB Pay изменится, структуры и методы в этом репозитории нужно будет обновить вручную.

## Разработка

```bash
gofmt -w client.go types.go
go test ./...
```

## Лицензия

MIT. См. [LICENSE](LICENSE).
