// Code generated by goa v3.12.4, DO NOT EDIT.
//
// game HTTP client CLI support package
//
// Command:
// $ goa gen game-service/design

package cli

import (
	"flag"
	"fmt"
	characterc "game-service/gen/http/character/client"
	frontc "game-service/gen/http/front/client"
	itemc "game-service/gen/http/item/client"
	"net/http"
	"os"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//	command (subcommand1|subcommand2|...)
func UsageCommands() string {
	return `front list-items
item (list|show|add|update|remove)
character (list|show|add|update|remove)
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` front list-items` + "\n" +
		os.Args[0] + ` item list` + "\n" +
		os.Args[0] + ` character list` + "\n" +
		""
}

// ParseEndpoint returns the endpoint and payload as specified on the command
// line.
func ParseEndpoint(
	scheme, host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restore bool,
) (goa.Endpoint, any, error) {
	var (
		frontFlags = flag.NewFlagSet("front", flag.ContinueOnError)

		frontListItemsFlags = flag.NewFlagSet("list-items", flag.ExitOnError)

		itemFlags = flag.NewFlagSet("item", flag.ContinueOnError)

		itemListFlags = flag.NewFlagSet("list", flag.ExitOnError)

		itemShowFlags    = flag.NewFlagSet("show", flag.ExitOnError)
		itemShowIDFlag   = itemShowFlags.String("id", "REQUIRED", "ID of item to show")
		itemShowViewFlag = itemShowFlags.String("view", "", "")

		itemAddFlags    = flag.NewFlagSet("add", flag.ExitOnError)
		itemAddBodyFlag = itemAddFlags.String("body", "REQUIRED", "")

		itemUpdateFlags    = flag.NewFlagSet("update", flag.ExitOnError)
		itemUpdateBodyFlag = itemUpdateFlags.String("body", "REQUIRED", "")
		itemUpdateIDFlag   = itemUpdateFlags.String("id", "REQUIRED", "ID of the item to be updated")

		itemRemoveFlags  = flag.NewFlagSet("remove", flag.ExitOnError)
		itemRemoveIDFlag = itemRemoveFlags.String("id", "REQUIRED", "ID of item to remove")

		characterFlags = flag.NewFlagSet("character", flag.ContinueOnError)

		characterListFlags = flag.NewFlagSet("list", flag.ExitOnError)

		characterShowFlags    = flag.NewFlagSet("show", flag.ExitOnError)
		characterShowIDFlag   = characterShowFlags.String("id", "REQUIRED", "ID of character to show")
		characterShowViewFlag = characterShowFlags.String("view", "", "")

		characterAddFlags    = flag.NewFlagSet("add", flag.ExitOnError)
		characterAddBodyFlag = characterAddFlags.String("body", "REQUIRED", "")

		characterUpdateFlags    = flag.NewFlagSet("update", flag.ExitOnError)
		characterUpdateBodyFlag = characterUpdateFlags.String("body", "REQUIRED", "")
		characterUpdateIDFlag   = characterUpdateFlags.String("id", "REQUIRED", "ID of the character to be updated")

		characterRemoveFlags  = flag.NewFlagSet("remove", flag.ExitOnError)
		characterRemoveIDFlag = characterRemoveFlags.String("id", "REQUIRED", "ID of character to remove")
	)
	frontFlags.Usage = frontUsage
	frontListItemsFlags.Usage = frontListItemsUsage

	itemFlags.Usage = itemUsage
	itemListFlags.Usage = itemListUsage
	itemShowFlags.Usage = itemShowUsage
	itemAddFlags.Usage = itemAddUsage
	itemUpdateFlags.Usage = itemUpdateUsage
	itemRemoveFlags.Usage = itemRemoveUsage

	characterFlags.Usage = characterUsage
	characterListFlags.Usage = characterListUsage
	characterShowFlags.Usage = characterShowUsage
	characterAddFlags.Usage = characterAddUsage
	characterUpdateFlags.Usage = characterUpdateUsage
	characterRemoveFlags.Usage = characterRemoveUsage

	if err := flag.CommandLine.Parse(os.Args[1:]); err != nil {
		return nil, nil, err
	}

	if flag.NArg() < 2 { // two non flag args are required: SERVICE and ENDPOINT (aka COMMAND)
		return nil, nil, fmt.Errorf("not enough arguments")
	}

	var (
		svcn string
		svcf *flag.FlagSet
	)
	{
		svcn = flag.Arg(0)
		switch svcn {
		case "front":
			svcf = frontFlags
		case "item":
			svcf = itemFlags
		case "character":
			svcf = characterFlags
		default:
			return nil, nil, fmt.Errorf("unknown service %q", svcn)
		}
	}
	if err := svcf.Parse(flag.Args()[1:]); err != nil {
		return nil, nil, err
	}

	var (
		epn string
		epf *flag.FlagSet
	)
	{
		epn = svcf.Arg(0)
		switch svcn {
		case "front":
			switch epn {
			case "list-items":
				epf = frontListItemsFlags

			}

		case "item":
			switch epn {
			case "list":
				epf = itemListFlags

			case "show":
				epf = itemShowFlags

			case "add":
				epf = itemAddFlags

			case "update":
				epf = itemUpdateFlags

			case "remove":
				epf = itemRemoveFlags

			}

		case "character":
			switch epn {
			case "list":
				epf = characterListFlags

			case "show":
				epf = characterShowFlags

			case "add":
				epf = characterAddFlags

			case "update":
				epf = characterUpdateFlags

			case "remove":
				epf = characterRemoveFlags

			}

		}
	}
	if epf == nil {
		return nil, nil, fmt.Errorf("unknown %q endpoint %q", svcn, epn)
	}

	// Parse endpoint flags if any
	if svcf.NArg() > 1 {
		if err := epf.Parse(svcf.Args()[1:]); err != nil {
			return nil, nil, err
		}
	}

	var (
		data     any
		endpoint goa.Endpoint
		err      error
	)
	{
		switch svcn {
		case "front":
			c := frontc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "list-items":
				endpoint = c.ListItems()
				data = nil
			}
		case "item":
			c := itemc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "list":
				endpoint = c.List()
				data = nil
			case "show":
				endpoint = c.Show()
				data, err = itemc.BuildShowPayload(*itemShowIDFlag, *itemShowViewFlag)
			case "add":
				endpoint = c.Add()
				data, err = itemc.BuildAddPayload(*itemAddBodyFlag)
			case "update":
				endpoint = c.Update()
				data, err = itemc.BuildUpdatePayload(*itemUpdateBodyFlag, *itemUpdateIDFlag)
			case "remove":
				endpoint = c.Remove()
				data, err = itemc.BuildRemovePayload(*itemRemoveIDFlag)
			}
		case "character":
			c := characterc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "list":
				endpoint = c.List()
				data = nil
			case "show":
				endpoint = c.Show()
				data, err = characterc.BuildShowPayload(*characterShowIDFlag, *characterShowViewFlag)
			case "add":
				endpoint = c.Add()
				data, err = characterc.BuildAddPayload(*characterAddBodyFlag)
			case "update":
				endpoint = c.Update()
				data, err = characterc.BuildUpdatePayload(*characterUpdateBodyFlag, *characterUpdateIDFlag)
			case "remove":
				endpoint = c.Remove()
				data, err = characterc.BuildRemovePayload(*characterRemoveIDFlag)
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
}

// frontUsage displays the usage of the front command and its subcommands.
func frontUsage() {
	fmt.Fprintf(os.Stderr, `The front service is the main HTTP API for the game
Usage:
    %[1]s [globalflags] front COMMAND [flags]

COMMAND:
    list-items: List all items

Additional help:
    %[1]s front COMMAND --help
`, os.Args[0])
}
func frontListItemsUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] front list-items

List all items

Example:
    %[1]s front list-items
`, os.Args[0])
}

