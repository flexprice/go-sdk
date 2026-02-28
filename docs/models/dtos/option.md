## Options

### WithServerURL

WithServerURL allows providing an alternative server URL.

```go
dtos.WithServerURL("http://api.example.com")
```

## WithTemplatedServerURL

WithTemplatedServerURL allows providing an alternative server URL with templated parameters.

```go
dtos.WithTemplatedServerURL("http://{host}:{port}", map[string]string{
    "host": "api.example.com",
    "port": "8080",
})
```

### WithRetries

WithRetries allows customizing the default retry configuration. Only usable with methods that mention they support retries.

```go
dtos.WithRetries(retry.Config{
    Strategy: "backoff",
    Backoff: retry.BackoffStrategy{
        InitialInterval: 500 * time.Millisecond,
        MaxInterval: 60 * time.Second,
        Exponent: 1.5,
        MaxElapsedTime: 5 * time.Minute,
    },
    RetryConnectionErrors: true,
})
```