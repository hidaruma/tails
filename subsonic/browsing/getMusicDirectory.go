/*
Copyright (C) 2014 Matt Layher
Released under the MIT license
http://opensource.org/licenses/mit-license.php
 */

package browsing

import (
	"log"
	"encoding/xml"
	"net/http"
	"context"
	"strings"
	"strconv"
)

type MusicdirectoryContainer struct {
	XMLName xml.Name `xml:"directory,omitempty"`

	ID string `xml:"id,attr"`
	Name string `xml:"name,attr"`

	Children []Child `xml:"child"`
}

type Child struct {
	XMLName xml.Name `xml:"child,omitempty"`

	ID string `xml:"id,attr"`
	Title string `xml:"title,attr"`
	Album string `xml:"album,attr"`
	Artist string `xml:"artist,attr"`
	IsDir string `xml:"isDir,attr"`
	CoverArt int `xml:"coverArt,attr"`
	Created string `xml:"created,attr"`
}

func GetMusicDirectory(w http.ResponseWriter, r *http.Request) {
	mux := http.NewServeMux()
	ctx := r.Context()
	var ctxKey struct{}
	w.WriteHeader(http.StatusOK)
	pID := r.URL.Query().Get("id")
	if pID == "" {

	}
	pair := strings.Split(pID, "_")
	if len(pair) < 2 {


	}
	prefix := pair[0]
	id, err := strconv.Atoi(pair[1])
	if err != nil {
		log.Println(err)


	}
	ctx = context.WithValue(ctx, ,)
}