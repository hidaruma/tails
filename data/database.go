package data

var DB db

type db interface{
	Open() error
	Close() error
	Setup() error
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

	AllAlbums() ([]Album, error)
	LimitAlbums(int, int) ([]Album, error)
	AlbumsForArtist(int) ([]Album, error)
	SearchAlbums(string) ([]Album, error)
	CountAlbums() (int64, error)
	PurgeOrphanAlbums() (int, error)
	DeleteAlbum(*Album) error
	LoadAlbum(*Album) error
	SaveAlbum(*Album) error

	AllFolders() ([]Folder, error)
	LimitFolders(int, int) ([]Folder, error)
	Subfolders(int) ([]Folder, error)
	FoldersInPath(string) ([]Folder, error)
	FoldersNotInPath(string) ([]Folder, error)
	SearchFolders(string) ([]Folder, error)
	CountFolders() (int64, error)
	DeleteFolder(*Folder) error
	LoadFolder(*Folder) error
	SaveFolder(*Folder) error

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

