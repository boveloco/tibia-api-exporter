package main

import (
	"log"
	"os"
)

func getEnv(key string, fallback ...string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	if len(fallback) == 0 {
		log.Panicf("Variable %s not found in environment.", key)
	}
	return fallback[0]
}
