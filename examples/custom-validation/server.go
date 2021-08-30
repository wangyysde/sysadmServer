package main

import (
	"net/http"
	"time"
	"fmt"

	"github.com/wangyysde/sysadmServer"
	"github.com/wangyysde/sysadmServer/binding"
	"github.com/go-playground/validator/v10"
)

// Booking contains binded and validated data.
type Booking struct {
	CheckIn  time.Time `form:"check_in" binding:"required,bookabledate" time_format:"2006-01-02"`
	CheckOut time.Time `form:"check_out" binding:"required,gtfield=CheckIn" time_format:"2006-01-02"`
}

var bookableDate validator.Func = func(fl validator.FieldLevel) bool {
	date, ok := fl.Field().Interface().(time.Time)
	if ok {
		today := time.Now()
		if today.After(date) {
			return false
		}
	}
	return true
}

func main() {
	route := sysadmServer.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("bookabledate", bookableDate)
		fmt.Print("This is runing")
	}

	route.GET("/bookable", getBookable)
	route.Run(":8085")
}

func getBookable(c *sysadmServer.Context) {
	var b Booking
	if err := c.ShouldBindWith(&b, binding.Query); err == nil {
		c.JSON(http.StatusOK, sysadmServer.H{"message": "Booking dates are valid!"})
	} else {
		c.JSON(http.StatusBadRequest, sysadmServer.H{"error": err.Error()})
	}
}
