package main

import "sync"

type Database interface {
	Init()
	UpdateDatabase() error
	WriteStatistics([]CreatureStatistic, string, *sync.WaitGroup)
	ValidateExecution() (bool, error)
	SetLastExecution()
	Close()
}
