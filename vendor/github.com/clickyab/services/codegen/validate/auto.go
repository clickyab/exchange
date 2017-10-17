package validate

import (
	"fmt"
	"html/template"

	"github.com/clickyab/services/codegen/annotate"
	"github.com/clickyab/services/codegen/plugins"

	"bytes"
	"io/ioutil"

	"path/filepath"

	"strings"

	"sort"

	"github.com/goraz/humanize"
	"golang.org/x/tools/imports"
)

type validatePlugin struct {
}

type validate struct {
	pkg  humanize.Package
	file humanize.File
	ann  annotate.Annotate
	typ  humanize.TypeName

	Map  []fieldMap
	Rec  string
	Type string
}

type fieldMap struct {
	Name string
	Json string
	Err  string
}

type context []validate

func (c context) Len() int {
	return len(c)
}

func (c context) Less(i, j int) bool {
	return strings.Compare(c[i].Type, c[j].Type) < 0
}

func (c context) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

var (
	validateFunc = `
// Code generated build with variable DO NOT EDIT.

package {{ .PackageName }}
// AUTO GENERATED CODE. DO NOT EDIT!
import (
	"gopkg.in/go-playground/validator.v9"
	"github.com/clickyab/services/framework/middleware"
	"context"
	"net/http"
)
	{{ range $m := .Data }}
	func ({{ $m.Rec }} *{{ $m.Type }}) Validate(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
		err := func(in interface{}) error {
			if v, ok := in.(interface {
				ValidateExtra(ctx context.Context, w http.ResponseWriter, r *http.Request) error
			}); ok {
				return v.ValidateExtra(ctx, w, r)
			}
			return nil
		}({{ $m.Rec }})
		if err != nil {
			return err
		}
		errs :=  validator.New().Struct({{ $m.Rec }})
		if errs == nil {
			return nil
		}
		res := middleware.GroupError{}
		for _, i := range errs.(validator.ValidationErrors) {
			switch i.Field() { {{ range $f := $m.Map }}
				case "{{ $f.Name }}":
					res["{{ $f.Json }}"] = trans.E("{{ $f.Err }}")
			{{ end }}
				default :
					logrus.Panicf("the field %s is not translated", i)
			}
		}
		if len(res) >0 {
			return res
		}
		return nil
	}
	{{ end }}
	`

	tpl = template.Must(template.New("validate").Parse(validateFunc))
)

// GetType return all types that this plugin can operate on
// for example if the result contain Route then all @Route sections are
// passed to this plugin
func (e validatePlugin) GetType() []string {
	return []string{"Validate"}
}

// Finalize is called after all the functions are done. the context is the one from the
// process
func (e validatePlugin) Finalize(c interface{}, p humanize.Package) error {
	var ctx context
	if c != nil {
		var ok bool
		ctx, ok = c.(context)
		if !ok {
			return fmt.Errorf("invalid context, need %T , got %T", ctx, c)
		}
	}

	buf := &bytes.Buffer{}
	sort.Sort(ctx)
	err := tpl.Execute(buf, struct {
		Data        context
		PackageName string
	}{
		Data:        ctx,
		PackageName: p.Name,
	})
	if err != nil {
		return err
	}
	f := filepath.Dir(p.Files[0].FileName)
	f = filepath.Join(f, "validators.gen.go")
	res, err := imports.Process("", buf.Bytes(), nil)
	if err != nil {
		fmt.Println(buf.String())
		return err
	}

	err = ioutil.WriteFile(f, res, 0644)
	if err != nil {
		return err
	}
	return nil
}

func (r *validatePlugin) ProcessStructure(
	c interface{},
	pkg humanize.Package,
	p humanize.File,
	f humanize.TypeName,
	a annotate.Annotate,
) (interface{}, error) {
	var ctx context
	if c != nil {
		var ok bool
		ctx, ok = c.(context)
		if !ok {
			return nil, fmt.Errorf("invalid context, need %T , got %T", ctx, c)
		}
	}

	dt := validate{
		pkg:  pkg,
		file: p,
		ann:  a,
		typ:  f,

		Type: f.Name,
		Rec:  "pl",
	}

	for _, field := range f.Type.(*humanize.StructType).Fields {
		if field.Tags.Get("validate") != "" {
			t := fieldMap{
				Name: field.Name,
				Json: field.Tags.Get("json"),
				Err:  field.Tags.Get("error"),
			}

			if t.Json == "" {
				t.Json = t.Name
			}

			if t.Err == "" {
				t.Err = "invalid value"
			}

			dt.Map = append(dt.Map, t)
		}
	}

bigLoop:
	for i := range pkg.Files {
		for _, fn := range pkg.Files[i].Functions {
			if fn.Receiver != nil {
				rec := fn.Receiver.Type
				if s, ok := rec.(*humanize.StarType); ok {
					rec = s.Target
				}
				if f.Name == rec.GetDefinition() {
					dt.Rec = fn.Receiver.Name
					break bigLoop
				}
			}
		}
	}

	ctx = append(ctx, dt)
	return ctx, nil
}

func (r *validatePlugin) StructureIsSupported(file humanize.File, fn humanize.TypeName) bool {
	return true
}

func (r *validatePlugin) GetOrder() int {
	return 5999
}

func init() {
	plugins.Register(&validatePlugin{})
}
