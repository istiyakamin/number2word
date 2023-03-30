# number2word

An simple package to convert numbers to words in Golang for South Asian numbering system. e.g. Two crore four lakh.

```
go get -u github.com/istiyakaminsanto/number2word
```
Import Packages 

```
import (
	...
	"github.com/istiyakaminsanto/number2word"
  ...
)
```

# Coded Example 

```
package main

import (
	"fmt"
	"github.com/istiyakaminsanto/number2word"
)

func main(){
	str := number2word.Convert(102547)
	fmt.Println(str)
}
```

# Example

```golang
1          ->  one
12         ->  twelve
123        ->  one hundred and twenty three
1234       ->  one thousand two hundred and thirty four
12345      ->  twelve thousand three hundred and forty five
123456     ->  one lakh twenty three thousand four hundred and fifty six


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
(watch on pkg.go.dev)[https://pkg.go.dev/github.com/istiyakaminsanto/number2word]
