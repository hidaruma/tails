package mediaretrieval

import (
	""
	"net/http"
	"context"
)


type HlsContainer{

	XMLName xml.Name `xml:"hls,omitempty"`

	id string `xml:"id,attr"`
	bitRate string `xml:"bitrate,attr"`
	audioTrack string `xml:"audiotrack,attr"`
}

func Hls(res http.ResponseWriter, req *http.Request) {

}