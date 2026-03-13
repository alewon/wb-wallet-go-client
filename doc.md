# Методы WB Pay API

Источник: [getting-started](https://docs.wbpay.ru/getting-started.html) и раздел [Документация API](https://docs.wbpay.ru/api.html) на `docs.wbpay.ru`.

Ниже для каждого метода приведено полное содержимое страницы документации: запросы, ответы, таблицы полей, enum-значения, примеры body и вложенные схемы без унификации.

## API Онлайн-оплаты / Платёжный токен плательщика / Генерация токена плательщика

Источник метода: <https://docs.wbpay.ru/api/online/Platyozhnyj-token-platelshika/apiv1userstokens-post.html>

# Генерация токена плательщика

Метод возвращает ID операции регистрации в статусе `pending`

## Request

POST

    
        ```
https://api.wbpay.ru/api/v1/users/tokens
```

        
    

Production

POST

    
        ```
https://sandbox.wbpay.ru/api/v1/users/tokens
```

        
    

Sandbox

### Headers

| **Name** | **Description** |
| --- | --- |
| Authorization* | **Type:** string<br><br>Токен авторизации<br>Example: `Bearer {access_token}` |
| Content-Type* | **Type:** string<br><br>Тип контента<br><br>*Enum:* `application/json` |
| X-Request-Country* | **Type:** string<br><br>Код страны сервера, октуда поступает запрос, в формате ISO 3166-1<br><br>*Pattern:* `^[a-zA-Z]{2}$` |
| X-Request-Region* | **Type:** string<br><br>Код региона сервера, октуда поступает запрос, в формате ISO 3166-2<br><br>*Pattern:* `^[a-zA-Z]{2}-[a-zA-Z]{2,3}$` |
| X-Signature* | **Type:** string<br><br>Подпись запроса на основе алгоритма ED25519 |

### Body

**application/json**

    
        ```
{
"terminal_id": "string",
"phone_number": "string",
"created_at": 0,
"client_id": "string"
}
```

        
    
| **Name** | **Description** |
| --- | --- |
| client_id* | **Type:** string<br><br>Уникальный ID пользователя в системе продавца<br><br>*Max length:* `2048` |
| created_at* | **Type:** integer<int64><br><br>Дата создания заказа в формате Unix timestamp |
| phone_number* | **Type:** string<br><br>Номер телефона плательщика (пример 71234567890) |
| terminal_id* | **Type:** string<br><br>ID терминала |

## Responses

## 200 OK

Успешно

### Body

**application/json**

    
        ```
{
"error_code": "ERR_NONE",
"error_description": "string",
"data": {
"registration_id": "string",
"deep_link": "wildberries://wbpay/agreement?confirmation_id=6f3bf8a8e2454876be478e3034316916_1765141200"
}
}
```

        
    
| **Name** | **Description** |
| --- | --- |
| data* | **Type:** [TokenGenerationData](api/online/Platyozhnyj-token-platelshika/apiv1userstokens-post.html#tokengenerationdata)<br><br>Данные ответа |
| error_code | **Type:** string<br><br>*Enum:* `ERR_NONE` |
| error_description | **Type:** string<br><br>Описание кода ответа |

### TokenGenerationData

Данные ответа

| **Name** | **Description** |
| --- | --- |
| registration_id* | **Type:** string<br><br>ID запроса для создания платёжного токена |
| deep_link | **Type:** string<br><br>Диплинк на экран подтверждения создания платёжного токена в приложении ВБ<br><br>*Example:* `wildberries://wbpay/agreement?confirmation_id=6f3bf8a8e2454876be478e3034316916_1765141200` |

## 400 Bad Request

Неправильный запрос

### Body

**application/json**

    
        ```
{
"error_code": "INVALID_REQUEST_ERROR",
"error_description": "string"
}
```

        
    
| **Name** | **Description** |
| --- | --- |
| error_code* | **Type:** string<br><br>*Enum:* `INVALID_REQUEST_ERROR`, `INVALID_PHONE_NUMBER` |
| error_description* | **Type:** string<br><br>Описание кода ответа |

## 403 Forbidden

Доступ запрещён

## API Онлайн-оплаты / Платёжный токен плательщика / Получение статуса генерации токена плательщика

Источник метода: <https://docs.wbpay.ru/api/online/Platyozhnyj-token-platelshika/apiv1userstokensregistration_idstatus-get.html>

# Получение статуса генерации токена плательщика

Метод возвращает статус операции регистрации

## Request

GET

    
        ```
https://api.wbpay.ru/api/v1/users/tokens/{registration_id}/status
```

        
    

Production

GET

    
        ```
https://sandbox.wbpay.ru/api/v1/users/tokens/{registration_id}/status
```

        
    

Sandbox

### Path parameters

| **Name** | **Description** |
| --- | --- |
| registration_id* | **Type:** string |

### Headers

| **Name** | **Description** |
| --- | --- |
| Authorization* | **Type:** string<br><br>Токен авторизации<br>Example: `Bearer {access_token}` |
| Content-Type* | **Type:** string<br><br>Тип контента<br><br>*Enum:* `application/json` |
| X-Wbpay-Id* | **Type:** string<br><br>ID запроса. Используется значение параметра `registration_id`, полученного ранее в методе генерации токена. Параметр необходим для осуществления корректной маршрутизации запроса |

## Responses

## 200 OK

Успешно

### Body

**application/json**

    
        ```
{
"error_code": "ERR_NONE",
"error_description": "string",
"data": {
"status": "pending",
"token": "string",
"fail_reason_code": "USER_NOT_APPROVE",
"fail_reason_description": "string"
}
}
```

        
    
| **Name** | **Description** |
| --- | --- |
| data* | **Type:** [TokenStatusData](api/online/Platyozhnyj-token-platelshika/apiv1userstokensregistration_idstatus-get.html#tokenstatusdata)<br><br>Данные ответа |
| error_code | **Type:** string<br><br>*Enum:* `ERR_NONE` |
| error_description | **Type:** string<br><br>Описание кода ответа |

### TokenStatusData

Данные ответа

| **Name** | **Description** |
| --- | --- |
| status* | **Type:** string<br><br>Статус операции.<br><br>- `pending` - Промежуточный статус. Операция в обработке, повторите запрос статуса позднее<br>- `failed` — Финальный статус. Не удалось обработать операцию<br>- `succeeded` — Финальный статус. Операция успешно обработана<br><br>*Enum:* `pending`, `failed`, `succeeded` |
| fail_reason_code | **Type:** string<br><br>Причина неуспешности операции.<br><br>- `USER_NOT_APPROVE` - пользователь отклонил запрос согласия<br>- `REQUEST_EXPIRED` - истек срок ожидания согласия пользователя<br>- `INTERNAL_SERVER_ERROR` - внутренняя ошибка<br><br>*Enum:* `USER_NOT_APPROVE`, `REQUEST_EXPIRED`, `INTERNAL_SERVER_ERROR` |
| fail_reason_description | **Type:** string |
| token | **Type:** string<br><br>Платёжный токен плательщика |

## 400 Bad Request

Неправильный запрос

### Body

**application/json**

    
        ```
{
"error_code": "INVALID_REQUEST_ERROR",
"error_description": "string"
}
```

        
    
| **Name** | **Description** |
| --- | --- |
| error_code* | **Type:** string<br><br>*Enum:* `INVALID_REQUEST_ERROR`, `NOT_FOUND` |
| error_description* | **Type:** string<br><br>Описание кода ответа |

## 403 Forbidden

Доступ запрещён

## API Онлайн-оплаты / Оплата по платёжному токену / Регистрация онлайн-оплаты по токену

Источник метода: <https://docs.wbpay.ru/api/online/Oplata-po-platyozhnomu-tokenu/apiv1ordersonlineregister-post.html>

# Регистрация онлайн-оплаты по токену

Метод возвращает ID заказа в статусе `created`

## Request

POST

    
        ```
https://api.wbpay.ru/api/v1/orders/online/register
```

        
    

Production

POST

    
        ```
https://sandbox.wbpay.ru/api/v1/orders/online/register
```

        
    

Sandbox

### Headers

| **Name** | **Description** |
| --- | --- |
| Authorization* | **Type:** string<br><br>Токен авторизации<br>Example: `Bearer {access_token}` |
| Content-Type* | **Type:** string<br><br>Тип контента<br><br>*Enum:* `application/json` |
| X-Request-Country* | **Type:** string<br><br>Код страны сервера, октуда поступает запрос, в формате ISO 3166-1<br><br>*Pattern:* `^[a-zA-Z]{2}$` |
| X-Request-Region* | **Type:** string<br><br>Код региона сервера, октуда поступает запрос, в формате ISO 3166-2<br><br>*Pattern:* `^[a-zA-Z]{2}-[a-zA-Z]{2,3}$` |

### Body

**application/json**

    
        ```
{
"terminal_id": "string",
"invoice_id": "string",
"token": "string",
"amount": 0,
"currency_code": 0,
"created_at": 0,
"client_id": "string",
"redirect_url": "wildberries://wbpay/agreement, https://wildberries.ru",
"positions": [
{
"name": "string",
"price": 0,
"count": 0
}
]
}
```

        
    
| **Name** | **Description** |
| --- | --- |
| amount* | **Type:** integer<int64><br><br>Сумма в копейках, без разделителя. `50 = "5000"` |
| client_id* | **Type:** string<br><br>Уникальный ID пользователя в системе продавца, должен совпадать с client_id, указанным при регистрации токена<br><br>*Max length:* `2048` |
| created_at* | **Type:** integer<int64><br><br>Дата создания заказа в формате Unix timestamp |
| currency_code* | **Type:** integer<int64><br><br>Код валюты. Для оплаты в рублях передавать значение `643` |
| terminal_id* | **Type:** string<br><br>ID терминала |
| token* | **Type:** string<br><br>Платёжный токен плательщика |
| invoice_id | **Type:** string<br><br>ID операции на стороне продавца (должен быть уникальным). Данный идентификатор будет отображаться в отчетах<br><br>*Max length:* `2048` |
| positions | **Type:** [Position](api/online/Oplata-po-platyozhnomu-tokenu/apiv1ordersonlineregister-post.html#position)[]<br><br>Корзина товаров<br><br>*Max items:* `1000` |
| redirect_url | **Type:** string<uri><br><br>Специальная ссылка, которая позволяет направить пользователя в конкретное место внутри мобильного приложения или на веб-страницу<br><br>*Example:* `wildberries://wbpay/agreement, https://wildberries.ru` |

### Position

| **Name** | **Description** |
| --- | --- |
| count | **Type:** integer<int64><br><br>Количество единиц товара |
| name | **Type:** string<br><br>Название товара<br><br>*Max length:* `2048` |
| price | **Type:** integer<int64><br><br>Стоимость единицы товара в копейках, без разделителя. `50 = "5000"` |

## Responses

## 200 OK

Успешно

### Body

**application/json**

    
        ```
{
"error_code": "ERR_NONE",
"error_description": "string",
"data": {
"order_id": "string"
}
}
```

        
    
| **Name** | **Description** |
| --- | --- |
| data* | **Type:** [OrderRegistrationData](api/online/Oplata-po-platyozhnomu-tokenu/apiv1ordersonlineregister-post.html#orderregistrationdata)<br><br>Данные ответа |
| error_code | **Type:** string<br><br>*Enum:* `ERR_NONE` |
| error_description | **Type:** string<br><br>Описание кода ответа |

### OrderRegistrationData

Данные ответа

| **Name** | **Description** |
| --- | --- |
| order_id* | **Type:** string<br><br>ID заказа |

## 400 Bad Request

Неправильный запрос

### Body

**application/json**

    
        ```
{
"error_code": "INVALID_PAYMENT_TOKEN",
"error_description": "string"
}
```

        
    
| **Name** | **Description** |
| --- | --- |
| error_code* | **Type:** string<br><br>*Enum:* `INVALID_PAYMENT_TOKEN`, `INVALID_REQUEST_ERROR` |
| error_description* | **Type:** string<br><br>Описание кода ответа |

## 403 Forbidden

Доступ запрещён

## API Онлайн-оплаты / Оплата по платёжному токену / Выполнение оплаты

Источник метода: <https://docs.wbpay.ru/api/online/Oplata-po-platyozhnomu-tokenu/apiv1ordersdo-post.html>

# Выполнение оплаты

Метод применяется для финализации оплаты

## Request

POST

    
        ```
https://api.wbpay.ru/api/v1/orders/do
```

        
    

Production

POST

    
        ```
https://sandbox.wbpay.ru/api/v1/orders/do
```

        
    

Sandbox

### Headers

| **Name** | **Description** |
| --- | --- |
| Authorization* | **Type:** string<br><br>Токен авторизации<br>Example: `Bearer {access_token}` |
| Content-Type* | **Type:** string<br><br>Тип контента<br><br>*Enum:* `application/json` |
| X-Wbpay-Id* | **Type:** string<br><br>ID запроса. Используется значение параметра `order_id`, полученного ранее в методе регистрации оплаты. Параметр необходим для осуществления корректной маршрутизации запроса |

### Body

**application/json**

    
        ```
{
"order_id": "alskdjfhalksjdhf_84765384765"
}
```

        
    
| **Name** | **Description** |
| --- | --- |
| order_id* | **Type:** string<br><br>ID заказа. Используется значение параметра `order_id`, полученного ранее в методе регистрации оплаты |

## Responses

## 200 OK

Успешно

### Body

**application/json**

    
        ```
{
"error_code": "ERR_NONE",
"error_description": "string",
"data": {
"deep_link": "wildberries://wbpay/agreement?confirmation_id=6f3bf8a8e2454876be478e3034316916_1765141200"
}
}
```

        
    
| **Name** | **Description** |
| --- | --- |
| data | **Type:** [DoOrderData](api/online/Oplata-po-platyozhnomu-tokenu/apiv1ordersdo-post.html#doorderdata)<br><br>Данные ответа |
| error_code | **Type:** string<br><br>*Enum:* `ERR_NONE` |
| error_description | **Type:** string<br><br>Описание кода ответа |

### DoOrderData

Данные ответа

| **Name** | **Description** |
| --- | --- |
| deep_link | **Type:** string<br><br>Диплинк на экран оплаты в приложении ВБ<br><br>*Example:* `wildberries://wbpay/agreement?confirmation_id=6f3bf8a8e2454876be478e3034316916_1765141200` |

## 400 Bad Request

Неправильный запрос

### Body

**application/json**

    
        ```
{
"error_code": "INVALID_REQUEST_ERROR",
"error_description": "string"
}
```

        
    
| **Name** | **Description** |
| --- | --- |
| error_code* | **Type:** string<br><br>*Enum:* `INVALID_REQUEST_ERROR`, `NOT_FOUND` |
| error_description* | **Type:** string<br><br>Описание кода ответа |

## 403 Forbidden

Доступ запрещён

## API Онлайн-оплаты / Оплата по платёжному токену / Получение статуса оплаты

Источник метода: <https://docs.wbpay.ru/api/online/Oplata-po-platyozhnomu-tokenu/apiv1ordersorder_idstatus-get.html>

# Получение статуса оплаты

Метод возвращает статус оплаты заказа

## Request

GET

    
        ```
https://api.wbpay.ru/api/v1/orders/{order_id}/status
```

        
    

Production

GET

    
        ```
https://sandbox.wbpay.ru/api/v1/orders/{order_id}/status
```

        
    

Sandbox

### Path parameters

| **Name** | **Description** |
| --- | --- |
| order_id* | **Type:** string<br><br>ID заказа |

### Headers

| **Name** | **Description** |
| --- | --- |
| Authorization* | **Type:** string<br><br>Токен авторизации<br>Example: `Bearer {access_token}` |
| Content-Type* | **Type:** string<br><br>Тип контента<br><br>*Enum:* `application/json` |
| X-Wbpay-Id* | **Type:** string<br><br>ID запроса. Используется значение параметра `order_id`, полученного ранее в методе регистрации оплаты. Параметр необходим для осуществления корректной маршрутизации запроса |

## Responses

## 200 OK

Успешно

### Body

**application/json**

    
        ```
{
"error_code": "ERR_NONE",
"error_description": "Запрос выполнен успешно",
"data": {
"status": "succeeded"
}
}
```

        
    
| **Name** | **Description** |
| --- | --- |
| data* | **Type:** [OrderStatusData](api/online/Oplata-po-platyozhnomu-tokenu/apiv1ordersorder_idstatus-get.html#orderstatusdata)<br><br>Данные ответа |
| error_code | **Type:** string<br><br>*Enum:* `ERR_NONE` |
| error_description | **Type:** string<br><br>Описание кода ответа |

### OrderStatusData

Данные ответа

| **Name** | **Description** |
| --- | --- |
| status* | **Type:** string<br><br>Статус заказа.<br><br>- `created` — Промежуточный статус. Операция создана, но обработка еще не началась<br>- `pending` — Промежуточный статус. Операция в обработке, повторите запрос статуса позднее<br>- `failed` — Финальный статус. Не удалось обработать операцию<br>- `succeeded` — Финальный статус. Операция успешно обработана<br><br>*Enum:* `created`, `pending`, `failed`, `succeeded` |
| fail_reason_code | **Type:** string<br><br>Код причины неуспешной обработки операции (присутствует при статусе `failed`):<br><br>- `UNABLE_TO_PROCESS` — Невозможно провести операцию.<br>- `NOT_ENOUGH_MONEY` — Недостаточно средств. Привяжите другую карту и попробуйте снова.<br>- `LIMIT_EXCEEDED` — Сумма оплаты превышает допустимый лимит.<br>- `NO_AVAILABLE_PAYMENT_METHODS` — Невозможно провести списание. Привяжите другую карту и попробуйте снова.<br>- `ORDER_EXPIRED` — Истек срок действия заказа.<br>- `CONFIRMATION_TIME_EXPIRED` — Время на подтверждение оплаты истекло.<br>- `CONFIRMATION_REJECTED` — Подтверждение было отклонено.<br>- `SYSTEM_ERROR` — Невозможно провести операцию.<br><br>*Enum:* `UNABLE_TO_PROCESS`, `NOT_ENOUGH_MONEY`, `LIMIT_EXCEEDED`, `NO_AVAILABLE_PAYMENT_METHODS`, `ORDER_EXPIRED`, `CONFIRMATION_TIME_EXPIRED`, `CONFIRMATION_REJECTED`, `SYSTEM_ERROR` |
| fail_reason_description | **Type:** string<br><br>Описание причины неуспешной обработки операции (присутствует при статусе `failed`) |
| token | **Type:** string<br><br>Платёжный токен плательщика<br><br>*Example:* `asjhdgfkajshgdkfjahgsd` |

## 400 Bad Request

Неправильный запрос

### Body

**application/json**

    
        ```
{
"error_code": "INVALID_REQUEST_ERROR",
"error_description": "string"
}
```

        
    
| **Name** | **Description** |
| --- | --- |
| error_code* | **Type:** string<br><br>*Enum:* `INVALID_REQUEST_ERROR`, `NOT_FOUND` |
| error_description* | **Type:** string<br><br>Описание кода ответа |

## 403 Forbidden

Доступ запрещён

## API Онлайн-оплаты / Оплата по номеру телефона / Регистрация оплаты по номеру телефона

Источник метода: <https://docs.wbpay.ru/api/online/Oplata-po-nomeru-telefona/apiv1ordersonlineregister-by-phone-post.html>

# Регистрация оплаты по номеру телефона

Метод возвращает ID заказа в статусе `created`

## Request

POST

    
        ```
https://api.wbpay.ru/api/v1/orders/online/register-by-phone
```

        
    

Production

POST

    
        ```
https://sandbox.wbpay.ru/api/v1/orders/online/register-by-phone
```

        
    

Sandbox

### Headers

| **Name** | **Description** |
| --- | --- |
| Authorization* | **Type:** string<br><br>Токен авторизации<br>Example: `Bearer {access_token}` |
| Content-Type* | **Type:** string<br><br>Тип контента<br><br>*Enum:* `application/json` |
| X-Request-Country* | **Type:** string<br><br>Код страны сервера, октуда поступает запрос, в формате ISO 3166-1<br><br>*Pattern:* `^[a-zA-Z]{2}$` |
| X-Request-Region* | **Type:** string<br><br>Код региона сервера, октуда поступает запрос, в формате ISO 3166-2<br><br>*Pattern:* `^[a-zA-Z]{2}-[a-zA-Z]{2,3}$` |
| X-Signature* | **Type:** string<br><br>Подпись запроса на основе алгоритма ED25519 |

### Body

**application/json**

    
        ```
{
"terminal_id": "string",
"invoice_id": "string",
"phone_number": "string",
"redirect_url": "wildberries://wbpay/agreement, https://wildberries.ru",
"amount": 0,
"currency_code": 0,
"created_at": 0,
"positions": [
{
"name": "string",
"price": 0,
"count": 0
}
]
}
```

        
    
| **Name** | **Description** |
| --- | --- |
| amount* | **Type:** integer<int64><br><br>Сумма в копейках, без разделителя. `50 = "5000"` |
| created_at* | **Type:** integer<int64><br><br>Дата создания заказа в формате Unix timestamp |
| currency_code* | **Type:** integer<int64><br><br>Код валюты. Для оплаты в рублях передавать значение `643` |
| phone_number* | **Type:** string<br><br>Номер телефона плательщика (пример 71234567890) |
| terminal_id* | **Type:** string<br><br>ID терминала |
| invoice_id | **Type:** string<br><br>ID операции на стороне продавца (должен быть уникальным). Данный идентификатор будет отображаться в отчетах<br><br>*Max length:* `2048` |
| positions | **Type:** [Position](api/online/Oplata-po-nomeru-telefona/apiv1ordersonlineregister-by-phone-post.html#position)[]<br><br>Корзина товаров<br><br>*Max length:* `1000` |
| redirect_url | **Type:** string<uri><br><br>Специальная ссылка, которая позволяет направить пользователя в конкретное место внутри мобильного приложения или на веб-страницу<br><br>*Example:* `wildberries://wbpay/agreement, https://wildberries.ru` |

### Position

| **Name** | **Description** |
| --- | --- |
| count | **Type:** integer<int64><br><br>Количество единиц товара |
| name | **Type:** string<br><br>Название товара<br><br>*Max length:* `2048` |
| price | **Type:** integer<int64><br><br>Стоимость единицы товара в копейках, без разделителя. `50 = "5000"` |

## Responses

## 200 OK

Успешно

### Body

**application/json**

    
        ```
{
"error_code": "ERR_NONE",
"error_description": "string",
"data": {
"order_id": "string"
}
}
```

        
    
| **Name** | **Description** |
| --- | --- |
| data* | **Type:** [OrderRegistrationData](api/online/Oplata-po-nomeru-telefona/apiv1ordersonlineregister-by-phone-post.html#orderregistrationdata)<br><br>Данные ответа |
| error_code | **Type:** string<br><br>*Enum:* `ERR_NONE` |
| error_description | **Type:** string<br><br>Описание кода ответа |

### OrderRegistrationData

Данные ответа

| **Name** | **Description** |
| --- | --- |
| order_id* | **Type:** string<br><br>ID заказа |

## 400 Bad Request

Неправильный запрос

### Body

**application/json**

    
        ```
{
"error_code": "INVALID_REQUEST_ERROR",
"error_description": "string"
}
```

        
    
| **Name** | **Description** |
| --- | --- |
| error_code* | **Type:** string<br><br>*Enum:* `INVALID_REQUEST_ERROR`, `INVALID_PHONE_NUMBER` |
| error_description* | **Type:** string<br><br>Описание кода ответа |

## 403 Forbidden

Доступ запрещён

## API Онлайн-оплаты / Оплата по номеру телефона / Выполнение оплаты

Источник метода: <https://docs.wbpay.ru/api/online/Oplata-po-nomeru-telefona/apiv1ordersdo-post.html>

# Выполнение оплаты

Метод применяется для финализации оплаты

## Request

POST

    
        ```
https://api.wbpay.ru/api/v1/orders/do
```

        
    

Production

POST

    
        ```
https://sandbox.wbpay.ru/api/v1/orders/do
```

        
    

Sandbox

### Headers

| **Name** | **Description** |
| --- | --- |
| Authorization* | **Type:** string<br><br>Токен авторизации<br>Example: `Bearer {access_token}` |
| Content-Type* | **Type:** string<br><br>Тип контента<br><br>*Enum:* `application/json` |
| X-Wbpay-Id* | **Type:** string<br><br>ID запроса. Используется значение параметра `order_id`, полученного ранее в методе регистрации оплаты. Параметр необходим для осуществления корректной маршрутизации запроса |

### Body

**application/json**

    
        ```
{
"order_id": "alskdjfhalksjdhf_84765384765"
}
```

        
    
| **Name** | **Description** |
| --- | --- |
| order_id* | **Type:** string<br><br>ID заказа. Используется значение параметра `order_id`, полученного ранее в методе регистрации оплаты |

## Responses

## 200 OK

Успешно

### Body

**application/json**

    
        ```
{
"error_code": "ERR_NONE",
"error_description": "string",
"data": {
"deep_link": "wildberries://wbpay/agreement?confirmation_id=6f3bf8a8e2454876be478e3034316916_1765141200"
}
}
```

        
    
| **Name** | **Description** |
| --- | --- |
| data | **Type:** [DoOrderData](api/online/Oplata-po-nomeru-telefona/apiv1ordersdo-post.html#doorderdata)<br><br>Данные ответа |
| error_code | **Type:** string<br><br>*Enum:* `ERR_NONE` |
| error_description | **Type:** string<br><br>Описание кода ответа |

### DoOrderData

Данные ответа

| **Name** | **Description** |
| --- | --- |
| deep_link | **Type:** string<br><br>Диплинк на экран оплаты в приложении ВБ<br><br>*Example:* `wildberries://wbpay/agreement?confirmation_id=6f3bf8a8e2454876be478e3034316916_1765141200` |

## 400 Bad Request

Неправильный запрос

### Body

**application/json**

    
        ```
{
"error_code": "INVALID_REQUEST_ERROR",
"error_description": "string"
}
```

        
    
| **Name** | **Description** |
| --- | --- |
| error_code* | **Type:** string<br><br>*Enum:* `INVALID_REQUEST_ERROR`, `NOT_FOUND` |
| error_description* | **Type:** string<br><br>Описание кода ответа |

## 403 Forbidden

Доступ запрещён

## API Онлайн-оплаты / Оплата по номеру телефона / Получение статуса оплаты

Источник метода: <https://docs.wbpay.ru/api/online/Oplata-po-nomeru-telefona/apiv1ordersorder_idstatus-get.html>

# Получение статуса оплаты

Метод возвращает статус оплаты заказа

## Request

GET

    
        ```
https://api.wbpay.ru/api/v1/orders/{order_id}/status
```

        
    

Production

GET

    
        ```
https://sandbox.wbpay.ru/api/v1/orders/{order_id}/status
```

        
    

Sandbox

### Path parameters

| **Name** | **Description** |
| --- | --- |
| order_id* | **Type:** string<br><br>ID заказа |

### Headers

| **Name** | **Description** |
| --- | --- |
| Authorization* | **Type:** string<br><br>Токен авторизации<br>Example: `Bearer {access_token}` |
| Content-Type* | **Type:** string<br><br>Тип контента<br><br>*Enum:* `application/json` |
| X-Wbpay-Id* | **Type:** string<br><br>ID запроса. Используется значение параметра `order_id`, полученного ранее в методе регистрации оплаты. Параметр необходим для осуществления корректной маршрутизации запроса |

## Responses

## 200 OK

Успешно

### Body

**application/json**

    
        ```
{
"error_code": "ERR_NONE",
"error_description": "Запрос выполнен успешно",
"data": {
"status": "succeeded"
}
}
```

        
    
| **Name** | **Description** |
| --- | --- |
| data* | **Type:** [OrderStatusData](api/online/Oplata-po-nomeru-telefona/apiv1ordersorder_idstatus-get.html#orderstatusdata)<br><br>Данные ответа |
| error_code | **Type:** string<br><br>*Enum:* `ERR_NONE` |
| error_description | **Type:** string<br><br>Описание кода ответа |

### OrderStatusData

Данные ответа

| **Name** | **Description** |
| --- | --- |
| status* | **Type:** string<br><br>Статус заказа.<br><br>- `created` — Промежуточный статус. Операция создана, но обработка еще не началась<br>- `pending` — Промежуточный статус. Операция в обработке, повторите запрос статуса позднее<br>- `failed` — Финальный статус. Не удалось обработать операцию<br>- `succeeded` — Финальный статус. Операция успешно обработана<br><br>*Enum:* `created`, `pending`, `failed`, `succeeded` |
| fail_reason_code | **Type:** string<br><br>Код причины неуспешной обработки операции (присутствует при статусе `failed`):<br><br>- `UNABLE_TO_PROCESS` — Невозможно провести операцию.<br>- `NOT_ENOUGH_MONEY` — Недостаточно средств. Привяжите другую карту и попробуйте снова.<br>- `LIMIT_EXCEEDED` — Сумма оплаты превышает допустимый лимит.<br>- `NO_AVAILABLE_PAYMENT_METHODS` — Невозможно провести списание. Привяжите другую карту и попробуйте снова.<br>- `ORDER_EXPIRED` — Истек срок действия заказа.<br>- `CONFIRMATION_TIME_EXPIRED` — Время на подтверждение оплаты истекло.<br>- `CONFIRMATION_REJECTED` — Подтверждение было отклонено.<br>- `SYSTEM_ERROR` — Невозможно провести операцию.<br><br>*Enum:* `UNABLE_TO_PROCESS`, `NOT_ENOUGH_MONEY`, `LIMIT_EXCEEDED`, `NO_AVAILABLE_PAYMENT_METHODS`, `ORDER_EXPIRED`, `CONFIRMATION_TIME_EXPIRED`, `CONFIRMATION_REJECTED`, `SYSTEM_ERROR` |
| fail_reason_description | **Type:** string<br><br>Описание причины неуспешной обработки операции (присутствует при статусе `failed`) |
| token | **Type:** string<br><br>Платёжный токен плательщика<br><br>*Example:* `asjhdgfkajshgdkfjahgsd` |

## 400 Bad Request

Неправильный запрос

### Body

**application/json**

    
        ```
{
"error_code": "INVALID_REQUEST_ERROR",
"error_description": "string"
}
```

        
    
| **Name** | **Description** |
| --- | --- |
| error_code* | **Type:** string<br><br>*Enum:* `INVALID_REQUEST_ERROR`, `NOT_FOUND` |
| error_description* | **Type:** string<br><br>Описание кода ответа |

## 403 Forbidden

Доступ запрещён

## API Онлайн-оплаты / Оплата по платёжной ссылке / Регистрация оплаты по платёжной ссылке

Источник метода: <https://docs.wbpay.ru/api/online/Oplata-po-platyozhnoj-ssylke/apiv1ordershppregister-post.html>

# Регистрация оплаты по платёжной ссылке

Метод возвращает ID заказа и ссылку на платежную страницу

## Request

POST

    
        ```
https://api.wbpay.ru/api/v1/orders/hpp/register
```

        
    

Production

POST

    
        ```
https://sandbox.wbpay.ru/api/v1/orders/hpp/register
```

        
    

Sandbox

### Headers

| **Name** | **Description** |
| --- | --- |
| Authorization* | **Type:** string<br><br>Токен авторизации<br>Example: `Bearer {access_token}` |
| Content-Type* | **Type:** string<br><br>Тип контента<br><br>*Enum:* `application/json` |
| X-Request-Country* | **Type:** string<br><br>Код страны сервера, октуда поступает запрос, в формате ISO 3166-1<br><br>*Pattern:* `^[a-zA-Z]{2}$` |
| X-Request-Region* | **Type:** string<br><br>Код региона сервера, октуда поступает запрос, в формате ISO 3166-2<br><br>*Pattern:* `^[a-zA-Z]{2}-[a-zA-Z]{2,3}$` |
| X-Signature* | **Type:** string<br><br>Подпись запроса на основе алгоритма ED25519 |

### Body

**application/json**

    
        ```
{
"terminal_id": "string",
"invoice_id": "string",
"phone_number": "string",
"redirect_url": "wildberries://wbpay/agreement, https://wildberries.ru",
"amount": 0,
"currency_code": 0,
"created_at": 0,
"positions": [
{
"name": "string",
"price": 0,
"count": 0
}
]
}
```

        
    
| **Name** | **Description** |
| --- | --- |
| amount* | **Type:** integer<int64><br><br>Сумма в копейках, без разделителя. `50 = "5000"` |
| created_at* | **Type:** integer<int64><br><br>Дата создания заказа в формате Unix timestamp |
| currency_code* | **Type:** integer<int64><br><br>Код валюты. Для оплаты в рублях передавать значение `643` |
| phone_number* | **Type:** string<br><br>Номер телефона плательщика (пример 71234567890) |
| redirect_url* | **Type:** string<uri><br><br>Специальная ссылка, которая позволяет направить пользователя в конкретное место внутри мобильного приложения или на веб-страницу<br><br>*Example:* `wildberries://wbpay/agreement, https://wildberries.ru` |
| terminal_id* | **Type:** string<br><br>ID терминала |
| invoice_id | **Type:** string<br><br>ID операции на стороне продавца (должен быть уникальным). Данный идентификатор будет отображаться в отчетах<br><br>*Max length:* `2048` |
| positions | **Type:** [Position](api/online/Oplata-po-platyozhnoj-ssylke/apiv1ordershppregister-post.html#position)[]<br><br>Корзина товаров<br><br>*Max length:* `1000` |

### Position

| **Name** | **Description** |
| --- | --- |
| count | **Type:** integer<int64><br><br>Количество единиц товара |
| name | **Type:** string<br><br>Название товара<br><br>*Max length:* `2048` |
| price | **Type:** integer<int64><br><br>Стоимость единицы товара в копейках, без разделителя. `50 = "5000"` |

## Responses

## 200 OK

Успешно

### Body

**application/json**

    
        ```
{
"error_code": "ERR_NONE",
"error_description": "string",
"data": {
"order_id": "string",
"payment_url": "string"
}
}
```

        
    
| **Name** | **Description** |
| --- | --- |
| data* | **Type:** [HPPOrderRegistrationData](api/online/Oplata-po-platyozhnoj-ssylke/apiv1ordershppregister-post.html#hpporderregistrationdata)<br><br>Данные ответа |
| error_code | **Type:** string<br><br>*Enum:* `ERR_NONE` |
| error_description | **Type:** string<br><br>Описание кода ответа |

### HPPOrderRegistrationData

Данные ответа

| **Name** | **Description** |
| --- | --- |
| order_id* | **Type:** string<br><br>ID заказа |
| payment_url* | **Type:** string<br><br>Платёжная ссылка |

## 400 Bad Request

Неправильный запрос

### Body

**application/json**

    
        ```
{
"error_code": "INVALID_REQUEST_ERROR",
"error_description": "string"
}
```

        
    
| **Name** | **Description** |
| --- | --- |
| error_code* | **Type:** string<br><br>*Enum:* `INVALID_REQUEST_ERROR`, `INVALID_PHONE_NUMBER` |
| error_description* | **Type:** string<br><br>Описание кода ответа |

## 403 Forbidden

Доступ запрещён

## API Онлайн-оплаты / Оплата по платёжной ссылке / Получение статуса оплаты

Источник метода: <https://docs.wbpay.ru/api/online/Oplata-po-platyozhnoj-ssylke/apiv1ordersorder_idstatus-get.html>

# Получение статуса оплаты

Метод возвращает статус оплаты заказа

## Request

GET

    
        ```
https://api.wbpay.ru/api/v1/orders/{order_id}/status
```

        
    

Production

GET

    
        ```
https://sandbox.wbpay.ru/api/v1/orders/{order_id}/status
```

        
    

Sandbox

### Path parameters

| **Name** | **Description** |
| --- | --- |
| order_id* | **Type:** string<br><br>ID заказа |

### Headers

| **Name** | **Description** |
| --- | --- |
| Authorization* | **Type:** string<br><br>Токен авторизации<br>Example: `Bearer {access_token}` |
| Content-Type* | **Type:** string<br><br>Тип контента<br><br>*Enum:* `application/json` |
| X-Wbpay-Id* | **Type:** string<br><br>ID запроса. Используется значение параметра `order_id`, полученного ранее в методе регистрации оплаты. Параметр необходим для осуществления корректной маршрутизации запроса |

## Responses

## 200 OK

Успешно

### Body

**application/json**

    
        ```
{
"error_code": "ERR_NONE",
"error_description": "Запрос выполнен успешно",
"data": {
"status": "succeeded"
}
}
```

        
    
| **Name** | **Description** |
| --- | --- |
| data* | **Type:** [OrderStatusData](api/online/Oplata-po-platyozhnoj-ssylke/apiv1ordersorder_idstatus-get.html#orderstatusdata)<br><br>Данные ответа |
| error_code | **Type:** string<br><br>*Enum:* `ERR_NONE` |
| error_description | **Type:** string<br><br>Описание кода ответа |

### OrderStatusData

Данные ответа

| **Name** | **Description** |
| --- | --- |
| status* | **Type:** string<br><br>Статус заказа.<br><br>- `created` — Промежуточный статус. Операция создана, но обработка еще не началась<br>- `pending` — Промежуточный статус. Операция в обработке, повторите запрос статуса позднее<br>- `failed` — Финальный статус. Не удалось обработать операцию<br>- `succeeded` — Финальный статус. Операция успешно обработана<br><br>*Enum:* `created`, `pending`, `failed`, `succeeded` |
| fail_reason_code | **Type:** string<br><br>Код причины неуспешной обработки операции (присутствует при статусе `failed`):<br><br>- `UNABLE_TO_PROCESS` — Невозможно провести операцию.<br>- `NOT_ENOUGH_MONEY` — Недостаточно средств. Привяжите другую карту и попробуйте снова.<br>- `LIMIT_EXCEEDED` — Сумма оплаты превышает допустимый лимит.<br>- `NO_AVAILABLE_PAYMENT_METHODS` — Невозможно провести списание. Привяжите другую карту и попробуйте снова.<br>- `ORDER_EXPIRED` — Истек срок действия заказа.<br>- `CONFIRMATION_TIME_EXPIRED` — Время на подтверждение оплаты истекло.<br>- `CONFIRMATION_REJECTED` — Подтверждение было отклонено.<br>- `SYSTEM_ERROR` — Невозможно провести операцию.<br><br>*Enum:* `UNABLE_TO_PROCESS`, `NOT_ENOUGH_MONEY`, `LIMIT_EXCEEDED`, `NO_AVAILABLE_PAYMENT_METHODS`, `ORDER_EXPIRED`, `CONFIRMATION_TIME_EXPIRED`, `CONFIRMATION_REJECTED`, `SYSTEM_ERROR` |
| fail_reason_description | **Type:** string<br><br>Описание причины неуспешной обработки операции (присутствует при статусе `failed`) |
| token | **Type:** string<br><br>Платёжный токен плательщика<br><br>*Example:* `asjhdgfkajshgdkfjahgsd` |

## 400 Bad Request

Неправильный запрос

### Body

**application/json**

    
        ```
{
"error_code": "INVALID_REQUEST_ERROR",
"error_description": "string"
}
```

        
    
| **Name** | **Description** |
| --- | --- |
| error_code* | **Type:** string<br><br>*Enum:* `INVALID_REQUEST_ERROR`, `NOT_FOUND` |
| error_description* | **Type:** string<br><br>Описание кода ответа |

## 403 Forbidden

Доступ запрещён

## API Онлайн-оплаты / Оплата с созданием платёжного токена / Регистрация оплаты с созданием платёжного токена

Источник метода: <https://docs.wbpay.ru/api/online/Oplata-s-sozdaniem-platyozhnogo-tokena/apiv1ordersonlineregister-with-create-token-post.html>

# Регистрация оплаты с созданием платёжного токена

Метод возвращает ID заказа в статусе `created`

## Request

POST

    
        ```
https://api.wbpay.ru/api/v1/orders/online/register-with-create-token
```

        
    

Production

POST

    
        ```
https://sandbox.wbpay.ru/api/v1/orders/online/register-with-create-token
```

        
    

Sandbox

### Headers

| **Name** | **Description** |
| --- | --- |
| Authorization* | **Type:** string<br><br>Токен авторизации<br>Example: `Bearer {access_token}` |
| Content-Type* | **Type:** string<br><br>Тип контента<br><br>*Enum:* `application/json` |
| X-Request-Country* | **Type:** string<br><br>Код страны сервера, октуда поступает запрос, в формате ISO 3166-1<br><br>*Pattern:* `^[a-zA-Z]{2}$` |
| X-Request-Region* | **Type:** string<br><br>Код региона сервера, октуда поступает запрос, в формате ISO 3166-2<br><br>*Pattern:* `^[a-zA-Z]{2}-[a-zA-Z]{2,3}$` |
| X-Signature* | **Type:** string<br><br>Подпись запроса на основе алгоритма ED25519 |

### Body

**application/json**

    
        ```
{
"terminal_id": "string",
"invoice_id": "string",
"phone_number": "string",
"redirect_url": "wildberries://wbpay/agreement, https://wildberries.ru",
"amount": 0,
"currency_code": 0,
"created_at": 0,
"client_id": "string",
"positions": [
{
"name": "string",
"price": 0,
"count": 0
}
]
}
```

        
    
| **Name** | **Description** |
| --- | --- |
| amount* | **Type:** integer<int64><br><br>Сумма в копейках, без разделителя. `50 = "5000"` |
| client_id* | **Type:** string<br><br>Уникальный ID пользователя в системе продавца<br><br>*Max length:* `2048` |
| created_at* | **Type:** integer<int64><br><br>Дата создания заказа в формате Unix timestamp |
| currency_code* | **Type:** integer<int64><br><br>Код валюты. Для оплаты в рублях передавать значение `643` |
| phone_number* | **Type:** string<br><br>Номер телефона плательщика (пример 71234567890) |
| terminal_id* | **Type:** string<br><br>ID терминала |
| invoice_id | **Type:** string<br><br>ID операции на стороне продавца (должен быть уникальным). Данный идентификатор будет отображаться в отчетах<br><br>*Max length:* `2048` |
| positions | **Type:** [Position](api/online/Oplata-s-sozdaniem-platyozhnogo-tokena/apiv1ordersonlineregister-with-create-token-post.html#position)[]<br><br>Корзина товаров<br><br>*Max length:* `1000` |
| redirect_url | **Type:** string<uri><br><br>Специальная ссылка, которая позволяет направить пользователя в конкретное место внутри мобильного приложения или на веб-страницу<br><br>*Example:* `wildberries://wbpay/agreement, https://wildberries.ru` |

### Position

| **Name** | **Description** |
| --- | --- |
| count | **Type:** integer<int64><br><br>Количество единиц товара |
| name | **Type:** string<br><br>Название товара<br><br>*Max length:* `2048` |
| price | **Type:** integer<int64><br><br>Стоимость единицы товара в копейках, без разделителя. `50 = "5000"` |

## Responses

## 200 OK

Успешно

### Body

**application/json**

    
        ```
{
"error_code": "ERR_NONE",
"error_description": "string",
"data": {
"order_id": "string"
}
}
```

        
    
| **Name** | **Description** |
| --- | --- |
| data* | **Type:** [OrderRegistrationData](api/online/Oplata-s-sozdaniem-platyozhnogo-tokena/apiv1ordersonlineregister-with-create-token-post.html#orderregistrationdata)<br><br>Данные ответа |
| error_code | **Type:** string<br><br>*Enum:* `ERR_NONE` |
| error_description | **Type:** string<br><br>Описание кода ответа |

### OrderRegistrationData

Данные ответа

| **Name** | **Description** |
| --- | --- |
| order_id* | **Type:** string<br><br>ID заказа |

## 400 Bad Request

Неправильный запрос

### Body

**application/json**

    
        ```
{
"error_code": "INVALID_REQUEST_ERROR",
"error_description": "string"
}
```

        
    
| **Name** | **Description** |
| --- | --- |
| error_code* | **Type:** string<br><br>*Enum:* `INVALID_REQUEST_ERROR`, `INVALID_PHONE_NUMBER` |
| error_description* | **Type:** string<br><br>Описание кода ответа |

## 403 Forbidden

Доступ запрещён

## API Онлайн-оплаты / Оплата с созданием платёжного токена / Выполнение оплаты

Источник метода: <https://docs.wbpay.ru/api/online/Oplata-s-sozdaniem-platyozhnogo-tokena/apiv1ordersdo-post.html>

# Выполнение оплаты

Метод применяется для финализации оплаты

## Request

POST

    
        ```
https://api.wbpay.ru/api/v1/orders/do
```

        
    

Production

POST

    
        ```
https://sandbox.wbpay.ru/api/v1/orders/do
```

        
    

Sandbox

### Headers

| **Name** | **Description** |
| --- | --- |
| Authorization* | **Type:** string<br><br>Токен авторизации<br>Example: `Bearer {access_token}` |
| Content-Type* | **Type:** string<br><br>Тип контента<br><br>*Enum:* `application/json` |
| X-Wbpay-Id* | **Type:** string<br><br>ID запроса. Используется значение параметра `order_id`, полученного ранее в методе регистрации оплаты. Параметр необходим для осуществления корректной маршрутизации запроса |

### Body

**application/json**

    
        ```
{
"order_id": "alskdjfhalksjdhf_84765384765"
}
```

        
    
| **Name** | **Description** |
| --- | --- |
| order_id* | **Type:** string<br><br>ID заказа. Используется значение параметра `order_id`, полученного ранее в методе регистрации оплаты |

## Responses

## 200 OK

Успешно

### Body

**application/json**

    
        ```
{
"error_code": "ERR_NONE",
"error_description": "string",
"data": {
"deep_link": "wildberries://wbpay/agreement?confirmation_id=6f3bf8a8e2454876be478e3034316916_1765141200"
}
}
```

        
    
| **Name** | **Description** |
| --- | --- |
| data | **Type:** [DoOrderData](api/online/Oplata-s-sozdaniem-platyozhnogo-tokena/apiv1ordersdo-post.html#doorderdata)<br><br>Данные ответа |
| error_code | **Type:** string<br><br>*Enum:* `ERR_NONE` |
| error_description | **Type:** string<br><br>Описание кода ответа |

### DoOrderData

Данные ответа

| **Name** | **Description** |
| --- | --- |
| deep_link | **Type:** string<br><br>Диплинк на экран оплаты в приложении ВБ<br><br>*Example:* `wildberries://wbpay/agreement?confirmation_id=6f3bf8a8e2454876be478e3034316916_1765141200` |

## 400 Bad Request

Неправильный запрос

### Body

**application/json**

    
        ```
{
"error_code": "INVALID_REQUEST_ERROR",
"error_description": "string"
}
```

        
    
| **Name** | **Description** |
| --- | --- |
| error_code* | **Type:** string<br><br>*Enum:* `INVALID_REQUEST_ERROR`, `NOT_FOUND` |
| error_description* | **Type:** string<br><br>Описание кода ответа |

## 403 Forbidden

Доступ запрещён

## API Онлайн-оплаты / Оплата с созданием платёжного токена / Получение статуса оплаты

Источник метода: <https://docs.wbpay.ru/api/online/Oplata-s-sozdaniem-platyozhnogo-tokena/apiv1ordersorder_idstatus-get.html>

# Получение статуса оплаты

Метод возвращает статус оплаты заказа

## Request

GET

    
        ```
https://api.wbpay.ru/api/v1/orders/{order_id}/status
```

        
    

Production

GET

    
        ```
https://sandbox.wbpay.ru/api/v1/orders/{order_id}/status
```

        
    

Sandbox

### Path parameters

| **Name** | **Description** |
| --- | --- |
| order_id* | **Type:** string<br><br>ID заказа |

### Headers

| **Name** | **Description** |
| --- | --- |
| Authorization* | **Type:** string<br><br>Токен авторизации<br>Example: `Bearer {access_token}` |
| Content-Type* | **Type:** string<br><br>Тип контента<br><br>*Enum:* `application/json` |
| X-Wbpay-Id* | **Type:** string<br><br>ID запроса. Используется значение параметра `order_id`, полученного ранее в методе регистрации оплаты. Параметр необходим для осуществления корректной маршрутизации запроса |

## Responses

## 200 OK

Успешно

### Body

**application/json**

    
        ```
{
"error_code": "ERR_NONE",
"error_description": "Запрос выполнен успешно",
"data": {
"status": "succeeded"
}
}
```

        
    
| **Name** | **Description** |
| --- | --- |
| data* | **Type:** [OrderStatusData](api/online/Oplata-s-sozdaniem-platyozhnogo-tokena/apiv1ordersorder_idstatus-get.html#orderstatusdata)<br><br>Данные ответа |
| error_code | **Type:** string<br><br>*Enum:* `ERR_NONE` |
| error_description | **Type:** string<br><br>Описание кода ответа |

### OrderStatusData

Данные ответа

| **Name** | **Description** |
| --- | --- |
| status* | **Type:** string<br><br>Статус заказа.<br><br>- `created` — Промежуточный статус. Операция создана, но обработка еще не началась<br>- `pending` — Промежуточный статус. Операция в обработке, повторите запрос статуса позднее<br>- `failed` — Финальный статус. Не удалось обработать операцию<br>- `succeeded` — Финальный статус. Операция успешно обработана<br><br>*Enum:* `created`, `pending`, `failed`, `succeeded` |
| fail_reason_code | **Type:** string<br><br>Код причины неуспешной обработки операции (присутствует при статусе `failed`):<br><br>- `UNABLE_TO_PROCESS` — Невозможно провести операцию.<br>- `NOT_ENOUGH_MONEY` — Недостаточно средств. Привяжите другую карту и попробуйте снова.<br>- `LIMIT_EXCEEDED` — Сумма оплаты превышает допустимый лимит.<br>- `NO_AVAILABLE_PAYMENT_METHODS` — Невозможно провести списание. Привяжите другую карту и попробуйте снова.<br>- `ORDER_EXPIRED` — Истек срок действия заказа.<br>- `CONFIRMATION_TIME_EXPIRED` — Время на подтверждение оплаты истекло.<br>- `CONFIRMATION_REJECTED` — Подтверждение было отклонено.<br>- `SYSTEM_ERROR` — Невозможно провести операцию.<br><br>*Enum:* `UNABLE_TO_PROCESS`, `NOT_ENOUGH_MONEY`, `LIMIT_EXCEEDED`, `NO_AVAILABLE_PAYMENT_METHODS`, `ORDER_EXPIRED`, `CONFIRMATION_TIME_EXPIRED`, `CONFIRMATION_REJECTED`, `SYSTEM_ERROR` |
| fail_reason_description | **Type:** string<br><br>Описание причины неуспешной обработки операции (присутствует при статусе `failed`) |
| token | **Type:** string<br><br>Платёжный токен плательщика<br><br>*Example:* `asjhdgfkajshgdkfjahgsd` |

## 400 Bad Request

Неправильный запрос

### Body

**application/json**

    
        ```
{
"error_code": "INVALID_REQUEST_ERROR",
"error_description": "string"
}
```

        
    
| **Name** | **Description** |
| --- | --- |
| error_code* | **Type:** string<br><br>*Enum:* `INVALID_REQUEST_ERROR`, `NOT_FOUND` |
| error_description* | **Type:** string<br><br>Описание кода ответа |

## 403 Forbidden

Доступ запрещён

## API Онлайн-оплаты / Возвраты / Регистрация возврата

Источник метода: <https://docs.wbpay.ru/api/online/Vozvraty/apiv1refundsregister-post.html>

# Регистрация возврата

Метод возвращает ID возврата в статусе `created`

## Request

POST

    
        ```
https://api.wbpay.ru/api/v1/refunds/register
```

        
    

Production

POST

    
        ```
https://sandbox.wbpay.ru/api/v1/refunds/register
```

        
    

Sandbox

### Headers

| **Name** | **Description** |
| --- | --- |
| Authorization* | **Type:** string<br><br>Токен авторизации<br>Example: `Bearer {access_token}` |
| Content-Type* | **Type:** string<br><br>Тип контента<br><br>*Enum:* `application/json` |
| X-Signature* | **Type:** string<br><br>Подпись запроса на основе алгоритма ED25519 |
| X-Wbpay-Id* | **Type:** string<br><br>ID запроса. Используется значение параметра `order_id`, полученного ранее в методе регистрации оплаты. Параметр необходим для осуществления корректной маршрутизации запроса |

### Body

**application/json**

    
        ```
{
"terminal_id": "terminal_id",
"invoice_id": "invoice_id",
"order_id": "asdkjfhgaskdjhfg_743568374",
"amount": 123,
"currency_code": 643,
"created_at": 13453453345,
"positions": [
{
"name": "Блузка",
"price": 5000,
"count": 1
}
]
}
```

        
    
| **Name** | **Description** |
| --- | --- |
| amount* | **Type:** integer<int64><br><br>Сумма к возврату в копейках, без разделителя. `50 = "5000"` |
| created_at* | **Type:** integer<int64><br><br>Дата создания возврата в формате Unix timestamp |
| currency_code* | **Type:** integer<int64><br><br>Код валюты. Для оплаты в рублях передавать значение `643` |
| order_id* | **Type:** string<br><br>ID заказа. Используется значение параметра `order_id`, полученного ранее в методе регистрации оплаты |
| terminal_id* | **Type:** string<br><br>ID терминала |
| invoice_id | **Type:** string<br><br>ID операции на стороне продавца (должен быть уникальным). Данный идентификатор будет отображаться в отчетах<br><br>*Max length:* `2048` |
| positions | **Type:** [Position](api/online/Vozvraty/apiv1refundsregister-post.html#position)[]<br><br>Корзина товаров<br><br>*Max length:* `1000` |

### Position

| **Name** | **Description** |
| --- | --- |
| count | **Type:** integer<int64><br><br>Количество единиц товара |
| name | **Type:** string<br><br>Название товара<br><br>*Max length:* `2048` |
| price | **Type:** integer<int64><br><br>Стоимость единицы товара в копейках, без разделителя. `50 = "5000"` |

## Responses

## 200 OK

Успешно

### Body

**application/json**

    
        ```
{
"error_code": "ERR_NONE",
"error_description": "Запрос выполнен успешно",
"data": {
"refund_id": "sjdhfgjshdgfjsdfasdfa_345345"
}
}
```

        
    
| **Name** | **Description** |
| --- | --- |
| data* | **Type:** [RegisterRefundData](api/online/Vozvraty/apiv1refundsregister-post.html#registerrefunddata)<br><br>Данные ответа |
| error_code | **Type:** string<br><br>*Enum:* `ERR_NONE` |
| error_description | **Type:** string<br><br>Описание кода ответа |

### RegisterRefundData

Данные ответа

| **Name** | **Description** |
| --- | --- |
| refund_id* | **Type:** string<br><br>ID возврата |

## 400 Bad Request

Неправильный запрос

### Body

**application/json**

    
        ```
{
"error_code": "INVALID_REQUEST_ERROR",
"error_description": "string"
}
```

        
    
| **Name** | **Description** |
| --- | --- |
| error_code* | **Type:** string<br><br>*Enum:* `INVALID_REQUEST_ERROR`, `NOT_FOUND` |
| error_description* | **Type:** string<br><br>Описание кода ответа |

## 403 Forbidden

Доступ запрещён

## API Онлайн-оплаты / Возвраты / Выполнение возврата

Источник метода: <https://docs.wbpay.ru/api/online/Vozvraty/apiv1refundsdo-post.html>

# Выполнение возврата

Метод применяется для финализации возврата

## Request

POST

    
        ```
https://api.wbpay.ru/api/v1/refunds/do
```

        
    

Production

POST

    
        ```
https://sandbox.wbpay.ru/api/v1/refunds/do
```

        
    

Sandbox

### Headers

| **Name** | **Description** |
| --- | --- |
| Authorization* | **Type:** string<br><br>Токен авторизации<br>Example: `Bearer {access_token}` |
| Content-Type* | **Type:** string<br><br>Тип контента<br><br>*Enum:* `application/json` |
| X-Wbpay-Id* | **Type:** string<br><br>ID запроса. Используется значение параметра `order_id`, полученного ранее в методе регистрации оплаты. Параметр необходим для осуществления корректной маршрутизации запроса |

### Body

**application/json**

    
        ```
{
"refund_id": "sjdhfgjshdgfjsdfasdfa_345345"
}
```

        
    
| **Name** | **Description** |
| --- | --- |
| refund_id* | **Type:** string<br><br>ID возврата. Используется значение параметра `refund_id`, полученного ранее в методе регистрации возврата |

## Responses

## 200 OK

Успешно

### Body

**application/json**

    
        ```
{
"error_code": "ERR_NONE",
"error_description": "Запрос выполнен успешно"
}
```

        
    
| **Name** | **Description** |
| --- | --- |
| error_code | **Type:** string<br><br>*Enum:* `ERR_NONE` |
| error_description | **Type:** string<br><br>Описание кода ответа |

## 400 Bad Request

Неправильный запрос

### Body

**application/json**

    
        ```
{
"error_code": "INVALID_REQUEST_ERROR",
"error_description": "string"
}
```

        
    
| **Name** | **Description** |
| --- | --- |
| error_code* | **Type:** string<br><br>*Enum:* `INVALID_REQUEST_ERROR`, `NOT_FOUND` |
| error_description* | **Type:** string<br><br>Описание кода ответа |

## 403 Forbidden

Доступ запрещён

## API Онлайн-оплаты / Возвраты / Получение статуса возврата

Источник метода: <https://docs.wbpay.ru/api/online/Vozvraty/apiv1refundsrefund_idstatus-get.html>

# Получение статуса возврата

Метод возвращает статус возврата

## Request

GET

    
        ```
https://api.wbpay.ru/api/v1/refunds/{refund_id}/status
```

        
    

Production

GET

    
        ```
https://sandbox.wbpay.ru/api/v1/refunds/{refund_id}/status
```

        
    

Sandbox

### Path parameters

| **Name** | **Description** |
| --- | --- |
| refund_id* | **Type:** string<br><br>ID возврата |

### Headers

| **Name** | **Description** |
| --- | --- |
| Authorization* | **Type:** string<br><br>Токен авторизации<br>Example: `Bearer {access_token}` |
| Content-Type* | **Type:** string<br><br>Тип контента<br><br>*Enum:* `application/json` |
| X-Wbpay-Id* | **Type:** string<br><br>ID запроса. Используется значение параметра `order_id`, полученного ранее в методе регистрации оплаты. Параметр необходим для осуществления корректной маршрутизации запроса |

## Responses

## 200 OK

Успешно

### Body

**application/json**

    
        ```
{
"error_code": "ERR_NONE",
"error_description": "string",
"data": {
"status": "created",
"fail_reason_code": "UNABLE_TO_PROCESS",
"fail_reason_description": "string"
}
}
```

        
    
| **Name** | **Description** |
| --- | --- |
| data* | **Type:** [RefundStatusData](api/online/Vozvraty/apiv1refundsrefund_idstatus-get.html#refundstatusdata)<br><br>Данные ответа |
| error_code | **Type:** string<br><br>*Enum:* `ERR_NONE` |
| error_description | **Type:** string<br><br>Описание кода ответа |

### RefundStatusData

Данные ответа

| **Name** | **Description** |
| --- | --- |
| status* | **Type:** string<br><br>Статус возврата.<br><br>- `created` — Промежуточный статус. Операция создана, но обработка еще не началась<br>- `pending` — Промежуточный статус. Операция в обработке, повторите запрос статуса позднее<br>- `failed` — Финальный статус. Не удалось обработать операцию<br>- `succeeded` — Финальный статус. Операция успешно обработана<br><br>*Enum:* `created`, `pending`, `failed`, `succeeded` |
| fail_reason_code | **Type:** string<br><br>Код причины неуспешной обработки операции (присутствует при статусе `failed`)<br><br>- `UNABLE_TO_PROCESS` — Невозможно провести операцию.<br>- `REFUND_NOT_POSSIBLE` — Невозможно выполнить возврат средств.<br>- `REFUND_EXPIRED` — Истек срок действия возврата.<br>- `SYSTEM_ERROR` — Невозможно провести операцию.<br><br>*Enum:* `UNABLE_TO_PROCESS`, `REFUND_NOT_POSSIBLE`, `REFUND_EXPIRED`, `SYSTEM_ERROR` |
| fail_reason_description | **Type:** string<br><br>Описание причины неуспешной обработки операции (присутствует при статусе `failed`) |

## 400 Bad Request

Неправильный запрос

### Body

**application/json**

    
        ```
{
"error_code": "INVALID_REQUEST_ERROR",
"error_description": "string"
}
```

        
    
| **Name** | **Description** |
| --- | --- |
| error_code* | **Type:** string<br><br>*Enum:* `INVALID_REQUEST_ERROR`, `NOT_FOUND` |
| error_description* | **Type:** string<br><br>Описание кода ответа |

## 403 Forbidden

Доступ запрещён

## API Оффлайн-оплаты / Оффлайн-оплата / Регистрация оффлайн оплаты

Источник метода: <https://docs.wbpay.ru/api/offline/Offlajn-oplata/apiv1ordersofflineregister-post.html>

# Регистрация оффлайн оплаты

Метод возвращает ID заказа в статусе `created`

## Request

POST

    
        ```
https://api.wbpay.ru/api/v1/orders/offline/register
```

        
    

Production

POST

    
        ```
https://sandbox.wbpay.ru/api/v1/orders/offline/register
```

        
    

Sandbox

### Headers

| **Name** | **Description** |
| --- | --- |
| Authorization* | **Type:** string<br><br>Токен авторизации<br>Example: `Bearer {access_token}` |
| Content-Type* | **Type:** string<br><br>Тип контента<br><br>*Enum:* `application/json` |
| X-Request-Country* | **Type:** string<br><br>Код страны сервера, октуда поступает запрос, в формате ISO 3166-1<br><br>*Pattern:* `^[a-zA-Z]{2}$` |
| X-Request-Region* | **Type:** string<br><br>Код региона сервера, октуда поступает запрос, в формате ISO 3166-2<br><br>*Pattern:* `^[a-zA-Z]{2}-[a-zA-Z]{2,3}$` |
| X-Signature* | **Type:** string<br><br>Подпись запроса на основе алгоритма ED25519 |

### Body

**application/json**

    
        ```
{
"terminal_id": "terminal_id",
"invoice_id": "invoice_id",
"qr_code": "qr_code",
"amount": 5000,
"currency_code": 643,
"created_at": 123435453,
"positions": [
{
"name": "Блузка",
"price": 5000,
"count": 1
}
]
}
```

        
    
| **Name** | **Description** |
| --- | --- |
| amount* | **Type:** integer<int64><br><br>Сумма в копейках, без разделителя. `50 = "5000"` |
| created_at* | **Type:** integer<int64><br><br>Дата создания заказа в формате Unix timestamp |
| currency_code* | **Type:** integer<int64><br><br>Код валюты. Для оплаты в рублях передавать значение `643` |
| qr_code* | **Type:** string<br><br>QR код |
| terminal_id* | **Type:** string<br><br>ID терминала |
| invoice_id | **Type:** string<br><br>ID операции на стороне продавца (должен быть уникальным). Данный идентификатор будет отображаться в отчетах<br><br>*Max length:* `2048` |
| positions | **Type:** [Position](api/offline/Offlajn-oplata/apiv1ordersofflineregister-post.html#position)[]<br><br>Корзина товаров<br><br>*Max items:* `1000` |

### Position

| **Name** | **Description** |
| --- | --- |
| count | **Type:** integer<int64><br><br>Количество единиц товара |
| name | **Type:** string<br><br>Название товара<br><br>*Max length:* `2048` |
| price | **Type:** integer<int64><br><br>Стоимость единицы товара в копейках, без разделителя. `50 = "5000"` |

## Responses

## 200 OK

Успешно

### Body

**application/json**

    
        ```
{
"error_code": "ERR_NONE",
"error_description": "string",
"data": {
"order_id": "string"
}
}
```

        
    
| **Name** | **Description** |
| --- | --- |
| data* | **Type:** [RegisterOfflineOrderData](api/offline/Offlajn-oplata/apiv1ordersofflineregister-post.html#registerofflineorderdata)<br><br>Данные ответа |
| error_code | **Type:** string<br><br>*Enum:* `ERR_NONE` |
| error_description | **Type:** string<br><br>Описание кода ответа |

### RegisterOfflineOrderData

Данные ответа

| **Name** | **Description** |
| --- | --- |
| order_id* | **Type:** string<br><br>ID заказа |

## 400 Bad Request

Неправильный запрос

### Body

**application/json**

    
        ```
{
"error_code": "INVALID_REQUEST_ERROR",
"error_description": "string"
}
```

        
    
| **Name** | **Description** |
| --- | --- |
| error_code* | **Type:** string<br><br>*Enum:* `INVALID_REQUEST_ERROR`, `DUPLICATE_QR_CODE`, `INVALID_QR_CODE`, `EXPIRED_QR_CODE` |
| error_description* | **Type:** string<br><br>Описание кода ответа |

## 403 Forbidden

Доступ запрещён

## API Оффлайн-оплаты / Оффлайн-оплата / Выполнение оффлайн оплаты

Источник метода: <https://docs.wbpay.ru/api/offline/Offlajn-oplata/apiv1ordersdo-post.html>

# Выполнение оффлайн оплаты

Метод применяется для финализации оплаты

## Request

POST

    
        ```
https://api.wbpay.ru/api/v1/orders/do
```

        
    

Production

POST

    
        ```
https://sandbox.wbpay.ru/api/v1/orders/do
```

        
    

Sandbox

### Headers

| **Name** | **Description** |
| --- | --- |
| Authorization* | **Type:** string<br><br>Токен авторизации<br>Example: `Bearer {access_token}` |
| Content-Type* | **Type:** string<br><br>Тип контента<br><br>*Enum:* `application/json` |
| X-Wbpay-Id* | **Type:** string<br><br>ID заказа. Используется значение параметра `order_id`, полученного ранее в методе регистрации оплаты. Параметр необходим для осуществления корректной маршрутизации запроса |

### Body

**application/json**

    
        ```
{
"order_id": "alskdjfhalksjdhf_84765384765"
}
```

        
    
| **Name** | **Description** |
| --- | --- |
| order_id* | **Type:** string<br><br>ID заказа. Используется значение параметра `order_id`, полученного ранее в методе регистрации оплаты |

## Responses

## 200 OK

Успешно

### Body

**application/json**

    
        ```
{
"error_code": "ERR_NONE",
"error_description": "Запрос выполнен успешно"
}
```

        
    
| **Name** | **Description** |
| --- | --- |
| error_code | **Type:** string<br><br>*Enum:* `ERR_NONE` |
| error_description | **Type:** string<br><br>Описание кода ответа |

## 400 Bad Request

Неправильный запрос

### Body

**application/json**

    
        ```
{
"error_code": "INVALID_REQUEST_ERROR",
"error_description": "string"
}
```

        
    
| **Name** | **Description** |
| --- | --- |
| error_code* | **Type:** string<br><br>*Enum:* `INVALID_REQUEST_ERROR`, `NOT_FOUND` |
| error_description* | **Type:** string<br><br>Описание кода ответа |

## 403 Forbidden

Доступ запрещён

## API Оффлайн-оплаты / Оффлайн-оплата / Запрос статуса оплаты

Источник метода: <https://docs.wbpay.ru/api/offline/Offlajn-oplata/apiv1ordersorder_idstatus-get.html>

# Запрос статуса оплаты

Метод возвращает статус оплаты заказа

## Request

GET

    
        ```
https://api.wbpay.ru/api/v1/orders/{order_id}/status
```

        
    

Production

GET

    
        ```
https://sandbox.wbpay.ru/api/v1/orders/{order_id}/status
```

        
    

Sandbox

### Path parameters

| **Name** | **Description** |
| --- | --- |
| order_id* | **Type:** string<br><br>ID заказа. Используется значение параметра `order_id`, полученного ранее в методе регистрации оплаты |

### Headers

| **Name** | **Description** |
| --- | --- |
| Authorization* | **Type:** string<br><br>Токен авторизации<br>Example: `Bearer {access_token}` |
| Content-Type* | **Type:** string<br><br>Тип контента<br><br>*Enum:* `application/json` |
| X-Wbpay-Id* | **Type:** string<br><br>ID заказа. Используется значение параметра `order_id`, полученного ранее в методе регистрации оплаты. Параметр необходим для осуществления корректной маршрутизации запроса |

## Responses

## 200 OK

Успешно

### Body

**application/json**

    
        ```
{
"error_code": "ERR_NONE",
"error_description": "Запрос выполнен успешно",
"data": {
"status": "succeeded"
}
}
```

        
    
| **Name** | **Description** |
| --- | --- |
| data* | **Type:** [OrderStatusData](api/offline/Offlajn-oplata/apiv1ordersorder_idstatus-get.html#orderstatusdata)<br><br>Данные ответа |
| error_code | **Type:** string<br><br>*Enum:* `ERR_NONE` |
| error_description | **Type:** string<br><br>Описание кода ответа |

### OrderStatusData

Данные ответа

| **Name** | **Description** |
| --- | --- |
| status* | **Type:** string<br><br>Статус заказа.<br><br>- `created` — Промежуточный статус. Операция создана, но обработка еще не началась<br>- `pending` — Промежуточный статус. Операция в обработке, повторите запрос статуса позднее<br>- `failed` — Финальный статус. Не удалось обработать операцию<br>- `succeeded` — Финальный статус. Операция успешно обработана<br><br>*Enum:* `created`, `pending`, `failed`, `succeeded` |
| fail_reason_code | **Type:** string<br><br>Код причины неуспешной обработки операции (присутствует при статусе `failed`):<br><br>- `UNABLE_TO_PROCESS` — Невозможно провести операцию.<br>- `NOT_ENOUGH_MONEY` — Недостаточно средств. Привяжите другую карту и попробуйте снова.<br>- `LIMIT_EXCEEDED` — Сумма оплаты превышает допустимый лимит.<br>- `NO_AVAILABLE_PAYMENT_METHODS` — Невозможно провести списание. Привяжите другую карту и попробуйте снова.<br>- `ORDER_EXPIRED` — Истек срок действия заказа.<br>- `CONFIRMATION_TIME_EXPIRED` — Время на подтверждение оплаты истекло.<br>- `CONFIRMATION_REJECTED` — Подтверждение было отклонено.<br>- `SYSTEM_ERROR` — Невозможно провести операцию.<br><br>*Enum:* `UNABLE_TO_PROCESS`, `NOT_ENOUGH_MONEY`, `LIMIT_EXCEEDED`, `NO_AVAILABLE_PAYMENT_METHODS`, `ORDER_EXPIRED`, `CONFIRMATION_TIME_EXPIRED`, `CONFIRMATION_REJECTED`, `SYSTEM_ERROR` |
| fail_reason_description | **Type:** string<br><br>Описание причины неуспешной обработки операции (присутствует при статусе `failed`) |

## 400 Bad Request

Неправильный запрос

### Body

**application/json**

    
        ```
{
"error_code": "INVALID_REQUEST_ERROR",
"error_description": "string"
}
```

        
    
| **Name** | **Description** |
| --- | --- |
| error_code* | **Type:** string<br><br>*Enum:* `INVALID_REQUEST_ERROR`, `NOT_FOUND` |
| error_description* | **Type:** string<br><br>Описание кода ответа |

## 403 Forbidden

Доступ запрещён

## API Оффлайн-оплаты / Возвраты / Регистрация возврата

Источник метода: <https://docs.wbpay.ru/api/offline/Vozvraty/apiv1refundsregister-post.html>

# Регистрация возврата

Метод возвращает ID возврата в статусе `pending`

## Request

POST

    
        ```
https://api.wbpay.ru/api/v1/refunds/register
```

        
    

Production

POST

    
        ```
https://sandbox.wbpay.ru/api/v1/refunds/register
```

        
    

Sandbox

### Headers

| **Name** | **Description** |
| --- | --- |
| Authorization* | **Type:** string<br><br>Токен авторизации<br>Example: `Bearer {access_token}` |
| Content-Type* | **Type:** string<br><br>Тип контента<br><br>*Enum:* `application/json` |
| X-Signature* | **Type:** string<br><br>Подпись запроса на основе алгоритма ED25519 |
| X-Wbpay-Id* | **Type:** string<br><br>ID заказа. Используется значение параметра `order_id`, полученного ранее в методе регистрации оплаты. Параметр необходим для осуществления корректной маршрутизации запроса |

### Body

**application/json**

    
        ```
{
"terminal_id": "terminal_id",
"invoice_id": "invoice_id",
"order_id": "asdkjfhgaskdjhfg_743568374",
"amount": 123,
"currency_code": 643,
"created_at": 13453453345,
"positions": [
{
"name": "Блузка",
"price": 5000,
"count": 1
}
]
}
```

        
    
| **Name** | **Description** |
| --- | --- |
| amount* | **Type:** integer<int64><br><br>Сумма к возврату в копейках, без разделителя. `50 = "5000"` |
| created_at* | **Type:** integer<int64><br><br>Дата создания возврата в формате Unix timestamp |
| currency_code* | **Type:** integer<int64><br><br>Код валюты. Для оплаты в рублях передавать значение `643` |
| order_id* | **Type:** string<br><br>ID заказа. Используется значение параметра `order_id`, полученного ранее в методе регистрации оплаты |
| terminal_id* | **Type:** string<br><br>ID терминала |
| invoice_id | **Type:** string<br><br>ID операции на стороне продавца (должен быть уникальным). Данный идентификатор будет отображаться в отчетах<br><br>*Max length:* `2048` |
| positions | **Type:** [Position](api/offline/Vozvraty/apiv1refundsregister-post.html#position)[]<br><br>Корзина товаров<br><br>*Max items:* `1000` |

### Position

| **Name** | **Description** |
| --- | --- |
| count | **Type:** integer<int64><br><br>Количество единиц товара |
| name | **Type:** string<br><br>Название товара<br><br>*Max length:* `2048` |
| price | **Type:** integer<int64><br><br>Стоимость единицы товара в копейках, без разделителя. `50 = "5000"` |

## Responses

## 200 OK

Успешно

### Body

**application/json**

    
        ```
{
"error_code": "ERR_NONE",
"error_description": "Запрос выполнен успешно",
"data": {
"refund_id": "sjdhfgjshdgfjsdfasdfa_345345"
}
}
```

        
    
| **Name** | **Description** |
| --- | --- |
| data* | **Type:** [RegisterRefundData](api/offline/Vozvraty/apiv1refundsregister-post.html#registerrefunddata)<br><br>Данные ответа |
| error_code | **Type:** string<br><br>*Enum:* `ERR_NONE` |
| error_description | **Type:** string<br><br>Описание кода ответа |

### RegisterRefundData

Данные ответа

| **Name** | **Description** |
| --- | --- |
| refund_id* | **Type:** string<br><br>ID возврата |

## 400 Bad Request

Неправильный запрос

### Body

**application/json**

    
        ```
{
"error_code": "INVALID_REQUEST_ERROR",
"error_description": "string"
}
```

        
    
| **Name** | **Description** |
| --- | --- |
| error_code* | **Type:** string<br><br>*Enum:* `INVALID_REQUEST_ERROR`, `NOT_FOUND` |
| error_description* | **Type:** string<br><br>Описание кода ответа |

## 403 Forbidden

Доступ запрещён

## API Оффлайн-оплаты / Возвраты / Выполнение возврата

Источник метода: <https://docs.wbpay.ru/api/offline/Vozvraty/apiv1refundsdo-post.html>

# Выполнение возврата

Метод применяется для финализации возврата

## Request

POST

    
        ```
https://api.wbpay.ru/api/v1/refunds/do
```

        
    

Production

POST

    
        ```
https://sandbox.wbpay.ru/api/v1/refunds/do
```

        
    

Sandbox

### Headers

| **Name** | **Description** |
| --- | --- |
| Authorization* | **Type:** string<br><br>Токен авторизации<br>Example: `Bearer {access_token}` |
| Content-Type* | **Type:** string<br><br>Тип контента<br><br>*Enum:* `application/json` |
| X-Wbpay-Id* | **Type:** string<br><br>ID заказа. Используется значение параметра `order_id`, полученного ранее в методе регистрации оплаты. Параметр необходим для осуществления корректной маршрутизации запроса |

### Body

**application/json**

    
        ```
{
"refund_id": "sjdhfgjshdgfjsdfasdfa_345345"
}
```

        
    
| **Name** | **Description** |
| --- | --- |
| refund_id* | **Type:** string<br><br>ID возврата. Используется значение параметра `refund_id`, полученного ранее в методе регистрации возврата |

## Responses

## 200 OK

Успешно

### Body

**application/json**

    
        ```
{
"error_code": "ERR_NONE",
"error_description": "Запрос выполнен успешно"
}
```

        
    
| **Name** | **Description** |
| --- | --- |
| error_code | **Type:** string<br><br>*Enum:* `ERR_NONE` |
| error_description | **Type:** string<br><br>Описание кода ответа |

## 400 Bad Request

Неправильный запрос

### Body

**application/json**

    
        ```
{
"error_code": "INVALID_REQUEST_ERROR",
"error_description": "string"
}
```

        
    
| **Name** | **Description** |
| --- | --- |
| error_code* | **Type:** string<br><br>*Enum:* `INVALID_REQUEST_ERROR`, `NOT_FOUND` |
| error_description* | **Type:** string<br><br>Описание кода ответа |

## 403 Forbidden

Доступ запрещён

## API Оффлайн-оплаты / Возвраты / Запрос статуса возврата

Источник метода: <https://docs.wbpay.ru/api/offline/Vozvraty/apiv1refundsrefund_idstatus-get.html>

# Запрос статуса возврата

Метод возвращает статус возврата

## Request

GET

    
        ```
https://api.wbpay.ru/api/v1/refunds/{refund_id}/status
```

        
    

Production

GET

    
        ```
https://sandbox.wbpay.ru/api/v1/refunds/{refund_id}/status
```

        
    

Sandbox

### Path parameters

| **Name** | **Description** |
| --- | --- |
| refund_id* | **Type:** string<br><br>ID возврата. Используется значение параметра `refund_id`, полученного ранее в методе регистрации возврата |

### Headers

| **Name** | **Description** |
| --- | --- |
| Authorization* | **Type:** string<br><br>Токен авторизации<br>Example: `Bearer {access_token}` |
| Content-Type* | **Type:** string<br><br>Тип контента<br><br>*Enum:* `application/json` |
| X-Wbpay-Id* | **Type:** string<br><br>ID заказа. Используется значение параметра `order_id`, полученного ранее в методе регистрации оплаты. Параметр необходим для осуществления корректной маршрутизации запроса |

## Responses

## 200 OK

Успешно

### Body

**application/json**

    
        ```
{
"error_code": "ERR_NONE",
"error_description": "string",
"data": {
"status": "created",
"fail_reason_code": "UNABLE_TO_PROCESS",
"fail_reason_description": "string"
}
}
```

        
    
| **Name** | **Description** |
| --- | --- |
| data* | **Type:** [RefundStatusData](api/offline/Vozvraty/apiv1refundsrefund_idstatus-get.html#refundstatusdata)<br><br>Данные ответа |
| error_code | **Type:** string<br><br>*Enum:* `ERR_NONE` |
| error_description | **Type:** string<br><br>Описание кода ответа |

### RefundStatusData

Данные ответа

| **Name** | **Description** |
| --- | --- |
| fail_reason_code | **Type:** string<br><br>Код причины неуспешной обработки операции (присутствует при статусе `failed`)<br><br>- `UNABLE_TO_PROCESS` — Невозможно провести операцию.<br>- `REFUND_NOT_POSSIBLE` — Невозможно выполнить возврат средств.<br>- `REFUND_EXPIRED` — Истек срок действия возврата.<br>- `SYSTEM_ERROR` — Невозможно провести операцию.<br><br>*Enum:* `UNABLE_TO_PROCESS`, `REFUND_NOT_POSSIBLE`, `REFUND_EXPIRED`, `SYSTEM_ERROR` |
| fail_reason_description | **Type:** string<br><br>Описание причины неуспешной обработки операции (присутствует при статусе `failed`) |
| status | **Type:** string<br><br>Статус возврата.<br><br>- `created` — Промежуточный статус. Операция создана, но обработка еще не началась<br>- `pending` — Промежуточный статус. Операция в обработке, повторите запрос статуса позднее<br>- `failed` — Финальный статус. Не удалось обработать операцию<br>- `succeeded` — Финальный статус. Операция успешно обработана<br><br>*Enum:* `created`, `pending`, `failed`, `succeeded` |

## 400 Bad Request

Неправильный запрос

### Body

**application/json**

    
        ```
{
"error_code": "INVALID_REQUEST_ERROR",
"error_description": "string"
}
```

        
    
| **Name** | **Description** |
| --- | --- |
| error_code* | **Type:** string<br><br>*Enum:* `INVALID_REQUEST_ERROR`, `NOT_FOUND` |
| error_description* | **Type:** string<br><br>Описание кода ответа |

## 403 Forbidden

Доступ запрещён

## API Программы лояльности / Программа лояльности / Получение настроек лояльности терминала

Источник метода: <https://docs.wbpay.ru/api/loyalty/Programma-loyalnosti/apiv1loyaltiessettings-post.html>

# Получение настроек лояльности терминала

Используется для получения актуальных настроек программы лояльности по терминалу

## Request

POST

    
        ```
https://api.wbpay.ru/api/v1/loyalties/settings
```

        
    

Production

POST

    
        ```
https://sandbox.wbpay.ru/api/v1/loyalties/settings
```

        
    

Sandbox

### Headers

| **Name** | **Description** |
| --- | --- |
| Authorization* | **Type:** string<br><br>Токен авторизации<br>Example: `Bearer {access_token}` |
| Content-Type* | **Type:** string<br><br>Тип контента<br><br>*Enum:* `application/json` |
| X-Request-Country* | **Type:** string<br><br>Код страны в формате ISO 3166-1<br><br>*Pattern:* `^[a-zA-Z]{2}$` |
| X-Request-Region* | **Type:** string<br><br>Код региона в формате ISO 3166-2<br><br>*Pattern:* `^[a-zA-Z]{2}-[a-zA-Z]{2,3}$` |

### Body

**application/json**

    
        ```
{
"terminal_id": "sandbox_a30fffcd4ry0g_1"
}
```

        
    
| **Name** | **Description** |
| --- | --- |
| terminal_id* | **Type:** string<br><br>ID терминала |

## Responses

## 200 OK

Успешно

### Body

**application/json**

    
        ```
{
"error_code": "ERR_NONE",
"error_description": "string",
"data": {
"loyalty_start_date": 1748773839,
"loyalty_end_date": null,
"status": "active",
"loyalty_rate": 5,
"min_cashback_amount": 10,
"max_cashback_amount": 200,
"monthly_cashback_limit": 1000,
"actual_date": 1750225539
}
}
```

        
    
| **Name** | **Description** |
| --- | --- |
| data* | **Type:** [LoyaltySettingsData](api/loyalty/Programma-loyalnosti/apiv1loyaltiessettings-post.html#loyaltysettingsdata)<br><br>Данные ответа<br><br>*Example:* `[object Object]` |
| error_code | **Type:** string<br><br>*Enum:* `ERR_NONE` |
| error_description | **Type:** string<br><br>Описание кода ответа |

### LoyaltySettingsData

Данные ответа

| **Name** | **Description** |
| --- | --- |
| actual_date | **Type:** integer<br><br>Дата, на которую актуальны настройки в формате Unix timestamp |
| loyalty_end_date | **Type:** integer<br><br>Дата окончания участия в формате Unix timestamp |
| loyalty_rate | **Type:** number<float><br><br>Процент кешбэка от суммы покупки (от 0 до 100) |
| loyalty_start_date | **Type:** integer<br><br>Дата начала участия в формате Unix timestamp |
| max_cashback_amount | **Type:** integer<br><br>Максимальная сумма кешбэка за одну операцию в копейках, без разделителя. `50 = "5000"`. Не может быть меньше min_cashback_amount |
| min_cashback_amount | **Type:** integer<br><br>Минимальная сумма кешбэка за одну операцию в копейках, без разделителя. `50 = "5000"`. Не может быть больше max_cashback_amout |
| monthly_cashback_limit | **Type:** integer<br><br>Лимит кешбэка в месяц на одного клиента по терминалу в копейках, без разделителя. `50 = "5000"` |
| status | **Type:** string<br><br>Статус программы.<br><br>- `active` — Программа активна, создаем зачисление кешбэка<br>- `blocked` — Программа временно заблокирована<br>- `inactive` - Программа завершена (по end_date)<br><br>*Enum:* `active`, `blocked`, `inactive` |

## 400 Bad Request

Неправильный запрос

### Body

**application/json**

    
        ```
{
"error_code": "INVALID_REQUEST_ERROR",
"error_description": "string"
}
```

        
    
| **Name** | **Description** |
| --- | --- |
| error_code* | **Type:** string<br><br>*Enum:* `INVALID_REQUEST_ERROR`, `NOT_FOUND` |
| error_description* | **Type:** string<br><br>Описание кода ответа |

## 403 Forbidden

Доступ запрещён

## API Программы лояльности / Программа лояльности / Предварительный расчет количества бонусов

Источник метода: <https://docs.wbpay.ru/api/loyalty/Programma-loyalnosti/apiv1loyaltiescalculate_cashback-post.html>

# Предварительный расчет количества бонусов

Используется для получения предварительного расчета количества ягодок на покупку

## Request

POST

    
        ```
https://api.wbpay.ru/api/v1/loyalties/calculate_cashback
```

        
    

Production

POST

    
        ```
https://sandbox.wbpay.ru/api/v1/loyalties/calculate_cashback
```

        
    

Sandbox

### Headers

| **Name** | **Description** |
| --- | --- |
| Authorization* | **Type:** string<br><br>Токен авторизации<br>Example: `Bearer {access_token}` |
| Content-Type* | **Type:** string<br><br>Тип контента<br><br>*Enum:* `application/json` |
| X-Request-Country* | **Type:** string<br><br>Код страны в формате ISO 3166-1<br><br>*Pattern:* `^[a-zA-Z]{2}$` |
| X-Request-Region* | **Type:** string<br><br>Код региона в формате ISO 3166-2<br><br>*Pattern:* `^[a-zA-Z]{2}-[a-zA-Z]{2,3}$` |

### Body

**application/json**

    
        ```
{
"terminal_id": "string",
"phone_number": "string",
"qr_code": "string",
"currency_code": 0,
"amount": 0,
"positions": [
{
"number": 1,
"name": "string",
"price": 0,
"count": 0
}
]
}
```

        
    
| **Name** | **Description** |
| --- | --- |
| amount* | **Type:** integer<br><br>Сумма в копейках, без разделителя. 50 = "5000" |
| currency_code* | **Type:** integer<br><br>Код валюты |
| terminal_id* | **Type:** string<br><br>ID терминала, в котором происходит покупка |
| phone_number | **Type:** string<br><br>Номер телефона плательщика |
| positions | **Type:** [Position](api/loyalty/Programma-loyalnosti/apiv1loyaltiescalculate_cashback-post.html#position)[]<br><br>Корзина товаров |
| qr_code | **Type:** string<br><br>QR код покупателя при оффлайн оплате |

### Position

| **Name** | **Description** |
| --- | --- |
| count* | **Type:** integer<int64><br><br>Количество единиц товара |
| name* | **Type:** string<br><br>Название товара<br><br>*Max length:* `2048` |
| number* | **Type:** integer<br><br>Порядковый номер<br><br>*Example:* `1` |
| price* | **Type:** integer<int64><br><br>Стоимость единицы товара в копейках, без разделителя. `50 = "5000"` |

## Responses

## 200 OK

Успешно

### Body

**application/json**

    
        ```
{
"error_code": "ERR_NONE",
"error_description": "string",
"data": {
"total_reward": 0,
"positions": [
{
"reward_per_unit": 0,
"total_unit_reward": 0
}
]
}
}
```

        
    
| **Name** | **Description** |
| --- | --- |
| data* | **Type:** [CalculateRewardData](api/loyalty/Programma-loyalnosti/apiv1loyaltiescalculate_cashback-post.html#calculaterewarddata) |
| error_code | **Type:** string<br><br>*Enum:* `ERR_NONE` |
| error_description | **Type:** string<br><br>Описание кода ответа |

### CalculateRewardData

| **Name** | **Description** |
| --- | --- |
| total_reward* | **Type:** integer<br><br>Количество ягодок, которые получит клиент за покупку с использованием WB Кошелька |
| positions | **Type:** [CalculatePosition](api/loyalty/Programma-loyalnosti/apiv1loyaltiescalculate_cashback-post.html#calculateposition)[] |

### CalculatePosition

| **Name** | **Description** |
| --- | --- |
| reward_per_unit* | **Type:** integer<br><br>Количество ягодок за единицу товара/услуги |
| total_unit_reward* | **Type:** integer<br><br>Количество ягодок за все количество товара/услуги |

## 400 Bad Request

Неправильный запрос

### Body

**application/json**

    
        ```
{
"error_code": "INVALID_REQUEST_ERROR",
"error_description": "string"
}
```

        
    
| **Name** | **Description** |
| --- | --- |
| error_code* | **Type:** string<br><br>*Enum:* `INVALID_REQUEST_ERROR`, `NOT_FOUND` |
| error_description* | **Type:** string<br><br>Описание кода ответа |

## 403 Forbidden

Доступ запрещён
