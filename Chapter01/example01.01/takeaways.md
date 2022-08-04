## Sectioon 1 of the code
```
package main
```
package declaration: All go codes must start with main

If you want to run code directly you have to name it main otherwise you have to use it as a library and import it in other Go codes

All Go files in the same directory are
considered part of the same package, which means all the files must have the same
package name.

## Sectio 2

```
import (
"errors"
"fmt"
"log"
"math/rand"
"strconv"
"time"
)
```

the above packages are all from Go's standard library

## Section 3

```
  var helloList = []string{
	"Hello, world",
	"Καλημέρα κόσμε",
	"こんにちは世界",
	"سلام دنیا",
	"Привет, мир",
}
```

Global variable declaration: A list of strings with initial data. 

In Go strings support multi-byte UTF-8 encoding