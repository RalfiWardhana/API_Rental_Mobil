package userType

import (
	"rental/domain"
	userType "rental/repository/user_type"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	cr userType.UserTypeRepository
}

func NewUserTypeController(cr userType.UserTypeRepository) UserTypeController {
	return &Controller{
		cr: cr,
	}
}

func (cr *Controller) CreateUserType(c *gin.Context) {

	var UserType domain.User_type

	if err := c.ShouldBind(&UserType); err != nil {
		c.JSON(400, map[string]string{
			"message": "invalid input",
		})
		return
	}

	if UserType.User_type == "" {
		c.JSON(400, map[string]string{
			"message": "UserType name required",
		})
		return
	}

	err := cr.cr.CreateUserType(UserType)

	if err != nil {
		c.JSON(500, map[string]any{
			"error": err.Error(),
		})
		return
	}

	c.JSON(201, map[string]any{
		"message": "success add UserType",
	})

}

func (cr *Controller) FindAllUserType(c *gin.Context) {
	UserTypes, err := cr.cr.FindAllUserType()
	if err != nil {
		c.JSON(500, map[string]any{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, map[string]any{
		"message": "success find all UserTypes",
		"data":    UserTypes,
	})
}

func (cr *Controller) FindByIDUserType(c *gin.Context) {
	id := c.Param("id")

	UserType, err := cr.cr.FindByIDUserType(id)
	if err != nil {
		c.JSON(500, map[string]any{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, map[string]any{
		"message": "success find by id UserType",
		"data":    UserType,
	})
}

func (cr *Controller) UpdateUserType(c *gin.Context) {

	var UserType domain.User_type
	// param id
	id := c.Param("id")
	// request body
	if err := c.ShouldBind(&UserType); err != nil {
		c.JSON(400, map[string]string{
			"message": "invalid input",
		})
		return
	}
	if UserType.User_type == "" {
		c.JSON(400, map[string]string{
			"message": "UserType name required",
		})
		return
	}

	err, message := cr.cr.UpdateUserType(id, UserType)
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

func (cr *Controller) DeleteUserType(c *gin.Context) {

	UserType := domain.User_type{}
	// param id
	id := c.Param("id")

	err, message := cr.cr.DeleteUserType(id, UserType)
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
