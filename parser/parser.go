package parser

import (
	"net/url"

	"golang.org/x/net/html"
)

func ExtractLinks(base string, node *html.Node) []string {
	var links []string

	baseURL, err := url.Parse(base)
	if err != nil {
		return links
	}

	var dfs func(*html.Node)
	dfs = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, attr := range n.Attr {
				if attr.Key == "href" {
					u, err := url.Parse(attr.Val)
					if err == nil {
						links = append(links, baseURL.ResolveReference(u).String())
					}
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			dfs(c)
		}
	}

	dfs(node)
	return links
}
