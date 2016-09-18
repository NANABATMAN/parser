package main

type Link struct {
	Url   string `json:"url"`
	Title string `json:"title"`
}

func (link *Link) setUrl(url string) {
	link.Url = url
}

func (link *Link) setTitle(title string) {
	link.Title = title
}
