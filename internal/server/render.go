package server

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/go-chi/render"

	"github.com/Chyroc/vlog/internal/common"
)

func RenderError(w http.ResponseWriter, r *http.Request, err error) {
	render.Status(r, http.StatusInternalServerError)
	render.JSON(w, r, err)
}

func RenderSuccess(w http.ResponseWriter, r *http.Request, body interface{}) {
	render.Status(r, http.StatusOK)
	render.JSON(w, r, body)
}

func RenderHTML(w http.ResponseWriter, r *http.Request, view string, data map[string]interface{}) error {
	b, err := ioutil.ReadFile("./internal/static/" + view + ".html")
	if err != nil {
		return err
	}

	html, err := common.ParseTmpl(string(b), data)
	if err != nil {
		return err
	}

	render.HTML(w, r, string(html))

	return nil
}

func BindJSON(r *http.Request, obj interface{}) error {
	defer io.Copy(ioutil.Discard, r.Body)
	return json.NewDecoder(r.Body).Decode(obj)
}
