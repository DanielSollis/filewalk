package main

import (
	"context"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"slices"

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

type sortedSlice []struct {
	path string
	size int
}

func Search(ctx context.Context, cmd *cli.Command) error {
	fileListSize := int(cmd.Int("size"))
	biggestFiles := &sortedSlice{}

	search := func(path string, dir fs.DirEntry, err error) error {
		file, fileErr := dir.Info()
		if fileErr != nil {
			return fileErr
		}

		if len(*biggestFiles) >= fileListSize {
			smallestFileSize := (*biggestFiles)[len(*biggestFiles)-1].size
			if smallestFileSize > int(file.Size()) {
				return nil
			}
		}

		inx, found := slices.BinarySearch(biggestFiles, file.Size())
		slices.Insert(biggestFiles, inx, file.Size())
		return nil
	}

	dirToWalk := cmd.String("directory")
	if err := filepath.WalkDir(dirToWalk, search); err != nil {
		log.Fatal(err)
	}
	fmt.Println(biggestFiles)
	return nil
}
