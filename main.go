package main

import (
	"github.com/go-playground/webhooks/v6/gitlab"
	"log"
	"net/http"
	"os"
	"time"

	tele "gopkg.in/telebot.v3"
)

func main() {
	pref := tele.Settings{
		Token:  os.Getenv("TOKEN"),
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
	}

	b.Handle("/hello", func(c tele.Context) error {
		return c.Send(c.Data())
	})

	hook, _ := gitlab.New()
	http.HandleFunc("/webhooks", func(w http.ResponseWriter, r *http.Request) {
		payload, err := hook.Parse(r, gitlab.IssuesEvents, gitlab.MergeRequestEvents)
		if err != nil {
			if err == gitlab.ErrEventNotFound {
				log.Println("Event not supported")
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}
		switch payload.(type) {
		case gitlab.IssueEventPayload:
			//content := payload.(gitlab.IssueEventPayload)
			//b.Send("", content)
		}
		w.WriteHeader(http.StatusOK)
	})
	if err != nil {
		log.Fatal(err)
	}

	b.Start()
	log.Fatal(http.ListenAndServe(":8980", nil))
}
