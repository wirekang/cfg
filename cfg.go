//Package cfg parses line separated key-value format file or string.
package cfg

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

type Config struct {
	m map[string]Value
}

type Value string

var ErrWrongFormat = fmt.Errorf("line parsing failed")

// Load parses config from string
func Load(str string) (rst Config, err error) {
	lines := strings.Split(str, "\n")
	rst.m = make(map[string]Value, len(lines))
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		if line[0] == '#' {
			continue
		}
		t := strings.Split(line, ":")
		if len(t) != 2 {
			err = ErrWrongFormat
			return
		}
		rst.m[strings.TrimSpace(t[0])] = Value(strings.TrimSpace(t[1]))
	}
	return
}

// LoadFile calls Load() with the files content.
func LoadFile(filename string) (rst Config, err error) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}
	return Load(string(bytes))
}

// Find returns value with the key.
func (c Config) Find(key string) Value {
	return c.m[key]
}

// IsExist returns true of value with the key is exists.
func (c Config) IsExist(key string) bool {
	return c.m[key] != ""
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

// Bool returns true if given value is "true"(case insensitive) or 1.
func (v Value) Bool() (b bool) {
	defer func() {
		r := recover()
		if r != nil {
			b = false
		}
	}()
	return v == "1" || strings.ToLower(string(v)) == "true"
}
