# Simple key-value parser

## Install

`go get github.com/wirekang/cfg`

## Usage

```
#file.txt
key1: value
key2 1, 2, 3, 4
key3: true
```

```golang
config,err := cfg.LoadFile("file.txt")
value, ok := config.Get("key1")
fmt.Println(ok) // true
fmt.Println(value) // "value"

value, ok = config.Get("key2")
arr, err := value.IntArray()
fmt.Println(arr) // [1 2 3 4]

_,ok = config.Get("non existing key")
fmt.Println(ok) // false
```
