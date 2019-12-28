package flyjapan

var (
	defaultWorkerOption = workerOptions{
		MaxWorkers: 1,
	}
)

type WorkerOption func(o *workerOptions)

type workerOptions struct {
	MaxWorkers int
}

func MaxWorkers(maxWorkers int) WorkerOption {
	return func(o *workerOptions) {
		o.MaxWorkers = maxWorkers
	}
}
