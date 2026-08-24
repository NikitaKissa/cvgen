package cli

import (
	"github.com/NikitaKissa/cvgen/internal/app/generate"
	"github.com/spf13/cobra"
)

func newGenerateCmd() *cobra.Command {
	var (
		inputPath    string
		templatePath string
		stylePath    string
		outputPath   string
	)

	cmd := &cobra.Command{
		Use:   "generate",
		Short: "Generate a self-contained HTML resume",
		RunE: func(cmd *cobra.Command, args []string) error {
			opts := generate.Options{
				InputPath:    inputPath,
				TemplatePath: templatePath,
				StylePath:    stylePath,
				OutputPath:   outputPath,
			}
			err := generate.Run(cmd.Context(), opts)
			return handleErr(err)
		},
	}

	cmd.Flags().StringVar(&inputPath, "input", "", "path to CV JSON file (required)")
	cmd.Flags().StringVar(&templatePath, "template", "", "path to custom HTML template (optional)")
	cmd.Flags().StringVar(&stylePath, "style", "", "path to custom CSS file (optional)")
	cmd.Flags().StringVar(&outputPath, "output", "resume.html", "output HTML file path")

	cmd.MarkFlagRequired("output")

	return cmd
}
