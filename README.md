# GO-Utils The Utility package for golang

## install 
`go get github.com/jerbe/go-utils`

## how to use

```go
import (
	"fmt"
	
    utils "github.com/jerbe/go-utils"
)


func main() {
	// print the local host name
    fmt.Println(utils.GetHostname())   
	// >> jerbe-mac
	
	// print local IP
	fmt.Println(utils.GetLocalIPv4())
	// >> 192.168.31.100
	
	...
}

```