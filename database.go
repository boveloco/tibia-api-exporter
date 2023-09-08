package main

import "sync"

type Database interface {
	Init()
	UpdateDatabase() bool
	WriteStatistics([]CreatureStatistic, string, chan bool, *sync.WaitGroup)
	Close()
}
