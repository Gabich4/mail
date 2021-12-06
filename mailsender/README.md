# mailsender

## Endpoints

### /send

`/send` применяется для отправки заполненного шаблона на сервис `profile`.

Запрос:
```sh
curl --location --request POST 'http://localhost:4000/api/v1/send' \
--header 'Content-Type: application/json' \
-d '
    {
        "receivers": [
            "x", "y", "z"
        ],
        "message": "<div><h1>{{ .paramValue1}}</h1><p>{{ .paramValue2}}</p></div>"
    }
'
```
При успехе возвращается: **200 OK**