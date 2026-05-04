package itemscmd

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/d-darac/lagra/cmd/lagra-mock/cmd/shared"
	"github.com/d-darac/lagra/internal/com"
	"github.com/d-darac/lagra/internal/database/sqlc"
	"github.com/d-darac/lagra/internal/services/items"
	"github.com/d-darac/lagra/pkg/currency"
	"github.com/d-darac/lagra/pkg/i32"
	"github.com/d-darac/lagra/pkg/id"
	"github.com/d-darac/lagra/pkg/str"
	"github.com/spf13/cobra"
)

var groupID id.NullID
var inventoryID id.NullID

// ItemsCmd represents the items command
var ItemsCmd = &cobra.Command{
	Use:               "items N",
	Short:             "Use to create items",
	PersistentPreRun:  shared.RequireConfig,
	PersistentPreRunE: shared.SetAccID,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		groupIDStr, err := cmd.Flags().GetString("group")
		if err != nil {
			return err
		}
		if groupIDStr != "" {
			gID, err := id.ParseNull(&groupIDStr)
			if err != nil {
				return fmt.Errorf("value of 'group' flag must be a valid ID")
			}
			groupID = gID
		}
		return nil
	},
	RunE: mockItems,
	Args: cobra.ExactArgs(1),
}

func init() {
	ItemsCmd.Flags().StringP("group", "g", "", "create all items under the specified group")
}

func mockItems(cmd *cobra.Command, args []string) error {
	n, err := strconv.Atoi(args[0])
	if err != nil {
		return errors.New("N must be an integer")
	}

	ctx := cmd.Context()

	db, err := shared.InitDatabase(ctx)
	if err != nil {
		return fmt.Errorf("error connecting to database; %v", err)
	}
	defer db.Close()

	tx, err := db.BeginTx(ctx)
	defer tx.Rollback(ctx)

	res := make([]id.ID, n)

	for i := range n {
		s := items.NewService(db)
		itemID := id.Generate(string(com.IDPrefixItem))

		itm, err := s.Create(ctx, tx, items.CreateParams{
			AccountID:     shared.AccountID,
			ItemID:        itemID,
			Inventory:     inventoryID,
			Group:         groupID,
			CreatedAt:     itemID.Time(),
			UpdatedAt:     itemID.Time(),
			Description:   str.ParseNull(new(fmt.Sprintf("Mock item %s", itemID))),
			Name:          fmt.Sprintf("Mock item %s", itemID),
			Type:          sqlc.ItemTypePRODUCT,
			PriceCurrency: currency.NullCurrency{Currency: sqlc.CurrencyEUR, Valid: true},
			PriceAmount:   i32.ParseNull(new(i)),
			Active:        true,
		})
		if err != nil {
			return err
		}
		res[i] = itm.ID
	}

	if err := tx.Commit(ctx); err != nil {
		return err
	}

	shared.PrintResult(res...)

	return nil
}
