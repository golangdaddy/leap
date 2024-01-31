
type ASYNCJOB struct {
	Meta Internals
	// pending:started:completed:failed
	Status    string
	Stage     int
	Stages    []*Stage
	Completed int64
}

func (job *ASYNCJOB) SetData(x interface{}) {
	job.Stages[job.Stage].Data = x
}

func (job *ASYNCJOB) Data() interface{} {
	return job.Stages[job.Stage].Data
}

func (job *ASYNCJOB) AddNote(note string) {
	job.Stages[job.Stage].Notes = append(
		job.Stages[job.Stage].Notes,
		note,
	)
}

func (job *ASYNCJOB) CompleteStage() {
	job.Stages[job.Stage].Completed = getTime()
	if job.Stage+1 == len(job.Stages) {
		if job.Completed == 0 {
			job.Completed = getTime()
		}
	} else {
		job.Stage++
	}
	if job.Stage < len(job.Stages) {
		job.Status = job.Stages[job.Stage].Name
	}
}

type Stage struct {
	Name      string
	Data      interface{}
	Notes     []string
	Started   int64
	Completed int64
	Failed    int64
}

func NewASYNCJOB(parent *Internals, stages ...*Stage) *ASYNCJOB {
	return &ASYNCJOB{
		Meta:   parent.NewInternals("asyncjobs"),
		Stages: stages,
	}
}
