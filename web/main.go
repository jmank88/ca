// +build js
package main

import (
	"bytes"
	"net/url"
	"strconv"

	"github.com/jmank88/ca/lib"
	"honnef.co/go/js/dom"
	"fmt"
	"strings"
	"html"
)

func main() {
	var config ca.Config = ca.Default

	if s := dom.GetWindow().Location().Search; s != "" {
		config = parseConfig(s[1:])
	}

	//TODO set input buttons

	dom.GetWindow().Document().
		GetElementByID("command").
		SetTextContent(asCommand(config))

	var b bytes.Buffer

	if err := config.Print(&b); err != nil {
		//TODO problem
		dom.GetWindow().Alert(err.Error())
		return
	}

	//TODO clear "output"

	//TODO set output
	switch config.Format {
	case "svg":
		dom.GetWindow().Document().
			GetElementByID("output").
			SetInnerHTML(b.String())
	case "gif","png","jpg","jpeg":
		//TODO set image somehow
	case "","log","txt":
		fallthrough
	default:
		dom.GetWindow().Document().
			GetElementByID("output").
			SetInnerHTML(strings.Replace(html.EscapeString(b.String()), "\n", "<br/>", -1))
	}
}

//TODO function for turning a url into ca params
func parseConfig(search string) ca.Config {
	q, err := url.ParseQuery(search)
	if err != nil {
		dom.GetWindow().Alert(err.Error())
	}
	var c ca.Config = ca.Default
	if rule, err := strconv.Atoi(q.Get("rule")); err == nil {
		c.Rule = rule
	}
	if rand, err := strconv.ParseBool(q.Get("random")); err == nil {
		c.Random = rand
	}
	if cells, err := strconv.Atoi(q.Get("cells")); err == nil {
		c.Cells = cells
	}
	if gens, err := strconv.Atoi(q.Get("generations")); err == nil {
		c.Generations = gens
	}
	c.Format = q.Get("format")
	return c
}

//TODO function for setting button inputs from ca params
func setInputs(c ca.Config) {
	return //TODO
}

//TODO function for ca params -> command line equivalent
func asCommand(c ca.Config) string {
	var b bytes.Buffer
	b.WriteString("ca")
	if c.Cells != ca.Default.Cells {
		fmt.Fprintf(&b, " -cells %d", c.Cells)
	}
	if c.Generations != ca.Default.Generations {
		fmt.Fprintf(&b, " -gens %d", c.Generations)
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
	return b.String()
}

//TODO function for turning button intputs into ca params
func getInputConfig() ca.Config {
	var c ca.Config
	//TODO
	return c
}

//TODO function for setting url from ca params
func getSearch(c ca.Config) string {
	var v url.Values
	v.Add("rule", strconv.Itoa(c.Rule))
	v.Add("random", strconv.FormatBool(c.Random))
	v.Add("cells", strconv.Itoa(c.Cells))
	v.Add("generations", strconv.Itoa(c.Generations))
	if c.Format != "" {
		v.Add("format", c.Format)
	}
	return v.Encode()
}