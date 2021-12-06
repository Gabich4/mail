# auth



## Логирование

Используется [logrus](https://github.com/sirupsen/logrus) обращения к сервису пишутся через него в stdout.
Так-же через него выполняется логирование различных ситуаций в работе сервиса.
``` json
{"level":"info","msg":"Service profile started. localhost:8000"}
{"http_method":"GET","http_proto":"HTTP/1.1","http_scheme":"http","level":"info","msg":"request started","remote_addr":"127.0.0.1:64917","req_id":"LAPTOP-NFG6QTU2/wA0XUpCfvv-000001","ts":"Sun, 14 Nov 2021 13:59:26 UTC","uri":"http://localhost:8000/login","user_agent":"PostmanRuntime/7.28.4"}
{"level":"info","msg":"Url Param 'redirect_uri' is missing"}
{"http_method":"GET","http_proto":"HTTP/1.1","http_scheme":"http","level":"info","msg":"request complete","remote_addr":"127.0.0.1:64917","req_id":"LAPTOP-NFG6QTU2/wA0XUpCfvv-000001","resp_bytes_length":0,"resp_elapsed_ms":75.5813,"resp_status":0,"ts":"Sun, 14 Nov 2021 13:59:26 UTC","uri":"http://localhost:8000/login","user_agent":"PostmanRuntime/7.28.4"}
{"http_method":"GET","http_proto":"HTTP/1.1","http_scheme":"http","level":"info","msg":"request started","remote_addr":"127.0.0.1:64917","req_id":"LAPTOP-NFG6QTU2/wA0XUpCfvv-000002","ts":"Sun, 14 Nov 2021 13:59:28 UTC","uri":"http://localhost:8000/i","user_agent":"PostmanRuntime/7.28.4"}
{"http_method":"GET","http_proto":"HTTP/1.1","http_scheme":"http","level":"info","msg":"request complete","remote_addr":"127.0.0.1:64917","req_id":"LAPTOP-NFG6QTU2/wA0XUpCfvv-000002","resp_bytes_length":21,"resp_elapsed_ms":0,"resp_status":200,"ts":"Sun, 14 Nov 2021 13:59:28 UTC","uri":"http://localhost:8000/i","user_agent":"PostmanRuntime/7.28.4"}
{"http_method":"GET","http_proto":"HTTP/1.1","http_scheme":"http","level":"info","msg":"request started","remote_addr":"127.0.0.1:64917","req_id":"LAPTOP-NFG6QTU2/wA0XUpCfvv-000003","ts":"Sun, 14 Nov 2021 13:59:31 UTC","uri":"http://localhost:8000/me","user_agent":"PostmanRuntime/7.28.4"}
{"http_method":"GET","http_proto":"HTTP/1.1","http_scheme":"http","level":"info","msg":"request complete","remote_addr":"127.0.0.1:64917","req_id":"LAPTOP-NFG6QTU2/wA0XUpCfvv-000003","resp_bytes_length":21,"resp_elapsed_ms":0.2847,"resp_status":200,"ts":"Sun, 14 Nov 2021 13:59:31 UTC","uri":"http://localhost:8000/me","user_agent":"PostmanRuntime/7.28.4"}
{"http_method":"GET","http_proto":"HTTP/1.1","http_scheme":"http","level":"info","msg":"request started","remote_addr":"127.0.0.1:64917","req_id":"LAPTOP-NFG6QTU2/wA0XUpCfvv-000004","ts":"Sun, 14 Nov 2021 13:59:33 UTC","uri":"http://localhost:8000/logout","user_agent":"PostmanRuntime/7.28.4"}
{"level":"info","msg":"Url Param 'redirect_uri' is missing"}
{"http_method":"GET","http_proto":"HTTP/1.1","http_scheme":"http","level":"info","msg":"request complete","remote_addr":"127.0.0.1:64917","req_id":"LAPTOP-NFG6QTU2/wA0XUpCfvv-000004","resp_bytes_length":0,"resp_elapsed_ms":0,"resp_status":0,"ts":"Sun, 14 Nov 2021 13:59:33 UTC","uri":"http://localhost:8000/logout","user_agent":"PostmanRuntime/7.28.4"}
```
## Профилирование

``` sh
# Профилирование работает, если в командную строку pprof включен параметр "enable=true"

# Память
go tool pprof "http://localhost:8000/debug/pprof/heap?enable=true"

# Процессор
go tool pprof "http://localhost:8000/debug/pprof/profile?seconds=5&enable=true"

#В случаях, когда параметр выключен (enable=false) или отсутствует, сервис вернет:
server response: 406 Not Acceptable
failed to fetch any source profiles
```

## Эндпоинты

### /login

Тест
``` sh
curl --location --request GET 'localhost:8000/login' \
--header 'Authorization: Basic YWRtaW46YWRtaW4='
```
Результат
``` 
HTTP/1.1 200 OK
Content-Type: application/json
Set-Cookie: accessToken=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFkbWluIiwiZXhwIjoxNjM2ODk4OTQ0fQ.OdC9v6YatYuK2zGvL3p1bGx5hEnl3TJNakSEj42QurU; Expires=Sun, 14 Nov 2021 14:09:04 GMT
Set-Cookie: refreshToken=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFkbWluIiwiZXhwIjoxNjM2OTAyNDg0fQ.y1-8bI-lNQ3qwJ_Y1Yt5PbGLak12o1Dq1xmeR3Te0KM; Expires=Sun, 14 Nov 2021 15:08:04 GMT
Date: Sun, 14 Nov 2021 14:08:04 GMT
Content-Length: 0
```

---

### /logout

Тест
``` sh
curl --location --request GET 'localhost:8000/logout' \
--header 'Authorization: Basic YWRtaW46YWRtaW4=' \
--header 'Cookie: accessToken=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFkbWluIiwiZXhwIjoxNjM2ODk5MjYwfQ.kSFRUMLkk-KgdbQO6w3Wb-gkhGyhkiZn57HL0129wrE; refreshToken=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFkbWluIiwiZXhwIjoxNjM2OTAyODAwfQ.k-x_ofoUyDfyZYbHsOI6q3NRy6wacqyLLFfOfvZo45M'
```
Результат
``` 
HTTP/1.1 200 OK
Content-Type: application/json
Set-Cookie: accessToken=; Expires=Sun, 14 Nov 2021 13:14:13 GMT
Set-Cookie: refreshToken=; Expires=Sun, 14 Nov 2021 13:14:13 GMT
Date: Sun, 14 Nov 2021 14:14:13 GMT
Content-Length: 0
```

---

### /i

Tест
``` sh
curl --location --request GET 'localhost:8000/i' \
--header 'Authorization: Basic YWRtaW46YWRtaW4=' \
--header 'Cookie: accessToken=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFkbWluIiwiZXhwIjoxNjM2ODk5MzY1fQ.NDDHu2Q8j4fG7i9YpdBKx7j17r1pJ4uVoV-xaHnWf24; refreshToken=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFkbWluIiwiZXhwIjoxNjM2OTAyOTA1fQ.oA4eyVGM7sIgtKdaAhkr5YYCGXSa_8fjMIrCgYHoCdU'
```
Результат
``` json
{
    "username": "admin"
}
```

---

### /me

Tест
``` sh
curl --location --request GET 'localhost:8000/me' \
--header 'Authorization: Basic YWRtaW46YWRtaW4=' \
--header 'Cookie: accessToken=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFkbWluIiwiZXhwIjoxNjM2ODk5MzY1fQ.NDDHu2Q8j4fG7i9YpdBKx7j17r1pJ4uVoV-xaHnWf24; refreshToken=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFkbWluIiwiZXhwIjoxNjM2OTAyOTA1fQ.oA4eyVGM7sIgtKdaAhkr5YYCGXSa_8fjMIrCgYHoCdU'
```
Результат
``` json
{
    "username": "admin"
}
```