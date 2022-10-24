# Housecall Pro Take Home Test

## How to run:

1. Follow the steps to install golang: https://go.dev/doc/install
2. Running from source:
   1. After installing go:
      1. Create spotify.json and changes.json files(it will error if those files don't exist)
      2. Execute `go run main.go -input=spotify.json -changes=changes.json -output=output.json`
3. Running from binary:
4. Execute `./main -input=spotify.json -changes=changes.json -output=output.json`

## Input structure:

ChangeType Enum:
AddSong = 1
RemoveSong = 2
AddPlaylist = 3
RemovePlaylist = 4

changes.json examples:

```
[{
  // Add a song to a playlist
  {
    changeType: 1,
    song_id: 1,
    playlist_id: 1
  },
  // Remove a playlist
  {
    changeType: 4,
    playlist_id: 1
  },
  // Add playlist
  {
    changeType: 3,
    user_id: 1,
    playlist: {
      song_ids: ["1"]
    }
  }
}]
```

## Output validation:

1. Manually verify list is updated in `output.json`.

Ideally, we would have test cases for the action functions(adding, removing, updating) to ensure functionality.

## Scalability:

- Use an API to handle updating data instead of manually editing json files.
- Instead of using json files, I'd incorporate a relational database and querying tool in the backend API and data layer service in order to handle larger data sets.
- Messaging system to handle incoming requests.

## Design decisions:

- Golang:
  1. Easy to install and develop server side projects(API, web server, etc).
  2. Build process and deployment is easy.
  3. Great for reading/writing files.
- Json files for database:
  1. Easy to work with for small projects.

## Final thoughts:

- Spent 3-4 hours on this project.
- It's been awhile since I worked in golang so there's many things I could have done to improve code efficiency and organization. I put everything in the main package when it could be separate packages in order to achieve a more modular project. I was also slow to implement since I had to look everything up.
- I didn't set up the project for an object oriented design.
- Many missing error check:
  - When adding playlists, make sure 1 song exists
  - When adding songs, make sure it exists
  - When adding playlist for user, make sure user exists
  - Throw error when any of the above fail
  - Make sure we're not duplicating records
