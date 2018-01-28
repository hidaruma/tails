/*
Copyright (C) 2014 Matt Layher
Released under the MIT license
http://opensource.org/licenses/mit-license.php
 */

package data

const (
	APE = iota
	FLAC
	M4A
	MP3
	MPC
	OGG
	WMA
	WV
)

var (
	ErrSongTags = errors.New("song: required tags could not be extracted from file.")

	ErrSongProperties = errors.New("song: required properties could not be extracted from file.")
)

var FileTypeMap = map[string]int{
	".ape": APE,
	".flac": FLAC,
	".m4a": M4A,
	".mp3": MP3,
	".mpc": MPC,
	".ogg": OGG,
	".wma": WMA,
	".wv": WV,
}

var CodecMap = map[int]string{
	APE: "APE",
	FLAC: "FLAC",
	M4A: "M4A",
	MP3: "MP3",
	OGG: "OGG",
	WMA: "WMA",
	WV: "WV",
}

var MIMEMap = map[int]string{
	APE: "audio/ape",
	FLAC: "audio/flac",
	M4A: "audio/aac",
	MP3: "audio/mpeg",
	MPC: "audio/mpc",
	OGG: "audio/ogg",
	WMA: "audio/wma",
	WV: "audio/wv",
}

type Song struct {
	ID int `json:"id"`
	Album string
	AlbumID int
	AlbumArtist string
	AlbumArtistID int
	CoverID int
	Artist []string
	ArtistID int
	BitRate int
	Composer []string
	Channels int
	Comment string

}