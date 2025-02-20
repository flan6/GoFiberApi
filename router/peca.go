package router

import (
	"api/controllers"
	"net/http"
)

var routePecas = []Route{
	{path: "/pecas",
		method:  http.MethodGet,
		handler: controllers.ListAllPecas,
	},
	{path: "/pecas/descricao/:id",
		method:  http.MethodGet,
		handler: controllers.FindPeca,
	},
	{path: "/pecas",
		method:  http.MethodPost,
		handler: controllers.CreatePeca,
	},
	{path: "/pecas/:id",
		method:  http.MethodPut,
		handler: controllers.EditPeca,
	},

	{path: "/pecas/:id",
		method:  http.MethodDelete,
		handler: controllers.DeletePeca,
	},
	{path: "/pecas/search",
		method:  http.MethodGet,
		handler: controllers.SaerchPeca},
}
