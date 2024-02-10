package main

import (
	"backend/internal/app/api/routes"
	"backend/internal/app/api/validations"

	"github.com/gin-gonic/gin/binding"
)

func main() {
	binding.Validator = validations.NewBindingValidator()
	routes.Run(":8000")
}
