package routes

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PaulWaldo/gomoney/models"
	"github.com/gin-gonic/gin"
)

type SearchData struct {
	PageData
	Results []models.Website
}

func (controller Controller) Search(c *gin.Context) {
	pd := SearchData{
		PageData: PageData{
			Title:           "Search",
			IsAuthenticated: isAuthenticated(c),
		},
	}
	search := c.PostForm("search")

	var results []models.Website

	log.Println(search)
	search = fmt.Sprintf("%s%s%s", "%", search, "%")

	log.Println(search)
	res := controller.db.Where("title LIKE ? OR description LIKE ?", search, search).Find(&results)

	if res.Error != nil || len(results) == 0 {
		pd.Messages = append(pd.Messages, Message{
			Type:    "error",
			Content: "No results found",
		})
		log.Println(res.Error)
		c.HTML(http.StatusOK, "search.html", pd)
		return
	}

	pd.Results = results

	c.HTML(http.StatusOK, "search.html", pd)
}
