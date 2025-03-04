package single_flight

import (
	"fmt"
	"golang.org/x/sync/singleflight"
	"sync"
	"time"
)

func main() {
	var g singleflight.Group
	var wg sync.WaitGroup

	expensiveOperation := func(key string) (any, error) {
		time.Sleep(2 * time.Second)
		return fmt.Sprintf("Result: %s", key), nil
	}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			key := "some-key"
			result, err, shared := g.Do(key, func() (any, error) {
				return expensiveOperation(key)
			})
			fmt.Printf("Goroutine %d: Result: %v, Error: %v, Shared: %v\n", i, result, err, shared)
		}()
	}

	wg.Wait()
}
