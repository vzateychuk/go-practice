# Data-puller

The program pulls different data feeds from the web and compares the content against a
search term. The content that matches is then displayed in the terminal window.

## Application architecture
![Application architecture](docs/Go%20in%20Action%20-%20puller%20-%20architecture.png)

## Application structure
```
- data
    data.json -- Contains a list of data feeds
- matcher
    rss.go -- Matcher for searching rss feeds
- search
    default.go -- Default matcher for searching data
    feed.go -- Support for reading the json data file
    match.go -- Interface support for using different matchers
    search.go -- Main program logic for performing search
main.go -- Programs entry point
```

## Links
[Inspired by materials of 'Go in action' book](https://www.manning.com/books/go-in-action)
