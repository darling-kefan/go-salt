package salt

type JobResponse struct {
	Info   []Job  `json:"info"`
	Return []Result `json:"return"`
}

type JobsResponse struct {
	Jobs []map[string]Job `json:"return"`
}

type ResultResponse struct {
	Return []Result `json:"return"`
}

type ExecutionResponse struct {
	Job []Job `json:"return"`
}

type Job struct {
	ID         string            `json:"jid"`
	Function   string            `json:"Function"`
	Target     string            `json:"Target"`
	User       string            `json:"User"`
	StartTime  string            `json:"StartTime"`
	TargetType string            `json:"Target-Type"`
	Arguments  []string          `json:"Arguments"`
	Minions    []string          `json:"Minions"`
	Result     Result            `json:"-"`
}

type Result map[string]interface{}

func (j *Job) Running() bool {
	if len(j.Minions) != len(j.Result) {
		return false
	}
	return true
}

func (j *Job) Successful() bool {
	return true
}
