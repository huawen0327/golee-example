package main

import (
	"github.com/huawen0327/golee"
)

func url(r *golee.Router) {
	r.Path("/aa", logger, index)
	r.Path("/hello", hello)
}
func urlGroup(g *golee.Group) {
	g.SetGroup("/g1", logger)
	g.Path("/hello", hello)
	g.Path("/hello1", hello)
	g.Path("/temp", temp)
}
