package main

import (
	"fmt"
	"net/http"
	"sync/atomic"
)

var number uint64 = 0

func main() {
	// m := sync.Mutex{}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// m.Lock()
		// m.Unlock()
		atomic.AddUint64(&number, 1)
		w.Write([]byte(fmt.Sprint("Voce é o visitante n:%d", number)))
	})

	http.ListenAndServe(":3000", nil)
}
