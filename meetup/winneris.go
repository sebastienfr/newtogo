package main

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"html/template"
	"math/big"
	"net/http"
	"os"
	"strings"
)

// query the query URL
var query string = "https://api.meetup.com/GDG-Lille/events/236825262/rsvps?key=%s&sign=true&photo-host=public"

// RSVP strucutre
type RSVP struct {
	Response string `json:"response"`
	Member   struct {
		Name  string `json:"name"`
		Photo struct {
			HiRes string `json:"photo_link"`
		} `json:"photo"`
		Context struct {
			Host bool `json:"host"`
		} `json:"event_context"`
	} `json:"Member"`
}

// RSVPArray array of RSVP
type RSVPArray []RSVP

// Winner simplified structure for template
type Winner struct {
	Name  string
	Photo string
}

// main
func main() {
	// new http client
	hc := &http.Client{}

	// get secret API KEY
	key := os.Getenv("MEETUP_APIKEY")

	// retrieve participant liste
	response, err := hc.Get(fmt.Sprintf(query, key))
	if err != nil {
		panic(err)
		return
	}

	// decode response
	rsvpArray := make(RSVPArray, 0)
	err = json.NewDecoder(response.Body).Decode(&rsvpArray)

	if err != nil {
		panic(err)
		return
	}

	// filter ok and not organizer
	rsvpArrayOk := make(RSVPArray, 0)
	for _, rsvp := range rsvpArray {

		if "yes" == strings.ToLower(rsvp.Response) &&
			!rsvp.Member.Context.Host {
			rsvpArrayOk = append(rsvpArrayOk, rsvp)
		}
	}

	// print result count
	fmt.Printf("Members %d\n", len(rsvpArrayOk))

	// build http server to display the random winner
	http.HandleFunc("/index.html", func(w http.ResponseWriter, r *http.Request) {

		t, err := template.ParseFiles("index.html")

		if err != nil {
			panic(err)
		}

		index, err := rand.Int(rand.Reader, big.NewInt(int64(len(rsvpArrayOk))))
		if err != nil {
			panic(err)
		}

		randomRsvp := rsvpArrayOk[index.Int64()]

		winner := Winner{
			Name:  randomRsvp.Member.Name,
			Photo: randomRsvp.Member.Photo.HiRes,
		}

		fmt.Printf("Winner is %+v \n", winner)

		err = t.ExecuteTemplate(w, "index.html", winner)
		if err != nil {
			panic(err)
		}
	})

	http.Handle("/", http.FileServer(http.Dir(".")))

	// serve content
	err = http.ListenAndServe(":3000", nil)
	if err != nil {
		panic(err)
	}

}
