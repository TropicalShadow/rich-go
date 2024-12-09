# rich-go [![Build Status](https://travis-ci.org/tropicalshadow/rich-go.svg?branch=master)](https://travis-ci.org/tropicalshadow/rich-go)

An implementation of Discord's rich presence in Golang for Linux, macOS and Windows

## Installation

Install `github.com/tropicalshadow/rich-go`:

```
$ go get github.com/tropicalshadow/rich-go
```

## Usage

First of all import rich-go
```golang
import "github.com/tropicalshadow/rich-go/client"
```

then login by sending the first handshake
```golang
rpcClient := client.NewClient()
err := rpcClient.Login("DISCORD_APP_ID")
if err != nil {
	panic(err)
}
```

and you can set the Rich Presence activity (parameters can be found :
```golang
err = rpcClient.SetActivity(client.Activity{
	State:      "Heyy!!!",
	Details:    "I'm running on rich-go :)",
	LargeImage: "largeimageid",
	LargeText:  "This is the large image :D",
	SmallImage: "smallimageid",
	SmallText:  "And this is the small image",
	Party: &client.Party{
		ID:         "-1",
		Players:    15,
		MaxPlayers: 24,
	},
	Timestamps: &client.Timestamps{
		Start: time.Now(),
	},
})

if err != nil {
	panic(err)
}
```

More details in the [example](https://github.com/ananagame/rich-go/blob/master/example/main.go)

## Contributing

1. Fork it (https://github.com/tropicalshadow/rich-go/fork)
2. Create your feature branch (git checkout -b my-new-feature)
3. Commit your changes (git commit -am 'Add some feature')
4. Push to the branch (git push origin my-new-feature)
5. Create a new Pull Request

## Contributors

# Upstream Contributors

- [hugolgst](https://github.com/hugolgst) - creator, maintainer
- [donovansolms](https://github.com/donovansolms) - contributor
- [heroslender](https://github.com/heroslender) - contributor

# Fork Contributors

- [tropicalshadow](https://github.com/TropicalShadow) - contributor