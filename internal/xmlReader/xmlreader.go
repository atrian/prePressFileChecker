package xmlReader

import (
	"encoding/xml"
	"io"
	"strings"
)

const pictPattern = "<Pict>"

type Sst struct {
	XMLName     xml.Name `xml:"sst"`
	Text        string   `xml:",chardata"`
	Xmlns       string   `xml:"xmlns,attr"`
	Count       string   `xml:"count,attr"`
	UniqueCount string   `xml:"uniqueCount,attr"`
	Si          []struct {
		Text string `xml:",chardata"`
		T    struct {
			Text  string `xml:",chardata"`
			Space string `xml:"space,attr"`
		} `xml:"t"`
		R []struct {
			Text string `xml:",chardata"`
			T    struct {
				Text  string `xml:",chardata"`
				Space string `xml:"space,attr"`
			} `xml:"t"`
			RPr struct {
				Text string `xml:",chardata"`
				B    string `xml:"b"`
				Sz   struct {
					Text string `xml:",chardata"`
					Val  string `xml:"val,attr"`
				} `xml:"sz"`
				Color struct {
					Text  string `xml:",chardata"`
					Rgb   string `xml:"rgb,attr"`
					Theme string `xml:"theme,attr"`
				} `xml:"color"`
				RFont struct {
					Text string `xml:",chardata"`
					Val  string `xml:"val,attr"`
				} `xml:"rFont"`
				Family struct {
					Text string `xml:",chardata"`
					Val  string `xml:"val,attr"`
				} `xml:"family"`
				Charset struct {
					Text string `xml:",chardata"`
					Val  string `xml:"val,attr"`
				} `xml:"charset"`
			} `xml:"rPr"`
		} `xml:"r"`
	} `xml:"si"`
}

func (s *Sst) ParseXML(file io.ReadCloser) error {
	defer file.Close()

	dataBytes, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	err = xml.Unmarshal(dataBytes, s)
	if err != nil {
		return err
	}

	return nil
}

func (s *Sst) GetPics() []string {
	pics := make([]string, 0, len(s.Si))

	for _, picture := range s.Si {
		if strings.Contains(picture.T.Text, pictPattern) {
			nameParts := strings.Split(picture.T.Text, "\\")
			clearPictureName := nameParts[len(nameParts)-1]
			pics = append(pics, clearPictureName)
		}
	}

	return pics
}
