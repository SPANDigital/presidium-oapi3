# Presidium OpenAPI 3

A Golang tool for importing your [OpenAPI 3](https://spec.openapis.org/oas/v3.0.3) spec into 
[Presidium](http://presidium.spandigital.net) documentation.

## Prerequisites
- npm 6+

## Generation
This tool can be run as a:
 - Standalone executable:
 - Part of your Presidium Project
    
 
## Standalone

Install presidium-oapi3 globally usin npm

```shell
npm install -g presidium-oapi-3
```

Execute presidium-oapi3 will print the usage:

```shell
presidium-oapi3 -h
```

```
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

To convert a file you simply:

```shell
presidium-oapi3 convert -f <YOUR_API_SPEC> -o <THE_OUTPUT_DIRECTORY> -r <THE_PRESIDIUM_REFERENCE_URL>
```

## Part Of Your Project

Include as part of the npm build building your Presidium site as in the following sample:

```json
    ...
    "scripts" : {
        "import-open-api-3" : "presidium-oapi3 convert -f <YOUR_API_SPEC> -o <THE_OUTPUT_DIRECTORY> -r <THE_PRESIDIUM_REFERENCE_URL>"
    },
    "devDependencies": {
        "presidium-oapi-3" : "#.#.#"
    }
    ...
```
