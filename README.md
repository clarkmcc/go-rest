# go-rest
A rest client pulled from the Kubernetes client-go package. The client is in it's basic form and can handle any marshaller instead of using the kubernetes GVK decoders (i.e. it can be used on apis that don't support the kubernetes GVK standards).

## Installation
```bash
$ go get github.com/clarkmcc/go-rest
```

## Usage
Create a new client that unmarshalls json when using the `Into` method, and has a rate limiter with 1QPS and 5 burst using the following
```go
client, err := GenericClientFor(&Config{
    Host: "http://127.0.0.1:8080",
    ContentConfig: ContentConfig{
        Marshaller: marshaller.NewJsonMarshaller(),
    },
    APIPath: "/",
    RateLimiter: flowcontrol.NewTokenBucketRateLimiter(1, 5),
})
```

Make a request using the client
```go
resp := map[string]interface{}{}
err = client.
    Get().
    Prefix("/").
    Do(context.Background()).
    Into(&resp)
```

If you want to return the raw bytes and deal with the response yourself
```go
bytes, err := client.
    Get().
    Prefix("/").
    Do(context.Background()).
    Raw()
```