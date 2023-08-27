package main

type Database interface {
	Init()
	Write(data *interface{}) bool
	UpdateDatabase()
	Close()
}
