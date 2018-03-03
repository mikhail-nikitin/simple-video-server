package main

import (
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	c := &BadPoemClient{}
	c.cb = &CircuitBreaker{}
	c.run()
}

func getSessionId() (string, error) {
	resp, err := http.Get("http://bad-poem.herokuapp.com/register/mikhail-nikitin")
	if err != nil {
		return "", err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	return string(body), nil
}

func getStatus(s string) (int, error) {
	resp, err := http.Get("http://bad-poem.herokuapp.com/get/" + s)
	if err != nil {
		return 0, err
	}
	return resp.StatusCode, nil
}

type CircuitBreakerState int

const (
	Closed CircuitBreakerState = iota
	HalfOpen
	Open
)

type CircuitBreaker struct {
	state        CircuitBreakerState
	failureCount int
}

func (r *CircuitBreaker) isAvailable() bool {
	return r.state == Closed || (r.state == HalfOpen && rand.Float32() < 0.5)
}

func (r *CircuitBreaker) registerFailure() {
	if r.state == Closed || r.state == HalfOpen {
		r.state = Open
		go func() {
			time.Sleep(time.Second * 3)
			r.state = HalfOpen
		}()
	}
}

func (r *CircuitBreaker) registerSuccess() {
	if r.state == HalfOpen {
		r.state = Closed
	}
}

type BadPoemClient struct {
	cb *CircuitBreaker
}

func (r *BadPoemClient) run() {
	isSessionToGet := true
	s := ""
	var err error
	for {
		if isSessionToGet {
			s, err = getSessionId()
			if err != nil {
				r.cb.registerFailure()
			}
			isSessionToGet = false
		}

		if r.cb.isAvailable() {
			status, err := getStatus(s)
			if err != nil || status == 500 {
				r.cb.registerFailure()
			}
			if err == nil && status == 404 {
				isSessionToGet = true
				time.Sleep(time.Millisecond * 5)
			}
		}
	}
}
