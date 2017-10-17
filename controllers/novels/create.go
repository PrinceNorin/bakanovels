package novelController

import (
	"net/http"

	"github.com/PrinceNorin/bakanovels/models/novels"
	"github.com/gin-gonic/gin"
)

func CreateNovelHandler(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	novel, err := novels.CreateNovel(c)
	if err == nil {
		c.JSON(http.StatusOK, gin.H{"novel": novel})
	} else {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err})
	}
}
