module github.com/NoistNT/go-first-practice

go 1.23.1

require github.com/gorilla/mux v1.8.1

require github.com/NoistNT/go-first-practice/handlers v0.0.0-00000000000000-000000000000

replace github.com/NoistNT/go-first-practice/handlers => ./handlers

require github.com/NoistNT/go-first-practice/api v0.0.0-00000000000000-000000000000 //indirect

replace github.com/NoistNT/go-first-practice/api => ./api
