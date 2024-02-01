type ASYNCJOB struct {
	Meta Internals
	// pending:started:completed:failed
	Status string
	Stage  int
	Stages []ASYNCJOBSTAGE
	Data   interface{}
}

type ASYNCJOBSTAGE struct {
	Name      string
	Notes     []string
	Started   int64
	Failed    int64
	Completed int64
}

func NewASYNCJOB(parent *Internals, stages ...ASYNCJOBSTAGE) *ASYNCJOB {
	return &ASYNCJOB{
		Meta:   parent.NewInternals("asyncjobs"),
		Stages: stages,
		Status: "PENDING",
	}
}

func NewASYNCJOBSTAGE(name string) ASYNCJOBSTAGE {
	return ASYNCJOBSTAGE{
		Name: name,
	}
}

func (job *ASYNCJOB) DataTo(dst interface{}) error {
	b, err := json.Marshal(job.Data)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, dst)
}

func (job *ASYNCJOB) AddNote(note string) {
	job.Stages[job.Stage].Notes = append(
		job.Stages[job.Stage].Notes,
		note,
	)
}

func (job *ASYNCJOB) StartStage() {
	job.Stages[job.Stage].Started = getTime()
	job.Status = "STARTED"
}

func (job *ASYNCJOB) FailStage(err error) {
	job.AddNote(err.Error())
	job.Stages[job.Stage].Failed = getTime()
	job.Status = "FAILED"
}

func (job *ASYNCJOB) CompleteStage() {
	job.Stages[job.Stage].Completed = getTime()
	if job.Stage+1 < len(job.Stages) {
		log.Println("JOB STAGE COMPLETED", job.Stage)
		job.AddNote("COMPLETED STAGE: " + strconv.Itoa(job.Stage))
		job.Stage++
	} else {
		job.Status = "COMPLETED"
		log.Println("JOB STATUS:", job.Status)
		job.AddNote("JOB COMPLETED")
	}
}
