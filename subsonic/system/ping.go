/*
Copyright (C) 2014 Matt Layher
Released under the MIT license
http://opensource.org/licenses/mit-license.php
 */

package system

import (
	"net/http"
	"github.com/hidaruma/tails/subsonic"

	"context"
)

func Ping(w http.ResponseWriter, r *http.Request) {
	c := subsonic.NewContainer()
	mux := http.NewServeMux()
	ctx := r.Context()
	var ctxKey struct {}
	ctx = context.WithValue(ctx, ctxKey, c.ToXML)
	w.WriteHeader(http.StatusOK)
	mux.ServeHTTP(w,r.WithContext(ctx))
}