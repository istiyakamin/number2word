# number2word

An simple module to convert numbers to words in Golang for South Asian numbering system. e.g. Two crore four lakh.

# Example

```golang
1          ->  one
12         ->  twelve
123        ->  one hundred and twenty three
1234       ->  one thousand two hundred and thirty four
12345      ->  twelve thousand three hundred and forty five
123456     ->  one lakh twenty three thousand four hundred and fifty six


str := num2words.Convert(17) // outputs "seventeen"

str := num2words.Convert(1024) // outputs "one thousand twenty four"

str := num2words.Convert(-123) // outputs "minus one hundred twenty three"
Convert number with " and " between number groups:

str := num2words.ConvertAnd(514) // outputs "five hundred and fourteen"

str := num2words.ConvertAnd(123) // outputs "one hundred and twenty three"
```

# Contributing

In case you notice a bug, please open an issue mentioning the input that has caused an incorrect conversion.