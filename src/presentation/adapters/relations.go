package adapters

import (
	"go.uber.org/dig"
)

type RelationsContainer struct {
	dig.In

	Book BookRelations
}
