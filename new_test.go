package clio

import (
	"bytes"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"testing"

	"github.com/markbates/pkger/here"
	"github.com/markbates/pkger/pkging/mem"
	"github.com/stretchr/testify/require"
)

func Test_New(t *testing.T) {
	t.SkipNow()
	r := require.New(t)

	ip := "github.com/declan/macmanus"
	pwd, err := os.Getwd()
	r.NoError(err)

	name := path.Base(ip)
	dir := filepath.Join(pwd, name)
	info := here.Info{
		ImportPath: ip,
		Name:       name,
		Dir:        dir,
		Module: here.Module{
			Path: ip,
			Dir:  dir,
		},
	}

	pkg, err := mem.New(info)
	r.NoError(err)

	err = New(pkg)
	r.NoError(err)

	pkgall := func(s string) []byte {
		t.Helper()
		f, err := pkg.Open(s)
		r.NoError(err)
		defer f.Close()

		b, err := ioutil.ReadAll(f)
		r.NoError(err, s)
		return bytes.TrimSpace(b)
	}

	osall := func(s string) []byte {
		t.Helper()
		fp := filepath.Join(pwd, "internal", "elvis")
		fp = filepath.Join(fp, s)
		b, err := ioutil.ReadFile(fp)
		r.NoError(err, fp)
		return bytes.TrimSpace(b)
	}

	exp := string(osall("go.mod"))
	act := string(pkgall("/go.mod"))
	r.NotEqual(exp, act)
	r.Contains(act, "module github.com/declan/macmanus")

	exp = string(osall("elvis.go"))
	act = string(pkgall("/macmanus.go"))
	r.NotEqual(exp, act)
	r.Contains(act, "package macmanus")

	exp = string(osall(filepath.Join("cmd", "elvis", "cli", "main.go")))
	act = string(pkgall("/cmd/macmanus/cli/main.go"))
	r.NotEqual(exp, act)
	r.Contains(act, "macmanus/cmd/internal/cmdx")
}
