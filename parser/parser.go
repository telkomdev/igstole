package parser

import (
	"encoding/json"
	"io"
	"strings"

	"golang.org/x/net/html"
)

const (
	Prefix = "window._sharedData"
)

// Parser interface
type Parser interface {
	Parse(reader io.Reader) ([]byte, error)
}

// IGParser struct
type IGParser struct {
	tokenizer *html.Tokenizer
}

// Parse
func Parse(reader io.Reader) ([]byte, error) {
	z := html.NewTokenizer(reader)

	for {
		tokenType := z.Next()

		switch tokenType {
		case html.ErrorToken:
			return nil, z.Err()
		case html.StartTagToken:
			token := z.Token()

			if token.Data != "script" {
				continue
			}

			for _, a := range token.Attr {
				if a.Key == "type" {

					z.Next()
					t := z.Token()

					if strings.HasPrefix(t.Data, Prefix) {
						data := t.Data[len(Prefix)+3 : len(t.Data)-1]
						return []byte(data), nil
					}
				}
			}
		}
	}
}

// GetIGProfile return Profile, and error otherwise
func GetIGProfile(data []byte) (*Profile, error) {
	var profile Profile
	err := json.Unmarshal(data, &profile)
	if err != nil {
		return nil, err
	}

	return &profile, nil
}
