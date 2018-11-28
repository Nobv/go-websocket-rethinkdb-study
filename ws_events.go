package main

import (
	"log"

	r "gopkg.in/gorethink/gorethink.v4"
)

func allChanges(ch chan interface{}) {
	go func() {
		for {
			res, err := r.DB("chat").Table("messages").Changes().Run(session)
			if err != nil {
				log.Fatalln(err)
			}

			var response interface{}
			for res.Next(&response) {
				ch <- response
			}

			if res.Err() != nil {
				log.Println(res.Err())
			}
		}
	}()
}
