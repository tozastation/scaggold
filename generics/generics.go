package generics

import (
	"fmt"
	"github.com/cheekybits/genny/parse"
	"github.com/rakyll/statik/fs"
	"golang.org/x/sync/errgroup"
	"io"
	"os"
	"strings"
)
func GenerateGenerics(domainName string) error {
	eg := errgroup.Group{}
	structName :=  strings.Title(domainName)
	var typeSets []map[string]string
	value := make(map[string]string)
	value["KeyType"] = structName
	typeSets = append(typeSets, value)
	eg.Go(func() error {
		serviceReader, err := read("/service_generics.go")
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
	//
	eg.Go(func() error {
		serviceReader, err := read("/repository_generics.go")
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
	out.Write(output)
	return nil
}

func read(filePath string) (io.ReadSeeker, error) {
	fileSystem, err := fs.New()
	if err != nil {
		return nil, err
	}
	f, err := fileSystem.Open(filePath)
	if err != nil {
		return nil, err
	}
	fmt.Println(f)
	return f, nil
}

func write(filePath string) (io.Writer, error) {
	file, err := os.Create(filePath)
	if err != nil {
		return nil, err
	}
	return file, nil
}