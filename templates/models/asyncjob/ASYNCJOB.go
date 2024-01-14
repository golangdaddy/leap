package main

type ASYNCJOB struct {
	Meta   Internals
	Fields FieldsASYNCJOB
}

type FieldsASYNCJOB struct {
	// pending:started:completed:failed
	Status   string `json:"status" firestore:"status"`
	Stage    string `json:"stage" firestore:"stage"`
	Prepare  Stage
	Generate Stage
}

type Stage struct {
	Data      interface{}
	Notes     []string
	Completed bool
}

func NewASYNCJOB(parent *Internals, fields FieldsASYNCJOB) *ASYNCJOB {
	return &ASYNCJOB{
		Meta:   parent.NewInternals("asyncjobs"),
		Fields: fields,
	}
}
