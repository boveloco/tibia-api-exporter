package main

import "sync"

type Database interface {
	Init()
	UpdateDatabase()
	WriteStatistics([]CreatureStatistic, string, chan bool, *sync.WaitGroup)
	Close()
}
