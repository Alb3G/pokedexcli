package main

type CliCommand struct {
	name        string
	description string
	callback    func() error
}

type Config struct {
	previousUrl string
	nextUrl     string
}
