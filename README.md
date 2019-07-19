streamable-go
======

streamable-go is a Go wrapper for working with [Streamable's API](https://streamable.com/documentation)

Use of this library is subject to Streamable's [Terms of Service](https://terms.streamable.com/)

## Installation
`go get github.com/fgoyer/streamable-go`

## Usage

To initialize a streamable client:
````Go
// Streamable utilizes Basic Authentication
streamable := New(email, password)
````

To get a video, pass the known 5 character shortcode or streamable url.
````Go
video, _ := streamable.GetVideoFromShortcode("...")

videoEmbed, _ := streamable.GetVideoFromURL("https://streamable.com/...")
````

To import a video, pass a url:
````Go
result, _ := streamable.Import("...")
````

To upload a video, pass the path:
````Go
result, _ := streamable.Upload("...")
````

## Important Note
All functionality and naming is subject to non-passive changes.

## Acknowledgements
The structure of this api wrapper is heavily inspired by [zmb3/spotify](https://github.com/zmb3/spotify).

