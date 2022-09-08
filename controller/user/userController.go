package user

import (
	"rental/domain"
	user "rental/repository/user"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	cr user.UserRepository
}

func NewUserController(cr user.UserRepository) UserController {
	return &Controller{
		cr: cr,
	}
}

func (cr *Controller) CreateUser(c *gin.Context) {

	var User domain.User

	if err := c.ShouldBind(&User); err != nil {
		c.JSON(400, map[string]string{
			"message": "invalid input",
		})
		return
	}

	if User.Username == "" {
		c.JSON(400, map[string]string{
			"message": "Username required",
		})
		return
	}
	if User.Email == "" {
		c.JSON(400, map[string]string{
			"message": "Email required",
		})
		return
	}
	if User.Password == "" {
		c.JSON(400, map[string]string{
			"message": "Password required",
		})
		return
	}
	if User.Id_user_type == 0 {
		c.JSON(400, map[string]string{
			"message": "Id type user required",
		})
		return
	}

	err := cr.cr.CreateUser(User)

	if err != nil {
		c.JSON(500, map[string]any{
			"error": err.Error(),
		})
		return
	}

	c.JSON(201, map[string]any{
		"message": "success add User",
	})

}

func (cr *Controller) FindAllUser(c *gin.Context) {
	Users, err := cr.cr.FindAllUser()
	if err != nil {
		c.JSON(500, map[string]any{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, map[string]any{
		"message": "success find all Users",
		"data":    Users,
	})
}

func (cr *Controller) FindByIDUser(c *gin.Context) {
	id := c.Param("id")

	User, err := cr.cr.FindByIDUser(id)
	if err != nil {
		c.JSON(500, map[string]any{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, map[string]any{
		"message": "success find by id User",
		"data":    User,
	})
}

func (cr *Controller) UpdateUser(c *gin.Context) {
	// domain user
	var User domain.User
	// param id
	id := c.Param("id")
	// request body
	if err := c.ShouldBind(&User); err != nil {
		c.JSON(400, map[string]string{
			"message": "invalid input",
		})
		return
	}

	if User.Username == "" {
		c.JSON(400, map[string]string{
			"message": "Username required",
		})
		return
	}
	if User.Email == "" {
		c.JSON(400, map[string]string{
			"message": "Email required",
		})
		return
	}
	if User.Password == "" {
		c.JSON(400, map[string]string{
			"message": "Password required",
		})
		return
	}
	if User.Id_user_type == 0 {
		c.JSON(400, map[string]string{
			"message": "Id type user required",
		})
		return
	}

	// update user
	err, message := cr.cr.UpdateUser(id, User)
	if err != nil {
		c.JSON(500, map[string]any{
			"message": message,
			"error":   err.Error(),
		})
		return
	}
	// return success update user
	c.JSON(200, map[string]any{
		"message": message,
	})
}

func (cr *Controller) DeleteUser(c *gin.Context) {
	// domain user
	User := domain.User{}
	// param id
	id := c.Param("id")

	// delete user
	err, message := cr.cr.DeleteUser(id, User)
	if err != nil {
		c.JSON(500, map[string]any{
			"message": message,
			"error":   err.Error(),
		})
		return
	}
	// return success delete user
	c.JSON(200, map[string]any{
		"message": message,
	})
}
