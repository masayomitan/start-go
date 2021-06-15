package delivery

import (
		"todo/domain"

		"github.com/gofiber/fiber/v2"
)

type todoSearchHandler struct {
		todoUseCase domain.TodoUsecase
}

func NewTodoSearchHandler ( c * fiber.App, th domain.TodoUsecase ) {
		handler := &todoSearchHandler {
			todoUseCase: th,
		}

		c.Post( "/todo/search", handler.Search )
}

func ( h * todoSearchHandler ) Search ( c * fiber.Ctx ) error {
		todo := new( domain.Todo )
		err := c.BodyParser( todo )
		if err != nil {
				c.Status(400)
				return c.JSON( fiber.Map {
					"message": "結果がありませんでした",
				})
		}
		
		// UseCaseのSearchを呼びだす
		todos, err := h.todoUseCase.Search(todo.Text)
		if err != nil {
			c.Status(500)
			return c.JSON(fiber.Map{
				"message": "サーバーエラーです",
			})
		}
		return c.JSON(todos)
}