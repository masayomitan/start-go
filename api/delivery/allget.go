package delivery

import(
	"todo/domain"

	"github.com/gofiber/fiber/v2"
)

type todoAllGetHandler struct {
    todoUseCase domain.TodoUseCase
}
