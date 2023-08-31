// Code generated by goa v3.12.4, DO NOT EDIT.
//
// game gRPC client CLI support package
//
// Command:
// $ goa gen game-service/design

package cli

import (
	"flag"
	"fmt"
	characterc "game-service/gen/grpc/character/client"
	inventoryc "game-service/gen/grpc/inventory/client"
	itemc "game-service/gen/grpc/item/client"
	"os"

	goa "goa.design/goa/v3/pkg"
	grpc "google.golang.org/grpc"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//	command (subcommand1|subcommand2|...)
func UsageCommands() string {
	return `character (list|show|add|update|remove)
inventory (show|add|remove)
item (list|show|add|update|remove)
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` character list` + "\n" +
		os.Args[0] + ` inventory show --message '{
      "id": "Qui enim modi magni doloremque fugiat impedit."
   }'` + "\n" +
		os.Args[0] + ` item list` + "\n" +
		""
}

// ParseEndpoint returns the endpoint and payload as specified on the command
// line.
func ParseEndpoint(cc *grpc.ClientConn, opts ...grpc.CallOption) (goa.Endpoint, any, error) {
	var (
		characterFlags = flag.NewFlagSet("character", flag.ContinueOnError)

		characterListFlags = flag.NewFlagSet("list", flag.ExitOnError)

		characterShowFlags       = flag.NewFlagSet("show", flag.ExitOnError)
		characterShowMessageFlag = characterShowFlags.String("message", "", "")
		characterShowViewFlag    = characterShowFlags.String("view", "", "")

		characterAddFlags       = flag.NewFlagSet("add", flag.ExitOnError)
		characterAddMessageFlag = characterAddFlags.String("message", "", "")

		characterUpdateFlags       = flag.NewFlagSet("update", flag.ExitOnError)
		characterUpdateMessageFlag = characterUpdateFlags.String("message", "", "")

		characterRemoveFlags       = flag.NewFlagSet("remove", flag.ExitOnError)
		characterRemoveMessageFlag = characterRemoveFlags.String("message", "", "")

		inventoryFlags = flag.NewFlagSet("inventory", flag.ContinueOnError)

		inventoryShowFlags       = flag.NewFlagSet("show", flag.ExitOnError)
		inventoryShowMessageFlag = inventoryShowFlags.String("message", "", "")

		inventoryAddFlags       = flag.NewFlagSet("add", flag.ExitOnError)
		inventoryAddMessageFlag = inventoryAddFlags.String("message", "", "")

		inventoryRemoveFlags       = flag.NewFlagSet("remove", flag.ExitOnError)
		inventoryRemoveMessageFlag = inventoryRemoveFlags.String("message", "", "")

		itemFlags = flag.NewFlagSet("item", flag.ContinueOnError)

		itemListFlags = flag.NewFlagSet("list", flag.ExitOnError)

		itemShowFlags       = flag.NewFlagSet("show", flag.ExitOnError)
		itemShowMessageFlag = itemShowFlags.String("message", "", "")
		itemShowViewFlag    = itemShowFlags.String("view", "", "")

		itemAddFlags       = flag.NewFlagSet("add", flag.ExitOnError)
		itemAddMessageFlag = itemAddFlags.String("message", "", "")

		itemUpdateFlags       = flag.NewFlagSet("update", flag.ExitOnError)
		itemUpdateMessageFlag = itemUpdateFlags.String("message", "", "")

		itemRemoveFlags       = flag.NewFlagSet("remove", flag.ExitOnError)
		itemRemoveMessageFlag = itemRemoveFlags.String("message", "", "")
	)
	characterFlags.Usage = characterUsage
	characterListFlags.Usage = characterListUsage
	characterShowFlags.Usage = characterShowUsage
	characterAddFlags.Usage = characterAddUsage
	characterUpdateFlags.Usage = characterUpdateUsage
	characterRemoveFlags.Usage = characterRemoveUsage

	inventoryFlags.Usage = inventoryUsage
	inventoryShowFlags.Usage = inventoryShowUsage
	inventoryAddFlags.Usage = inventoryAddUsage
	inventoryRemoveFlags.Usage = inventoryRemoveUsage

	itemFlags.Usage = itemUsage
	itemListFlags.Usage = itemListUsage
	itemShowFlags.Usage = itemShowUsage
	itemAddFlags.Usage = itemAddUsage
	itemUpdateFlags.Usage = itemUpdateUsage
	itemRemoveFlags.Usage = itemRemoveUsage

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
		case "character":
			svcf = characterFlags
		case "inventory":
			svcf = inventoryFlags
		case "item":
			svcf = itemFlags
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

		case "inventory":
			switch epn {
			case "show":
				epf = inventoryShowFlags

			case "add":
				epf = inventoryAddFlags

			case "remove":
				epf = inventoryRemoveFlags

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
		case "character":
			c := characterc.NewClient(cc, opts...)
			switch epn {
			case "list":
				endpoint = c.List()
				data = nil
			case "show":
				endpoint = c.Show()
				data, err = characterc.BuildShowPayload(*characterShowMessageFlag, *characterShowViewFlag)
			case "add":
				endpoint = c.Add()
				data, err = characterc.BuildAddPayload(*characterAddMessageFlag)
			case "update":
				endpoint = c.Update()
				data, err = characterc.BuildUpdatePayload(*characterUpdateMessageFlag)
			case "remove":
				endpoint = c.Remove()
				data, err = characterc.BuildRemovePayload(*characterRemoveMessageFlag)
			}
		case "inventory":
			c := inventoryc.NewClient(cc, opts...)
			switch epn {
			case "show":
				endpoint = c.Show()
				data, err = inventoryc.BuildShowPayload(*inventoryShowMessageFlag)
			case "add":
				endpoint = c.Add()
				data, err = inventoryc.BuildAddPayload(*inventoryAddMessageFlag)
			case "remove":
				endpoint = c.Remove()
				data, err = inventoryc.BuildRemovePayload(*inventoryRemoveMessageFlag)
			}
		case "item":
			c := itemc.NewClient(cc, opts...)
			switch epn {
			case "list":
				endpoint = c.List()
				data = nil
			case "show":
				endpoint = c.Show()
				data, err = itemc.BuildShowPayload(*itemShowMessageFlag, *itemShowViewFlag)
			case "add":
				endpoint = c.Add()
				data, err = itemc.BuildAddPayload(*itemAddMessageFlag)
			case "update":
				endpoint = c.Update()
				data, err = itemc.BuildUpdatePayload(*itemUpdateMessageFlag)
			case "remove":
				endpoint = c.Remove()
				data, err = itemc.BuildRemovePayload(*itemRemoveMessageFlag)
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
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
	fmt.Fprintf(os.Stderr, `%[1]s [flags] character show -message JSON -view STRING

Show character by ID
    -message JSON: 
    -view STRING: 

Example:
    %[1]s character show --message '{
      "id": "Atque tenetur non dicta dolores."
   }' --view "default"
`, os.Args[0])
}

func characterAddUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] character add -message JSON

Add new character and return its ID
    -message JSON: 

Example:
    %[1]s character add --message '{
      "description": "A grizzled wizard with a penchant for mayhem and mead",
      "experience": 85126,
      "health": 667,
      "name": "Arvish the Wise"
   }'
