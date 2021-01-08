**Notice** This is just a fork for myself to add proxy from enviorment. The all code is from https://gitlab.com/toby3d/telegraph

# GoLang bindings for the Telegraph API [![discord](https://discordapp.com/api/guilds/208605007744860163/widget.png)](https://discord.gg/QJ8z5BN)
> This project is just to provide a wrapper around the API without any additional features.

All methods and types available and this library (possibly) is ready for use in production. Yaay!

## Start using telegraph
Download and install it:
`$ go get -u gitlab.com/toby3d/telegraph`

Import it in your code:
`import "gitlab.com/toby3d/telegraph"`

Use Proxy
`telegraph.SetSocksDialer("localhost:7890")`
`telegraph.SetHttpDialer("username:password@localhost:9050")`


## Examples
See [GoDoc examples section](https://godoc.org/gitlab.com/toby3d/telegraph#pkg-examples) or check [example_test.go](/example_test.go).

## Need help?
- [Open new issue](https://gitlab.com/toby3d/telegraph/issues/new)
- [Discuss in Discord](https://discord.gg/QJ8z5BN)