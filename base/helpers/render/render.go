package render

import (
	"fmt"
	"html/template"
	"net/url"
	"strings"

	"github.com/unrolled/render"
)

var (
	Template *render.Render
)

var tmplfuncs = []template.FuncMap{
	template.FuncMap{
		"html": func(s string) template.HTML {
			return template.HTML(s)
		},
		"htmlattr": func(s string) template.HTMLAttr {
			return template.HTMLAttr(s)
		},
		"sprintf": func(f string, a ...interface{}) string {
			return fmt.Sprintf(f, a...)
		},
		"urlencode": func(s string) string {
			return url.QueryEscape(s)
		},
		"ifnotempty": func(value, def string) string {
			if strings.TrimSpace(value) != "" {
				return value
			}
			return def
		},
		"renderopts": func(min, max int, format string) template.HTML {
			var elems []string
			for i := min; i <= max; i++ {
				text := fmt.Sprintf("%v", i)

				if format != "" && strings.Contains(format, "#") {
					text = strings.Replace(format, "#", text, -1)
				}

				elems = append(
					elems,
					fmt.Sprintf(`<option value="%v">%v</option>`, i, text),
				)
			}
			return template.HTML(strings.Join(elems, "\n"))
		},
		"renderoptmap": func(values map[string]string) template.HTML {
			var elems []string
			for key, value := range values {
				elems = append(
					elems,
					fmt.Sprintf(`<option value="%v">%v</option>`, key, value),
				)
			}
			return template.HTML(strings.Join(elems, "\n"))
		},
	},
}

func init() {
	Template = render.New(render.Options{
		Directory:     "views",
		Layout:        "layout",
		Extensions:    []string{".tmpl"},
		Funcs:         tmplfuncs,
		Charset:       "UTF-8",
		IsDevelopment: true,
	})
}
