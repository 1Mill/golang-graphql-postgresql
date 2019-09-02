package schema

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

// String exports the graphql schema as a string to be consumed by graphql-go
func String() string {
	var paths []string

	root := "./gql"
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// * Only include non-directory files
		if info.IsDir() == false {
			paths = append(paths, path)
			return nil
		}

		return nil
	})
	if err != nil {
		log.Panic(err)
	}

	var content [][]byte
	for _, path := range paths {
		b, err := ioutil.ReadFile(path)
		if err != nil {
			log.Panic(err)
		}

		content = append(content, b)
	}

	var s []byte
	for _, b := range content {
		s = append(s, b...)
	}

	return string(s)
}
