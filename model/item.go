package model

// Item contains shorten url 's info
type Item struct {
	URL          string `json:"url"`
	VisitedTimes int    `json:"visited_times" gorethink:"visited_times"`
}

// NewItem return new item
func NewItem(url string) *Item {
	return &Item{
		URL:          url,
		VisitedTimes: 0,
	}
}
