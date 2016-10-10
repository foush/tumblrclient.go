# Tumblr API Client Library

This is a concrete implementation of the `ClientInterface` defined in the [Tumblr API](https://github.com/foush/tumblr.go) library.

This project utilizes an external OAuth1 library you can find at [https://github.com/dghubble/oauth1](github.com/dghubble/oauth1) 

Install by running `go get github.com/foush/tumblrclient.go`

You can instantiate a client by running

```
client := tumblrclient.NewClient(
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
