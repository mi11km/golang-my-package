package templates

import (
	"app/pkg/chat/middleware"
	"log"
	"net/http"
	"path/filepath"
	"sync"
	"text/template"
)

type TemplateHandler struct {
	once     sync.Once
	Filename string
	templ    *template.Template
}

func (t *TemplateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("pkg/chat/templates", t.Filename)))
	})
	data := map[string]interface{}{
		"Host": r.Host,
	}
	userData := middleware.DecodeUserInfo(r)
	if userData != nil {
		data["UserData"] = userData
	}
	if err := t.templ.Execute(w, data); err != nil {
		log.Fatal(err)
	}
}
