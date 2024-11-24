package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rifkyrizkita/book_management/controllers"
	"github.com/rifkyrizkita/book_management/middlewares"
)

func UserRouters(user fiber.Router) {
	// post routers
	user.Post("/", middlewares.ValidatorRegister, controllers.Register)
	user.Post("/login", middlewares.ValidatorLogin, controllers.Login)
	// 	// patch routers
	user.Patch("/update-profile", middlewares.VerifyToken, middlewares.ValidatorUpdateProfile, controllers.UpdateProfile)
	user.Patch("/update-password", middlewares.VerifyToken, middlewares.ValidatorUpdatePassword, controllers.UpdatePassword)
	user.Patch("/profile-picture", middlewares.VerifyToken, middlewares.UploadFile("PIMG", ""), controllers.ProfilePicture)
	user.Patch("/reset-password", middlewares.VerifyToken, middlewares.ValidatorResetPassword, controllers.ResetPassword)
	// 	// put routers
	user.Put("/forget-password", middlewares.ValidatorForgetPassword, controllers.ForgetPassword)
	// 	// get routers
	user.Get("/", middlewares.VerifyToken, controllers.Validation)
}
