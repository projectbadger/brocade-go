package test_data

var ContentType []testData = []testData{
	{
		Name:        "String output for ContentTypeJSON",
		InputStr:    []string{"application/yang-data+json"},
		ExpectedStr: []string{"application/yang-data+json"},
		ExpectedBool: []bool{
			true, // Is Valid
		},
	},
	{
		Name:        "String output for ContentTypeXML",
		InputStr:    []string{"application/yang-data+xml"},
		ExpectedStr: []string{"application/yang-data+xml"},
		ExpectedBool: []bool{
			true, // Is Valid
		},
	},
	{
		Name:        "String output for invalid",
		InputStr:    []string{"application/invalid"},
		ExpectedStr: []string{""},
		ExpectedBool: []bool{
			false, // Is Valid
		},
	},
}
var ContentTypeMarshalResponse []testData = []testData{
	{
		Name: "Valid XML response",
		InputStr: []string{
			"application/yang-data+xml",
			string(`<?xml version="1.0"?>
<Response>
  <test-field-top>
    <test-string>STRING</test-string>
    <test-int>10</test-int>
    <test-bool>true</test-bool>
  </test-field-top>
</Response>

`)},
		ExpectedStr: []string{"STRING"},
		ExpectedInt: []int{10},
		ExpectedBool: []bool{
			true, // Is Valid
		},
		ExpectedErr: []string{
			"",
		},
	},
	{
		Name: "Valid JSON response",
		InputStr: []string{
			"application/yang-data+json",
			string(`{
  "Response": {
    "test-field-top": {
      "test-string": "STRING",
      "test-int": 10,
      "test-bool": true
    }
  }
}`)},
		ExpectedStr: []string{"STRING"},
		ExpectedInt: []int{10},
		ExpectedBool: []bool{
			true, // Is Valid
		},
		ExpectedErr: []string{
			"",
		},
	},
	{
		Name: "Valid JSON struct",
		InputStr: []string{
			"application/yang-data+json",
			string(`{
  "test-field-top": {
    "test-string": "STRING",
    "test-int": 10,
    "test-bool": true
  }
}`)},
		ExpectedStr: []string{"STRING"},
		ExpectedInt: []int{10},
		ExpectedBool: []bool{
			true, // Is Valid
		},
		ExpectedErr: []string{
			"",
		},
	},
	{
		Name: "Invalid XML response",
		InputStr: []string{
			"application/yang-data+xml",
			string(`<*xml version="1.0"
<Response
	<test-field-top>
		<test-string>STRING</test-string>
		<test-int>10</test-int>
		<test-bool>true</test-bool>
	</test-field-top>
</Response>`)},
		ExpectedStr: []string{""},
		ExpectedInt: []int{0},
		ExpectedBool: []bool{
			false, // Is Valid
		},
		ExpectedErr: []string{
			"0: no parsable content found",
		},
	},
	{
		Name: "Incomplete XML response",
		InputStr: []string{
			"application/yang-data+xml",
			string(`<?xml version="1.0"?>
<Response>
	<test-field-top>
		<test-string>STRING</test-string>
		<test-bool>true</test-bool>
	</test-field-top>
</Response>`)},
		ExpectedStr: []string{"STRING"},
		ExpectedInt: []int{0},
		ExpectedBool: []bool{
			true, // Is Valid
		},
		ExpectedErr: []string{
			"",
		},
	},
	{
		Name: "Invalid JSON response",
		InputStr: []string{
			"application/yang-data+json",
			string(`
"Response": {
	"test-field-top": {
		"test-string": "STRING",
		"test-int": 10,
		"test-bool": true
	}
}`)},
		ExpectedStr: []string{""},
		ExpectedInt: []int{0},
		ExpectedBool: []bool{
			false, // Is Valid
		},
		ExpectedErr: []string{
			"500: invalid character ':' after top-level value",
		},
	},
	{
		Name: "Incomplete JSON response",
		InputStr: []string{
			"application/yang-data+json",
			string(`{
"Response": {
	"test-field-top": {
		"test-string": "STRING",
		"test-bool": true
	}
}
}`)},
		ExpectedStr: []string{"STRING"},
		ExpectedInt: []int{0},
		ExpectedBool: []bool{
			true, // Is Valid
		},
		ExpectedErr: []string{
			"",
		},
	},
	{
		Name: "Incomplete JSON struct",
		InputStr: []string{
			"application/yang-data+json",
			string(`{
	"test-field-top": {
		"test-string": "STRING",
		"test-bool": true
	}
}`)},
		ExpectedStr: []string{"STRING"},
		ExpectedInt: []int{0},
		ExpectedBool: []bool{
			true, // Is Valid
		},
		ExpectedErr: []string{
			"",
		},
	},
	{
		Name: "Incomplete JSON struct",
		InputStr: []string{
			"application/yang-data+json",
			string(`{
	"test-field-top": {
		"test-string": "STRING",
		"test-int": 10
	}
}`)},
		ExpectedStr: []string{"STRING"},
		ExpectedInt: []int{10},
		ExpectedBool: []bool{
			false, // Is Valid
		},
		ExpectedErr: []string{
			"",
		},
	},
	{
		Name: "Valid XML error",
		InputStr: []string{
			"application/yang-data+xml",
			string(`<?xml version="1.0"?>
<errors>
  <error>
    <error-message>STRING</error-message>
    <error-info>
		<error-code>10</error-code>
	</error-info>
  </error>
</errors>

`)},
		ExpectedStr: []string{""},
		ExpectedInt: []int{0},
		ExpectedBool: []bool{
			false, // Is Valid
		},
		ExpectedErr: []string{
			"10: STRING",
		},
	},
	{
		Name: "Valid JSON error",
		InputStr: []string{
			"application/yang-data+json",
			string(`{
	"errors": {
		"error": [{
			"error-message": "STRING",
			"error-info": {
					"error-code": 10,
					"error-module": "rest"
			}
		}]
	}
}`)},
		ExpectedStr: []string{""},
		ExpectedInt: []int{0},
		ExpectedBool: []bool{
			false, // Is Valid
		},
		ExpectedErr: []string{
			"10: STRING",
		},
	},
}
