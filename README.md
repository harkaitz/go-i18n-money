# INTERNATIONALIZABLE MONETARY VALUE PARSING IN GO

This library provides a function "MoneyParse" that parses monetary
values using the *spanish convention*.

By this convention all commas (,) dots (.) and apostrophes (') can be
used to separate thousands and decimals. The distintion relies on the
amount of digits after the separator. If there are two or less digits
after the separator, it is considered a decimal separator, otherwise
it is considered a thousands separator.

Examples:

    1.23€  ->  123    cents
    1.234€ ->  123400 cents
    1,234€ ->  123400 cents
    1'234€ ->  123400 cents
    €1'23  ->  123    cents
    €1'6   ->  160    cents
    €1,02  ->  102    cents

## TODO

Allow parenthesis to indicate negative values. Example:

    (1.23€) -> -123 cents

## Go documentation

    package i18nm // import "github.com/harkaitz/go-i18n-money"
    
    type Cents uint64
        func ParseMoney(s string) (value Cents, negative bool, err error)
        func Value(a, b int) Cents

## Go programs

    Usage: parse-monetary [-a] -- <money>
    
    Parse a monetary value from the command line arguments. The value
    is expected to be a string in the spanish format, with a comma or
    a dot as the decimal or thousands separator.
    
    The return value is always in cents, if "-a" is set it prints
    the absolute value.
    
    Copyright (c) 2023 - Harkaitz Agirre - MIT License

## Collaborating

For making bug reports, feature requests and donations visit
one of the following links:

1. [gemini://harkadev.com/oss/](gemini://harkadev.com/oss/)
2. [https://harkadev.com/oss/](https://harkadev.com/oss/)
