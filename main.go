package main

// mkdir myapp && cd myapp
// $ go mod init myapp
// $ go get github.com/labstack/echo/v4

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"github.com/labstack/echo/v4"
)

func main (){
	fmt.Println("Start Server")
	s := echo.New()
	s.Use(MW) // инициализоровано глобально. Будет применим ко всем обработчикам, далее мы его применяем только к одному разработчику
	s.GET("/status", Handler, MW) // MV тут, что был применем только к этому обработчику
	err := s.Start(":8080")
	if err != nil {
		log.Fatal(err)
	}
}

func Handler (ctx echo.Context) error {
		d := time.Date(2025, time.January, 1,0,0,0,0,time.UTC)
		dure := time.Until(d)
		s := fmt.Sprintf("Количество дней: %d", int64(dure.Hours())/24)
	 err := ctx.String(http.StatusOK, s)
	 if err != nil {
		return nil
	 }
	 return nil
}

func MW (next echo.HandlerFunc) echo.HandlerFunc{
	return func (ctx echo.Context) error{
		val := ctx.Request().Header.Get("User-Role")
		if val == "admin" {
			log.Println("red button user datected")
		}
		err := next(ctx)
		if err != nil {
			return err
		}
		return nil
	}

	
}