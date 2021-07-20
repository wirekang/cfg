package cfg

import (
	"testing"
	"time"
)

func TestConfig(t *testing.T) {
	con := "num: 3\n#hello\ntitle : The TITLE  \ntags: javascript, golang , c sharp\npi : 3.14\n\ndate: 2018-1-02\nflag : true"
	c, err := Load(con)
	if err != nil {
		t.Fatal(err)
	}

	v, ok := c.Get("num")
	i, err := v.Int()
	if !ok || err != nil || i != 3 {
		t.FailNow()
	}
	v, ok = c.Get("title")
	s := v.String()
	if !ok || s != "The TITLE" {
		t.FailNow()
	}

	v, ok = c.Get("tags")
	sa := v.StringArray()
	if !ok || len(sa) != 3 || sa[0] != "javascript" || sa[1] != "golang" || sa[2] != "c sharp" {
		t.FailNow()
	}

	v, ok = c.Get("pi")
	f, err := v.Float()
	if !ok || err != nil || f != 3.14 {
		t.FailNow()
	}

	v, ok = c.Get("date")
	d, err := v.Date()
	if !ok || err != nil || d.Year() != 2018 || d.Month() != time.January || d.Day() != 2 {
		t.FailNow()
	}

	v, ok = c.Get("flag")
	b := v.Bool()
	if !ok || !b {
		t.FailNow()
	}

	_, ok = c.Get("not exists")
	if ok {
		t.FailNow()
	}
}
