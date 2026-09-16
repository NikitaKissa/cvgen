# cvgen

`cvgen` is a CLI-first CV document compiler written in Go.

It takes structured CV data and generates a **self-contained HTML document** using a Go HTML template, CSS, and optional assets.

```text
CV data
   ↓
cvgen
   ↓
self-contained HTML
```

The generated HTML can then be opened in a browser or printed to PDF using any external browser or print engine.

## Installation

Clone the repository and enter the project directory:

```bash
git clone https://github.com/NikitaKissa/cvgen.git
cd cvgen
```

Build and install `cvgen`:

```bash
sudo make cvgen-install
```

This builds the binary and installs it to `/usr/local/bin/cvgen`.

Alternatively, you can build the binary without installing it:

```bash
make cvgen-build
```

The resulting binary will be placed in the configured build directory.

Verify the installation:

```bash
cvgen version
```

## Usage

Generate a CV from an input file:

```bash
cvgen generate input.json
```

By default, `cvgen` uses its embedded template and styles.

Custom template:

```bash
cvgen generate input.json --template ./template.html
```

Custom template and CSS:

```bash
cvgen generate input.json \
  --template ./template.html \
  --style ./style.css
```

Specify the output file:

```bash
cvgen generate input.json --output ./resume.html
```

Validate your input and resources without writing output:

```bash
cvgen verify input.json
cvgen verify input.json \
  --template ./template.html \
  --style ./style.css
```

Get an embedded default template or style to customize:

```bash
cvgen get template > template.html
cvgen get style > style.css
```

## Input Format

`cvgen` uses JSON as its input format.

For a complete reference of the supported CV JSON structure, see the [CV JSON documentation](./docs/cv-json.md).

The input JSON contains the CV data used to generate the resume, including personal information, experience, education, skills, projects, certificates, languages, and other supported fields.

## Commands

```text
cvgen completion <shell>
cvgen generate <input> [--output <file>]
cvgen get <template | style>
cvgen verify <input>
cvgen version
```

* `completion` — generates the autocompletion script for the specified shell.
* `generate` — generates a self-contained HTML resume.
* `get` — prints an embedded default resource (`template` or `style`).
* `verify` — validates the CV JSON, template, and style, and renders the result without saving it.
* `version` — shows the installed `cvgen` version.


### Generate options

```text
--template <path>   Custom HTML template
--style <path>      Custom CSS
--output <path>     Output HTML file (stdout by default)
```

All presentation options are optional. When they are not provided, `cvgen` uses embedded defaults.

## Images

Images are specified in the CV input JSON.

Local image files are embedded into the generated HTML as Base64 data URIs, making the output fully self-contained.

Images specified by URL remain external references in the generated HTML. This produces a smaller HTML file, but requires the image to remain available at that URL.

## Templates

`cvgen` uses standard Go `html/template` files. A custom template controls the HTML structure and presentation of the generated document.

CSS is embedded into the resulting HTML, and local assets can be embedded as well.

The result is a single portable HTML file with no dependency on the original template, CSS, or asset files.

## Output

Both `generate` and `get` support two canonical output modes:

1. **File output** — pass `--output` (or `-o`) to write the generated HTML directly to a file.
2. **stdout output** — omit `--output` to write the generated HTML to stdout.

Both modes are first-class and equivalent ways of using `cvgen`.

For example:

```bash
# Write directly to a file
cvgen generate cv.json --output resume.html
cvgen get style --output style.css

# Write to stdout
cvgen generate cv.json
cvgen get style

# Redirect stdout to a file
cvgen generate cv.json > resume.html
cvgen get style > style.css
```

The stdout mode is useful for shell pipelines and for users who prefer to control output redirection themselves.

