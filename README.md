# Errorcodes

A very dumb webserver that returns whatever http status code you pass in the path: i.e /500 will return an HTTP 500

```
❯ curl -v  http://localhost:8080/500

> GET /500 HTTP/2
> Host: localhost:8080
> user-agent: curl/7.85.0
> accept: */*
>
< HTTP/2 500
```

Use it for ad-hoc testing some error handling in your web app without writting mocks
