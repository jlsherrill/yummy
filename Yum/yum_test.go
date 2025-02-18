package Yum

import (
	"fmt"
	"log"
	"os"
	"testing"
)

// Check that the parser can decompress a gzip file and read the correct number of packages
func TestParseCompressedXMLData(t *testing.T) {
	xmlFile, err := os.Open("mocks/yum_test.xml.gz")
	if err != nil {
		log.Fatal(err)
	}
	defer xmlFile.Close()
	result, err := ParseCompressedXMLData(xmlFile)
	if err != nil {
		t.Errorf("Error in test: %v", err)
	}
	if len(result) != 2 {
		t.Errorf("Error - Expected to return 2 packages but received: %v", len(result))
	}
}

func TestGetPrimaryURLFromRepomdXML(t *testing.T) {
	var url string = "gator/stickhat"
	xmlFile, err := os.Open("mocks/yum_test_repomd.xml")
	if err != nil {
		log.Fatal(err)
	}
	defer xmlFile.Close()
	result, err := GetPrimaryURLFromRepomdXML(xmlFile, url)
	if err != nil {
		t.Errorf("Error in test: %v", err)
	}
	expect := fmt.Sprintf("%s/repodata/primary.xml.gz", url)
	if result != expect {
		t.Errorf("Error -  Expected: %v, received: %v", expect, result)
	}
}
