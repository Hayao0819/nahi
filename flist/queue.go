package flist

import "sync"

type Queue struct {
	name string
	path string
	opts []Option
}

func NewQueue(name string, path string, opts ...Option) *Queue {
	return &Queue{
		name: name,
		path: path,
		opts: opts,
	}
}

type getResult struct {
	name string
	list *[]string
	err  error
}

func GetAll(queues ...Queue) (map[string]*[]string, []error) {
	rtn := map[string]*[]string{}

	var wg sync.WaitGroup
	wg.Add(len(queues))
	reses := []getResult{}
	for _, q := range queues {
		go func(q Queue) {
			defer wg.Done()
			list, err := Get(q.path, q.opts...)
			reses = append(reses, getResult{name: q.name, list: list, err: err})

		}(q)
	}

	wg.Wait()

	for _, r := range reses {
		if r.err != nil {
			return nil, []error{r.err}
		}
		// rtn = append(rtn, r.list)
		rtn[r.name] = r.list
	}

	return rtn, nil
}
