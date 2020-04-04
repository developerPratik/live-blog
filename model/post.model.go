package model

type Post struct {
	ID     string `rethinkdb:"id, omitempty"`
	Title  string `rethinkdb:"title"`
	Author string `rethinkdb:"author"`
	Body   string `rethinkdb:"body"`
}