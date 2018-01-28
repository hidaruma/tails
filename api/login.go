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

type LoginResponse struct {
	Error *Error `json:"error"`
	Session *data.Session `json:"session"`
}

func PostLogin(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	ctx = context.WithValue(ctx, )

	user := new(data.User)
	if

	out := LoginResponse{}

	if version, ok

	out.Session = session
	r
}