package flyjapan

import (
	"context"
	"sync"
)

type Worker interface {
	Work(ctx context.Context) (<-chan *Result, <-chan error)
}

type worker struct {
	workerOptions
	queries []*Query
	jobs    chan *job
}

func NewWorker(queries []*Query, setters ...WorkerOption) Worker {
	worker := &worker{
		workerOptions: defaultWorkerOption,
		queries:       queries,
		jobs:          make(chan *job),
	}

	for _, setter := range setters {
		setter(&worker.workerOptions)
	}

	for i := 0; i < worker.MaxWorkers; i++ {
		go worker.worker()
	}
	return worker
}

func (w *worker) Work(ctx context.Context) (<-chan *Result, <-chan error) {
	resCh, errCh := make(chan *Result), make(chan error)

	var wg sync.WaitGroup
	wg.Add(len(w.queries))

	go func() {
		wg.Wait()
		close(resCh)
		close(errCh)
	}()

	for _, query := range w.queries {
		go w.addJob(newJob(ctx, &wg, query, resCh, errCh))

	}
	return resCh, errCh
}

func (w *worker) addJob(job *job) {
	w.jobs <- job
}

func (w *worker) worker() {
	for job := range w.jobs {
		job.do()
	}
}

type job struct {
	ctx      context.Context
	wg       *sync.WaitGroup
	query    *Query
	resultCh chan *Result
	errorCh  chan error
}

func newJob(ctx context.Context, wg *sync.WaitGroup, query *Query, resultCh chan *Result, errorCh chan error) *job {
	return &job{
		ctx:      ctx,
		wg:       wg,
		query:    query,
		resultCh: resultCh,
		errorCh:  errorCh,
	}
}

func (job *job) do() {
	defer job.wg.Done()

	res, err := job.query.query(job.ctx)
	if err != nil {
		job.errorCh <- err
		return
	}

	job.resultCh <- res
}
