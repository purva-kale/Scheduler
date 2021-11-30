package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/go-co-op/gocron"
)

var s = gocron.NewScheduler(time.UTC)
var wg sync.WaitGroup
var count = 0
//var count1 = 0
var count2 = 0


func task() {
	count += 1
	fmt.Println("Welcome to the first Schedule->",count)
	if count == 4 {
		wg.Done()
		s.RemoveByTag("tag1")
		
	}
}

func task1() {
	//count1 += 1
	fmt.Println("Here's Second Schedule")

	//if count1 == 7 {
		wg.Done()
	s.RemoveByTag("tag2")
		
	
}

func task2() {
	count2 += 1
	fmt.Println(" The 3rd Schedule -->",count2)

	if count2 == 5 {
		wg.Done()
		s.RemoveByTag("tag3")
		
	}
}

func main() {
	s.Every(2).Seconds().Tag("tag1").Do(task)
	s.Every(1).Seconds().Tag("tag2").Do(task1)
	s.Every(2).Seconds().Tag("tag3").Do(task2)
	
	wg.Add(3)
	go s.StartBlocking()
	wg.Wait()
	fmt.Println("----Finally done!----")
	
	}
