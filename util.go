package main

import (
	"strconv"
	"strings"
)

func AddSongToPlaylist(songId int, playlistId int) {
	songIdStr := strings.ToLower(strconv.FormatInt(int64(songId), 10))
	playlistIdStr := strings.ToLower(strconv.FormatInt(int64(playlistId), 10))

	for i, playlist := range SpotifyObj.Playlists {
		if strings.EqualFold(playlist.Id, playlistIdStr) && !Contains(playlist.Song_Ids, songIdStr) {
			playlist.Song_Ids = append(playlist.Song_Ids, songIdStr)
			SpotifyObj.Playlists[i] = playlist
		}
	}
}

func AddPlaylistForUser(userId int, newPlaylist Playlist) {
	userIdStr := strings.ToLower(strconv.FormatInt(int64(userId), 10))

	// I'm assuming list is ordered by id
	// Increment last id by 1 and set to newPlaylist.Owner_Id
	lastPlaylistId, _ := strconv.Atoi(SpotifyObj.Playlists[len(SpotifyObj.Playlists)-1].Id)
	newPlaylistId := lastPlaylistId + 1
	newPlaylist.Id = strconv.FormatInt(int64(newPlaylistId), 10)
	newPlaylist.Owner_Id = userIdStr

	// Add new playlist to playlists
	SpotifyObj.Playlists = append(SpotifyObj.Playlists, newPlaylist)
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
