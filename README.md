![gorge](gorge.jpg)

# gorge

gorge is a web development DSL for go. It is simple and easy-to-use. 

## Installation

Use the `go get` CLI command to install gorge.

```bash
go get -u github.com/cococats/gorge
```

## Usage

### Connections

All functions will post to `localhost:6789` unless otherwise specified through `end`

### Functions

`say`: Use say to print text. The URL will parse to: `localhost:6789`.

`get`: Use get to get a HTTP request. The URL will parse to: `localhost:6789/get`.

`put`: Use put to put a HTTP request. The URL will parse to: `localhost:6789/put`.

`post`: Use post to post a HTTP Request.

`method`: Use method to get and post a HTTP method such as `patch`.

`end`: Use end to start a connection to `localhost:6789`. This is required for your application to work!

 

### Example

Example using `get` and `say`:

```go

package main

import (
"github.com/cococats/gorge"
)

func main() {
   get("http://localhost:6789")
   say("hello world")
}

```

### Full List of Command Usage

```go
import "GitHub.com/cococats/gorge"

func main() {

say("Hello World")
get("/")
post("localhost:6789", "img/jpg", nil)
method("PATCH", "http://localhost:6789", nil)
end()
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[Apache 2.0](https://choosealicense.com/licenses/apache-2.0/)
