package publicapi

import (
	"net/http"

	"github.com/d-darac/lagra/internal/api/public_api/v1/groups"
	"github.com/d-darac/lagra/internal/com"
	"github.com/d-darac/lagra/internal/database"
	"github.com/d-darac/lagra/internal/models"
)

var FieldNames = make(com.FieldNames)
var ExpansionConfigs = make(com.ExpansionConfigs)

type publicapi struct {
	groups groups.Handlers
	// groups groups.Service
	// identifiers itemidentifiers.Service
	// inventories inventories.Service
	// items       items.Service
	// person      services.PersonService
	// types       services.TypeService
	// kinds       services.KindService
}

func RegisterHandlers(mux *http.ServeMux, db database.Database) {
	api := publicapi{
		groups: groups.New(db, ExpansionConfigs, FieldNames),
		// groups: groups.NewService(db),
		// identifiers: itemidentifiers.NewService(db),
		// inventories: inventories.NewService(db),
		// items:       items.NewService(db),
	}
	api.buildFieldNames()
	api.buildExpansionConfigs()
	api.registerHandlers(mux)
}

func (api publicapi) registerHandlers(mux *http.ServeMux) {
	// TODO: Implement routes
	mux.HandleFunc("POST /groups", api.groups.Create)
	mux.HandleFunc("DELETE /groups/{id}", api.groups.Delete)
	mux.HandleFunc("GET /groups", api.groups.List)
	mux.HandleFunc("GET /groups/{id}", api.groups.Get)
	mux.HandleFunc("PATCH /groups/{id}", api.groups.Update)

	// mux.HandleFunc("GET /people/{id}", v1.GetPerson)
	// mux.HandleFunc("POST /items", v1.CreateItem)
	// mux.HandleFunc("DELETE /items/{id}", v1.DeleteItem)
	// mux.HandleFunc("GET /items", v1.ListItems)
	// mux.HandleFunc("GET /items/{id}", v1.RetrieveItem)
	// mux.HandleFunc("PATCH /items/{id}", v1.UpdateItem)

	// mux.HandleFunc("POST /inventories", v1.CreateInventory)
	// mux.HandleFunc("DELETE /inventories/{id}", v1.DeleteInventory)
	// mux.HandleFunc("GET /inventories", v1.ListInventories)
	// mux.HandleFunc("GET /inventories/{id}", v1.RetrieveInventory)
	// mux.HandleFunc("PATCH /inventories/{id}", v1.UpdateInventory)

	// mux.HandleFunc("POST /item_identifiers", v1.CreateIdentifiers)
	// mux.HandleFunc("DELETE /item_identifiers/{id}", v1.DeleteIdentifiers)
	// mux.HandleFunc("GET /item_identifiers", v1.ListIdentifiers)
	// mux.HandleFunc("GET /item_identifiers/{id}", v1.RetrieveIdentifiers)
	// mux.HandleFunc("PATCH /item_identifiers/{id}", v1.UpdateIdentifiers)
}

func (api publicapi) buildExpansionConfigs() {
	ExpansionConfigs["group"] = map[string]com.ExpansionConfig{
		"parent_group": {
			Resolver: api.groups,
			IsArray:  false,
		},
	}
}

func (api publicapi) buildFieldNames() {
	com.FieldNamesFromTags(FieldNames, "json", models.Group{})
	// com.FieldNamesFromTags(FieldNames, "json", models.Inventory{})
	// com.FieldNamesFromTags(FieldNames, "json", models.ItemIdentifier{})
	// com.FieldNamesFromTags(FieldNames, "json", models.Item{})
}
