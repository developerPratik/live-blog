package main

import (
	"fmt"
	rethink "gopkg.in/rethinkdb/rethinkdb-go.v6"
	"log"
	"real_time_go/config"
	"real_time_go/database"
	"real_time_go/model"
	"time"
)

func init() {
	if err := config.Load(); err != nil {
		log.Fatal(err)
	}

}

func InsertData(ch chan<- string, session *rethink.Session) {
	for  {
		post := model.Post{
			Author: "pratik ",
			Body:   time.Now().String(),
			Title: "post title",
		}

		_, err := rethink.Table("posts").Insert(&post).RunWrite(session)

		if err != nil {
			ch <- err.Error()
		}

		ch <- "Created a new record"

		time.Sleep(time.Second * 2)
	}

}

func RetrieveData(ch chan<- string, session *rethink.Session) {
	cursor, err := rethink.DB("real_time_test").Table("posts").Changes().Run(session)

	if err != nil {
		ch <- err.Error()
		return
	}

	if cursor == nil {
		ch <- "something went wrong"
		return
	}

	var post model.Post

	for cursor.Next(&post) {
		fmt.Println(post.Title)
		ch <- "got a new record"
	}



}

func main() {
	db := database.Get()
	err := db.CreateConnection()

	if err != nil {
		log.Fatal(err)
	}

	log.Println("connected to database")

	insertChannel := make(chan string, 20)
	retrieveChannel := make(chan string, 20)

	response, err := rethink.Table("posts").Delete().Run(db.Session)

	fmt.Println(response, err)


	go InsertData(insertChannel, db.Session)
	go RetrieveData(retrieveChannel, db.Session)

	for {
		select {
		case insertEvent := <-insertChannel:

			fmt.Println("from insert", insertEvent)
		case getEvent := <-retrieveChannel:
			fmt.Println("from get", getEvent)
		}
	}
}
