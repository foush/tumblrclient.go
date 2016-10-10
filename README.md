# tumblrapiclient

This is a concrete implementation of the `ClientInterface` defined in the [Tumblr API](https://github.com/foush/tumblrapi) library.

Install by running `go get github.com/foush/tumblrapiclient`

You can instantiate a client by running

```
client := tumblrapi.NewClient(
        "CONSUMER KEY",
        "CONSUMER SECRET",
    )
    // or
    client := tumblr_go.NewClientWithToken(
        "CONSUMER KEY",
        "CONSUMER SECRET",
        "USER TOKEN",
        "USER TOKEN SECRET",
    )
```

From there you can use the convenience methods.
