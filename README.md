# number2word

A simple package to convert numbers to words in Golang.

[![](https://github.com/istiyakamin/number2word/actions/workflows/go.yml/badge.svg "test status")](https://github.com/istiyakamin/number2word/actions) 
![test status](https://img.shields.io/github/go-mod/go-version/istiyakamin/number2word?label=Go%20Version)
[![MIT license](https://img.shields.io/badge/license-MIT-brightgreen.svg)](https://opensource.org/licenses/MIT)
[![Go.Dev reference](https://img.shields.io/badge/go.dev-reference-blue?logo=go&logoColor=white)](https://pkg.go.dev/github.com/istiyakamin/number2word@v1.0.1)
```
go get -u github.com/istiyakamin/number2word
```
Import Packages 
```
import (
	...
	"github.com/istiyakamin/number2word"
  	...
)
```

# Coded Example 

```
package main

import (
	"fmt"
	"github.com/istiyakamin/number2word"
)

func main(){
	str := number2word.Convert(102547)
	fmt.Println(str)
}
```

# Example

```golang

str := number2word.Convert(17) // outputs "seventeen"

str := number2word.Convert(1024) // outputs "one thousand twenty four"

str := number2word.Convert(-123) // outputs "minus one hundred twenty three"
Convert number with " and " between number groups:

str := number2word.ConvertAnd(514) // outputs "five hundred and fourteen"

str := number2word.ConvertAnd(123) // outputs "one hundred and twenty three"
```

# Contributing

In case you notice a bug, please open an issue mentioning the input that has caused an incorrect conversion.

## Go to 
(watch on pkg.go.dev)[https://pkg.go.dev/github.com/istiyakamin/number2word]
