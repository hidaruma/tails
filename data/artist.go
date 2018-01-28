/*
Copyright (C) 2014 Matt Layher
Released under the MIT license
http://opensource.org/licenses/mit-license.php
 */

package data


type Artist struct {
	ID int
	Title string
}

func ArtistFromSong(song *Song) *Artist {
	return &Artist{
		Title: song.Artist[0],
	}
}

func ArtistsFromSong(song *Song) []*Artist {
	artists := make([]*Artist,0)
	for _, artist := range song.Artist {
		artists = append(artists, &Artist{Title: artist})
	}
	return artists
}

func (a *Artist) Save() error {
	return DataBase.SaveArtist(a)
}

func (a *Artist) Delete() error {
	return DataBase.DeleteArtist(a)
}

func (a *Artist) Load() error {
	return DataBase.LoadArtist(a)
}

type SameArtist struct {
	ID int
	AnotherNames []Artist
}

func SameArtistFromArtist(a1 *Artist, a2 *Artist) *SameArtist {
	var sameartist *SameArtist
	sameartist.AnotherNames = append(sameartist.AnotherNames, a1)
	sameartist.AnotherNames = append(sameartist.AnotherNames, a2)

	return sameartist
}

func (sa *SameArtist) Save() error {
	return DataBase.SaveSameArtist(sa)
}

func (sa *SameArtist) Delete() error {
	return DataBase.DeleteSameArtist(sa)
}

func (sa *SameArtist) Load() error {
	return DataBase.LoadSameArtist(sa)
}