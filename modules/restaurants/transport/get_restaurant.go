package transport

import (
	"food-delivery/common"
	"food-delivery/component"
	"food-delivery/modules/restaurants/restaurantbiz"
	"food-delivery/modules/restaurants/storage"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetRestaurant(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(401, gin.H{
				"error": err.Error(),
			})
			return
		}

		store := storage.NewSqlStore(appCtx.GetMainDBConnection())
		biz := restaurantbiz.NewGetRestaurantBiz(store)

		data, err := biz.GetRestaurant(c.Request.Context(), id)

		if err != nil {
			c.JSON(401, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(200, common.SimpleSuccessResponse(data))
	}
}
