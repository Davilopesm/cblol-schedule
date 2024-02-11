package handler

import (
	"context"
	"os"

	"github.com/a-h/templ"
)

func render(c context.Context, component templ.Component) error {
	return component.Render(context.Background(), os.Stdout)
}

func renderToString(c context.Context, component templ.Component) string {
	html, _ := templ.ToGoHTML(context.Background(), component)
	return string(html)
}
