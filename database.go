package main

type Database interface {
	Init()
	UpdateDatabase()
	WriteStatistics([]CreatureStatistic, string) bool
	Close()
}
