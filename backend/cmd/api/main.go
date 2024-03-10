package main

import (
	"backend/internal/app/api/pkg/routes"
	"backend/internal/app/api/pkg/validations"

	"github.com/gin-gonic/gin/binding"
)

func main() {
	binding.Validator = validations.NewBindingValidator()
	routes.Run(":8000")
}
