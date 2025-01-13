package mocaprotocol

import (
	"encoding/xml"
	"testing"

	"github.com/matryer/is"
)

// TestMocaRequest_Serialize tests the serialization of a MocaRequest to be sent to the server
func TestMocaRequest_Serialize(t *testing.T) {
	is := is.New(t)
	payload := `<moca-request autocommit="true">
	<environment>
		<var name="USR_ID" value="SUPER01"></var>
		<var name="LOCALE_ID" value="US_ENGLISH"></var>
		<var name="SESSION_KEY" value="srv26-dev;uid=SUPER01|sid=7706a9f7-4954-48f9-83c5-01d0c7bd59bb|dt=luv9fp8k|sec=ALL;ZAofk27.c7mjtHb1VPyShjoKH2"></var>
	</environment>
	<query><![CDATA[set return status where status = 9 and message = "hello"]]></query>
</moca-request>`

	want := MocaRequest{
		XMLName:    xml.Name{Local: "moca-request"},
		Autocommit: "true",
		Query: Query{
			Text: `set return status where status = 9 and message = "hello"`,
		},
		Environment: Environment{
			Vars: []Var{
				{Name: "USR_ID", Value: "SUPER01"},
				{Name: "LOCALE_ID", Value: "US_ENGLISH"},
				{Name: "SESSION_KEY", Value: "srv26-dev;uid=SUPER01|sid=7706a9f7-4954-48f9-83c5-01d0c7bd59bb|dt=luv9fp8k|sec=ALL;ZAofk27.c7mjtHb1VPyShjoKH2"},
			},
		},
	}
	b, err := xml.MarshalIndent(want, "", "\t")
	is.NoErr(err)
	is.Equal(payload, string(b))
}

// TestMocaRequest_Deserialize tests the deserialization of a MocaRequest.
// This is just for completeness, as we are not expecting to be receiving these.
func TestMocaRequest_Deserialize(t *testing.T) {
	is := is.New(t)
	payload := `<moca-request autocommit="true">
  <environment>
    <var name="USR_ID" value="SUPER01" />
    <var name="LOCALE_ID" value="US_ENGLISH" />
    <var name="SESSION_KEY" value="srv26-dev;uid=SUPER01|sid=7706a9f7-4954-48f9-83c5-01d0c7bd59bb|dt=luv9fp8k|sec=ALL;ZAofk27.c7mjtHb1VPyShjoKH2" />
  </environment>
  <query><![CDATA[set return status where status = 9 and message = "hello"]]></query>
</moca-request>`

	want := MocaRequest{
		XMLName:    xml.Name{Local: "moca-request"},
		Autocommit: "true",
		Query: Query{
			Text: `set return status where status = 9 and message = "hello"`,
		},
		Environment: Environment{
			Vars: []Var{
				{Name: "USR_ID", Value: "SUPER01"},
				{Name: "LOCALE_ID", Value: "US_ENGLISH"},
				{Name: "SESSION_KEY", Value: "srv26-dev;uid=SUPER01|sid=7706a9f7-4954-48f9-83c5-01d0c7bd59bb|dt=luv9fp8k|sec=ALL;ZAofk27.c7mjtHb1VPyShjoKH2"},
			},
		},
	}
	var got MocaRequest
	err := xml.Unmarshal([]byte(payload), &got)
	is.NoErr(err)
	is.Equal(want.Autocommit, got.Autocommit)
	is.Equal(want.Query.Text, got.Query.Text)
	is.Equal(want.Environment.Vars, got.Environment.Vars)
}

// TestMocaRequest_DeserializeNoEnvironment tests the deserialization of a MocaRequest.
// This is just for completeness, as we are not expecting to be receiving these.
func TestMocaRequest_DeserializeNoEnvironment(t *testing.T) {
	is := is.New(t)
	payload := `<moca-request autocommit="true">
  <query><![CDATA[login user]]></query>
</moca-request>`

	want := MocaRequest{
		XMLName:    xml.Name{Local: "moca-request"},
		Autocommit: "true",
		Query: Query{
			Text: `login user`,
		},
	}
	var got MocaRequest
	err := xml.Unmarshal([]byte(payload), &got)
	is.NoErr(err)
	is.Equal(want.Autocommit, got.Autocommit)
	is.Equal(want.Query.Text, got.Query.Text)
	is.Equal(want.Environment.Vars, got.Environment.Vars)
}
