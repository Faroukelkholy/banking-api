package server

import (
	"fmt"
	"net/http"

	"github.com/faroukelkholy/bank/internal/service/models"
	"github.com/faroukelkholy/bank/internal/service/transaction"
	"github.com/faroukelkholy/bank/internal/storage/postgres"
	"github.com/labstack/echo/v4"
)

func CTHandler(srv transaction.Service) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		var t models.Transfer
		if err = c.Bind(&t); err != nil {
			fmt.Println("err bind account ", err)
			return c.JSON(http.StatusBadRequest, HTTPResponse{
				Data: nil,
				Err: HTTPError{
					Title:       "account data is not valid",
					Description: "",
				},
			})
		}

		if err = srv.CreateTransaction(&t); err != nil {
			fmt.Println("err execute service ", err)
			if err.Error() == postgres.NoAccountID {
				return c.JSON(http.StatusNotFound, HTTPResponse{
					Data: nil,
					Err: HTTPError{
						Title:       postgres.NoAccountID,
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
