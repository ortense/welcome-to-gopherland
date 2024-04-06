package ping

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func Ping(url string) {
	for {
		start := time.Now()

		res, err := http.Get(url)

		// erros devem ser tratados explicidamente
		if err != nil {
			fmt.Println(err)
			return
		}

		duration := time.Since(start)

		log.Println(res.StatusCode, url, "response time", duration)

		if duration < time.Second {
			time.Sleep(time.Second - duration)
		}
	}
}
