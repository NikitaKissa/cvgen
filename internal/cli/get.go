package cli

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/NikitaKissa/cvgen/internal/template"
	"github.com/spf13/cobra"
)

// getResources maps the argument accepted by `cvgen get <resource>` to
// the embedded content it returns. Add an entry here (e.g. "json" for
// an example CV) to extend the command — no other changes needed.
var getResources = map[string]string{
	"template": template.DefaultTemplateHTML,
	"style":    template.DefaultStyleCSS,
}

func newGetCmd() *cobra.Command {
	var outputPath string

	validArgs := make([]string, 0, len(getResources))
	for name := range getResources {
		validArgs = append(validArgs, name)
	}
	sort.Strings(validArgs)

	cmd := &cobra.Command{
		Use:   "get [" + strings.Join(validArgs, "|") + "]",
		Short: "Print an embedded default resource",
		Long: "get prints one of cvgen's embedded default resources to stdout " +
			"(or to a file with --output).\n\n" +
			"Available arguments:\n" +
			"  template   the embedded default HTML template\n" +
			"  style      the embedded default CSS\n",
		ValidArgs: validArgs,
		Args:      cobra.ExactValidArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			content := getResources[args[0]]

			if outputPath == "" {
				fmt.Fprint(cmd.OutOrStdout(), content)
				return nil
			}

			if err := os.WriteFile(outputPath, []byte(content), 0664); err != nil {
				return fmt.Errorf("writing output `%s`: %w", outputPath, err)
			}

			return nil
		},
	}

	cmd.Flags().StringVar(&outputPath, "output", "", "write to this file instead of stdout")

	return cmd
}
