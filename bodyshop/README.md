# bodyshop

## Endpoints

### /send

`/send` применяется для отправки заполненного шаблона на сервис `profile`.

Запрос:
```sh
curl --location --request POST 'http://localhost:5000/send' \
--header 'Content-Type: application/json' \
-d '
    {
        "receivers": [
            "x",
            "z",
            "y"
        ],
        "template_id": "testTemplateId1",
        "parameters": [
            {
                "parameter_name": "testParamName1",
                "parameter_value": "paramValue1"
            },
            {
                "parameter_name": "testParamName2",
                "parameter_value": "paramValue2"
            }
        ]
    }
'
```

При успехе возвращается: **200 OK**