// itemUsage displays the usage of the item command and its subcommands.
func itemUsage() {
	fmt.Fprintf(os.Stderr, `The item service is the service for managing items
Usage:
    %[1]s [globalflags] item COMMAND [flags]

COMMAND:
    list: List all items
    show: Show item by ID
    add: Add new item and return its ID
    update: Update an item with the given ID
    remove: Remove an item

Additional help:
    %[1]s item COMMAND --help
`, os.Args[0])
}
func itemListUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] item list

List all items

Example:
    %[1]s item list
`, os.Args[0])
}

func itemShowUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] item show -id STRING -view STRING

Show item by ID
    -id STRING: ID of item to show
    -view STRING: 

Example:
    %[1]s item show --id "Ex est omnis qui." --view "default"
`, os.Args[0])
}

func itemAddUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] item add -body JSON

Add new item and return its ID
    -body JSON: 

Example:
    %[1]s item add --body '{
      "damage": 2,
      "description": "A magnificent sword which grants the bearer +2 wisdom",
      "healing": 75,
      "name": "Sword of Wisdom",
      "protection": 10
   }'
`, os.Args[0])
}

func itemUpdateUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] item update -body JSON -id STRING

Update an item with the given ID
    -body JSON: 
    -id STRING: ID of the item to be updated

Example:
    %[1]s item update --body '{
      "item": {
         "damage": 49,
         "description": "A magnificent sword which grants the bearer +2 wisdom",
         "healing": 103,
         "name": "Sword of Wisdom",
         "protection": 12
      }
   }' --id "Quod qui aut id."
`, os.Args[0])
}

func itemRemoveUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] item remove -id STRING

Remove an item
    -id STRING: ID of item to remove

Example:
    %[1]s item remove --id "Velit facilis."
`, os.Args[0])
}

// characterUsage displays the usage of the character command and its
// subcommands.
func characterUsage() {
	fmt.Fprintf(os.Stderr, `The character service is the service for managing characters
Usage:
    %[1]s [globalflags] character COMMAND [flags]

COMMAND:
    list: List all characters
    show: Show character by ID
    add: Add new character and return its ID
    update: Update a character with the given ID
    remove: Remove a character

Additional help:
    %[1]s character COMMAND --help
`, os.Args[0])
}
func characterListUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] character list

List all characters

Example:
    %[1]s character list
`, os.Args[0])
}

func characterShowUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] character show -id STRING -view STRING

Show character by ID
    -id STRING: ID of character to show
    -view STRING: 

Example:
    %[1]s character show --id "Expedita ex nihil consequatur quia harum earum." --view "tiny"
`, os.Args[0])
}

func characterAddUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] character add -body JSON

Add new character and return its ID
    -body JSON: 

Example:
    %[1]s character add --body '{
      "description": "A grizzled wizard with a penchant for mayhem and mead",
      "experience": 83494,
      "health": 157,
      "name": "Arvish the Wise"
   }'
`, os.Args[0])
}

func characterUpdateUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] character update -body JSON -id STRING

Update a character with the given ID
    -body JSON: 
    -id STRING: ID of the character to be updated

Example:
    %[1]s character update --body '{
      "character": {
         "description": "A grizzled wizard with a penchant for mayhem and mead",
         "experience": 44185,
         "health": 1543,
         "name": "Arvish the Wise"
      }
   }' --id "Dolore ut qui in omnis."
`, os.Args[0])
}

func characterRemoveUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] character remove -id STRING

Remove a character
    -id STRING: ID of character to remove

Example:
    %[1]s character remove --id "Sequi qui."
`, os.Args[0])
}
