package api

import (
	"time"

	mo "github.com/EmilioCliff/portfolio.git/models"
	"github.com/gin-gonic/gin"
)

//	type Server struct{
//		store *pgx
//	}
var BlogList []mo.Blog

func StartServer(address string) error {
	createdTime, _ := time.Parse("2006-01-02", "2022-02-22")

	blog1 := mo.Blog{
		Title:     "First Blog",
		Category:  "Software Engineering",
		Content:   "This will be my Software Engineering content",
		Author:    "Emilio Cliff",
		ImagePath: "static/images/piano.jpg",
		CreatedAt: createdTime.Format("2006-01-02"),
	}
	blog2 := mo.Blog{
		Title:     "Second Blog",
		Category:  "Cryptography",
		Content:   "This will be my Cryptography content",
		Author:    "Emilio Cliff",
		ImagePath: "static/images/piano2.jpg",
		CreatedAt: createdTime.Format("2006-01-02"),
	}
	blog3 := mo.Blog{
		Title:     "Second Blog",
		Category:  "Cryptography",
		Content:   "This will be my Cryptography content",
		Author:    "Emilio Cliff",
		ImagePath: "static/images/piano2.jpg",
		CreatedAt: createdTime.Format("2006-01-02"),
	}
	blog4 := mo.Blog{
		Title:     "Second Blog",
		Category:  "Cryptography",
		Content:   "This will be my Cryptography content",
		Author:    "Emilio Cliff",
		ImagePath: "static/images/piano2.jpg",
		CreatedAt: createdTime.Format("2006-01-02"),
	}
	blog5 := mo.Blog{
		Title:     "Second Blog",
		Category:  "Cryptography",
		Content:   "This will be my Cryptography content",
		Author:    "Emilio Cliff",
		ImagePath: "static/images/piano2.jpg",
		CreatedAt: createdTime.Format("2006-01-02"),
	}
	blog6 := mo.Blog{
		Title:     "Second Blog",
		Category:  "Cryptography",
		Content:   "This will be my Cryptography content",
		Author:    "Emilio Cliff",
		ImagePath: "static/images/piano2.jpg",
		CreatedAt: createdTime.Format("2006-01-02"),
	}
	blog7 := mo.Blog{
		Title:     "Second Blog",
		Category:  "Cryptography",
		Content:   "This will be my Cryptography content",
		Author:    "Emilio Cliff",
		ImagePath: "static/images/piano2.jpg",
		CreatedAt: createdTime.Format("2006-01-02"),
	}
	blog8 := mo.Blog{
		Title:     "Second Blog",
		Category:  "Cryptography",
		Content:   "This will be my Cryptography content",
		Author:    "Emilio Cliff",
		ImagePath: "static/images/piano2.jpg",
		CreatedAt: createdTime.Format("2006-01-02"),
	}
	blog9 := mo.Blog{
		Title:     "Second Blog",
		Category:  "Cryptography",
		Content:   "This will be my Cryptography content",
		Author:    "Emilio Cliff",
		ImagePath: "static/images/piano2.jpg",
		CreatedAt: createdTime.Format("2006-01-02"),
	}

	BlogList = append(BlogList, blog1, blog2, blog3, blog4, blog5, blog6, blog7, blog8, blog9)

	r := gin.Default()

	r.Static("/static", "./static")

	r.GET("/", LoadPosts)

	return r.Run(address)
}
