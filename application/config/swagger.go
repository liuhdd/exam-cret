package config

import (
	"github.com/swaggo/swag"
)

func SetUpSwagger() {
    
	swag.Register(swag.Name, &swag.Spec{
        Title:       "Swagger Example API",
        Description: "This is a sample server for Swagger.",
        Version:     "1.0",
    })
    
}