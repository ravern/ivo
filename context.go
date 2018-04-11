package ivo

import "log"

type Context interface {
	Logger() *log.Logger
}
