# Spritetool

A CLI tool for building sprite sheets.

## Installation

```bash
go install github.com/dswisher/gamekit/cmd/spritetool@latest
```

## Usage

```bash
# Show the help for the tool
spritetool --help

# Take all image files in imagefolder and pack them into a sprite sheet
# called sheet1.png, and generate a metadata file called sheet1.json
spritetool --sheet sheet1.png --data sheet1.json imagefolder

# Generate a sprite sheet based on a YAML config for fine-grained control;
# see below for more info on the YAML format
spritetool sheet2.yaml
```

## Features

- Pack multiple images into sprite sheets
- Generate sprite sheet metadata

## Configuration File Format

TBD

## License

MIT

