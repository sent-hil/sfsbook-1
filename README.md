# sfsbook

STOP VIOLENCE!

## Structure

We use [Gin http framework](https://github.com/gin-gonic/gin) for routing and
rendering of paths.

main.go - contains the Gin initializer and sets up various routes and runs the
server

route_index.go - contains the index handler.

templates/ - contains various Go templates.

static/ - contains js and css files; this folder is server to the public, so
BEWARE!

refguides/ - contains the resources in various formats.

vendor/ - contains vendored dependencies; if you add a new dep, be sure to add
it to this folder with `govendor add +external`

## Development

Install [gin](https://github.com/codegangsta/gin) which is a live reload
utility.

    go get github.com/codegangsta/gin

Run server with:

    gin run main.go

This should start server at http://localhost:3000.
