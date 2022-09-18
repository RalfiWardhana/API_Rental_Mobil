package user

import (
	"rental/domain"
	user "rental/repository/user"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
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
	emails, _ := cr.cr.FindByEmailUser(User.Email)
	if emails.Email == User.Email {
		c.JSON(400, map[string]string{
			"message": "Email already registered",
		})
		return
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(User.Password), bcrypt.DefaultCost)
	User.Password = string(hashedPassword)

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

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(User.Password), bcrypt.DefaultCost)
	User.Password = string(hashedPassword)
	err, message := cr.cr.UpdateUser(id, User)
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

func (cr *Controller) DeleteUser(c *gin.Context) {

	User := domain.User{}
	// param id
	id := c.Param("id")

	err, message := cr.cr.DeleteUser(id, User)
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

func (cr *Controller) Login(c *gin.Context) {
	var userRequest domain.User
	err := c.ShouldBind(&userRequest)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "invalid input",
		})
		return
	}
	if userRequest.Email == "" || userRequest.Password == "" {
		c.JSON(400, gin.H{
			"message": "email/password required",
		})
		return
	}

	User, err := cr.cr.FindByEmailUser(userRequest.Email)
	if err != nil || User.Id == 0 {
		c.JSON(400, gin.H{
			"message": "wrong email/password",
		})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(User.Password), []byte(userRequest.Password))
	if err != nil {
		c.JSON(400, gin.H{
			"message": "wrong email/password",
		})
		return
	}

	claims := jwt.MapClaims{
		"user_id": User.Id,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
		"iss":     "edspert",
	}
	tkn := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, _ := tkn.SignedString(domain.PrivateKey)
	c.JSON(200, gin.H{
		"token": token,
	})
}
