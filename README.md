UIM API in Go [![GoDoc](https://godoc.org/github.com/xopenapi/uim-api-go2?status.svg)](https://godoc.org/github.com/xopenapi/uim-api-go2) [![Build Status](https://travis-ci.org/xopenapi/uim-api-go2.svg)](https://travis-ci.org/xopenapi/uim-api-go2)
===============
This is the original UIM library for Go created by Norberto Lopez, transferred to a Github organization.

[![Join the chat at https://gitter.im/xopenapi/Lobby](https://badges.gitter.im/go-uim/Lobby.svg)](https://gitter.im/xopenapi/Lobby?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)

This library supports most if not all of the `api.uim.com` REST
calls, as well as the Real-Time Messaging protocol over websocket, in
a fully managed way.




## Changelog

[CHANGELOG.md](https://github.com/xopenapi/uim-api-go2/blob/master/CHANGELOG.md) is available. Please visit it for updates.

## Installing

### *go get*

    $ go get -u github.com/xopenapi/uim-api-go2

## Example

### Getting all groups

```golang
import (
	"fmt"

	"github.com/xopenapi/uim-api-go2"
)

func main() {
	api := uim.New("YOUR_TOKEN_HERE")
	// If you set debugging, it will log all requests to the console
	// Useful when encountering issues
	// uim.New("YOUR_TOKEN_HERE", uim.OptionDebug(true))
	groups, err := api.GetGroups(false)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	for _, group := range groups {
		fmt.Printf("ID: %s, Name: %s\n", group.ID, group.Name)
	}
}
```

### Getting User Information

```golang
import (
    "fmt"

    "github.com/xopenapi/uim-api-go2"
)

func main() {
    api := uim.New("YOUR_TOKEN_HERE")
    user, err := api.GetUserInfo("U023BECGF")
    if err != nil {
	    fmt.Printf("%s\n", err)
	    return
    }
    fmt.Printf("ID: %s, Fullname: %s, Email: %s\n", user.ID, user.Profile.RealName, user.Profile.Email)
}
```

## Minimal RTM usage:

See https://github.com/xopenapi/uim-api-go2/blob/master/examples/websocket/websocket.go


## Minimal EventsAPI usage:

See https://github.com/xopenapi/uim-api-go2/blob/master/examples/eventsapi/events.go


## Contributing

You are more than welcome to contribute to this project.  Fork and
make a Pull Request, or create an Issue if you see any problem.

Before making any Pull Request please run the following:

```
make pr-prep
```

This will check/update code formatting, linting and then run all tests

## License

BSD 2 Clause license
