package main

import "sync"

type Database interface {
	Init()
	UpdateDatabase() error
	WriteStatistics([]CreatureStatistic, string, chan bool, *sync.WaitGroup)
	Close()
}
