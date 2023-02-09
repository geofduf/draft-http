package templates

import (
	"bytes"
	"errors"
	"html/template"
	"io/fs"
	"net/http"
	"regexp"
)

type Store struct {
	Globals   globals
	templates map[string]*template.Template
}

func NewStore(fsys fs.FS) (*Store, error) {
	s := &Store{
		globals{data: make(map[string]any)},
		make(map[string]*template.Template),
	}
	if err := s.initTemplates(fsys, nil); err != nil {
		return nil, err
	}
	return s, nil
}

func NewStoreWithFuncMap(fsys fs.FS, funcMap template.FuncMap) (*Store, error) {
	s := &Store{
		globals{data: make(map[string]any)},
		make(map[string]*template.Template),
	}
	if err := s.initTemplates(fsys, funcMap); err != nil {
		return nil, err
	}
	return s, nil
}

func (s *Store) initTemplates(fsys fs.FS, funcMap template.FuncMap) error {
	base := template.New("base.html")
	if funcMap != nil {
		_ = base.Funcs(funcMap)
	}
	_, err := base.ParseFS(fsys, "base.html")
	if err != nil {
		return err
	}
	re := regexp.MustCompile(`^[a-zA-Z0-9]+\.html$`)
	files, err := fs.ReadDir(fsys, ".")
	if err != nil {
		return err
	}
	for _, file := range files {
		if re.MatchString(file.Name()) && file.Name() != "base.html" {
			s.templates[file.Name()], err = base.Clone()
			if err != nil {
				return err
			}
			_, err = s.templates[file.Name()].ParseFS(fsys, file.Name())
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (s *Store) Render(w http.ResponseWriter, contentType string, name string, data map[string]any) error {
	t, ok := s.templates[name]
	if !ok {
		return errors.New("template not found")
	}

	if data == nil {
		data = make(map[string]any)
	}

	s.Globals.mu.RLock()
	for k, v := range s.Globals.data {
		if _, ok := data[k]; !ok {
			data[k] = v
		}
	}
	s.Globals.mu.RUnlock()

	var buf bytes.Buffer

	err := t.ExecuteTemplate(&buf, "base.html", data)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", contentType)
	buf.WriteTo(w)
	return nil
}
