package main

import (
	"context"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"

	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:  "files",
		Usage: "",
		Commands: []*cli.Command{
			{
				Name:    "search",
				Aliases: []string{"c"},
				Action:  Search,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "directory",
						Value: "",
					},
					&cli.IntFlag{
						Name:  "size",
						Value: 10,
					},
				},
			},
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}

func Search(ctx context.Context, cmd *cli.Command) error {
	maxSize := int(cmd.Int("size"))
	sortedSlice := []struct {
		path string
		size int
	}{}

	search := func(path string, dir fs.DirEntry, err error) error {
		if len(sortedSlice) >= maxSize {
			//
			fmt.Println("foo")
		}
		return nil
	}

	dirToWalk := cmd.String("directory")
	if err := filepath.WalkDir(dirToWalk, search); err != nil {
		log.Fatal(err)
	}
	return nil
}

func binarySearch() error {
	return nil
}
