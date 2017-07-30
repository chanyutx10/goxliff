package main

import "encoding/xml"
import "io/ioutil"
import "log"
import "fmt"
import "strings"

type XliffObject struct {
	XMLName           xml.Name    `xml:"xliff"`
	XMLNs             string      `xml:"xmlns,attr"`
	XMLNsXSI          string      `xml:"xsi,attr"`
	Version           string      `xml:"version,attr"`
	XSISchemeLocation string      `xml:"schemaLocation,attr"`
	Files             []XLIFFFile `xml:"file"`
}

type XLIFFFile struct {
	XMLName        xml.Name        `xml:"file"`
	Original       string          `xml:"original,attr"`
	SourceLanguage string          `xml:"source-language,attr"`
	DataType       string          `xml:"datatype,attr"`
	TargetLanguage string          `xml:"target-language,attr"`
	Header         XLIFFFileHeader `xml:"header"`
	Body           XLIFFFileBody   `xml:"body"`
}

type XLIFFFileHeader struct {
	XMLName xml.Name  `xml:"header"`
	Tool    XLIFFTool `xml:"tool"`
}

type XLIFFFileBody struct {
	XMLName    xml.Name         `xml:"body"`
	TransUnits []XLIFFTransUnit `xml:"trans-unit"`
}

type XLIFFTool struct {
	XMLName     xml.Name `xml:"tool"`
	ID          string   `xml:"tool-id,attr"`
	Name        string   `xml:"tool-name,attr"`
	Version     string   `xml:"tool-version,attr,omitempty"`
	BuildNumber string   `xml:"build-num,attr"`
}

type XLIFFTransUnit struct {
	XMLName xml.Name              `xml:"trans-unit"`
	ID      string                `xml:"id,attr"`
	Source  *XLIFFTransUnitSource `xml:"source"`
	Target  *XLIFFTransUnitTraget `xml:"target,omitempty"`
	Note    *XLIFFTransUnitNote   `xml:"note,omitempty"`
}

type XLIFFTransUnitSource struct {
	XMLName xml.Name `xml:"source"`
	Value   string   `xml:",innerxml"`
}

type XLIFFTransUnitTraget struct {
	XMLName xml.Name `xml:"target"`
	Value   string   `xml:",innerxml"`
}

type XLIFFTransUnitNote struct {
	XMLName xml.Name `xml:"note"`
	Value   string   `xml:",innerxml"`
}

func main() {
	b, err := ioutil.ReadFile("/Users/chanyut/Desktop/IdeaColor/th.xliff")
	if err != nil {
		log.Fatalf("failed to open file due to error: %v", err)
	}

	xliffObj := new(XliffObject)
	err = xml.Unmarshal(b, xliffObj)
	if err != nil {
		log.Fatalf("failed to unmarshal xml due to error: %v", err)
	}

	fmt.Printf("unmarshaled xml\n%v", xliffObj)

	xliffXML, err := xml.MarshalIndent(xliffObj, "", "  ")
	if err != nil {
		log.Fatalf("failed to marshal back to xml due to error: %v", err)
	}

	// unescape &#39 to '
	xml := strings.Replace(string(xliffXML), "&#39;", "'", -1)

	ioutil.WriteFile("/Users/chanyut/Desktop/IdeaColor/th-compared.xliff", []byte(xml), 0644)
}
