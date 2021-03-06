# Strings-go

---

Strings-go is a string extension library for Golang.

## License

Apache License, Version 2.0

## Implements

### StringBuilder

* Append
* AppendNil
* AppendRune
* AppendRuneArray
* AppendString
* AppendObject
* AppendCodePoint
* Delete
* DeleteCharAt
* Replace
* SetCharAt
* Insert
* IndexOf
* IndexOfFrom
* LastIndexOf
* LastIndexOfFrom
* Reverse


## Usage
### Installation
```shell
go get -u github.com/guangzhou-meta/go-lib/strings
```

### Demo
#### StringBuilder

```go
package main

import (
	"fmt"
)

import (
	strBuilder "github.com/guangzhou-meta/go-lib/strings"
)

func main() {
	stringBuilder := strBuilder.NewStringBuilder()
	stringBuilder.Append("test")
	fmt.Println(stringBuilder)

	stringBuilder = strBuilder.NewStringBuilderStr("stringBuilder")
	fmt.Println(stringBuilder)
}
```
