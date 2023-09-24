package arrow

type Arrows []*Arrow

func (arrows Arrows) ConnectBranches() {
	for i := len(arrows) - 1; i >= 0; i-- {
		arrow := arrows[i]
		if srcArrId := arrow.Source.BranchArrowId; srcArrId != 0 {
			srcArrow := arrows[srcArrId-1]
			arrow.SourceArrow = srcArrow
		}
	}
}
