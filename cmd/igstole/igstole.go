package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/telkomdev/igstole/httplib"
	"github.com/telkomdev/igstole/parser"
)

const Instagram = "https://instagram.com/"

func main() {
	args := os.Args[1:]
	if len(args) <= 0 {
		fmt.Println("url empty or invalid")
		os.Exit(1)
	}

	instagramURL := fmt.Sprintf("%s%s", Instagram, args[0])

	response, err := httplib.HTTPGet(instagramURL)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if response.StatusCode == 404 {
		fmt.Println("user not found")
		os.Exit(1)
	}

	defer response.Body.Close()
	reader := response.Body

	// f, err := openFile("index.html")
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }

	p := &parser.IGParser{}

	downloader := &Downloader{Parser: p, Source: reader, Response: make(chan *http.Response), Done: make(chan bool)}

	downloader.Download("")

	for {
		select {
		case r := <-downloader.Response:
			fmt.Println(r.Header)
		case <-downloader.Done:
			return
		}
	}

	// for r := range downloader.Response {
	// 	fmt.Println(r.StatusCode)
	// }

}

func openFile(path string) (io.Reader, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	return f, nil
}

type Downloader struct {
	Parser   parser.Parser
	Source   io.Reader
	Response chan *http.Response
	Done     chan bool
}

func (d *Downloader) Download(path string) error {
	data, err := d.Parser.Parse(d.Source)
	if err != nil {
		return err
	}

	profile, err := d.Parser.ParseProfile(data)
	if err != nil {
		return err
	}

	if len(profile.EntryData.ProfilePage) <= 0 {
		return errors.New("empty edge")
	}

	edges := profile.EntryData.ProfilePage[0].Graphql.User.EdgeOwnerToTimelineMedia.Edges

	for _, e := range edges {
		node := e.Node

		url := strings.Split(node.DisplayURL, "\\")
		go d.Execute(url[0])
	}

	return nil
}

func (d *Downloader) Execute(url string) {
	response, err := httplib.HTTPGet(url)
	if err != nil {
		d.Done <- true
	}

	if response.StatusCode != 200 {
		d.Done <- true
	}

	d.Response <- response
}
