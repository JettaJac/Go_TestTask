package endpoint

import (
	"fmt"
	"net/http"
	// "time"
	"github.com/labstack/echo/v4"
)
type Service interface {
	DayLeft() int64
}
type Endpoint struct {
	s Service
}

func New(s Service) *Endpoint {
	return &Endpoint{
		s:s,
	}
}

func (e *Endpoint) Status (ctx echo.Context) error {

d := e.s.DayLeft()	

// 	d := time.Date(2025, time.January, 1,0,0,0,0,time.UTC)
// 	dure := time.Until(d) 

s := fmt.Sprintf("Количество дней: %d", d)
 
err := ctx.String(http.StatusOK, s)

 if err != nil {
	return nil
 }
 return nil
}