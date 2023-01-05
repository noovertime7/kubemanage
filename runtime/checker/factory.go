package checker

import (
	"reflect"
	"sync"
)

type SharedCheckerFactory interface {
	CheckerFor(ck Checker) Checker
	Start()
}

type SharedCheckerOption func(*checkerFactory) *checkerFactory

func NewSharedCheckerFactory(options ...SharedCheckerOption) SharedCheckerFactory {
	factory := &checkerFactory{
		startedCheckers: make(map[reflect.Type]bool),
		checkers:        make(map[reflect.Type]Checker),
	}

	// Apply all options
	for _, opt := range options {
		factory = opt(factory)
	}

	return factory
}

type checkerFactory struct {
	lock            sync.Mutex
	startedCheckers map[reflect.Type]bool
	checkers        map[reflect.Type]Checker
}

func (c *checkerFactory) CheckerFor(ck Checker) Checker {
	checkerType := reflect.TypeOf(ck)
	checkerExists, exists := c.checkers[checkerType]
	if exists {
		return checkerExists
	}
	c.checkers[checkerType] = ck
	return ck
}

func (c *checkerFactory) Start() {
	c.lock.Lock()
	defer c.lock.Unlock()

	for checkerType, checker := range c.checkers {
		if !c.startedCheckers[checkerType] {
			go checker.Run()
			c.startedCheckers[checkerType] = true
		}
	}
}
