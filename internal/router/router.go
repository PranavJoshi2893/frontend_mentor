package router

import (
	"net/http"

	qrcode "github.com/PranavJoshi2893/frontend_mentor/internal/qr-code"
)

func Routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /qr-code", qrcode.QRCodeHandler)

	fs := http.FileServer(http.Dir("./web/static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	mux.HandleFunc("/", qrcode.PageNotFoundHandler)

	return mux
}
