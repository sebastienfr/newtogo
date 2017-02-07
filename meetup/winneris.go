package main

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"os"
	"strings"
)

var query string = "https://api.meetup.com/GDG-Lille/events/236825262/rsvps?key=%s&sign=true&photo-host=public"

type RSVP struct {
	Response string `json:"response"`
	Member   struct {
		Name  string `json:"name"`
		Photo struct {
			HiRes string `json:"highres_link"`
		} `json:"photo"`
		Context struct {
			Host bool `json:"host"`
		} `json:"event_context"`
	} `json:"Member"`
}

type RSVPArray []RSVP

func main() {
	hc := &http.Client{}

	key := os.Getenv("MEETUP_APIKEY")
	response, err := hc.Get(fmt.Sprintf(query, key))
	if err != nil {
		panic(err)
		return
	}

	rsvpArray := make(RSVPArray, 0)
	err = json.NewDecoder(response.Body).Decode(&rsvpArray)

	if err != nil {
		panic(err)
		return
	}

	//fmt.Printf("Members %+v\n", rsvpArray)

	rsvpArrayOk := make(RSVPArray, 0)
	for _, rsvp := range rsvpArray {

		if "yes" == strings.ToLower(rsvp.Response) &&
			!rsvp.Member.Context.Host {
			//fmt.Printf("adding %+v\n", rsvp)
			rsvpArrayOk = append(rsvpArrayOk, rsvp)
		}
	}

	fmt.Printf("Members %d\n", len(rsvpArrayOk))
	//fmt.Printf("Members %+v\n", rsvpArrayOk)

	index, err := rand.Int(rand.Reader, big.NewInt(int64(len(rsvpArrayOk))))
	if err != nil {
		panic(err)
	}

	fmt.Printf("Winner is # %d => %+v", index.Int64(), rsvpArrayOk[index.Int64()])

}
