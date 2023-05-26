package cmd

import (
	"log"

	"cuelang.org/go/cue/cuecontext"
	"github.com/spf13/cobra"
)

// upCmd represents the up command
var upCmd = &cobra.Command{
	Use:   "up",
	Short: "Sync CUE definitions to application configurations",
	Long:  ``,
	RunE:  runUp,
}

func setupUpCommand(cmd *cobra.Command) {
	cmd.AddCommand(upCmd)
}

func runUp(cmd *cobra.Command, args []string) error {
	dir := defaultPkg
	if len(args) > 0 {
		dir = args[0]
	}

	ctx := cuecontext.New()
	store := newStore(dir, ctx)
	manifest, err := store.manifest()
	if err != nil {
		return err
	}

	if err = manifest.Validate(); err != nil {
		log.Fatal(err)
	}

	log.Println(manifest)

	return nil
}
