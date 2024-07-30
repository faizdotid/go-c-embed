# Go C Embed

Embedding c languange in golang

## Example
We will try to use the c library `<math.h>`, below is a little sample code

```go
package main

// #include <math.h>
import "C"

func main() {
    // Trying using function log from math.h library
    result := C.log(1)


	println(result)
}
```