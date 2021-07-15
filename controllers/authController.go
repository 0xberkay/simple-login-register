package controllers

import (
	"strconv"
	"time"

	"../databases"
	emailC "../mail"
	"../models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber"
	"golang.org/x/crypto/bcrypt"
)

const SecretKey = "secret"

func EmailVerification(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	var mail models.User
	if emailC.IsEmailValid(data["email"]) {

		databases.DB.Find(&mail, "Email = ?", data["email"])
		if !mail.Verified {
			if mail.Id != 0 {
				a := emailC.SendVerifyCode(data["email"])
				mail.Code = a
				databases.DB.Save(&mail)
				return c.JSON(fiber.Map{
					"message": "success",
				})
			} else {
				return c.JSON(fiber.Map{
					"message": "User not found",
				})
			}
		} else {
			return c.JSON(fiber.Map{
				"message": "Email already verified",
			})
		}

	} else {
		return c.JSON(fiber.Map{
			"message": "Incorrect email",
		})
	}
}
func CheckEmailVerification(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	var mail models.User
	databases.DB.Find(&mail, "Email = ?", data["email"])
	if data["code"] == mail.Code {
		mail.Verified = true

		if len(mail.Referance) != 0 {
			var name models.User
			databases.DB.Find(&name, "Name = ?", mail.Referance)
			name.Referances += 1
			name.Points = name.Points + 7
			databases.DB.Save(&name)
		}

		databases.DB.Save(&mail)
	}

	if !mail.Verified {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Email not verified",
		})
	}

	return c.JSON(fiber.Map{
		"message": "success",
	})
}
func ForgetPassword(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	var mail models.User
	if emailC.IsEmailValid(data["email"]) {

		databases.DB.Find(&mail, "Email = ?", data["email"])
		if mail.Verified == true {
			if mail.Id != 0 {
				a := emailC.SendForgetCode(data["email"])
				mail.Code = a
				databases.DB.Save(&mail)
				return c.JSON(fiber.Map{
					"message": "success",
				})
			} else {
				return c.JSON(fiber.Map{
					"message": "User not found",
				})
			}
		} else {
			return c.JSON(fiber.Map{
				"message": "Email already verified",
			})
		}

	} else {
		return c.JSON(fiber.Map{
			"message": "Incorrect email",
		})
	}
}
func ForgetChange(c *fiber.Ctx) error {

	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var user models.User

	databases.DB.Where("Code = ?", data["token"]).First(&user)


	if len(data["newPassword"]) > 6 {

		user.Password, _ = bcrypt.GenerateFromPassword([]byte(data["newPassword"]), 14)

		databases.DB.Save(&user)

	} else {
		return c.JSON(fiber.Map{
			"message": "short password",
		})
	}

	return c.JSON(fiber.Map{
		"message": "success",
	})
}

