package concurrency

import "time"

type WebsiteChecker func(string) bool
type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range urls {
		go func(u string) {
			resultChannel <- result{u, wc(u)} // send statement for channel
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		result := <-resultChannel // receive expression for channel
		results[result.string] = result.bool
	}

	time.Sleep(1 * time.Second)
	return results
}
