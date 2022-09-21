package redisx

import (
	"fmt"
	"testing"
	"time"
)

func TestR(t *testing.T) {
	conf := Conf{
		DSN: "redis://:123456@0.0.0.0:6379/1",
	}
	rds := Register("test", conf)

	s := rds.Set("test", 1, time.Minute)
	fmt.Println("s", s.Err())
	g := rds.Get("test")
	fmt.Println("s", g.Val())

	g2 := rds.Get("test2")
	fmt.Println(g2.Err())

}

func TestZard(t *testing.T) {
	conf := Conf{
		DSN: "redis://:123456@0.0.0.0:6379/1",
	}
	rds := Register("test", conf)

	rds.LLen()

}
