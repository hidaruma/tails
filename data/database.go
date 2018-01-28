/*
Copyright (C) 2014 Matt Layher
Released under the MIT license
http://opensource.org/licenses/mit-license.php
 */

package data

var DataBase dataBase

type dataBase interface{
	Open() error
	Close() error
	InitDB() error
	DSN(string)

	CoverInPath(string) ([]Cover, error)
	CoverNotInPath(string) ([]Cover, error)
	CountCovers() (int64, error)
	DeleteCover(*Cover) error
	LoadCover(*Cover) error
	SaveCover(*Cover) error

	AllArtists() ([]Artist, error)
	AllArtistsByTitle() ([]Artist, error)
	LimitArtists(int, int) ([]Artist, error)
	SearchArtists(string) ([]Artist, error)
	CountArtists() (int64, error)
	PurgeOrphanArtists() (int, error)
	DeleteArtist(*Artist) error
	LoadArtist(*Artist) error
	SaveArtist(*Artist) error

	AllSameArtists() ([]SameArtist, error)
	LimitSameArtist(int, int) ([]SameArtist, error)
	SearchSameArtist(string) ([]SameArtist, error)
	CountSameArtist() (int64, error)
	SaveSameArtist(*SameArtist) error
	DeleteSameArtist(*SameArtist) error
	LoadSameArtist(*SameArtist) error


	AllAlbums() ([]Album, error)
	LimitAlbums(int, int) ([]Album, error)
	AlbumsForArtist(int) ([]Album, error)
	SearchAlbums(string) ([]Album, error)
	CountAlbums() (int64, error)
	PurgeOrphanAlbums() (int, error)
	DeleteAlbum(*Album) error
	LoadAlbum(*Album) error
	SaveAlbum(*Album) error

	AllDirectories() ([]Directory, error)
	LimitDirectory(int, int) ([]Directory, error)
	SubDirectories(int) ([]Directory, error)
	DirectoryInPath(string) ([]Directory, error)
	DirectoryNotInPath(string) ([]Directory, error)
	SearchDirectories(string) ([]Directory, error)
	CountDirectories() (int64, error)
	DeleteDirectory(*Directory) error
	LoadDirectory(*Directory) error
	SaveDirectory(*Directory) error

	AllSongs() ([]Song, error)
	LimitSongs(int, int) ([]Song, error)
	RandomSongs(int) ([]Song, error)
	SearchSongs(string) ([]Song, error)
	SongsForAlbum(int) ([]Song, error)
	SongsForArtist(int) ([]Song, error)
	SongsForFolder(int) ([]Song, error)
	SongsInPath(string) ([]Song, error)
	SongsNotInPath(string) ([]Song, error)
	CountSongs() (int64, error)
	DeleteSong(*Song) error
	LoadSong(*Song) error
	SaveSong(*Song) error
	UpdateSong(*Song) error

	AllUsers() ([]User, error)
	DeleteUser(*User) error
	LoadUser(*User) error
	SaveUser(*User) error
	UpdateUser(*User) error

	SessionsForUser(int) ([]Session, error)
	DeleteSession(*Session) error
	LoadSession(*Session) error
	SaveSession(*Session) error
	UpdateSession(*Session) error

