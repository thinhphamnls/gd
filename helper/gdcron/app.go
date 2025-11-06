package gdcron

type ICronApp interface {
	Start() error
	Stop() error
	Register(ICronJob)
}

type cronApp struct {
	jobs []ICronJob
}

func NewCronApplication() ICronApp {
	return &cronApp{
		jobs: make([]ICronJob, 0),
	}
}

func (a *cronApp) Start() error {
	for _, e := range a.jobs {
		e.Start()
	}
	return nil
}

func (a *cronApp) Stop() error {
	for _, e := range a.jobs {
		e.Stop()
	}
	return nil
}

func (a *cronApp) Register(job ICronJob) {
	a.jobs = append(a.jobs, job)
}
