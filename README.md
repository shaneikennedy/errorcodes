# Errorcodes

A very dumb webserver that returns whatever http status code you pass in the path: i.e /500 will return an HTTP 500

```
❯ curl -v  https://errorcodes-bhzed7jexa-lm.a.run.app/500

> GET /500 HTTP/2
> Host: errorcodes-bhzed7jexa-lm.a.run.app
> user-agent: curl/7.85.0
> accept: */*
>
< HTTP/2 500
```

Hosted at https://errorcodes-bhzed7jexa-lm.a.run.app

Use it for ad-hoc testing some error handling in your web app without writting mocks
