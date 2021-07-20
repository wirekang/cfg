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
	if !c.IsExist("num") {
		t.FailNow()
	}
	if i, e := c.Find("num").Int(); e != nil || i != 3 {
		t.FailNow()
	}
	if c.Find("title").String() != "The TITLE" {
		t.FailNow()
	}
	tags := c.Find("tags").StringArray()
	if tags[0] != "javascript" || tags[1] != "golang" || tags[2] != "c sharp" {
		t.FailNow()
	}
	if f, e := c.Find("pi").Float(); e != nil || f != 3.14 {
		t.FailNow()
	}

	if d, e := c.Find("date").Date(); e != nil || d.Year() != 2018 || d.Month() != time.January || d.Day() != 2 {
		t.FailNow()
	}

	if !c.Find("flag").Bool() {
		t.FailNow()
	}

	if c.Find("notexists").Bool() {
		t.FailNow()
	}
}
