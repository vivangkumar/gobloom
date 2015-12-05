# gobloom

Simple Bloom Filter implementation using [murmur3]([http://code.google.com/p/smhasher/wiki/MurmurHash3)

## Getting started

`go get "github.com/vivangkumar/gobloom"`

## Example

You can use gobloom like this

```go
package main

import (
	"github.com/vivangkumar/gobloom"
	"fmt"
)

func main() {
	bf := gobloom.NewBloomFilter(10, 3)
	bf.Add("one", "two")
	fmt.Print(bf.Contains("one"))
}
```

## API

The API (for now) is pretty simple.

- `NewBloomFilter(size, hashFuncs int)`: Returns a reference to a new BloomFilter initialized with the `size`
(number of bits) and `HashFuncs` (number of times to Hash)

- `(bf *BloomFilter) Add(items ...string)`: Add new items to the filter.

- `(bf *BloomFilter) Contains(item string) bool`: Returns a `bool` value depending on existence in the filter.
Increasing the size of the filter can reduce false positives. However, this is a pretty basic implementation, so be
aware.

## Tests

Run tests using `go test`

(Just simple cases, I'll add more later)

## License

This code is licenced under the [MIT Licence](http://opensource.org/licenses/mit-license.php)
