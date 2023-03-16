# Go-Pointer

For all of you out there, struggling to take a pointer of a constant primitive inline, there is a simple, zero dependency solution for you.

## Problem

It's just not possible to take a pointer of a primitive directly inline. Currently, I only know two solutions to workaround this problem.

### Initialize a variable

```go
foo := func(bar *string) {}

bar := "baz"

foo(&bar)
```

### Use k8s.io/utils/pointer

This solution is possible, but has the disadvantage of downloading the whole k8s Package, just to take a pointer of a primitive. Also, this attempt does not use generics.

```go
import "k8s.io/utils/pointer"

func main() {
    foo := func(bar *string, baz *bool) {}
    foo(pointer.String("baz"), pointer.Bool(true))
}
````

## Solution

It's as simple as that.

```go
import "github.com/freisenhauer/go-pointer/pkg/pointer"

func main() {
    foo := func(bar *string, baz *bool) {}
    foo(pointer.Of("baz"), pointer.Of(true))
}
```
