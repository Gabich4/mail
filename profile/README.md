# profile



## Логирование

Используется [logrus](https://github.com/sirupsen/logrus) обращения к сервису пишутся через него в stdout.
Так-же через него выполняется логирование различных ситуаций в работе сервиса.
``` json
{"level":"info","msg":"Service profile started. localhost:3000"}
{"http_method":"GET","http_proto":"HTTP/1.1","http_scheme":"http","level":"info","msg":"request started","remote_addr":"127.0.0.1:54952","req_id":"LAPTOP-NFG6QTU2/krzAa7jhF2-000001","ts":"Sat, 13 Nov 2021 10:26:47 UTC","uri":"http://localhost:3000/i","user_agent":"PostmanRuntime/7.28.4"}
{"level":"error","msg":"Get \"http://127.0.0.1:8000/me\": dial tcp 127.0.0.1:8000: connectex: No connection could be made because the target machine actively refused it."}
{"http_method":"GET","http_proto":"HTTP/1.1","http_scheme":"http","level":"info","msg":"request complete","remote_addr":"127.0.0.1:54952","req_id":"LAPTOP-NFG6QTU2/krzAa7jhF2-000001","resp_bytes_length":144,"resp_elapsed_ms":2052.5483,"resp_status":400,"ts":"Sat, 13 Nov 2021 10:26:47 UTC","uri":"http://localhost:3000/i","user_agent":"PostmanRuntime/7.28.4"}
{"http_method":"GET","http_proto":"HTTP/1.1","http_scheme":"http","level":"info","msg":"request started","remote_addr":"127.0.0.1:54952","req_id":"LAPTOP-NFG6QTU2/krzAa7jhF2-000002","ts":"Sat, 13 Nov 2021 10:26:51 UTC","uri":"http://localhost:3000/receivers","user_agent":"PostmanRuntime/7.28.4"}
{"level":"info","msg":"mailing list is empty"}
{"http_method":"GET","http_proto":"HTTP/1.1","http_scheme":"http","level":"info","msg":"request complete","remote_addr":"127.0.0.1:54952","req_id":"LAPTOP-NFG6QTU2/krzAa7jhF2-000002","resp_bytes_length":22,"resp_elapsed_ms":0,"resp_status":404,"ts":"Sat, 13 Nov 2021 10:26:51 UTC","uri":"http://localhost:3000/receivers","user_agent":"PostmanRuntime/7.28.4"}
{"http_method":"POST","http_proto":"HTTP/1.1","http_scheme":"http","level":"info","msg":"request started","remote_addr":"127.0.0.1:54952","req_id":"LAPTOP-NFG6QTU2/krzAa7jhF2-000003","ts":"Sat, 13 Nov 2021 10:26:54 UTC","uri":"http://localhost:3000/receivers/user3","user_agent":"PostmanRuntime/7.28.4"}
{"http_method":"POST","http_proto":"HTTP/1.1","http_scheme":"http","level":"info","msg":"request complete","remote_addr":"127.0.0.1:54952","req_id":"LAPTOP-NFG6QTU2/krzAa7jhF2-000003","resp_bytes_length":0,"resp_elapsed_ms":0.7174,"resp_status":0,"ts":"Sat, 13 Nov 2021 10:26:54 UTC","uri":"http://localhost:3000/receivers/user3","user_agent":"PostmanRuntime/7.28.4"}
{"http_method":"GET","http_proto":"HTTP/1.1","http_scheme":"http","level":"info","msg":"request started","remote_addr":"127.0.0.1:54952","req_id":"LAPTOP-NFG6QTU2/krzAa7jhF2-000004","ts":"Sat, 13 Nov 2021 10:26:59 UTC","uri":"http://localhost:3000/receivers/user3","user_agent":"PostmanRuntime/7.28.4"}
{"http_method":"GET","http_proto":"HTTP/1.1","http_scheme":"http","level":"info","msg":"request complete","remote_addr":"127.0.0.1:54952","req_id":"LAPTOP-NFG6QTU2/krzAa7jhF2-000004","resp_bytes_length":9,"resp_elapsed_ms":0.0872,"resp_status":200,"ts":"Sat, 13 Nov 2021 10:26:59 UTC","uri":"http://localhost:3000/receivers/user3","user_agent":"PostmanRuntime/7.28.4"}
{"http_method":"DELETE","http_proto":"HTTP/1.1","http_scheme":"http","level":"info","msg":"request started","remote_addr":"127.0.0.1:54952","req_id":"LAPTOP-NFG6QTU2/krzAa7jhF2-000005","ts":"Sat, 13 Nov 2021 10:27:04 UTC","uri":"http://localhost:3000/receivers/user3","user_agent":"PostmanRuntime/7.28.4"}
{"http_method":"DELETE","http_proto":"HTTP/1.1","http_scheme":"http","level":"info","msg":"request complete","remote_addr":"127.0.0.1:54952","req_id":"LAPTOP-NFG6QTU2/krzAa7jhF2-000005","resp_bytes_length":0,"resp_elapsed_ms":0,"resp_status":0,"ts":"Sat, 13 Nov 2021 10:27:04 UTC","uri":"http://localhost:3000/receivers/user3","user_agent":"PostmanRuntime/7.28.4"}
{"http_method":"DELETE","http_proto":"HTTP/1.1","http_scheme":"http","level":"info","msg":"request started","remote_addr":"127.0.0.1:54952","req_id":"LAPTOP-NFG6QTU2/krzAa7jhF2-000006","ts":"Sat, 13 Nov 2021 10:27:44 UTC","uri":"http://localhost:3000/receivers/user3","user_agent":"PostmanRuntime/7.28.4"}
{"level":"info","msg":"username user3 not found"}
{"http_method":"DELETE","http_proto":"HTTP/1.1","http_scheme":"http","level":"info","msg":"request complete","remote_addr":"127.0.0.1:54952","req_id":"LAPTOP-NFG6QTU2/krzAa7jhF2-000006","resp_bytes_length":25,"resp_elapsed_ms":0.0828,"resp_status":400,"ts":"Sat, 13 Nov 2021 10:27:44 UTC","uri":"http://localhost:3000/receivers/user3","user_agent":"PostmanRuntime/7.28.4"}
{"http_method":"POST","http_proto":"HTTP/1.1","http_scheme":"http","level":"info","msg":"request started","remote_addr":"127.0.0.1:57843","req_id":"MB444705/Nhi3GlpHxO-000002","ts":"Mon, 15 Nov 2021 06:30:55 UTC","uri":"http://127.0.0.1:3000/status","user_agent":"Go-http-client/1.1"}
{"level":"info","msg":"received SendRequestStatus: {[test1 test2] test message success}"}
{"http_method":"POST","http_proto":"HTTP/1.1","http_scheme":"http","level":"info","msg":"request complete","remote_addr":"127.0.0.1:57843","req_id":"MB444705/Nhi3GlpHxO-000002","resp_bytes_length":0,"resp_elapsed_ms":0.058762,"resp_status":0,"ts":"Mon, 15 Nov 2021 06:30:55 UTC","uri":"http://127.0.0.1:3000/status","user_agent":"Go-http-client/1.1"}
```
## Профилирование

``` sh
# Профилирование работает, если в командную строку pprof включен параметр "enable=true"

# Память
go tool pprof "http://localhost:3000/debug/pprof/heap?enable=true"

# Процессор
go tool pprof "http://localhost:3000/debug/pprof/profile?seconds=5&enable=true"

#В случаях, когда параметр выключен (enable=false) или отсутствует, сервис вернет:
server response: 406 Not Acceptable
failed to fetch any source profiles
```

## Эндпоинты

### /i

Tест
``` sh
curl --location --request GET 'localhost:3000/i' \
--header 'Cookie: accessToken=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InVzZXJuYW1lMSIsImV4cCI6MTYzNjY5OTk4Nn0.1RB-PUxqEg8ZoLhE2tofbHiYaF58PHfI8NNtgYQ43lk; refreshToken=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InVzZXJuYW1lMSIsImV4cCI6MTYzNjcwMzUyNn0.OKeWy9SvRxw-UANDgDgyT9LhosiXZ7rLI7rBYL_8v74'
```
Результат
``` json
{
    "username": "username1"
}
```

---

### /upload_template

Tест
``` sh
curl --location --request POST 'localhost:8000/upload_template' \
--header 'Content-Type: text/html' \
--data-raw '<!doctype html>
<html lang='\''ru'\''>
    <head>
        <meta charset='\''utf-8'\''>
        <title>{{.Title}}</title>
    </head>
    
    <body>
        <header>
            <h1>{{.Title}}</h1>
        </header>
        
        <main>
            <h2>Добрый день {{.FIO}}</h2>
            <p>Материалы для ознакомления: <a href={{.ContentURL}}>{{.ContentTitle}}</a></p>
        </main>
    </body>
</html>'  
```
Результат
``` json
[
    "{{.Title}}",
    "{{.FIO}}",
    "{{.ContentURL}}",
    "{{.ContentTitle}}"
]
```

---

### /status
Получение статуса отправки писем из сервиса mailsender
``` sh
curl --location --request POST 'http://127.0.0.1:3000/status' \
--header 'Content-Type: application/json' \
--data-raw '{
    "Receivers": [
        "test1",
        "test2"
    ],
    "Message": "test message",
    "Status": "success"
}'
```
Результат
```
200 OK
```

---

### /receivers/{user_id}

#### CREATE
``` sh
curl --location --request POST 'localhost:3000/receivers/user3' \
--header 'Content-Type: text/plain' \
--data-raw 'z,x,c,v'     
```
При успехе возвращается: **200 OK**  
При ошибке возвращается: **username user3 already exists**

#### READ
``` sh
curl --location --request GET 'localhost:3000/receivers/user3'     
```
При успехе возвращается: **"z,x,c,v"**  
При ошибке возвращается: **username user3 not found**

#### UPDATE
``` sh
curl --location --request PUT 'localhost:3000/receivers/user3' \
--header 'Content-Type: text/plain' \
--data-raw 'z,x,c,v,b,n,m'    
```
При успехе возвращается: **200 OK**  
При ошибке возвращается: **username user3 not found**

#### DELETE
``` sh
curl --location --request DELETE 'localhost:3000/receivers/user3'    
```
При успехе возвращается: **200 OK**  
При ошибке возвращается: **username user3 not found**

#### LIST
``` sh
curl --location --request GET 'localhost:3000/receivers'    
```

``` json
# При успехе возвращается
[
    {
        "username": "user3",
        "receivers": "z,x,c,v"
    },
    {
        "username": "user1",
        "receivers": "q,w,e,r,t"
    }
]
```
При ошибке возвращается: **mailing list is empty**

