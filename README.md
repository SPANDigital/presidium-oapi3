# Presidium OpenAPI 3

A Golang tool for importing your [OpenAPI 3](https://spec.openapis.org/oas/v3.0.3) spec into 
[Presidium](http://presidium.spandigital.net) documentation.

## Generation

This tool can be run as a:

- Standalone executable:
- Part of your Presidium Project

## Install

Install presidium-oapi3 globally

With Brew - Recommended for ARM64
```shell
brew tap SPANDigital/homebrew-tap https://github.com/SPANDigital/homebrew-tap.git
brew install presidium-oapi3

```

With NPM
```shell
npm install -g presidium-oapi-3
```

## Run

Execute presidium-oapi3 will print the usage:

```shell
presidium-oapi3 -h
```

```text
A Presidium tool that converts OAPI3 spec's to markdown

Usage:
  presidium-oapi3 [flags]
  presidium-oapi3 [command]

Available Commands:
  convert     Converts an OpenAPI 3 spec to markdown
  help        Help about any command

Flags:
  -h, --help   help for presidium-oapi3

Use "presidium-oapi3 [command] --help" for more information about a command.
```

To convert a file you can use the convert command:

```text
Usage:
  presidium-oapi3 convert [flags]

Flags:
  -e, --allowExternalRefs     Allow external references in the OpenAPI spec. 
  -n, --apiName string        The name under which the generated docs will be grouped
  -f, --file string           OpenAPI 3 spec file
      --includeExamples       Include a column on the schema for examples
      --includeRestrictions   Include a column on the schema for restrictions (default true)
  -i, --inlineProperties      Inline properties in the request and response schemas
  -o, --outputDir string      The output directory
  -r, --referenceURL string   The reference URL (default "reference")
  -s, --sortFilePath          Sort by filepath by prefixing a weight to the filename. Default is to use front matter weight
  -t, --titleFormat string    The template format used to create the title for each operation. 
                              Valid options are: 
                                - operationId: (Default) Uses the value of the operationId field.
                                - MethodURL: Uses a combination of the Method property and the URL.
```

Sample usage:
```shell
presidium-oapi3 convert -f api-spec.yaml -o /project/root/path
```

The converter will store the Markdown under the `/reference` directory by default, but you can change it by using the `-r, --referenceURL` flag.
```shell
presidium-oapi3 convert -f example.yaml -o /project/root/path -r /custom/path
```

## Part Of Your Project

Include as part of the npm build building your Presidium site as in the following sample:

```json
"scripts" : {
    "import-open-api-3" : "presidium-oapi3 convert -f <YOUR_API_SPEC> -o <THE_OUTPUT_DIRECTORY> -r <THE_PRESIDIUM_REFERENCE_URL>"
},
"devDependencies": {
    "presidium-oapi-3" : "#.#.#"
}
```
