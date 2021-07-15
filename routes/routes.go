package routes

import (
	"../controllers"
	"github.com/gofiber/fiber"
)

func Setup(app *fiber.App) {

	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)
	app.Get("/api/user", controllers.User)
	app.Post("/api/logout", controllers.Logout)
	app.Post("/api/change", controllers.ChangePassword)
	app.Post("/api/send", controllers.SendPoint)
	app.Get("/api/point", controllers.Point)
	app.Get("/api/:market/:withdraw", controllers.Market)
	app.Get("/api/upgrade", controllers.Upgrade)
	app.Get("/api/most", controllers.MostUser)
	app.Post("/api/check", controllers.EmailVerification)
	app.Post("/api/setcheck", controllers.CheckEmailVerification)
	app.Post("/api/forget", controllers.ForgetPassword)
	app.Post("/api/setforget", controllers.ForgetChange)
}
