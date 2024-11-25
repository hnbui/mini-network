package migrations

import (
	"database/sql"
	"errors"

	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/models"
)

func register0001Migration(app core.App) error {
	if _, err := app.Dao().FindCollectionByNameOrId("ep_users"); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			app.Logger().With("error", err).Error("failed to find ep_users collection")
			return err
		}
	}

	collection := &models.Collection{
		Name: "ep_users",
		Type: models.CollectionTypeBase,
	}

	return nil
}
