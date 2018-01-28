/*
Copyright (C) 2014 Matt Layher
Released under the MIT license
http://opensource.org/licenses/mit-license.php
 */
package system

import (
	"net/http"
	"encoding/xml"
	"context"
	"github.com/hidaruma/tails/subsonic"
)

type License struct {
	XMLName xml.Name `xml:"license,omitempty"`

	Valid bool `xml:"valid,attr"`
	Email string `xml:"email,attr"`
	Key string `xml:"key,attr"`
	Date string `xml:"date,attr"`
}

func GetLicense(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	mux := http.NewServeMux()
	var ctxKey struct{}
	w.WriteHeader(http.StatusOK)
	c := subsonic.NewContainer()
	c.License = &License{
		Valid: true,
	}
	ctx = context.WithValue(ctx, ctxKey, c.ToXML)
	mux.ServeHTTP(w, r.WithContext(ctx))
}