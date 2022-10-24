package main

type User struct {
	Id   string
	Name string
}

type Playlist struct {
	Id       string
	Owner_id string
	Song_ids []string
}

type Song struct {
	Id     string
	Artist string
	Title  string
}

type Spotify struct {
	Users     []User
	Playlists []Playlist
	Songs     []Song
}

type ChangeType int

const (
	AddSong ChangeType = iota + 1
	RemoveSong
	AddPlaylist
	RemovePlaylist
)

type Change struct {
	ChangeType ChangeType
	UserId     int      `json:"user_id,omitempty"`
	PlaylistId int      `json:"playlist_id,omitempty"`
	SongId     int      `json:"song_id,omitempty"`
	Playlist   Playlist `json:"playlist,omitempty"`
}
