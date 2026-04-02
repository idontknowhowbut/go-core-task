package wg

type waitGroup struct {
	counter int
	mu      chan struct{}
	done    chan struct{}
}

func New() *waitGroup {
	wg := waitGroup{
		counter: 0,
		mu:      make(chan struct{}, 1),
		done:    make(chan struct{}),
	}
	wg.mu <- struct{}{}
	close(wg.done)

	return &wg
}

func (wg *waitGroup) Add(n int) {
	<-wg.mu
	defer func() {
		wg.mu <- struct{}{}
	}()

	if wg.counter+n < 0 {
		panic("WaitGroup can't be < 0")
	}

	if wg.counter == 0 && wg.counter+n > 0 {
		wg.done = make(chan struct{})
	}

	wg.counter = wg.counter + n

	if wg.counter == 0 {
		close(wg.done)
	}

}

func (wg *waitGroup) Done() {
	wg.Add(-1)
}

func (wg *waitGroup) Wait() {
	<-wg.mu
	done := wg.done
	wg.mu <- struct{}{}

	<-done
}
