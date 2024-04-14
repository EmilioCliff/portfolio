package mo

import "time"

type Blog struct {
	Title     string
	Category  string
	Content   string
	Author    string
	ImagePath string
	// change the layout class everytime i add a new blog
	LayoutClass string
	Duration    time.Duration
	CreatedAt   string
}

type Link struct {
	Title string
	Path  string
}
