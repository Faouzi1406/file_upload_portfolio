package main

import (
	"encoding/base64"
	"fmt"
	"os"

	env "github.com/Faouzi1406/learning/env"
	"github.com/gofiber/fiber/v2"
)

func check_err(e error) {
  if e != nil {
   panic(e);
  }
}

type File struct {
  FileName  string
  FileBlob  string
  ApiKey    string
}

func fileUploading(c *fiber.Ctx) error {
  //Only accept json
  c.Accepts("application.json");

  file := new (File);
  apiKey := os.Getenv("apikey");


  if err := c.BodyParser(file); err != nil  {
    return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
      "error": err.Error(),
    })
  }

  if file.ApiKey != apiKey {
    return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
      "auth":"Unauthorized",
    })
  }
  
  img, err := base64.StdEncoding.DecodeString(file.FileBlob);

  if err != nil {
    return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
      "error": err.Error(),
    });
  } 

  fileName := fmt.Sprint("images/",file.FileName);
  f, err := os.Create(fileName);
  
  if err != nil {
    return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
      "error": err.Error(),
    })
  }

  defer f.Close()

  if _, err := f.Write(img); err != nil {
    return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
      "error": err.Error(),
    })
  }

  if  err = f.Sync(); err != nil {
    return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
      "error": err.Error(),
    })
  }
  
  return c.JSON(fiber.Map{
    "error": "none",
    "succes": "uploaded file",
  });
}

func main() {
  //Env variables 
  env.Loadenv();

  app := fiber.New(fiber.Config{
    AppName: "File uploading portfolio",
  });

  app.Get("/", func(c *fiber.Ctx) error {
    return c.SendString("Some string");
  });

  app.Get("/getfile/:slug", func(c *fiber.Ctx) error {
    param := c.Params("slug");
    folder := os.Getenv("FOLDER");
    param_with_value := fmt.Sprint(folder, param);
    fmt.Println(param_with_value)

    return c.SendFile(param_with_value); 
  });

  app.Post("/uploadFile", func(c *fiber.Ctx) error {
    return fileUploading(c);
  });

  app.Listen(":3000")
}
