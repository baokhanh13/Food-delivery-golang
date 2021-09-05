package transport

import (
	"food-delivery/common"
	"food-delivery/component"
	"food-delivery/modules/restaurants/model"
	"food-delivery/modules/restaurants/restaurantbiz"
	"food-delivery/modules/restaurants/storage"
	"github.com/gin-gonic/gin"
)

func ListRestaurant(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		var filter model.Filter

		if err := c.ShouldBind(&filter); err != nil {
			c.JSON(401, gin.H{
				"error": err.Error(),
			})
			return
		}

		var paging common.Paging

		if err := c.ShouldBind(&paging); err != nil {
			c.JSON(401, gin.H{
				"error": err.Error(),
			})
			return
		}

		var result []model.Restaurant

		store := storage.NewSqlStore(appCtx.GetMainDBConnection())
		biz := restaurantbiz.NewListRestaurantBiz(store)
		result, err := biz.ListRestaurant(c.Request.Context(), &paging, &filter)

		if err != nil {
			c.JSON(401, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(200, common.NewSuccesResponse(result, paging, filter))
	}
}
