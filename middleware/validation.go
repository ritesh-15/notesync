package middleware

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/ritesh-15/notesync-backend/global"
	"github.com/ritesh-15/notesync-backend/utils"
)

func Validation(data interface{}, parseFrom string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var err error

		switch parseFrom {
		case "BODY":
			err = c.BodyParser(data)
		case "COOKIE":
			err = c.CookieParser(data)
		case "QUERY":
			err = c.QueryParser(data)
		case "PARAM":
			err = c.ParamsParser(data)
		}

		if err != nil {
			return c.Status(http.StatusUnprocessableEntity).JSON(utils.NewApiError("unprocessable entity", nil))
		}

		if errs := global.MyValidator.Validate(data); len(errs) > 0 && errs[0].Error {
			errMsgs := make([]string, 0)

			for _, err := range errs {
				errMsgs = append(errMsgs, fmt.Sprintf(
					"validation failed on field %s, condition: %s",
					err.FailedField,
					err.Tag,
				))
			}

			return c.Status(http.StatusUnprocessableEntity).JSON(
				utils.NewApiError("unprocessable entity", errMsgs),
			)
		}

		return c.Next()
	}
}
