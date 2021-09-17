package main

func main() {

	//algoritms for clean cache (capasity = 2)
	lru := &lru{}
	cache := initCache(lru)

	// set algoritm LRU (Least Recently Used)
	cache.setEvictionAlgo(lru)
	cache.add("a", "1")
	cache.add("b", "2")
	cache.add("c", "3")

	// set algoritm FIFO
	fifo := &fifo{}
	cache.setEvictionAlgo(fifo)

	cache.add("d", "4")
	/*Evicting by lru strtegy
	  Evicting by fifo strtegy*/

}
