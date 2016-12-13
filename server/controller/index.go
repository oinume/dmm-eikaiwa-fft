package controller

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/oinume/lekcije/server/config"
	"github.com/oinume/lekcije/server/errors"
	"github.com/oinume/lekcije/server/model"
)

var _ = fmt.Print

func Static(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, r.URL.Path[1:])
}

func Index(w http.ResponseWriter, r *http.Request) {
	if _, err := model.GetLoggedInUser(r.Context()); err == nil {
		http.Redirect(w, r, "/me", http.StatusFound)
	} else {
		indexLogout(w, r)
	}
}

func indexLogout(w http.ResponseWriter, r *http.Request) {
	t := ParseHTMLTemplates(TemplatePath("index.html"))
	type Data struct {
		commonTemplateData
	}
	data := &Data{
		commonTemplateData: getCommonTemplateData(r, false),
	}

	if err := t.Execute(w, data); err != nil {
		InternalServerError(w, errors.InternalWrapf(err, "Failed to template.Execute()"))
		return
	}
}

func RobotsTxt(w http.ResponseWriter, r *http.Request) {
	content := fmt.Sprintf(`
User-agent: *
Allow: /
Sitemap: %s/sitemap.xml
`, config.WebURL())
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, strings.TrimSpace(content))
}

func SitemapXML(w http.ResponseWriter, r *http.Request) {
	content := fmt.Sprintf(`
<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">
  <url>
    <loc>%s/</loc>
    <priority>1.0</priority>
  </url>
  <url>
    <loc>%s/terms</loc>
    <priority>1.0</priority>
  </url>
</urlset>
	`, config.WebURL(), config.WebURL())
	w.Header().Set("Content-Type", "text/xml; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, strings.TrimSpace(content))
}

func Terms(w http.ResponseWriter, r *http.Request) {
	t := ParseHTMLTemplates(TemplatePath("terms.html"))
	type Data struct {
		commonTemplateData
	}
	data := &Data{
		commonTemplateData: getCommonTemplateData(r, false),
	}

	if err := t.Execute(w, data); err != nil {
		InternalServerError(w, errors.InternalWrapf(err, "Failed to template.Execute()"))
		return
	}
}
