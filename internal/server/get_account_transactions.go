package server

import (
	"fmt"
	"net/http"

	"github.com/faroukelkholy/bank/internal/service/account"
	"github.com/labstack/echo/v4"
)

func GATsHandler(srv account.Service) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		result, err := srv.GetTransactions(c.Param("id"))
		if err != nil {
			fmt.Println("err execute service ", err)
			return c.JSON(http.StatusInternalServerError, HTTPResponse{
				Data: result,
				Err: HTTPError{
					Title:       "internal error",
					Description: "",
				},
			})
		}

		if result == nil {
			return c.JSON(http.StatusNotFound, HTTPResponse{
				Data: result,
				Err: HTTPError{
					Title:       "account not found",
					Description: "",
				},
			})
		}

		return c.JSON(http.StatusOK, HTTPResponse{
			Data: result,
			Err:  HTTPError{},
		})
	}
}

