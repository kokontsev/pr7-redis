# Фактический прогон curl

```text
$ curl -sS -i --max-time 15 -X GET http://127.0.0.1:18007/set?key=practice7&value=hello
HTTP/1.1 200 OK
Date: Sat, 31 Oct 2026 08:43:57 GMT
Content-Length: 30
Content-Type: text/plain; charset=utf-8

OK: practice7=hello (TTL 10s)


$ curl -sS -i --max-time 15 -X GET http://127.0.0.1:18007/get?key=practice7
HTTP/1.1 200 OK
Date: Sat, 31 Oct 2026 08:43:57 GMT
Content-Length: 23
Content-Type: text/plain; charset=utf-8

VALUE: practice7=hello


$ curl -sS -i --max-time 15 -X GET http://127.0.0.1:18007/ttl?key=practice7
HTTP/1.1 200 OK
Date: Sat, 31 Oct 2026 08:43:57 GMT
Content-Length: 23
Content-Type: text/plain; charset=utf-8

TTL for practice7: 10s


$ curl -sS -i --max-time 15 -X GET http://127.0.0.1:18007/get?key=practice7
HTTP/1.1 404 Not Found
Content-Type: text/plain; charset=utf-8
X-Content-Type-Options: nosniff
Date: Sat, 31 Oct 2026 08:44:08 GMT
Content-Length: 14

key not found


$ curl -sS -i --max-time 15 -X GET http://127.0.0.1:18007/get
HTTP/1.1 400 Bad Request
Content-Type: text/plain; charset=utf-8
X-Content-Type-Options: nosniff
Date: Sat, 31 Oct 2026 08:44:08 GMT
Content-Length: 13

key required

```
