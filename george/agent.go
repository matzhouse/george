package george

import (
	"fmt"
	"github.com/matzhouse/go-metrics"
	"github.com/matzhouse/go-metrics/http"
	"time"
)

func Start() {

	fmt.Println("Waking george up..")

	registry := metrics.NewRegistry()

	c := metrics.NewCounter()
	registry.Register("test", c)
	c.Inc(1)

	go http.Start(registry)

	for i := 0; i < 10; i++ {
		go func() {
			time.Sleep(10 * time.Second)
		}()
	}

	time.Sleep(60 * time.Second)

}
