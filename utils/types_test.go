package utils

import (
	"encoding/xml"
	"testing"

	"github.com/projectbadger/brocade-go/utils/test_data"
)

func TestContentType(t *testing.T) {
	t.Run("Properties and value", func(t *testing.T) {
		for _, test := range test_data.ContentType {
			t.Run(test.Name, func(t *testing.T) {
				ct := ContentType(test.InputStr[0])
				if ct.Valid() != test.ExpectedBool[0] {
					t.Errorf("'%s' returns valid='%v' when it should have returned valid='%v'", test.InputStr[0], ct.Valid(), test.ExpectedBool[0])
				}
				if ct.String() != test.ExpectedStr[0] {
					t.Errorf("'%s' returns string='%v' when it should have returned string='%v'", test.InputStr[0], ct.Valid(), test.ExpectedStr[0])
				}
				// if test.InputByte != nil {
				// 	// Test marshalling && Unmarshalling
				// }
			})
		}
	})
	t.Run("Marshal response", func(t *testing.T) {
		for _, test := range test_data.ContentTypeMarshalResponse {
			t.Run(test.Name, func(t *testing.T) {
				ct := ContentType(test.InputStr[0])
				var testStruct struct {
					XMLName    xml.Name `json:"-" xml:"test-field-top"`
					TestString string   `json:"test-string" xml:"test-string"`
					TestInt    int      `json:"test-int,omitempty" xml:"test-int"`
					TestBool   bool     `json:"test-bool" xml:"test-bool"`
				}
				err := ct.UnmarshalResponse([]byte(test.InputStr[1]), &testStruct)
				if err != nil && test.ExpectedErr[0] == "" {
					t.Errorf("error='%v' when it should have returned nil", err.Error())
				} else if err != nil && err.Error() != test.ExpectedErr[0] {
					t.Errorf("error='%s' when it should have returned '%s'\n", err.Error(), test.ExpectedErr[0])
				} else if err == nil && test.ExpectedErr[0] != "" {
					t.Errorf("error is nil when it should have returned '%s'", test.ExpectedErr[0])
				}
				if err != nil {
					return
				}

				if testStruct.TestString != test.ExpectedStr[0] {
					t.Errorf("'%s' returns string='%v' when it should have returned string='%v'\nStruct: '%#v'", test.InputStr[1], testStruct.TestString, test.ExpectedStr[0], testStruct)
				}
				if testStruct.TestInt != test.ExpectedInt[0] {
					t.Errorf("'%s' returns int='%d' when it should have returned int='%v'\nStruct: '%#v'", test.InputStr[0], testStruct.TestInt, test.ExpectedInt[0], testStruct)
				}
				// if test.InputByte != nil {
				// 	// Test marshalling && Unmarshalling
				// }
			})
		}
	})
}
