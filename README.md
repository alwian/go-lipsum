# lipsum

Lipsum allows you to retrieve Lorem Ipsum text. It makes use of the https://lispum.com API to do this.

## Install

```go
import "github.com/alwian/go-lipsum/pkg/lipsum"
```

## Usage

```go
// Generate bytes
text, err := lipsum.Bytes(100)

// Generrate words
text, err := lipsum.Words(100)

// Generate paragraphs
text, err := lipsum.Paragraphs(100)
```
