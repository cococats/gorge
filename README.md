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

`say`: Use say to print text. 

`get`: Use get to get a HTTP request.

`post`: Use post to post a HTTP Request

`method`: Use method to get and post a HTTP method such as patch.

`end`: Use end to parse a server. This is unnecessary unless you are just creating a localhost connection.

`auth`: Provides Authorization for your service 

### Example

Example using `get` and `say`:

```go

package main

import (
"github.com/cococats/gorge"
)

func main() {
   get("/")
   say("hello world")
}

```

### Full List of Command Usage

```go
import "GitHub.com/cococats/gorge"

func main() {

say("Hello World")
get("/")
post("localhost:6789")
method("PATCH", "/", "hello world")
end()
auth("<username>", "<password>")
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[Apache 2.0](https://choosealicense.com/licenses/apache-2.0/)
