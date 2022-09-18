package statusType

import (
	"rental/domain"
	statusType "rental/repository/status_type"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	cr statusType.StatusTypeRepository
}

func NewStatusTypeController(cr statusType.StatusTypeRepository) StatusTypeController {
	return &Controller{
		cr: cr,
	}
}

func (cr *Controller) CreateStatusType(c *gin.Context) {

	var StatusType domain.Status_type

	if err := c.ShouldBind(&StatusType); err != nil {
		c.JSON(400, map[string]string{
			"message": "invalid input",
		})
		return
	}

	if StatusType.Status_type == "" {
		c.JSON(400, map[string]string{
			"message": "StatusType name required",
		})
		return
	}

	err := cr.cr.CreateStatusType(StatusType)

	if err != nil {
		c.JSON(500, map[string]any{
			"error": err.Error(),
		})
		return
	}

	c.JSON(201, map[string]any{
		"message": "success add StatusType",
	})

}

func (cr *Controller) FindAllStatusType(c *gin.Context) {
	StatusTypes, err := cr.cr.FindAllStatusType()
	if err != nil {
		c.JSON(500, map[string]any{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, map[string]any{
		"message": "success find all StatusTypes",
		"data":    StatusTypes,
	})
}

func (cr *Controller) FindByIDStatusType(c *gin.Context) {
	id := c.Param("id")

	StatusType, err := cr.cr.FindByIDStatusType(id)
	if err != nil {
		c.JSON(500, map[string]any{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, map[string]any{
		"message": "success find by id StatusType",
		"data":    StatusType,
	})
}

func (cr *Controller) UpdateStatusType(c *gin.Context) {

	var StatusType domain.Status_type
	// param id
	id := c.Param("id")
	// request body
	if err := c.ShouldBind(&StatusType); err != nil {
		c.JSON(400, map[string]string{
			"message": "invalid input",
		})
		return
	}
	if StatusType.Status_type == "" {
		c.JSON(400, map[string]string{
			"message": "StatusType name required",
		})
		return
	}

	err, message := cr.cr.UpdateStatusType(id, StatusType)
	if err != nil {
		c.JSON(500, map[string]any{
			"message": message,
			"error":   err.Error(),
		})
		return
	}

	c.JSON(200, map[string]any{
		"message": message,
	})
}

func (cr *Controller) DeleteStatusType(c *gin.Context) {

	StatusType := domain.Status_type{}
	// param id
	id := c.Param("id")

	err, message := cr.cr.DeleteStatusType(id, StatusType)
	if err != nil {
		c.JSON(500, map[string]any{
			"message": message,
			"error":   err.Error(),
		})
		return
	}

	c.JSON(200, map[string]any{
		"message": message,
	})
}