func Register(c *fiber.Ctx) error {
	var data map[string]string

	var nameCheck models.User

	var mailCheck models.User

	var referanceCheck models.User

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	databases.DB.Find(&nameCheck, "Name = ?", data["name"])
	databases.DB.Find(&mailCheck, "Email = ?", data["email"])
	databases.DB.Find(&referanceCheck, "Name = ?", data["referance"])

	if len(nameCheck.Name) != 0 {
		return c.JSON(fiber.Map{
			"message": "Already registered",
		})
	} else if len(mailCheck.Email) != 0 {
		return c.JSON(fiber.Map{
			"message": "Already registered",
		})

	} else if len(referanceCheck.Name) == 0 && len(data["referance"]) != 0 {
		return c.JSON(fiber.Map{
			"message": "Unkown referance",
		})
	} else if len(data["password"]) < 6 {
		return c.JSON(fiber.Map{
			"message": "Short password",
		})
	} else if !emailC.IsEmailValid(data["email"]) {
		return c.JSON(fiber.Map{
			"message": "Incorrect email",
		})
	} else if (len(data["password"]) == 0) || (len(data["name"]) == 0) || (len(data["email"]) == 0) {
		return c.JSON(fiber.Map{
			"message": "Empty data",
		})

	} else {
		password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)
		referancePoint := 0
		if len(data["referance"]) > 0 {

			referancePoint = 5

		}

		user := models.User{
			Name:       data["name"],
			Email:      data["email"],
			Points:     referancePoint,
			Referance:  data["referance"],
			Referances: 0,
			Verified:   false,
			Created:    time.Now(),
			Password:   password,
		}

		databases.DB.Create(&user)

		return c.JSON(user)
	}
}
func ChangePassword(c *fiber.Ctx) error {

	cookie := c.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	claims := token.Claims.(*jwt.StandardClaims)

	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var user models.User

	databases.DB.Where("id = ?", claims.Issuer).First(&user)

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["oldPassword"])); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Incorrect password",
		})
	}

	if len(data["newPassword"]) > 6 {

		user.Password, _ = bcrypt.GenerateFromPassword([]byte(data["newPassword"]), 14)

		databases.DB.Save(&user)

	} else {
		return c.JSON(fiber.Map{
			"message": "short password",
		})
	}

	return c.JSON(fiber.Map{
		"message": "success",
	})
}

func Login(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var user models.User

	databases.DB.Where("email = ?", data["email"]).First(&user)

	if user.Id == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "User not found",
		})
	}

	if user.Verified == false {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "User not verified",
		})
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Incorrect password",
		})
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(user.Id)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), //1 day
	})

	token, err := claims.SignedString([]byte(SecretKey))

	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "could not login",
		})
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "success",
	})
}

func User(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	claims := token.Claims.(*jwt.StandardClaims)

	var user models.User

	databases.DB.Where("id = ?", claims.Issuer).First(&user)

	return c.JSON(user)
}

func Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "success",
	})
}

func Point(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	claims := token.Claims.(*jwt.StandardClaims)

	var user models.User

	databases.DB.Where("id = ?", claims.Issuer).First(&user)

	user.Points = user.Points + 2

	databases.DB.Save(&user)

	return c.JSON(user)
}

func Market(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	claims := token.Claims.(*jwt.StandardClaims)

	var user models.User

	databases.DB.Where("id = ?", claims.Issuer).First(&user)

	withdraw, _ := strconv.Atoi(c.Params("withdraw"))

	if withdraw > user.Points {
		return c.JSON(fiber.Map{
			"message": "failed",
		})
	} else {
		user.Points = user.Points - withdraw

		databases.DB.Save(&user)

		veri := ("Sipariş veren e posta : " + user.Email + " Sipariş miktarı : " + c.Params("withdraw") + " Sipariş Çeşidi : " + c.Params("market"))

		var data map[string]string

		if err := c.BodyParser(&data); err != nil {
			return err
		}

		order := models.Order{
			Email: user.Email,
			Order: c.Params("market"),
			Price: c.Params("withdraw"),
		}

		databases.DB.Create(&order)

		emailC.SendMail(veri)

		return c.JSON(fiber.Map{
			"message": "success",
		})
	}

}
func SendPoint(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	cookie := c.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	claims := token.Claims.(*jwt.StandardClaims)

	var user models.User
	var nameCheck models.User

	databases.DB.Where("id = ?", claims.Issuer).First(&user)

	i1, err := strconv.Atoi(data["sendingPoints"])
	if err != nil {
		return err
	}
	if user.Points > i1 {
		databases.DB.Find(&nameCheck, "Name = ?", data["receiver"])
		if nameCheck.Id != 0 {
			nameCheck.Points = nameCheck.Points + i1

			databases.DB.Save(&nameCheck)

			user.Points = user.Points - i1

			databases.DB.Save(&user)
			return c.JSON(fiber.Map{
				"message": "success",
			})

		} else {
			return c.JSON(fiber.Map{
				"message": "unkown user",
			})

		}

	} else {
		return c.JSON(fiber.Map{
			"message": "insufficient balance",
		})
	}

}
