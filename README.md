Lyric API written in Golang [![GoDoc](https://godoc.org/github.com/rhnvrm/lyric-api-go?status.svg)](https://godoc.org/github.com/rhnvrm/lyric-api-go) 
===============

This library provides an API to search for lyrics from various providers.

## Installing

### go get
```sh
    $ go get github.com/rhnvrm/lyric-api-go
```

## Usage Example

More examples can be found in the examples directory.

```go
package main

import (
	"fmt"

	"github.com/rhnvrm/lyric-api-go"
)

func main() {
	var (
		artist = "John Lennon"
		song   = "Imagine"
	)

	l := lyrics.New()
	lyric, err := l.Search(artist, song)

	if err != nil {
		fmt.Printf("Lyrics for %v-%v were not found", artist, song)
	}
	fmt.Println(lyric)
}
```


## Contributing

You are more than welcome to contribute to this project.  Fork and
make a Pull Request, or create an Issue if you see any problem.
