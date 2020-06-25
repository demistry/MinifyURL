package api

import (
	jsonserializer "MinifyURL/serializer/json"
	"MinifyURL/shortener"
	"log"
	"net/http"

	"github.com/gofiber/fiber"
)


type RedirectHandler interface{
	GET(app *fiber.App)
	POST(app *fiber.App)
}
type handler struct{
	redirectService shortener.RedirectService
}

func NewHandler(redirectService shortener.RedirectService) RedirectHandler{
	return &handler{redirectService: redirectService}
}

func (h *handler) GET(app *fiber.App){
	app.Get("/:code", func (c *fiber.Ctx){
		code := c.Params("code")
		redirect,err := h.redirectService.Find(code)
		if err != nil{
			return
		}
		c.Redirect(redirect.URL, http.StatusMovedPermanently)
	})
}

func (h *handler) POST(app *fiber.App){
	app.Post("/", func (c *fiber.Ctx){
		// red := new(shortener.Redirect)
		// if err := c.BodyParser(red); err != nil{
		// 	log.Fatal("Please pass in correct body")
		// }
		redirect,err := h.serializer().Decode([]byte(c.Body()))
		if err != nil{
			log.Fatal("Please pass in correct body")
		}
		err = h.redirectService.Store(redirect)
		if err != nil{
			log.Fatal("Could not save redirect code")
		}
		responseBody,err := h.serializer().Encode(redirect)
		c.JSON(responseBody)
	})
}

func (h *handler) serializer() shortener.RedirectSerializer{
	return &jsonserializer.Redirect{}
}