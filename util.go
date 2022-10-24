package main

import (
	"strconv"
	"strings"
)

func AddSongToPlaylist(songId int, playlistId int) {
	songIdStr := strings.ToLower(strconv.FormatInt(int64(songId), 10))
	playlistIdStr := strings.ToLower(strconv.FormatInt(int64(playlistId), 10))

	for i, playlist := range SpotifyObj.Playlists {
		if strings.EqualFold(playlist.Id, playlistIdStr) && !Contains(playlist.Song_ids, songIdStr) {
			playlist.Song_ids = append(playlist.Song_ids, songIdStr)
			SpotifyObj.Playlists[i] = playlist
		}

	}
}

func RemovePlaylistFromSpotify(playlistId int) {
	playlistIdStr := strings.ToLower(strconv.FormatInt(int64(playlistId), 10))

	index := 0
	for i, playlist := range SpotifyObj.Playlists {
		if playlist.Id == playlistIdStr {
			index = i
		}
	}
	SpotifyObj.Playlists = append(SpotifyObj.Playlists[:index], SpotifyObj.Playlists[index+1:]...)
}

// https://play.golang.org/p/Qg_uv_inCek
// contains checks if a string is present in a slice
func Contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}
