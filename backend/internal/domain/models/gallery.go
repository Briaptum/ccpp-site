package models

import "time"

type FacebookPhoto struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_time"`
	Images    []struct {
		Height int    `json:"height"`
		Width  int    `json:"width"`
		Source string `json:"source"`
	} `json:"images"`
	Name string `json:"name,omitempty"`
}

type FacebookAlbum struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	PhotoCount  int    `json:"count"`
}

type FacebookPhotosResponse struct {
	Data []FacebookPhoto `json:"data"`
	Paging struct {
		Cursors struct {
			Before string `json:"before"`
			After  string `json:"after"`
		} `json:"cursors"`
		Next string `json:"next"`
	} `json:"paging"`
}

type GalleryImage struct {
	ID        string    `json:"id"`
	Source    string    `json:"source"`
	Thumbnail string    `json:"thumbnail"`
	Title     string    `json:"title"`
	Alt       string    `json:"alt"`
	CreatedAt time.Time `json:"created_at"`
}
