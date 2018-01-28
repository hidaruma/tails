/*
Copyright (C) 2014 Matt Layher
Released under the MIT license
http://opensource.org/licenses/mit-license.php
 */

package subsonic

import (
	"encoding/xml"
	"net/http"
	"context"
	"path"
	"strconv"
	"time"
	"fmt"
	"bytes"

	"github.com/hidaruma/tails/subsonic/songlists"
	"github.com/hidaruma/tails/subsonic/playlists"
	"github.com/hidaruma/tails/subsonic/browsing"
	"github.com/hidaruma/tails/subsonic/system"
	"github.com/hidaruma/tails/subsonic/searching"
	"github.com/hidaruma/tails/subsonic/mediaannotation"
	"github.com/hidaruma/tails/subsonic/mediaretrieval"

	"github.com/hidaruma/tails/data"
	"path/filepath"
)

const (
	XMLName = "subsonic-response"

	XMLNS = "http://subsonic.org/restapi"

	Version = "1.16.0"
)

type Container struct {
	XMLName xml.Name `xml:"subsonic-response"`

	XMLNS string `xml:"xmlns,attr"`
	Status string `xml:"status,attr"`
	Version string `xml:"version,attr"`

	SubError *Error

	Album []Album `xml:"album"`

	AlbumList *songlists.AlbumListContainer `xml:"albumList"`

	AlbumList2 *songlists.AlbumList2Container `xml:"albumList2"`

	NowPlaying *songlists.NowPlayingContainer `xml:"nowPlaying"`

	Indexes *IndexesContainer `xml:"indexes"`

	License *system.License `xml:"license"`



	MusicDirectory *MusicdirectoryContainer

	MusicFolders *MusicFoldersContainer

	Playlist *Playlist `xml:"playlist"`

	Playlists *playlists.PlaylistsContainer `xml:"playlists"`

	RandomSongs *songlists.RandomSongsContainer

	Starred *songlists.Starred `xml:"starred"`

	Starred2 *songlists.Starred2 `xml:"starred2"`

}

func NewContainer() *Container {
	return &Container{
		XMLNS: XMLNS,
		Status: "ok",
		Version: Version,
	}
}

func (c *Container) ToXML() (string, error) {
	buf := bytes.NewBuffer(nil)
	buf, err := xml.MarshalIndent(c, "", " ")
	if err != nil {
		return nil, err
	}

	return string(buf.Bytes()), err
}

func (c Container) Error() string {
	return fmt.Sprintf("%d: %s", c.SubError.Code, c.SubError.Message)
}

type Error struct {
	XMLName xml.Name `xml:"error,omitempty"`

	Code int `xml:"code,attr"`
	Message string `xml:"message,attr"`
}

type SubsonicArtist struct {
	XMLName xml.Name `xml:"artist,omitempty"`

	Name string `xml:"name"`
	ID string `xml:"id,attr"`
}

type SubsonicAlbum struct {
	ID int `xml:"id,attr"`
	Name string `xml:"name,attr"`
	Artist string `xml:"artist,attr"`
	ArtistID int `xml:"artistId,attr"`
	CoverArt string `xml:"coverArt,attr"`
	SongCount int `xml:"songCount,attr"`
	Duration int `xml:"duration,attr"`
	Created string `xml:"created,attr"`

	Songs []SubsonicSong `xml:"song"`
}

func SubAlbum(album data.Album, songs []data.Song) SubsonicAlbum {
	return SubsonicAlbum{
		ID: album.ID,
		Name: album.Title,
		Artist: album.Artist,
		ArtistID: album.ArtistID,
		CoverArt: strconv.Itoa(songs[0].ArtID),
		SongCount: len(songs),
		Duration: songs.Length(),
		Created: subTime(songs[0].LastModified),
	}
}

type SubsonicSong struct {
	ID int `xml:"id,attr"`
	Parent int `xml:"parent,attr"`
	Title string `xml:"title,attr"`
	Album string `xml:"album,attr"`
	Artist string `xml:"artist,attr"`
	IsDir bool `xml:"isDir,attr"`
	CoverArt string `xml:"coverArt,attr"`
	Created string `xml:"created,attr"`
	Duration int `xml:"duration,attr"`
	BitRate int `xml:"bitRate,attr"`
	Track int `xml:"track,attr"`
	DiscNumber int `xml:"discNumber,attr"`
	Year int `xml:"year,attr"`
	Genre string `xml:"genre,attr"`
	Size int64 `xml:"size,attr"`
	Suffix string `xml:"suffix,attr"`
	ContentType string `xml:"contentType,attr"`
	IsVideo bool `xml:"isVideo,attr"`
	Path string `xml:"path,attr"`
	AlbumID int `xml:"albumId,attr"`
	ArtistID int `xml:"artistId,attr"`
	Type string `xml:"type,attr"`
}

func SubSong(song data.Song) SubsonicSong {
	return SubsonicSong{
		ID: song.ID,
		Parent: 0,
		Title: song.Title,
		Album: song.Album,
		Artist: song.Artist,
		IsDir: false,
		CoverArt: strconv.Itoa(song.CoverID),
		Created: subTime(song.LastModified),
		Duration: song.Length,
		BitRate: song.BitRate,
		Track: song.Track,
		DiscNumber: song.DiscNumber,
		Year: song.Year,
		Genre: song.Genre,
		Size: song.FileSize,
		Suffix: filepath.Ext(song.FileName)[1:],
		ContentType: data.MIMEMap[song.FileTypeID],
		IsVideo: false,
		Path: song.FileName,
		AlbumID: song.AlbumID,
		ArtistID: song.ArtistID,
		Type: "music",
	}
}

type MusicFolder struct {
	ID int `xml:"id,attr"`
	Name string `xml:"name,attr"`
}

func subTime(unix int64) string {
	return time.Unix(unix, 0).Format("2006-01-02T15:04:05")
}