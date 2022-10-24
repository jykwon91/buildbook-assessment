package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
)

var SpotifyObj Spotify
var Changes []Change

func main() {

	// In order to not include file path logic, we're just assuming all files are in
	// current working directory
	inputPtr := flag.String("input", "spotify.json", "Specify input file path.")
	changesPtr := flag.String("changes", "changes.json", "Specify changes file path.")
	outputPtr := flag.String("output", "output.json", "Specify output file path.")
	flag.Parse()

	bytes, err := ioutil.ReadFile(*inputPtr)
	if err != nil {
		fmt.Println(err.Error())
	}
	json.Unmarshal(bytes, &SpotifyObj)

	bytes, err = ioutil.ReadFile(*changesPtr)
	if err != nil {
		fmt.Println(err.Error())
	}
	json.Unmarshal(bytes, &Changes)

	for _, change := range Changes {
		switch change.ChangeType {
		case AddSong:
			AddSongToPlaylist(change.SongId, change.PlaylistId)
			break
		case AddPlaylist:
			AddPlaylistForUser(change.UserId, change.Playlist)
			break
		case RemovePlaylist:
			RemovePlaylistFromSpotify(change.PlaylistId)
			break
		default:
			break
		}
	}

	bytes, err = json.Marshal(SpotifyObj)
	if err != nil {
		fmt.Println(err.Error())
	}

	err = ioutil.WriteFile(*outputPtr, bytes, 0644)
	if err != nil {
		fmt.Println(err.Error())
	}
}
