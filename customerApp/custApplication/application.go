package custApplication

import "github.com/gin-gonic/gin"

var (
	router = gin.Default()
)

func StartCustApplication() {
	MapUrls()
	router.Run(":8080")
}