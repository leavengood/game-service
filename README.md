### Game Service

This is a Goa based project which hosts multiple HTTP and GRPC APIs to represent
an online multiplayer game.

There are 3 backend services which handle characters, items and then character inventories of items.

One frontend service offers the main API which delegates to the backend services.

## Code Generation

If the files in the `design` package change, code can be re-generated as so:

    $ goa gen game-service/design

It may also be useful to regenerate the example code if large changes are made.
Anything which exists will not be replaced so some code may need to be deleted first.
Then this command can be run:

    $ goa example game-service/design

## How to Build

Server:

    $ go build ./cmd/game

Client:

    $ go build ./cmd/game-cli

## How to Run

Server:

    $ ./game

Client (as one example):

    $ ./game-cli -url="http://localhost:8000" front list-characters

## How to Test

There are currently no automated tests. The CLI can be useful to test as well as applications like Insomnia or Postman.