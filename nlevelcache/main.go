package main

import "fmt"

func main() {

	cm := NewCacheManager(3)

	err := cm.CacheLevels[2].Write("kamlesh", "sahu")
	if err != nil {
		fmt.Printf("err while writing to level %d cache, %s = %s", 0, "kamlesh", "sahu")
	}
	value, err := cm.Read("kamlesh")
	if err != nil {
		fmt.Printf("err while reading from level %d cache, %s\n", 0, "kamlesh")
	} else {
		fmt.Printf("found %s = %s\n", "kamlesh", *value)
	}

	value, err = cm.Read("nilesh")
	if err != nil {
		fmt.Printf("err while reading from cache, %s, %s\n", "nilesh", err)
	} else {
		fmt.Printf("found %s = %s\n", "kamlesh", *value)
	}
	err = cm.CacheLevels[0].Write("nilesh", "bhai")
	if err != nil {
		fmt.Printf("err while writing to level %d cache, %s = %s", 1, "nilesh", "bhai")
	}
	value, err = cm.Read("nilesh")
	if err != nil {
		fmt.Printf("err while reading from cache, %s, %s\n", "nilesh", err)
	} else {
		fmt.Printf("found %s = %s\n", "nilesh", *value)
	}
}
