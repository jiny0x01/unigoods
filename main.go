package main

import "net/http"

const notionURL = "https://unigoods.notion.site/6fe54557b45b41d99e4f5f5512d2b86e?v=32a3b81a340e4553ae52523a872389c7&pvs=4"

func redirectToFreshman(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, notionURL, http.StatusSeeOther)
}

func main() {
	mux := http.NewServeMux()

	mux.Handle("/", http.RedirectHandler(notionURL, http.StatusSeeOther))

	http.ListenAndServe(":80", mux)
}
