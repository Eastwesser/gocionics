package main

// Точка входа в приложение
import (
	"fmt"
	userRepo "gocionics/internal/repositories/user"
	userUC "gocionics/internal/usecases/user"
)

func main() {
	fmt.Println("Hello Socionics on Go")

	repo := userRepo.INewUserRepository()
	useCase := userUC.NewUserUseCase(repo)
	controller := userCont.NewUserController(useCase)
	router.POST("/register", controller.Register)
}
