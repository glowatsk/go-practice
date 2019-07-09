package main

import (
	"fmt"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(strings.NewReader(page))
	if err != nil {
		fmt.Println("error")
	}

	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode {
			switch n.Data {
			case "a":
				for _, a := range n.Attr {
					fmt.Println(a)
					if a.Key == "href" {
						fmt.Println(a.Val)
						break
					}
				}
			case "p":
				fmt.Println(n.Attr)
				for _, p := range n.Attr {
					if p.Key == "class" {
						fmt.Println(p.Val)
						break
					}
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}

	f(doc)
}

var page = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Document</title>
</head>
<body>
	<p class="test">Hello World</p>
	<a href="foo">Foo</a>
</body>
</html>
`
