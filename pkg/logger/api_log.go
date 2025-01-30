package log

import (
	"log"
	"time"

	"github.com/fatih/color"
	"github.com/gofiber/fiber/v2"
)

func ApiLog() fiber.Handler {
	return func(c *fiber.Ctx) error {

		start := time.Now()

		err := c.Next()

		duration := time.Since(start)

		blue := color.New(color.FgBlue).SprintFunc()
		green := color.New(color.FgGreen).SprintFunc()
		yellow := color.New(color.FgYellow).SprintFunc()
		red := color.New(color.FgRed).SprintFunc()

		// log ทั้งหมดในบรรทัดเดียว
		log.Printf("%s %s %s %s",
			blue(">>>> Method: "+c.Method()),
			green(", Path: "+c.Path()),
			yellow(", Time: "+duration.String()),
			func() string {
				if err != nil {
					return red(", Error: " + err.Error())
				}
				return ""
			}(),
		)

		return err
	}
}
