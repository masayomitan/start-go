package main

import ("github.com/gofiber/fiber/v2"
				"database/sql"
				"fmt"
			
_       "github.com/go-sql-driver/mysql"
)
func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World !")
	})

	db, err := sql.Open("mysql", "root:@/start_go")
  if err != nil {
    panic(err.Error())
  }
  defer db.Close() // 関数がリターンする直前に呼び出される

	// users, err := db.Query("SELECT * FROM users") // 
	todos, err := db.Query("SELECT * FROM todos") // 
  if err != nil {
    panic(err.Error())
  }

	columns, err := todos.Columns() // カラム名を取得
  if err != nil {
    panic(err.Error())
  }

  values := make([]sql.RawBytes, len(columns))

  //  rows.Scan は引数に `[]interface{}`が必要.

  scanArgs := make([]interface{}, len(values))
  for i := range values {
    scanArgs[i] = &values[i]
  }

  for todos.Next() {
    err = todos.Scan(scanArgs...)
    if err != nil {
      panic(err.Error())
    }

		var value string
		
    for i, col := range values {
      // Here we can check if the value is nil (NULL value)
      if col == nil {
        value = "NULL"
      } else {
        value = string(col)
      }
      fmt.Println(columns[i], ": ", value)
    }
    fmt.Println("-----------------------------------")
  }

	app.Listen(":80")
}