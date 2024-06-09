package AuthController

import (
	"fmt"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-auth/controller/model"
	"github.com/go-auth/database"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

	user := model.User{
		Name: data["name"],
		Email: data["email"],
		Password: hashPassword,
	}

	database.DB.Create(&user)

	

	return c.JSON(fiber.Map{
		"Message": "User was registered!",
		"User": model.UserResponse{
			Id: user.Id,
			Name: user.Name,
			Email: user.Email,
		},
		"Status": 200,
	})
}

const SecretKey = "secret"
func Login(c *fiber.Ctx) error {
	var data map[string]string
	var user model.User

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	database.DB.Where("email=?", data["email"]).First(&user)
	
	err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"]))
	fmt.Println("value", err)


	if user.Email == "" || err != nil {
		c.Status(fiber.StatusBadRequest)

		return c.JSON(fiber.Map{
			"Message": "Invalid email or password.",
			"Status": 404,
		})
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer: strconv.Itoa(int(user.Id)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})

	token, err := claims.SignedString([]byte(SecretKey));
	
	if  err != nil {
		c.Status(fiber.StatusBadRequest)

		return c.JSON(fiber.Map{
			"Message": "Could not login.",
		})
	}

	cookie := fiber.Cookie{
		Name: "jwt",
		Value: token,
		Expires: time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"Message": "Login success!",
		"accessToken": token,
		"User": model.UserResponse{
			Id: user.Id,
			Name: user.Name,
			Email: user.Email,
		},
		"Status": 200,
	})
}
