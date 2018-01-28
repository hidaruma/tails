/*
Copyright (C) 2014 Matt Layher
Released under the MIT license
http://opensource.org/licenses/mit-license.php
 */

package songlists

import (
	"encoding/xml"

	"net/http"
	"strconv"
	"log"
	"context"
	"github.com/hidaruma/tails/data"
	"github.com/hidaruma/tails/subsonic"
)

type RandomSongsContainer struct {
	XMLName xml.Name `xml:"randomSongs,omitempty"`

	Songs []Song `xml:"song"`
}

func GetRandomSongs(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	mux := http.NewServeMux()
	var ctxKey struct{}
	w.WriteHeader(http.StatusOK)

	size := 10

	if pSize := r.URL.Query().Get("size"); pSize != "" {

		tmp, err := strconv.Atoi(pSize)
		if err != nil {
			log.Println(err)
			ctx = context.WithValue(ctx, ctxKey, subsonic.ErrGeneric.ToXML)
			mux.ServeHTTP(w, r.WithContext(ctx))
			return
		}
		size = tmp
	}

	songs, err := data.DataBase.RandomSongs(size)
	if err != nil {
		log.Println(err)
		ctx = context.WithValue(ctx, ctxKey, subsonic.ErrGeneric.ToXML)
		mux.ServeHTTP(w, r.WithContext(ctx))
		return
	}

	outSongs := make([]Song, 0)
	for _, s := range songs {
		outSongs = append(outSongs, subSong(s))
	}

	c := subsonic.NewContainer()
	c.RandomSongs = &RandomSongsContainer{Songs: outSongs}

	ctx = context.WithValue(ctx, ctxKey, c.ToXML)
	mux.ServeHTTP(w, r.WithContext(ctx))
}