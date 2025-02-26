package mandosjsontest

import (
	"io/ioutil"
	"os"
	"testing"

	fr "github.com/ElrondNetwork/arwen-wasm-vm/v1_4/mandos-go/fileresolver"
	mjparse "github.com/ElrondNetwork/arwen-wasm-vm/v1_4/mandos-go/json/parse"
	mjwrite "github.com/ElrondNetwork/arwen-wasm-vm/v1_4/mandos-go/json/write"
	"github.com/stretchr/testify/require"
)

func loadExampleFile(path string) ([]byte, error) {
	// Open our jsonFile
	var jsonFile *os.File
	var err error
	jsonFile, err = os.Open(path)
	// if we os.Open returns an error then handle it
	if err != nil {
		return nil, err
	}

	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	return ioutil.ReadAll(jsonFile)
}

func TestWriteTest(t *testing.T) {
	contents, err := loadExampleFile("example.test.json")
	require.Nil(t, err)

	p := mjparse.NewParser(
		fr.NewDefaultFileResolver().ReplacePath(
			"smart-contract.wasm",
			"exampleFile.txt"))

	testTopLevel, parseErr := p.ParseTestFile(contents)
	require.Nil(t, parseErr)

	serialized := mjwrite.TestToJSONString(testTopLevel)

	// good for debugging:
	_ = ioutil.WriteFile("serialized.test.json", []byte(serialized), 0644)

	require.Equal(t, contents, []byte(serialized))
}
