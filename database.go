package main

type Database interface {
	Init()
	Close()
	Write(interface{})
}
