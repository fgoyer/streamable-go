package streamable

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

const baseAddress = "https://api.streamable.com/"

// Streamable is a client for working with the Streamable Web API
// TODO document these fields
type Streamable struct {
	http    *http.Client
	baseURL string

	email    string
	password string
}

// Mp4 is a sub struct of the Video struct
// TODO should this be exported?
type mp4 struct {
	Status    int     `json:"status"`
	Width     int     `json:"width"`
	URL       string  `json:"url"`
	Bitrate   int     `json:"bitrate"`
	Duration  float64 `json:"duration"`
	Size      int     `json:"size"`
	Framerate int     `json:"framerate"`
	Height    int     `json:"height"`
}

// Video contains file information and metadata about a video.
type Video struct {
	Status int `json:"status"`
	Files  struct {
		Mp4       mp4 `json:"mp4"`
		Mp4Mobile mp4 `json:"mp4-mobile"`
	} `json:"files"`
	EmbedCode    string `json:"embed_code"`
	Source       string `json:"source"`
	ThumbnailURL string `json:"thumbnail_url"`
	URL          string `json:"url"`
	Message      string `json:"message"`
	Title        string `json:"title"`
	Percent      int    `json:"percent"`
}

// VideoEmbed contains the embed code and metadata about a video.
type VideoEmbed struct {
	ProviderURL  string `json:"provider_url"`
	HTML         string `json:"html"`
	Version      string `json:"version"`
	Title        string `json:"title"`
	Type         string `json:"type"`
	ProviderName string `json:"provider_name"`
	ThumbnailURL string `json:"thumbnail_url"`
	Width        int    `json:"width"`
	Height       int    `json:"height"`
}

// Result contains the response of a video import or upload request.
type Result struct {
	Status    int    `json:"status"`
	Shortcode string `json:"shortcode"`
}

// New creates a Streamable Client that will use the specified
// email and password for its API requests.
func New(email string, pass string) *Streamable {
	client := &http.Client{}
	return &Streamable{
		http:     client,
		baseURL:  baseAddress,
		email:    email,
		password: pass,
	}
}

// GetVideoFromShortcode retrieves the raw mp4 video files
// using the specified Streamable video shortcode.
func (c *Streamable) GetVideoFromShortcode(shortcode string) (*Video, error) {
	url := fmt.Sprintf("%svideos/%s", c.baseURL, shortcode)

	var video Video
	err := c.get(url, &video)
	if err != nil {
		return nil, err
	}

	return &video, nil
}

// GetVideoFromURL retrieves the video embed information
// using the specified Streamable URL.
func (c *Streamable) GetVideoFromURL(url string) (*VideoEmbed, error) {
	requestURL := fmt.Sprintf("%soembed.json?url=%s", c.baseURL, url)

	var video VideoEmbed
	err := c.get(requestURL, &video)
	if err != nil {
		return nil, err
	}

	return &video, nil
}

// Import uses the specified video url to import a video
// to Streamable.
func (c *Streamable) Import(videoURL string) (*Result, error) {
	url := fmt.Sprintf("%simport?url=%s", c.baseURL, videoURL)

	var result Result
	err := c.get(url, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// Upload uses the specified path to upload a video to Streamable.
func (c *Streamable) Upload(path string) (*Result, error) {
	url := fmt.Sprintf("%supload", c.baseURL)

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Write the bytes of the opened file into a form-data header.
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", filepath.Base(file.Name()))
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(part, file)
	if err != nil {
		return nil, err
	}
	err = writer.Close()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", writer.FormDataContentType())

	var result Result
	err = c.do(req, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Streamable) get(url string, result interface{}) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	return c.do(req, result)
}

func (c *Streamable) do(req *http.Request, result interface{}) error {
	req.SetBasicAuth(c.email, c.password)

	resp, err := c.http.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(result)
	if err != nil {
		return err
	}

	return nil
}
