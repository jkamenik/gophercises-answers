package linkParser

import (
	"errors"
	"fmt"
	"io"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

type Link struct {
	Href string
	Text string
}

func Parse(doc io.Reader) ([]Link, error) {
	rootNode, err := html.Parse(doc)
	if err != nil {
		return nil, err
	}

	ls := New()
	links := ls.CollectLinks(rootNode)
	fmt.Printf("Done: %+v\n", links)

	fmt.Println("--- Links ")
	for _, link := range links {
		fmt.Println(link)
	}

	return nil, errors.New("TBD")
}

type linkState struct {
	currentLink *Link
	links       []Link
}

func New() linkState {
	return linkState{nil, make([]Link, 0)}
}

func (ls linkState) GetLinks() []Link {
	return ls.links
}

func (ls linkState) String() string {
	return fmt.Sprintf("linkState{currentLink: %+v, links: %+v}", ls.currentLink, ls.links)
}

func (ls linkState) CollectLinks(node *html.Node) []Link {
	ls.collectLinks(node)
	if ls.currentLink != nil {
		fmt.Println("One last link to collect: %+v\n", ls.currentLink)
		ls.links = append(ls.links, *ls.currentLink)
	}

	return ls.links
}

func (ls linkState) collectLinks(node *html.Node) {
	fmt.Printf("%s\n", ls)
	fmt.Printf("Visiting node, Type: %v, DataAtom: %v, Data: %s\n", node.Type, node.DataAtom, node.Data)

	if ls.currentLink != nil {
		fmt.Println("Building existing link")
		// If there is a current node then we are collecting text
		if node.Data != "" {
			fmt.Printf("Appending %s to link\n", node.Data)
			ls.currentLink.Text = fmt.Sprintf("%s %s", ls.currentLink.Text, node.Data)
		}
	} else {
		if node.DataAtom == atom.A {
			ls.currentLink = &Link{}
			for _, attr := range node.Attr {
				if attr.Key == "href" {
					ls.currentLink.Href = attr.Val
				}
			}
			fmt.Printf("Found a link: %+v\n", ls.currentLink)
		} else {
			fmt.Printf("%s is not %s\n", node.DataAtom, atom.A)
		}
	}

	// If I have a link then I need to collect any text data into the existing
	// link.
	if node.FirstChild != nil {
		fmt.Println("First Child")
		ls.collectLinks(node.FirstChild)
	}

	if node.NextSibling != nil {
		fmt.Println("Next Sibling")
		// Siblings need to start a new link, so append the current, and start fresh
		if ls.currentLink != nil {
			fmt.Printf("Appending %+v to links list\n", ls.currentLink)
			ls.links = append(ls.links, *ls.currentLink)
			fmt.Printf("%+v\n", ls)
			ls.currentLink = nil
		}
		ls.collectLinks(node.NextSibling)
	}

	fmt.Printf("Finished for Type: %v, DataAtom: %v, Data: %s\n", node.Type, node.DataAtom, node.Data)
	fmt.Printf("%+v\n", ls)
}
