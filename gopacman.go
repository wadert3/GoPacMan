package main

import (
	"GoPacMan.wadert3/pacman"
)

//HTTP server main, logging to console for now

//func main() {
//	http.HandleFunc("/", indexHandler)
//	port := os.Getenv("PORT")
//	if port == "" {
//		port = "8080"
//		log.Printf("Defaulting to port %s", port)
//	}
//
//	log.Printf("Listening on port %s", port)
//	log.Printf("Open http://localhost:%s in the browser", port)
//	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
//}

//func indexHandler(w http.ResponseWriter, r *http.Request) {
//	if r.URL.Path != "/" {
//		http.NotFound(w, r)
//		return
//	}
//	_, err := fmt.Fprint(w, buildBoard(gameMap))
//	if err != nil {
//		w.WriteHeader(http.StatusInternalServerError)
//	}
//}

func main() {

	screen := pacman.InitializeScreen()
	game := pacman.InitializeGame(screen)

	defer screen.Fini()
	go game.HandleEvents()
	game.RunGame()
}
