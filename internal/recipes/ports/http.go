package ports

import (
	"net/http"

	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"
)

type HttpServer struct {
}

func NewHttpServer() HttpServer {
	return HttpServer{}
}

func (h HttpServer) GetRecipes(w http.ResponseWriter, r *http.Request, params GetRecipesParams) {
}

func (h HttpServer) CreateRecipe(w http.ResponseWriter, r *http.Request) {

}

func (h HttpServer) DeleteRecipe(w http.ResponseWriter, r *http.Request, recipeUUID openapi_types.UUID) {

}

func (h HttpServer) UpdateRecipe(w http.ResponseWriter, r *http.Request, recipeUUID openapi_types.UUID) {

}
