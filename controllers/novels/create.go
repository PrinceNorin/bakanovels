package novelController

import (
	"net/http"

	"github.com/PrinceNorin/bakanovels/models"
	"github.com/gin-gonic/gin"
)

func CreateNovelHandler(c *gin.Context) {
	// create novel logic
	var novel models.Novel
	models.DB.First(&novel)
	c.JSON(http.StatusOK, gin.H{"novel": novel})
}
