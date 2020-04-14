package model

type Link struct {
	// webpage target URL (required)
	Href string `json:"href,omitempty"`
	// Short description of the linked page
	Title string `json:"title,omitempty"`
}

