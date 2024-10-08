package main

import (
	"container/list"
	"errors"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	tw := NewTimeWheel(WithInterval(time.Second), WithTotalSlots(10), WithErrorChan(make(chan error, 100)), WithName("demo"))
	go func() {
		for err := range tw.errChan {
			fmt.Println(err.Error())
		}
	}()
	fun1s := func() error {
		fmt.Printf("running after 1 second with %s\n", time.Now().Format("2006-01-02 15:04:05.999"))
		return nil
	}
	fun7s := func() error {
		fmt.Printf("running after 7 second with %s\n", time.Now().Format("2006-01-02 15:04:05.999"))
		return nil
	}
	fun10s := func() error {
		fmt.Printf("running after 10 second with %s\n", time.Now().Format("2006-01-02 15:04:05.999"))
		return nil
	}
	fun15s := func() error {
		fmt.Printf("running after 15 second with %s\n", time.Now().Format("2006-01-02 15:04:05.999"))
		return nil
	}
	if err := tw.Run(); err != nil {
		fmt.Println(err.Error())
	}
	if _, err := tw.AddTask(time.Second, fun1s); err != nil {
		fmt.Println(err.Error())
	}
	if _, err := tw.AddTask(time.Second*7, fun7s); err != nil {
		fmt.Println(err.Error())
	}
	if _, err := tw.AddTask(time.Second*17, fun7s); err != nil {
		fmt.Println(err.Error())
	}
	if _, err := tw.AddTask(time.Second*10, fun10s); err != nil {
		fmt.Println(err.Error())
	}
	if _, err := tw.AddTask(time.Second*20, fun10s); err != nil {
		fmt.Println(err.Error())
	}
	if _, err := tw.AddTask(time.Second*30, fun10s); err != nil {
		fmt.Println(err.Error())
	}
	if _, err := tw.AddTask(time.Second*15, fun15s); err != nil {
		fmt.Println(err.Error())
	}
	// if err := tw.BQuit(); err != nil {
	// 	fmt.Println(err.Error())
	// }
	time.Sleep(time.Second * 60)
}

const (
	DEFAULT_TOTAL_SLOTS            = 3600
	DEFAULT_TIMEWHEEL_STEPDURATION = time.Microsecond
	DEFAULT_TIMEWHEEL_ERRORSIZE    = 1024 // 1.6kb
)

// timewheel running status
const (
	TIMEWHEEL_STATUS_IDEL = iota
	TIMEWHEEL_STATUS_RUNNING
	TIMEWHEEL_STATUS_END
)

type TimeWheel struct {
	name          string
	startTime     time.Time
	stepInterval  time.Duration
	totalSlot     int
	errChan       chan error
	status        int
	statusLock    sync.RWMutex
	quit          chan struct{}
	slots         []*slot
	taskMapper    map[uint64]*HandlerFunc
	currentID     *uint64
	cycleInterval time.Duration
	anchor        int // current slot index
}

type slot struct {
	mutex sync.RWMutex
	tasks *list.List
}
type HandlerFunc func() error

type task struct {
	id       uint64
	cycleNum int
}

// name string, totalSlots int, stepSize time.Duration, errChanBufferSize int

type Option func(*TimeWheel)

func WithTotalSlots(total int) Option {
	return func(tw *TimeWheel) {
		tw.totalSlot = total
	}
}

func WithInterval(size time.Duration) Option {
	return func(tw *TimeWheel) {
		tw.stepInterval = size
	}
}

func WithErrorChan(ch chan error) Option {
	return func(tw *TimeWheel) {
		tw.errChan = ch
	}
}

func WithName(name string) Option {
	return func(tw *TimeWheel) {
		tw.name = name
	}
}

func NewTimeWheel(opts ...Option) (tw *TimeWheel) {
	mapper := make(map[uint64]*HandlerFunc)
	currentID := uint64(0)
	errChan := make(chan error, DEFAULT_TIMEWHEEL_ERRORSIZE)
	tw = &TimeWheel{
		taskMapper: mapper,
		startTime:  time.Now(),
		currentID:  &currentID,
		totalSlot:  DEFAULT_TOTAL_SLOTS,
		quit:       make(chan struct{}),
		errChan:    errChan,
	}
	for _, opt := range opts {
		opt(tw)
	}
	slots := make([]*slot, tw.totalSlot)
	for i := range slots {
		slots[i] = &slot{
			tasks: list.New(),
		}
	}
	tw.slots = slots

	tw.cycleInterval = time.Duration(tw.totalSlot) * tw.stepInterval

	return
}

func (tw *TimeWheel) AddTask(delay time.Duration, handler HandlerFunc) (id uint64, err error) {
	if delay <= 0 {
		err = errors.New("delay must be greater than 0")
		return
	}
	slotIndex := int((delay % tw.cycleInterval) / tw.stepInterval)
	cycleNum := delay / tw.cycleInterval
	id = atomic.AddUint64(tw.currentID, 1)
	tw.slots[slotIndex].mutex.Lock()
	defer tw.slots[slotIndex].mutex.Unlock()
	tw.slots[slotIndex].tasks.PushBack(&task{
		id:       id,
		cycleNum: int(cycleNum),
	})
	tw.taskMapper[id] = &handler
	return
}

func (tw *TimeWheel) Run() (err error) {
	tw.statusLock.RLock()
	defer tw.statusLock.RUnlock()
	if tw.status == TIMEWHEEL_STATUS_RUNNING {
		err = errors.New("timewheel is already running")
		return
	}
	go func() {
		tw.statusLock.Lock()
		tw.status = TIMEWHEEL_STATUS_RUNNING
		tw.statusLock.Unlock()
		ticker := time.NewTicker(tw.stepInterval)
		for {
			select {
			case <-ticker.C:
				if tw.anchor >= tw.totalSlot {
					// new cycle
					tw.anchor = 0
				}
				tw.RunTask(tw.slots[tw.anchor])
				tw.anchor++
			case <-tw.quit:
				tw.status = TIMEWHEEL_STATUS_END
				return
			}

		}
	}()
	return
}

func (tw *TimeWheel) RunTask(s *slot) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	for v := s.tasks.Front(); v != nil; v = v.Next() {
		n := v // copy the val, so it can used in go func
		if t := n.Value.(*task); t.cycleNum == 0 {
			go func() {
				defer func() {
					// clean up function
					s.mutex.Lock()
					fmt.Println("clean up task id", t.id)
					s.tasks.Remove(n)
					s.mutex.Unlock()
				}()
				handler, ok := tw.taskMapper[t.id]
				if !ok {
					tw.errChan <- fmt.Errorf("task id %v not found", t.id)
					return
				}
				err := (*handler)()
				if err != nil {
					tw.errChan <- fmt.Errorf("task id %v run failed: %v", t.id, err)
					return
				}
			}()
		} else {
			t.cycleNum--
		}

	}
}
