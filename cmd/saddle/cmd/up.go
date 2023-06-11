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

type manifest struct {
	Files map[string]file
}

type file struct {
	Path    string
	Format  string
	Content string
}

func runUp(cmd *cobra.Command, args []string) error {
	dir := defaultPkg
	if len(args) > 0 {
		dir = args[0]
	}

	ctx := cuecontext.New()
	store := newStore(dir, ctx)
	mn, err := store.manifest()
	if err != nil {
		return err
	}

	if err = mn.Validate(); err != nil {
		log.Fatal(err)
	}

	// log.Println(mn)
	var manifest manifest
	err = mn.Decode(&manifest)
	if err != nil {
		log.Fatal(err)
	}
	// log.Println(manifest)

	return nil
}
