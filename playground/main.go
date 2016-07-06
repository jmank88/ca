// +build js
package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"html"
	"net/url"
	"strings"
	"strconv"

	"github.com/jmank88/ca/lib"

	"github.com/gopherjs/gopherjs/js"
	"github.com/gopherjs/jquery"
	"honnef.co/go/js/dom"
)

var WebDefault ca.Config = ca.Default

func init() {
	WebDefault.Format = "svg"
}

func main() {
	js.Global.Set("render", js.MakeFunc(render))

	var config ca.Config = WebDefault

	if s := dom.GetWindow().Location().Search; s != "" {
		config = parseConfig(s[1:])
	}

	setInputs(config)

	dom.GetWindow().Document().
		GetElementByID("command").
		SetTextContent(asCommand(config))

	var b bytes.Buffer

	if err := config.Print(&b); err != nil {
		dom.GetWindow().Alert(err.Error())
		return
	}

	switch config.Format {
	case "svg":
		dom.GetWindow().Document().
			GetElementByID("output").
			SetInnerHTML(b.String())
	case "gif","png","jpg","jpeg":
		img := dom.GetWindow().Document().CreateElement("img")
		b64 := base64.StdEncoding.EncodeToString(b.Bytes())
		img.Underlying().Set("src", "data:image/png;base64," + b64)
		dom.GetWindow().Document().
			GetElementByID("output").
			AppendChild(img)
	case "","log","txt":
		fallthrough
	default:
		dom.GetWindow().Document().
			GetElementByID("output").
			SetInnerHTML(strings.Replace(html.EscapeString(b.String()), "\n", "<br/>", -1))
	}
}

// The parseConfig function parses a ca.Config from a search string.
func parseConfig(search string) ca.Config {
	q, err := url.ParseQuery(search)
	if err != nil {
		dom.GetWindow().Alert(err.Error())
	}
	var c ca.Config = WebDefault
	if cells, err := strconv.Atoi(q.Get("cells")); err == nil {
		c.Cells = cells
	}
	if gens, err := strconv.Atoi(q.Get("generations")); err == nil {
		c.Generations = gens
	}
	if f := q.Get("format"); f != "" {
		c.Format = f
	}
	_, c.Random = q["random"]
	if rule, err := strconv.Atoi(q.Get("rule")); err == nil {
		c.Rule = rule
	}
	if size, err := strconv.Atoi(q.Get("size")); err == nil {
		c.Size = size
	}
	return c
}

// The setInputs function sets the form inputs to the values from c.
func setInputs(c ca.Config) {
	if c.Cells != WebDefault.Cells {
		setInput("cells", strconv.Itoa(c.Cells))
	}
	if c.Generations != WebDefault.Generations {
		setInput("generations", strconv.Itoa(c.Generations))
	}
	if c.Format != "" && c.Format != WebDefault.Format {
		setInput("format", c.Format)
	}
	if c.Random != WebDefault.Random {
		dom.GetWindow().Document().GetElementByID("random").
			Underlying().Set("checked", "checked")
	}
	if c.Rule != WebDefault.Rule {
		setInput("rule", strconv.Itoa(c.Rule))
	}
	if c.Size != WebDefault.Size {
		setInput("size", strconv.Itoa(c.Size))
	}
}

// The setInput function sets the value of element id.
func setInput(id, value string) {
	dom.GetWindow().Document().GetElementByID(id).
		Underlying().Set("value", value)
}

// The asCommand function converts a ca.Config into a command line statement.
func asCommand(c ca.Config) string {
	var b bytes.Buffer
	b.WriteString("ca")
	if c.Cells != ca.Default.Cells {
		fmt.Fprintf(&b, " -cells %d", c.Cells)
	}
	if c.Generations != ca.Default.Generations {
		fmt.Fprintf(&b, " -gens %d", c.Generations)
	}
	if c.Format == "" {
		c.Format = WebDefault.Format
	}
	if c.Format != ca.Default.Format {
		fmt.Fprintf(&b, " -format %s", c.Format)
	}
	if c.Random != ca.Default.Random {
		fmt.Fprintf(&b, " -rand %t", c.Random)
	}
	if c.Rule != ca.Default.Rule {
		fmt.Fprintf(&b, " -r %d", c.Rule)
	}
	if c.Size != ca.Default.Size {
		fmt.Fprintf(&b, " -s %d", c.Size)
	}
	return b.String()
}


// The render function submits the form after disabling empty inputs.
// Satifies js.MakeFunc().
func render(this *js.Object, arguments []*js.Object) interface{} {
	form := jquery.NewJQuery("#form")
	form.Find(`:input`).
		Filter(js.MakeFunc(func(this *js.Object, arguments []*js.Object) interface{} {
			return this.Get("value").String() == ""
		})).
		SetProp("disabled", true)
	format := jquery.NewJQuery("#format")
	if format.Val() == WebDefault.Format {
		format.SetProp("disabled", true)
	}
	form.Submit()

	return true
}