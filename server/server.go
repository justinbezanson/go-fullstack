package main

import (
	"net/http"
	"sync"
	"strings"
	"strconv"
)

type Person struct {
	Name string
	Age  int
}

type PersonController struct {
	sync.Mutex
	dataStore map[int]Person
}

/**
 * GET /api/person/:id
 */
func (c *PersonController) show(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) != 4 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	id, err := strconv.Atoi(parts[3])
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	c.Lock()
	defer c.Unlock()
	person, ok := c.dataStore[id]
	if !ok {
		http.Error(w, "Person not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(person.Name))
}

func newPersonController() *PersonController {
	return &PersonController{
		dataStore: map[int]Person{
			1: Person{Name: "Justin", Age: 30},
			2: Person{Name: "John", Age: 40},
			3: Person{Name: "Jack", Age: 50},
		},
	}
}

func main() {
	controller := newPersonController()
	http.HandleFunc("/api/person/", controller.show)
	
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}