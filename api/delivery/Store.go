package delivery

import (
    "todo/domain"
		"github.com/gofiber/fiber/v2"
) 

type todoStoreHandler struct {
    todoUseCase domain.TodoUsecase
}

func NewTodoStoreHandler( c *fiber.App, th domain.TodoUsecase ) {
	handler := &todoStoreHandler{
		todoUseCase: th,
	}
	c.Post( "/todo/store", handler.Store )
}

func (h * todoStoreHandler ) Store( c * fiber.Ctx ) error {
	todo := new( domain.Todo )
	err := c.BodyParser( todo )
	if err != nil {
		c.Status( 400 )
		return c.JSON( fiber.Map {
			"message": "エラーってるよ〜",
		})
	}
	err = h.todoUseCase.Store( * todo )
	if err != nil {
		c.Status( 500 )
		return c.JSON( fiber.Map {
			"message": "サーバーがおかしいよ〜",
		})
	}
	return c.JSON( "Store OK" )
}