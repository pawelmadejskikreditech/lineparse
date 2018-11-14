# Lineparse

Helps to extract line values from internal log structure.

Example log:

       campaign_manager | 2018-11-14T10:45:31Z INFO 127.0.0.1 GET /campaigns/1 403 user "guest" is not allowed to show campaign with id: '1'

## Usage

    substring := lineparse.Parse(logLine, ParserStructInstance)

## Features

* can extract IP from given string using - `IPParser`

## TODO

* add option to extract Path from given string

## Development

### Testing

    go test

### Linting

    golint

### Contributing

* each feature is done within a custom feature branch and merged to master after code review
* each change should have a meaningful commit message
* keep highest possible coverage
* keep code clean and well documented
* notify team members about breaking changes
* README.md is up to date with codebase
