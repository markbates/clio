package clio

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/markbates/hepa"
	"github.com/markbates/pkger"
	"github.com/markbates/pkger/pkging"
)

func New(pkg pkging.Pkger) error {

	cur, err := pkg.Current()
	if err != nil {
		return err
	}

	// declare root to be used later
	root := "github.com/markbates/clio:/internal/elvis"

	// use explicit string so pkger can find it.
	err = pkger.Walk("github.com/markbates/clio:/internal/elvis", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		pt, err := cur.Parse(strings.TrimPrefix(path, root))
		if err != nil {
			return err
		}
		pt.Name = strings.ReplaceAll(pt.Name, "elvis", cur.Name)

		if err := pkg.MkdirAll(filepath.Dir(pt.Name), 0755); err != nil {
			return err
		}

		nf, err := pkg.Create(pt.String())
		if err != nil {
			return err
		}
		defer nf.Close()

		of, err := pkger.Open(path)
		if err != nil {
			return err
		}
		defer of.Close()

		hep := hepa.New()
		hep = hepa.Rinse(hep, []byte("elvis"), []byte(fmt.Sprintf("%s", cur.Name)))
		hep = hepa.Rinse(hep, []byte("github.com/costello/elvis"), []byte(fmt.Sprintf("%s", cur.ImportPath)))
		hep = hepa.Rinse(hep, []byte("package elvis"), []byte(fmt.Sprintf("package %s", cur.Name)))

		b, err := hep.Clean(of)
		if err != nil {
			return err
		}

		_, err = nf.Write(b)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}
	return nil
}
