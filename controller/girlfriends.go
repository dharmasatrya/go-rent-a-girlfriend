package controller

import (
	"rent-a-girlfriend/service"
	"strconv"

	"github.com/labstack/echo/v4"
)

// roomController is the controller for room-related operations.
type girlfriendController struct {
	girlfriendService service.GirlfriendService
}

// NewRoomController creates a new instance of roomController.
func NewGirlfriendController(girlfriendService service.GirlfriendService) *girlfriendController {
	return &girlfriendController{girlfriendService}
}

func (h *girlfriendController) GetGirlById(c echo.Context) error {
	girlIdStr := c.Param("id")
	girlId, _ := strconv.Atoi(girlIdStr)
	status, webResponse := h.girlfriendService.GetGirlfriendById(girlId)

	return c.JSON(status, webResponse)
}
