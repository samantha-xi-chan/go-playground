package play110_sse

import (
	"fmt"
	"github.com/antage/eventsource"
	"log"
	"net/http"
	"strconv"
)

func Play() {
	http.Handle("/", http.FileServer(http.Dir("./static"))) // Serve static HTML/JS files
	http.HandleFunc("/events", handleEvents)                // Handle EventSource requests
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleEvents(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	xStr := r.URL.Query().Get("x")
	x, err := strconv.Atoi(xStr)
	if err != nil {
		http.Error(w, "Invalid input. Please provide a valid numeric value for x.", http.StatusBadRequest)
		return
	}

	// Set the proper Content-Type header for SSE
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	es := eventsource.New(
		eventsource.DefaultSettings(),
		func(req *http.Request) [][]byte {
			// This function returns event data to be sent to the client.
			// You can implement your logic here to send the appropriate events.
			x++
			return [][]byte{[]byte(fmt.Sprintf("data: %d\n\n", x))}
		},
	)
	defer es.Close()

}
