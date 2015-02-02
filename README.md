# Golang binding for wordexp(3), fnmatch(3)

[API reference](http://godoc.org/github.com/wtnb75/go-wordexp)

## Usage

- go get github.com/wtnb75/go-wordexp

### WordExp

```go
package main

import (
	"github.com/wtnb75/go-wordexp"
	"log"
	"os"
)

func main() {
	for _, v := range os.Args[1:] {
		res, err := wordexp.WordExp(v, wordexp.WRDE_UNDEF|wordexp.WRDE_NOCMD)
		log.Println(v, err, res)
	}
}
```

see wordexp(3) man page.

### FnMatch

```go
package main

import (
	"github.com/wtnb75/go-wordexp"
	"log"
)

func main() {
	// prints "YYYY/MM/DD hh:mm:ss fnmatch: String does not match"
	if err := wordexp.FnMatch("pattern", "target", 0); err != nil {
		log.Println("fnmatch:", err)
	}
}
```

see fnmatch(3) man page.
