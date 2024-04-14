package api

import (
	mo "github.com/EmilioCliff/portfolio.git/models"
	"github.com/EmilioCliff/portfolio.git/templates"
	"github.com/gin-gonic/gin"
)

func LoadPosts(ctx *gin.Context) {
	links := []mo.Link{
		{Title: "HOME", Path: "/"},
		{Title: "ABOUT", Path: "/about"},
		{Title: "PROJECTS", Path: "Projects"},
		{Title: "GITHUB", Path: "#"},
		{Title: "LINKEDIN", Path: "#"},
	}

	var templateLinks []mo.Link
	for _, link := range links {
		templateLinks = append(templateLinks, mo.Link(link))
	}

	templateBlogs := calculateLayout(BlogList)
	t := templates.IndexPage(templateLinks, templateBlogs)
	t.Render(ctx, ctx.Writer)
}

func calculateLayout(posts []mo.Blog) []mo.Blog {
	var templateBlog []mo.Blog
	for i, post := range posts {
		templateBlog = append(templateBlog, mo.Blog(post))
		switch i % 6 {
		case 0:
			posts[i].LayoutClass = "blog1"
		case 1, 2:
			posts[i].LayoutClass = "blog2"
		case 3, 4, 5:
			posts[i].LayoutClass = "blog3"
		}
	}
	return templateBlog
}
