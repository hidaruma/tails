/*
Copyright (C) 2014 Matt Layher
Released under the MIT license
http://opensource.org/licenses/mit-license.php
 */

package api

import (
	"log"
	"net/http"
	"context"

	"github.com/hidaruma/tails/data"

	"github.com/unrolled/render"
)

func GetStream(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	ctx = context.WithValue(ctx, )

	w.Header().Set("Accept-Ranges", "bytes")


}