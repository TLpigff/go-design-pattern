/*
  @Author :     lyb
  @File :       filter
  @Description:
*/
package filter

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// Request is the input of the filter
type Request interface{}
// Response is the output of the filter
type Response interface{}

type Filter interface {
	Process(data Request) (Response,error)
}

type SplitFilter struct {
	delimiter string
}

func NewSplitFilter(delimiter string) *SplitFilter {
	return &SplitFilter{delimiter:delimiter}
}

func (sf *SplitFilter) Process(data Request) (Response,error) {
	str,ok := data.(string)	//检测数据格式类型，是否可以处理

	if !ok {
		return nil,fmt.Errorf("input data must be string")
	}

	parts := strings.Split(str,sf.delimiter)
	return parts,nil
}

type ToIntFilter struct {}

func NewToIntFilter() *ToIntFilter {
	return &ToIntFilter{}
}

func (tlf *ToIntFilter) Process(data Request) (Response,error) {
	parts,ok := data.([]string)
	if !ok {
		return nil,fmt.Errorf("input data should be []string")
	}

	ret := make([]int,0)
	for _,part := range parts {
		s,err := strconv.Atoi(part)
		if err != nil {
			return nil,err
		}

		ret = append(ret,s)
	}
	return ret,nil
}

type SumFilter struct {
}

func NewSumFilter() *SumFilter {
	return &SumFilter{}
}

func (sf *SumFilter) Process(data Request) (Response, error) {
	elms, ok := data.([]int)
	if !ok {
		return nil, errors.New("input data should be []int")
	}
	ret := 0
	for _, elem := range elms {
		ret += elem
	}
	return ret, nil
}

type Pipeline struct {
	Name    string
	Filters *[]Filter
}

func NewPipeline(name string, filters ...Filter) *Pipeline {
	return &Pipeline{
		Name:    name,
		Filters: &filters,
	}
}

// call each filter's process function
func (p *Pipeline) Process(data Request) (Response, error) {
	var ret interface{}
	var err error
	for _, filter := range *p.Filters {
		ret, err = filter.Process(data)
		if err != nil {
			return ret, err
		}
		data = ret
	}
	return ret, err
}
