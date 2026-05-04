package itemscmd

import (
	"context"
	"fmt"

	"github.com/d-darac/lagra/cmd/lagra-mock/cmd/shared"
	"github.com/d-darac/lagra/internal/services/items"
	"github.com/d-darac/lagra/pkg/id"
	"github.com/spf13/cobra"
)

var ids = []id.ID{}

// wipeCmd represents the wipe command
var wipeCmd = &cobra.Command{
	Use:   "wipe [all]",
	Short: "Use to wipe items created by lagra-mock",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		IDs, _ := cmd.Flags().GetStringSlice("ids")
		if len(IDs) != 0 && len(args) != 0 {
			return fmt.Errorf("cannot use the 'all' arg and the 'ids' flag at the same time")
		}
		if len(IDs) == 0 && len(args) == 0 {
			return fmt.Errorf("must provide either 'all' arg or 'ids' flag")
		}

		for _, s := range IDs {
			id, err := id.Parse(s)
			if err != nil {
				return fmt.Errorf("%s is not a valid ID", s)
			}
			ids = append(ids, id)
		}

		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 0 {
			return wipeAll(cmd.Context())
		}
		return wipe(cmd.Context(), ids)
	},
	Args:      cobra.MatchAll(cobra.RangeArgs(0, 1), cobra.OnlyValidArgs),
	ValidArgs: []cobra.Completion{"all"},
}

func init() {
	ItemsCmd.AddCommand(wipeCmd)
	wipeCmd.Flags().StringSlice("ids", []string{}, "list of items IDs to delete")

}

func wipe(ctx context.Context, IDs []id.ID) error {
	db, err := shared.InitDatabase(ctx)
	if err != nil {
		return fmt.Errorf("error connecting to database; %v", err)
	}
	defer db.Close()

	tx, err := db.BeginTx(ctx)
	defer tx.Rollback(ctx)

	s := items.NewService(db)

	for _, itemID := range IDs {
		_, err := s.Get(ctx, items.GetParams{AccountID: shared.AccountID, ItemID: itemID})
		if err != nil {
			return err
		}
		if err := s.Delete(ctx, tx, items.DeleteParams{
			AccountID: shared.AccountID,
			ItemID:    itemID,
		}); err != nil {
			return err
		}
	}

	if err := tx.Commit(ctx); err != nil {
		return err
	}

	return nil
}

func wipeAll(ctx context.Context) error {
	db, err := shared.InitDatabase(ctx)
	if err != nil {
		return fmt.Errorf("error connecting to database; %v", err)
	}
	defer db.Close()

	tx, err := db.BeginTx(ctx)
	defer tx.Rollback(ctx)

	sql := "delete from items where account_id = $1 and name ~~* CONCAT('%', $2::text, '%')"
	if _, err := tx.Exec(ctx, sql, shared.AccountID.UUID(), "mock"); err != nil {
		return err
	}

	if err := tx.Commit(ctx); err != nil {
		return err
	}

	return nil
}
