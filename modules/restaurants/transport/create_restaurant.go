package transport

import (
	"food-delivery/common"
	"food-delivery/component"
	"food-delivery/modules/restaurants/model"
	"food-delivery/modules/restaurants/restaurantbiz"
	"food-delivery/modules/restaurants/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateRestaurant(appContext component.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {
		var data model.RestaurantCreate

		if err := context.ShouldBind(&data); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		store := storage.NewSqlStore(appContext.GetMainDBConnection())
		biz := restaurantbiz.NewCreateRestaurantBiz(store)

		if err := biz.CreateRestaurant(context.Request.Context(), &data); err != nil {
			context.JSON(401, gin.H{
				"error": err.Error(),
			})

			return
		}
		context.JSON(200, common.SimpleSuccessResponse(data))
	}
}
