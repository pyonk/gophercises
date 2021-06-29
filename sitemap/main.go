package main

import (
	"encoding/xml"
	"flag"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"golang.org/x/net/html"
)

type UrlSet struct {
	XMLName xml.Name `xml:"urlset"`
	Xmlns   string   `xml:"xmlns,attr"`
	Urls    []Url
}

func (urlset *UrlSet) Contains(url string) bool {
	for _, u := range urlset.Urls {
		if u.Loc == url {
			return true
		}
	}
	return false
}
func (urlset *UrlSet) Append(url string) {
	for _, u := range urlset.Urls {
		if u.Loc == url {
			return
		}
	}
	urlset.Urls = append(urlset.Urls, Url{
		Loc: url,
	})
}

type Url struct {
	XMLName xml.Name `xml:"url"`
	Loc     string   `xml:"loc"`
}

func main() {
	urlFlag := flag.String("url", "https://gophercises.com/", "Output sitemap created by this url")
	maxDepth := flag.Int("depth", 1, "The number of depth to create sitemap")
	flag.Parse()
	if *urlFlag == "" {
		log.Fatal("URL Required.")
	}
	targetUrl, err := url.Parse(*urlFlag)
	if err != nil {
		log.Fatalf("Failed to parse URL: %s\n%s", *urlFlag, err)
	}

	seen := make(map[string]bool)

	var f func(*url.URL) []*url.URL
	f = func(u *url.URL) []*url.URL {
		if _, ok := seen[u.String()]; ok {
			return nil
		}
		seen[u.String()] = true
		time.Sleep(2 * time.Second)
		res, err := http.Get(u.String())
		println("GET", u.String())
		if err != nil {
			log.Fatalf("Failed to access to %s\n%s", *urlFlag, err)
		}
		urls := extractUrls(*res)
		urls = filterSameDomainUrls(urls, *targetUrl)
		return urls
	}

	urls := f(targetUrl)

	urlset := UrlSet{
		Urls:  make([]Url, 0),
		Xmlns: "http://www.sitemaps.org/schemas/sitemap/0.9",
	}

	for i := 1; i <= *maxDepth; i++ {
		var nextUrls []*url.URL
		for _, u := range urls {
			nextUrls = append(nextUrls, f(u)...)
			urlset.Append(u.String())
		}
		urls = make([]*url.URL, 0)
		for _, nu := range nextUrls {
			if urlset.Contains(nu.String()) == false {
				urls = append(urls, nu)
			}
		}
	}

	output, err := xml.MarshalIndent(urlset, "", "  ")
	if err != nil {
		log.Fatalf("Failed to marshal XML\n%s", err)
	}
	log.Printf("%s%s", xml.Header, output)
}

func extractUrls(res http.Response) (result []*url.URL) {
	m := map[string]bool{}
	node, err := html.Parse(res.Body)
	if err != nil {
		log.Fatalf("Failed to parse HTML\n%s", err)
	}
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, attr := range n.Attr {
				if attr.Key == "href" {
					if m[attr.Val] != true {
						m[attr.Val] = true
						u, err := url.Parse(attr.Val)
						if err != nil {
							log.Printf("Failed to parse URL: %s\n%s\n", attr.Val, err)
						}
						result = append(result, u)
					}
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(node)
	return result
}

func filterSameDomainUrls(urls []*url.URL, targetUrl url.URL) (result []*url.URL) {
	for _, u := range urls {
		if u.Scheme == "https" || u.Scheme == "http" {
			if u.Host == targetUrl.Host {
				result = append(result, u)
			}
		}
		if u.Scheme == "" && strings.HasPrefix(u.Path, "/") {
			u.Scheme = targetUrl.Scheme
			u.Host = targetUrl.Host
			result = append(result, u)
		}
	}
	return result
}
