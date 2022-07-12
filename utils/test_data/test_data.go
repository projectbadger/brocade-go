package test_data

type testData struct {
	Name         string
	InputStr     []string
	InputByte    [][]byte
	InputBool    []bool
	InputInt     []int
	ExpectedStr  []string
	ExpectedByte [][]byte
	ExpectedBool []bool
	ExpectedErr  []string
	ExpectedInt  []int
}
