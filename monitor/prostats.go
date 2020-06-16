package monitor

type ProStats struct {
	PageFaults int
	StartStep  float32
	EndStep    float32
}

func NewProStats(start float32) *ProStats {

	return &ProStats{
		StartStep: start,
		EndStep:   -1.0,
	}

}
