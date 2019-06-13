package generics

import (
	"fmt"
	"github.com/cheekybits/genny/parse"
	"golang.org/x/sync/errgroup"
	"io"
	"os"
	"strings"
)

func GenerateGenerics(domainName string) error {
	eg := errgroup.Group{}
	// Replace Struct Name
	structName :=  strings.Title(domainName)
	var typeSets []map[string]string
	value := make(map[string]string)
	value["KeyType"] = structName
	typeSets = append(typeSets, value)
	// ---
	eg.Go(func() error {
		serviceReader, err := read(os.Getenv("GOPATH") + "/src/github.com/tozastation/gogoenv/boiler/services/service_generics.go")
		if err != nil {
			return err
		}
		serviceWriter , err := write("./domain/services/"+domainName+"_service.go")
		if err != nil {
			return err
		}
		err = gen(domainName + "_service.go", "", serviceReader, typeSets, serviceWriter)
		if err != nil {
			return err
		}
		fmt.Println("Generated Service Template")
		return nil
	})
	// ---
	eg.Go(func() error {
		serviceReader, err := read(os.Getenv("GOPATH") + "/src/github.com/tozastation/gogoenv/boiler/repositories/repository_generics.go")
		if err != nil {
			return err
		}
		serviceWriter , err := write("./infrastructure/persistence/repositories/"+domainName+"_repository.go")
		if err != nil {
			return err
		}
		err = gen(domainName + "_service.go", "", serviceReader, typeSets, serviceWriter)
		if err != nil {
			return err
		}
		fmt.Println("Generated Repository Template")
		return nil
	})
	if err := eg.Wait(); err != nil {
		return err
	}
	return nil
}

// gen performs the generic generation.
func gen(filename, pkgName string, in io.ReadSeeker, typesets []map[string]string, out io.Writer) error {
	var output []byte
	var err error
	output, err = parse.Generics(filename, pkgName, in, typesets)
	if err != nil {
		return err
	}
	_, err = out.Write(output)
	if err != nil {
		return err
	}
	return nil
}

func read(filePath string) (io.ReadSeeker, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	fmt.Printf("%s is opened", file.Name())
	return file, nil
}

func write(filePath string) (io.Writer, error) {
	file, err := os.Create(filePath)
	if err != nil {
		return nil, err
	}
	return file, nil
}