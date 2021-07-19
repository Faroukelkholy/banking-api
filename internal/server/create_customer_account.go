package server

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/faroukelkholy/bank/internal/service/customer"
	"github.com/faroukelkholy/bank/internal/service/models"
	"github.com/faroukelkholy/bank/internal/storage/postgres"
)

func CCAHandler(srv customer.Service) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		var a models.Account
		if err = c.Bind(&a); err != nil {
			fmt.Println("err bind account ", err)
			return c.JSON(http.StatusBadRequest, HTTPResponse{
				Data: nil,
				Err: HTTPError{
					Title:       "account data is not valid",
					Description: "",
				},
			})
		}

		if err = srv.CreateCustomerAccount(c.Param("id"), &a); err != nil {
			fmt.Println("err execute service ", err)
			if err.Error() == postgres.NoCustomerID {
				return c.JSON(http.StatusNotFound, HTTPResponse{
					Data: nil,
					Err: HTTPError{
						Title:       postgres.NoCustomerID,
						Description: "",
					},
				})
			}

			return c.JSON(http.StatusInternalServerError, HTTPResponse{
				Data: nil,
				Err: HTTPError{
					Title:       "internal error",
					Description: "",
				},
			})
		}

		return c.JSON(http.StatusCreated, HTTPResponse{
			Data: "created",
			Err:  HTTPError{},
		})
	}
}
