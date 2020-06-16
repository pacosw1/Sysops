package monitor

//ProStats saves stats for individual processes
type ProStats struct {
	PageFaults int
	StartStep  float32
	EndStep    float32
}

//NewProStats initialize a new Stat
func NewProStats(start float32) *ProStats {

	return &ProStats{
		StartStep: start,
		EndStep:   -1.0,
	}

}
