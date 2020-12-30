# id

A lightweight wrapper for generating common identifiers.

```go
package main

import (
    "fmt"

    "github.com/bndw/id"
)

func main() {
    _ = id.Short()              // => "tUn5nCxGR"
    _ = id.ShortPrefix("foo_")  // => "foo_tUn5nCxGR"
    _ = id.UUID()               // => "5c5aa201-e466-4f64-bac8-5cacd39160b5"
    _ = id.UUIDPrefix("foo_")   // => "foo_5c5aa201-e466-4f64-bac8-5cacd39160b5"
    _ = id.Name()               // => "happy_dubinsky"
    _ = id.NamePrefix("foo_")   // => "foo_happy_dubinsky"
}
```

| Function | Description | Example | Source |
| -------- | ----------- | ------- | ------ |
| `id.Short` | Short, fully unique, non-sequential, URL friendly Ids | `tUn5nCxGR` | [teris-io/shortid] |
| `id.ShortPrefix` | Prefixed, short, fully unique, non-sequential, URL friendly Ids | `foo_tUn5nCxGR` | [teris-io/shortid] |
| `id.UUID` | UUID (v4) | `5c5aa201-e466-4f64-bac8-5cacd39160b5`| [google/uuid] |
| `id.UUIDPrefix` | Prefixed UUID (v4) | `foo_5c5aa201-e466-4f64-bac8-5cacd39160b5`| [google/uuid] |
| `id.Name` | Docker-style names | `happy_dubinsky`| [moby/namegenerator] |
| `id.NamePrefix` | Docker-style names | `foo_happy_dubinsky`| [moby/namegenerator] |

[teris-io/shortid]: https://github.com/teris-io/shortid
[google/uuid]: https://github.com/google/uuid
[moby/namegenerator]: https://github.com/moby/moby/tree/master/pkg/namesgenerator
