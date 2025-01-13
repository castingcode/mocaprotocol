package mocaprotocol

import (
	"encoding/xml"
	"io"
	"os"
	"testing"

	"github.com/matryer/is"
)

// TestMocaResponse_Serialize tests the serialization of a MocaResponse received from the server
func TestMocaResponse_Deserialize(t *testing.T) {
	is := is.New(t)
	file, err := os.Open("testdata/sample_response_success.xml")
	if err != nil {
		t.Fatalf("failed to open file: %v", err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		t.Fatalf("failed to read file: %v", err)
	}

	var response MocaResponse
	err = xml.Unmarshal(data, &response)
	is.NoErr(err)
	is.Equal(response.Status, 0)
	is.Equal(response.Message, "")
	is.Equal(len(response.MocaResults.Metadata.Columns), 55)
	is.Equal(len(response.MocaResults.Data.Rows), 2)
	is.Equal(len(response.MocaResults.Data.Rows[0].Fields), 55)
	is.Equal(response.MocaResults.Data.Rows[0].Fields[0].Value, "WMD1")
	is.Equal(response.MocaResults.Metadata.Columns[0], Column{Name: "wh_id", Type: MocaString, Length: "32", Nullable: "false"})
}
