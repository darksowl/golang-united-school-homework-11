package batch

//package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"golang.org/x/sync/errgroup"
)

type user struct {
	ID int64
}

func getOne(id int64) user {
	time.Sleep(time.Millisecond * 100)
	return user{ID: id}
}

func getBatch(n int64, pool int64) (res []user) {
	//const parallel = 5
	//var gopherID []user
	errG, _ := errgroup.WithContext(context.Background())
	errG.SetLimit(int(pool))

	var mu sync.Mutex
	//c := make(chan user)
	for i := 0; i < int(n); i++ {
		j := i
		//c := make(chan user)
		errG.Go(func() error {
			//c <- getOne(int64(j))
			user := getOne(int64(j))

			mu.Lock()
			res = append(res, user)
			mu.Unlock()
			//c <- res
			return nil
		})
		//gopherID = <-c
		//res = append(res, gopherID)
		//err := errG.Wait()
		//if err := errG.Wait(); err != nil {
		//	fmt.Printf("Error group error: %s", err.Error())
		//}
		//if err != nil {
		//	fmt.Println(err)
		//	return nil
		//}
	}
	//for i := 0; i < int(pool); i++ {
	//	gopherID := <-c
	//	res = append(res, gopherID)
	//}
	if err := errG.Wait(); err != nil {
		fmt.Printf("Error group error: %s", err.Error())
	}
	//fmt.Println(gopherID)
	//err := errG.Wait()
	//fmt.Println(err)
	//return nil
	return res
}

func main() {
	res := getBatch(10, 10)
	fmt.Println(res)
}
