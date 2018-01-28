/*
Copyright (C) 2014 Matt Layher
Released under the MIT license
http://opensource.org/licenses/mit-license.php
 */

package data

type Album struct {
	ID int
	Artist string
	ArtistID int
	AlbumArtist string
	AlbumArtistID int
	Covers []Cover
	Title string
	Year int
	TotalDuration int
	TotalDiscs int
	DiscNumber int
}

func AlbumFromSong(song *Song) *Album {
	album := &Album{
		Artist: song.Artist,
		Title: song.Title,
		Year: song.Year,
	}
	return album
}

func (a *Album) Save() error {
	return DataBase.SaveAlbum(a)
}

func (a *Album) Delete() error {
	return DataBase.DeleteAlbum(a)
}

func (a *Album) Load() error {
	return DataBase.LoadAlbum(a)
}