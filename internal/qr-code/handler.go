package qrcode

import (
	"bytes"
	"html/template"
	"net/http"
)

func QRCodeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./web/templates/layouts/qr_code.html"))

	var buf bytes.Buffer
	err := tmpl.ExecuteTemplate(&buf, "qr_code.html", nil)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(buf.Bytes())

}

func PageNotFoundHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./web/templates/layouts/not_found.html"))

	var buf bytes.Buffer
	err := tmpl.ExecuteTemplate(&buf, "not_found.html", nil)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusNotFound)
	w.Write(buf.Bytes())
}
