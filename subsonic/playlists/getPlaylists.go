/*
Copyright (C) 2014 Matt Layher
Released under the MIT license
http://opensource.org/licenses/mit-license.php
 */

package playlists

import (
	"context"
	"net/http"
	"encoding/xml"
)

typePlaylists struct {
	XMLName xml.Name `xml:"playlists.omitempty"`
}

func GetPlaylists(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

}