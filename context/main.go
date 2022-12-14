package main

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

func main() {
	p := fmt.Println
	app := fiber.New()

	//accept
	app.Get("/accept", func(c *fiber.Ctx) error {
		c.Accepts("html")
		return c.SendString("accept html")
	})

	app.Get("/abc", func(c *fiber.Ctx) error {
		return c.SendString("htllo world")
	})

	// all perameters
	app.Get("/user/:name", func(c *fiber.Ctx) error {
		par := c.AllParams()
		p(par)
		return c.SendString(par["name"])
	})

	//App stack for all routes
	app.Get("/stack", func(c *fiber.Ctx) error {
		return c.JSON(c.App().Stack())
	})

	// Append will add to header section
	app.Get("/append", func(c *fiber.Ctx) error {
		c.Append("Links", "https://google.com", "http://localhost:3001/abc")
		c.Append("other", "hey this is other in head section ")
		return c.JSON("link append in head section")
	})

	// attachement
	app.Get("/attach", func(c *fiber.Ctx) error {
		c.Attachment()
		c.Attachment("./upload/img/logo.png")
		return c.JSON("attachment")
	})

	//base url
	app.Get("/baseurl", Baseurl)

	// body
	app.Post("/post", post)

	// person
	app.Post("/person", person)

	// cookie
	app.Get("/cookie", cookie)

	//client hello info
	app.Get("/hello", hello)

	// download
	app.Get("dl", download)

	// format
	app.Get("/format", format)

	// form value
	app.Post("/formval", formval)

	// get content type
	app.Get("/getcontent", getcontenttype)

	// get ip
	app.Get("/ip", func(c *fiber.Ctx) error {
		return c.JSON(c.IP())
	})
	// get hostname
	app.Get("/hostname", func(c *fiber.Ctx) error {
		return c.JSON(c.Hostname())
	})

	// json & map
	app.Get("/json", json)

	// jsonp
	app.Get("/jsonp", jsonp)

	// location
	app.Get("/loc", location)

	// method
	app.Post("/method", func(c *fiber.Ctx) error {
		return c.JSON(c.Method())
	})

	// next
	app.Get("/next", func(c *fiber.Ctx) error {
		fmt.Println("1st route!")
		return c.Next()
	})

	app.Get("/next", func(c *fiber.Ctx) error {
		fmt.Println("2nd route!")
		return c.Next()
	})

	app.Get("/next", func(c *fiber.Ctx) error {
		fmt.Println("3rd route!")
		return c.SendString("Hello, World!")
	})

	// path
	app.Get("/path", func(c *fiber.Ctx) error {
		return c.JSON("path name is : " + c.Path())
	})

	// protocol
	app.Get("/protocol", func(c *fiber.Ctx) error {
		return c.JSON("protocol is : " + c.Protocol())
	})

	// body writer
	app.Get("/bodywriter", bodywriter)

	// xml
	app.Get("/xml", xml)

	app.Listen(":3000")
}

func Baseurl(c *fiber.Ctx) error {
	return c.JSON(c.BaseURL()) // "http://localhost:3000"
}

// post
func post(c *fiber.Ctx) error {
	return c.Send(c.Body())
}

type Person struct {
	Name string `json:"name" form:"name" xml:"name"`
	Pass string `json:"pass" form:"pass" xml:"pass"`
}

func person(c *fiber.Ctx) error {
	per := new(Person)
	if err := c.BodyParser(per); err != nil {
		return err
	}
	return c.JSON(per)
}

func cookie(c *fiber.Ctx) error {
	c.ClearCookie()
	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    "lorem text the cookie value is here",
		Expires:  time.Now().Add(-(time.Second * 2)),
		HTTPOnly: true,
		SameSite: "lax",
	})
	return c.JSON("cookies")
}

func hello(c *fiber.Ctx) error {
	chi := c.ClientHelloInfo()
	return c.JSON(chi)
}

func download(c *fiber.Ctx) error {
	return c.Download("../apk/file.txt")
}

func format(c *fiber.Ctx) error {
	return c.Format("hello world")
}
func formval(c *fiber.Ctx) error {
	name := c.FormValue("name")
	email := c.FormValue("email")
	data := Person{
		Name: name,
		Pass: email,
	}
	return c.JSON(data)
}
func getcontenttype(c *fiber.Ctx) error {
	return c.JSON(c.GetReqHeaders())
}

func json(c *fiber.Ctx) error {
	data := Person{
		Name: "smith",
		Pass: "123456",
	}
	// return c.JSON(data)
	return c.JSON(fiber.Map{
		"username": "something it is",
		"userpass": data.Pass,
	})

}

func jsonp(c *fiber.Ctx) error {
	data := Person{
		Name: "Garame",
		Pass: "1233",
	}
	return c.JSONP(data, "customFunc")
}

func location(c *fiber.Ctx) error {
	c.Location("http://example.com")
	c.Location("/foo/bar")
	return nil
}

func bodywriter(c *fiber.Ctx) error {
	c.Response().BodyWriter().Write([]byte("hello world!"))
	return nil
}
func xml(c *fiber.Ctx) error {
	data := Person{
		Name: "Garame",
		Pass: "1233",
	}
	return c.XML(data)
}
