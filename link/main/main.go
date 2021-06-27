package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/pyonk/gophercises/link"
	"golang.org/x/net/html"
)

func main() {
	file := flag.String("file", "ex1.html", "the HTML file extracted Link")
	flag.Parse()
	fmt.Printf("Extract link in %s\n", *file)

	htmlFile, err := os.Open(*file)
	if err != nil {
		log.Fatalf("Failed to read file\n %s", err)
	}

	n, err := html.Parse(htmlFile)
	if err != nil {
		log.Fatalf("Failed to parse HTML\n %s", err)
	}

	var links []link.Link

	linkNodes := extractLinkNodes(n)
	for _, ln := range linkNodes {
		var href, text string
		for _, attr := range ln.Attr {
			if attr.Key == "href" {
				href = attr.Val
			}
		}
		text = extractText(ln)
		link := link.Link{
			Href: href,
			Text: text,
		}
		links = append(links, link)
	}

	for _, l := range links {
		j, _ := json.MarshalIndent(l, "", "  ")
		fmt.Printf("Link%s\n", j)
	}
	return
}

func extractText(parent *html.Node) (result string) {
	var f func(*html.Node)
	f = func(n *html.Node) {
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			if c.Type == html.TextNode {
				result += c.Data
			}
			f(c)
		}
	}
	f(parent)
	return strings.TrimSpace(result)
}

func extractLinkNodes(parent *html.Node) (result []*html.Node) {
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			result = append(result, n)
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(parent)
	return result
}
