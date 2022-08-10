package pipefilter
// Package pipe-filter is to define the interfaces and the structures for pipe-filter style implementation

// Request is the input of the filter
type Request interface{}

// Response is the output of the filter
type Response interface{}

// Filter interface is the definition of the data processing components
// Pipe-Filter structure
type Filter interface {
	Process(data Request) (Response, error)
}
