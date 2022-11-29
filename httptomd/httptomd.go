package httptomd

import (
	"crypto/md5"
	"fmt"
	"io"
	"net/http"
	netURL "net/url"
	"strings"
	"sync"
)

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error) // Do is a mock for http.Client.Do
}

var Client HTTPClient

func init() {
	Client = &http.Client{}
}

// Run pushes the URLs to the jobs channel and starts the workers
func Run(parallel int, urls []string) {
	var wg sync.WaitGroup
	jobs := make(chan string, len(urls))

	for i := 0; i < parallel; i++ {
		wg.Add(1)
		go runJob(jobs, &wg) // start the workers
	}

	for _, url := range urls {
		jobs <- url // push the URLs to the jobs channel
	}
	close(jobs) // close the jobs channel

	wg.Wait()
}

// runJob is a worker function that gets a URL from the jobs channel,
func runJob(jobs <-chan string, wg *sync.WaitGroup) {
	for url := range jobs {
		newURL, err := validateURL(url) // validate the URL
		if err != nil {
			fmt.Printf("[ERROR] %v\n", err)
			continue
		}

		checksum, err := getMD5(newURL)
		if err != nil {
			fmt.Printf("[ERROR] %v\n", err)
		} else {
			fmt.Printf("%s %x\n", newURL, checksum)
		}
	}

	wg.Done()
}

// get MD5 checksum of the URL
func getMD5(url string) (checksum [16]byte, err error) {
	req, err := http.NewRequest(http.MethodGet, url, nil) // make a request
	if err != nil {
		err = fmt.Errorf("invalid http request - %v", err)
		return
	}

	resp, err := Client.Do(req)
	if err != nil {
		err = fmt.Errorf("failed to get URL %s - %v", url, err)
		return
	}

	if resp.StatusCode != http.StatusOK { // 200
		err = fmt.Errorf("request %s failed with StatusCode=%d", url, resp.StatusCode)
		return
	}

	body, err := io.ReadAll(resp.Body) // read the body
	if err != nil {
		err = fmt.Errorf("failed to parse body - %v", err)
		return
	}

	checksum = md5.Sum(body)                 // calculate the checksum
	md5String := fmt.Sprintf("%x", checksum) // convert checksum to string
	if len(md5String) > 0 {                  // if the checksum is not empty
		return checksum, nil // return the checksum
	}
	return
}

// validate new URL and add http:// if needed
func validateURL(oldURL string) (newURL string, err error) {
	newURL = oldURL
	if !strings.HasPrefix(oldURL, "http://") { // add http:// if needed
		newURL = "http://" + oldURL
	}

	u, err := netURL.ParseRequestURI(newURL) // validate the URL
	if err != nil || u.Host == "" {          // if the URL is invalid or the host is empty
		err = fmt.Errorf("invalid URL '%s'", newURL) // return an error
	}
	return
}
