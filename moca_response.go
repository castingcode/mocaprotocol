package mocaprotocol

import "encoding/xml"

type MocaDataType string

const (
	MocaInt         MocaDataType = "I"
	MocaString      MocaDataType = "S"
	MocaDate        MocaDataType = "D"
	MocaFlag        MocaDataType = "O"
	MocaFloat       MocaDataType = "F"
	MocaBase64      MocaDataType = "V"
	MocaUnknownType MocaDataType = "U"
)

const (
	MocaStatusOK             int64 = 0
	MocaStatusLoginRequired  int64 = 523
	MocaStatusDBNoDataFound  int64 = -1403
	MocaStatusSrvNoDataFound int64 = 510
)

type MocaResponse struct {
	XMLName     xml.Name    `xml:"moca-response"`
	SessionID   string      `xml:"session-id"`
	Status      int         `xml:"status"`
	MocaResults MocaResults `xml:"moca-results"`
	Message     string      `xml:"message,omitempty"`
}

type MocaResults struct {
	Metadata Metadata `xml:"metadata"`
	Data     Data     `xml:"data"`
}

type Metadata struct {
	Columns []Column `xml:"column"`
}

type Column struct {
	Name     string       `xml:"name,attr"`
	Type     MocaDataType `xml:"type,attr"`
	Length   string       `xml:"length,attr"`
	Nullable string       `xml:"nullable,attr"`
}

type Data struct {
	Rows []Row `xml:"row"`
}

type Row struct {
	Fields []Field `xml:"field"`
}

type Field struct {
	Value string `xml:",chardata"`
	Null  string `xml:"null,attr,omitempty"`
}
