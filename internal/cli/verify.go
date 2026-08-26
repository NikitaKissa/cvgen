package cli

import (
	"fmt"

	"github.com/NikitaKissa/cvgen/internal/app/verify"
	"github.com/spf13/cobra"
)

func newVerifyCmd() *cobra.Command {
	var (
		inputPath    string
		templatePath string
		stylePath    string
	)

	cmd := &cobra.Command{
		Use:   "verify [flags] <file-path>",
		Short: "Validate CV JSON, template and style without writing output",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			inputPath = args[0]

			opts := verify.Options{
				InputPath:    inputPath,
				TemplatePath: templatePath,
				StylePath:    stylePath,
			}

			report, err := verify.Run(cmd.Context(), opts)
			if err != nil {
				return handleErr(err)
			}

			printVerifyReport(cmd, report)

			return nil
		},
	}

	cmd.Flags().StringVar(&templatePath, "template", "", "path to custom HTML template (optional)")
	cmd.Flags().StringVar(&stylePath, "style", "", "path to custom CSS file (optional)")

	return cmd
}

func printVerifyReport(cmd *cobra.Command, r verify.Report) {
	out := cmd.OutOrStdout()

	fmt.Fprintln(out, "OK: input, template and style are valid")
	fmt.Fprintf(out, "  experience:    %d\n", r.Experience)
	fmt.Fprintf(out, "  education:     %d\n", r.Education)
	fmt.Fprintf(out, "  skill groups:  %d\n", r.SkillGroups)
	fmt.Fprintf(out, "  certificates:  %d\n", r.Certificates)
	fmt.Fprintf(out, "  projects:      %d\n", r.Projects)
	fmt.Fprintf(out, "  languages:     %d\n", r.Languages)
	fmt.Fprintf(out, "  clause:        %t\n", r.Clause)
	fmt.Fprintf(out, "  rendered size: %d bytes (not saved)\n", r.OutputSize)
}
