package mocaprotocol

import "encoding/xml"

type MocaRequest struct {
	XMLName     xml.Name    `xml:"moca-request"`
	Environment Environment `xml:"environment,omitempty"`
	Autocommit  string      `xml:"autocommit,attr"`
	Query       Query
}

type Query struct {
	XMLName xml.Name `xml:"query"`
	Text    string   `xml:",cdata"`
}

type Environment struct {
	Vars []Var `xml:"var,omitempty"`
}

type Var struct {
	Name  string `xml:"name,attr"`
	Value string `xml:"value,attr"`
}

func NewMocaRequest(query string, vars ...Var) MocaRequest {
	return MocaRequest{
		Autocommit:  "true",
		Query:       Query{Text: query},
		Environment: Environment{Vars: vars},
	}
}
