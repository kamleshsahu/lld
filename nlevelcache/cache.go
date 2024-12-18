package main

import (
	"errors"
	"fmt"
)

type cache struct {
	Level int
	data  map[string]string
	next  ICache
}

func (c *cache) Read(key string) (*string, error) {
	if _, found := c.data[key]; !found {
		if c.next == nil {
			return nil, errors.New("data not found till last level")
		}
		value, err := c.next.Read(key)
		if err != nil {
			return nil, err
		}
		c.data[key] = *value
	}
	value := c.data[key]
	fmt.Printf("found %s = %s at level%d cache\n", key, value, c.Level)
	return &value, nil
}

func (c *cache) Write(key, value string) error {
	if _, found := c.data[key]; found {
		fmt.Printf("%s = %s already present at level%d cache\n", key, value, c.Level)
		return nil
	}
	c.data[key] = value
	if c.next != nil {
		err := c.next.Write(key, value)
		if err != nil {
			return err
		}
	}
	fmt.Printf("%s = %s written at level%d cache\n", key, value, c.Level)
	return nil
}

func (c *cache) Next(iCache ICache) {
	c.next = iCache
}

func NewCache(level int) ICache {
	return &cache{data: make(map[string]string), Level: level}
}
