package frontend

import "net/http"

const PATH = "frontend/svelte-app/public"

func InitFrontEnd() {
	fs := http.FileServer(http.Dir(PATH))
	http.Handle("/", fs)
}
