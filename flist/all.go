package flist

import "sync"

type Queue struct {
	path string
	opts []Option
}

func NewQueue(path string, opts ...Option) *Queue {
	return &Queue{
		path: path,
		opts: opts,
	}
}

type getResult struct {
	list *[]string
	err  error
}

func GetAll(queues ...Queue) ([]*[]string, []error) {
	rtn := []*[]string{}

	var wg sync.WaitGroup
	wg.Add(len(queues))
	reses := []getResult{}
	for _, q := range queues {
		go func(q Queue) {
			defer wg.Done()
			list, err := Get(q.path, q.opts...)
			reses = append(reses, getResult{list: list, err: err})

		}(q)
	}

	wg.Wait()

	for _, r := range reses {
		if r.err != nil {
			return nil, []error{r.err}
		}
		rtn = append(rtn, r.list)
	}

	return rtn, nil
}
