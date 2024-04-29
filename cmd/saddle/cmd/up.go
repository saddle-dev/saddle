package cmd

import (
	"errors"
	"fmt"
	"io/ioutil"
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
	Paths map[string]config
}

type config struct {
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

	err = sync(manifest)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func sync(mn manifest) error {
	for _, file := range mn.Paths {
		var err error
		// log.Println(file)

		switch file.Format {
		case "JSON":
			err = writeJSON(file)
		case "YAML":
			err = writeYAML(file)
		default:
			// Should not happen if CUE validates inputs
			return errors.New("Invalid file type")
		}

		if err != nil {
			return err
		}
	}

	return nil
}

func writeJSON(file config) error {
	fmt.Println("Writing JSON file to", file.Path)

	err := ioutil.WriteFile(file.Path, []byte(file.Content), 0644)
	if err != nil {
		return err
	}

	return nil
}

func writeYAML(file config) error {
	fmt.Println("Writing YAML file to", file.Path)

	err := ioutil.WriteFile(file.Path, []byte(file.Content), 0644)
	if err != nil {
		return err
	}

	return nil
}
