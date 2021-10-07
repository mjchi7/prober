package run

type RunConfig struct {
	ConcurrentWorkers   int              `json:"concurrentWorkers"`
	Duration            string           `json:"duration"`
	ConnectionTimeoutMs int              `json:"connectionTimeoutMs"`
	ResultConfig        RunResultConfig  `json:"result"`
	RequestConfig       RunRequestConfig `json:"request"`
}

type RunResultConfig struct {
	SlidingWindowDuration string `json:"slidingWindowDuration"`
	SamplingInterval      string `json:"samplingInterval"`
}

type RunRequestConfig struct {
	Url     string            `json:"url"`
	Method  string            `json:"method"`
	Headers map[string]string `json:"headers"`
	Body    string            `json:"body"`
}
