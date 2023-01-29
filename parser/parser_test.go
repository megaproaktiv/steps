package parser_test

import (
	"os"
	"steps/parser"
	"testing"

	"gotest.tools/assert"
)

func TestCount(t *testing.T){
	_,err := os.ReadFile("./testdata/step1.txt")
	if err != nil{
		t.Log(err)
		panic( "Testdata not supplied")
	}
	startFile := "testdata/given.txt"
	createdFile := "testdata/test.txt"
	count,err := parser.Parse(&startFile, &createdFile,1)
	assert.NilError(t, err, "Parse should give not errors")
	assert.Equal(t, 3, count, "Should count all begin/ends")
}
func TestStep1(t *testing.T){
	targetContent,err := os.ReadFile("./testdata/step1.txt")
	if err != nil{
		t.Log(err)
		panic( "Testdata not supplied")
	}

	startFile := "testdata/given.txt"
	createdFile := "testdata/test.txt"
	count,err := parser.Parse(&startFile, &createdFile,1)
	assert.NilError(t, err, "Parse should give not errors")
	assert.Equal(t, 3, count, "Should count all begin/ends")

	content, err := os.ReadFile(createdFile)
	assert.NilError(t, err, "Parse should give not errors")
	

	assert.Equal(t, string(targetContent), string(content), "Content should match target")

}