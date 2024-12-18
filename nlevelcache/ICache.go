package main

type ICache interface {
	Read(key string) (*string, error)
	Write(key, value string) error
	Next(cache ICache)
}
