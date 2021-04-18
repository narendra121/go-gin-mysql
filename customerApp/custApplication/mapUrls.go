package custApplication

import "github.com/narendra121/go-gin-mysql/customerApp/controller/custController"

func MapUrls() {
	urls:=router.Group("/customers")
	urls.POST("/create", custController.Create)
	urls.GET("display/:custid",custController.Search)
	urls.PUT("update/:custid",custController.Update)
	urls.DELETE("delete/:custid",custController.Delete)

}