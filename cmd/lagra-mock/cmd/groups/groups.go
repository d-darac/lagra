package groupscmd

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/d-darac/lagra/cmd/lagra-mock/cmd/shared"
	"github.com/d-darac/lagra/internal/com"
	"github.com/d-darac/lagra/internal/services/groups"
	"github.com/d-darac/lagra/pkg/id"
	"github.com/d-darac/lagra/pkg/str"
	"github.com/spf13/cobra"
)

var parentGroupID id.NullID

// GroupsCmd represents the groups command
var GroupsCmd = &cobra.Command{
	Use:               "groups N",
	Short:             "Use to create groups",
	PersistentPreRun:  shared.RequireConfig,
	PersistentPreRunE: shared.SetAccID,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		pgIDStr, err := cmd.Flags().GetString("parent-group")
		if err != nil {
			return err
		}
		if pgIDStr != "" {
			pgID, err := id.ParseNull(&pgIDStr)
			if err != nil {
				return fmt.Errorf("value of 'parent-group' flag must be a valid ID")
			}
			parentGroupID = pgID
		}
		return nil
	},
	RunE: mockGroups,
	Args: cobra.ExactArgs(1),
}

func init() {
	GroupsCmd.Flags().StringP("parent-group", "p", "", "create all groups under the specified parent group")
}

func mockGroups(cmd *cobra.Command, args []string) error {
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
		s := groups.NewService(db)
		groupID := id.Generate(string(com.IDPrefixGroup))

		grp, err := s.Create(ctx, tx, groups.CreateParams{
			AccountID:   shared.AccountID,
			GroupID:     groupID,
			CreatedAt:   groupID.Time(),
			UpdatedAt:   groupID.Time(),
			Description: str.ParseNull(new(fmt.Sprintf("Mock group %s", groupID))),
			Name:        fmt.Sprintf("Mock group %s", groupID),
			ParentGroup: parentGroupID,
		})
		if err != nil {
			return err
		}
		res[i] = grp.ID
	}

	if err := tx.Commit(ctx); err != nil {
		return err
	}

	shared.PrintResult(res...)

	return nil
}
