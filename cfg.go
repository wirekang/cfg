//Package cfg parses simple config format file or string.
package cfg

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

// Config is map of key-Value.
type Config map[string]Value

// Value is value of a key.
type Value string

// Load parses config from string by following format:
//
// - - -
//
// key= string value
//
// key2= 32
//
//
// key 3 =this,is, string, array
//
// key 4 = 2, 4, 6,8,10 , 12
//
// key5 = 2003-4-15
//
// #comment
func Load(str string) (Config, error) {
	configs := make(Config)
	lines := strings.Split(str, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		if line[0] == '#' {
			continue
		}
		t := strings.Split(line, "=")
		if len(t) != 2 {
			return nil, fmt.Errorf("\"%s\" is not a config", line)
		}
		configs[strings.TrimSpace(t[0])] = Value(strings.TrimSpace(t[1]))
	}
	return configs, nil
}

// LoadFile calls Load() with the files content.
func LoadFile(filename string) (Config, error) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return Load(string(bytes))
}

// Find returns value with the key.
func (c Config) Find(key string) Value {
	return c[key]
}

// IsExist returns true of value with the key is exists.
func (c Config) IsExist(key string) bool {
	return c[key] != ""
}

func (v Value) String() string {
	return string(v)
}

func (v Value) Int() (int, error) {
	return strconv.Atoi(string(v))
}

func (v Value) Float() (float64, error) {
	return strconv.ParseFloat(string(v), 64)
}

func (v Value) Date() (time.Time, error) {
	return time.Parse("2006-1-2", string(v))
}

func (v Value) StringArray() []string {
	arr := strings.Split(string(v), ",")
	for i, s := range arr {
		arr[i] = strings.TrimSpace(s)
	}
	return arr
}

func (v Value) IntArray() ([]int, error) {
	sa := strings.Split(string(v), ",")
	ia := make([]int, len(sa))
	for i, s := range sa {
		in, err := strconv.Atoi(s)
		if err != nil {
			return nil, err
		}
		ia[i] = in
	}
	return ia, nil
}

func (v Value) Bool() (b bool) {
	defer func() {
		r := recover()
		if r != nil {
			b = false
		}
	}()
	return strings.ToLower(string(v)) == "true"
}