`, os.Args[0])
}

func characterUpdateUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] character update -message JSON

Update a character with the given ID
    -message JSON: 

Example:
    %[1]s character update --message '{
      "character": {
         "description": "A grizzled wizard with a penchant for mayhem and mead",
         "experience": 13813,
         "health": 35,
         "name": "Arvish the Wise"
      },
      "id": "Velit aut velit."
   }'
`, os.Args[0])
}

func characterRemoveUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] character remove -message JSON

Remove a character
    -message JSON: 

Example:
    %[1]s character remove --message '{
      "id": "Eum et nihil enim."
   }'
`, os.Args[0])
}

// inventoryUsage displays the usage of the inventory command and its
// subcommands.
func inventoryUsage() {
	fmt.Fprintf(os.Stderr, `The inventory service is the service for managing character inventories
Usage:
    %[1]s [globalflags] inventory COMMAND [flags]

COMMAND:
    show: Show the inventory for a character as a list of item IDs
    add: Add an item ID to a character's inventory
    remove: Remove an item ID from a character's inventory

Additional help:
    %[1]s inventory COMMAND --help
`, os.Args[0])
}
func inventoryShowUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] inventory show -message JSON

Show the inventory for a character as a list of item IDs
    -message JSON: 

Example:
    %[1]s inventory show --message '{
      "id": "Qui enim modi magni doloremque fugiat impedit."
   }'
`, os.Args[0])
}

func inventoryAddUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] inventory add -message JSON

Add an item ID to a character's inventory
    -message JSON: 

Example:
    %[1]s inventory add --message '{
      "id": "Ipsa vitae molestias itaque animi.",
      "item_id": "Qui consequatur ut nisi expedita et."
   }'
`, os.Args[0])
}

func inventoryRemoveUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] inventory remove -message JSON

Remove an item ID from a character's inventory
    -message JSON: 

Example:
    %[1]s inventory remove --message '{
      "id": "Fuga quis aut et et pariatur eos.",
      "item_id": "Vitae quae dolores."
   }'
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
	fmt.Fprintf(os.Stderr, `%[1]s [flags] item show -message JSON -view STRING

Show item by ID
    -message JSON: 
    -view STRING: 

Example:
    %[1]s item show --message '{
      "id": "Cumque sapiente sapiente quaerat nisi quo eaque."
   }' --view "tiny"
`, os.Args[0])
}

func itemAddUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] item add -message JSON

Add new item and return its ID
    -message JSON: 

Example:
    %[1]s item add --message '{
      "damage": 85,
      "description": "A magnificent sword which grants the bearer +2 wisdom",
      "healing": 42,
      "name": "Sword of Wisdom",
      "protection": 15
   }'
`, os.Args[0])
}

func itemUpdateUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] item update -message JSON

Update an item with the given ID
    -message JSON: 

Example:
    %[1]s item update --message '{
      "id": "Quae sint sequi est repudiandae magni.",
      "item": {
         "damage": 54,
         "description": "A magnificent sword which grants the bearer +2 wisdom",
         "healing": 154,
         "name": "Sword of Wisdom",
         "protection": 19
      }
   }'
`, os.Args[0])
}

func itemRemoveUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] item remove -message JSON

Remove an item
    -message JSON: 

Example:
    %[1]s item remove --message '{
      "id": "Sint quos odit at quia officiis esse."
   }'
`, os.Args[0])
}
