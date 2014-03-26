# go-twitch

[![travis-ci status](https://api.travis-ci.org/mrshankly/go-twitch.png)](https://travis-ci.org/mrshankly/go-twitch)

Go library for accessing the [Twitch-API](https://github.com/justintv/Twitch-API).

**This is still a work in progress.**

Check the progress [here](https://github.com/mrshankly/go-twitch/issues/1).

## Usage

To install `go-twitch` run the command:

```bash
$ go get github.com/mrshankly/go-twitch/twitch
```

Here's an example program that gets the top 10 twitch games:

```go
package main

import (
	"fmt"
	"github.com/mrshankly/go-twitch/twitch"
	"log"
	"net/http"
)

func main() {
	client := twitch.NewClient(&http.Client{})
	opt := &twitch.ListOptions{
		Limit:  10,
		Offset: 0,
	}

	games, err := client.Games.Top(opt)
	if err != nil {
		log.Fatal(err)
	}

	for i, s := range games.Top {
		fmt.Printf("%d - %s (%d)\n", i+1, s.Game.Name, s.Viewers)
	}
}
```

### Authentication

**TODO**

## License

All files under this repository fall under the MIT License (see the file LICENSE).
