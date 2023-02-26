package link

import (
	"io"
	"strings"

	"golang.org/x/net/html"
)

type Link struct {
	Href string
	Text string
}

func Parse(f io.Reader) ([]Link, error) {
	node, err := html.Parse(f)
	if err != nil {
		return nil, err
	}

	var links []Link
	for _, node := range LinkNodes(node) {
		links = append(links, buildLink(node))
	}

	return links, err
}

func buildLink(n *html.Node) Link {
	var l Link
	for _, a := range n.Attr {
		if a.Key == "href" {
			l.Href = a.Val
			break
		}
	}
	l.Text = strings.Join(strings.Fields(text(n)), " ")
	return l
}

func text(n *html.Node) string {
	if n.Type == html.TextNode {
		return n.Data
	}
	if n.Type != html.ElementNode {
		return ""
	}
	var ret string
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ret += text(c) + " "
	}
	return ret
}

func LinkNodes(n *html.Node) []*html.Node {
	var nodes []*html.Node
	if n.Type == html.ElementNode && n.Data == "a" {
		nodes = append(nodes, n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		nodes = append(nodes, LinkNodes(c)...)
	}
	return nodes
}
