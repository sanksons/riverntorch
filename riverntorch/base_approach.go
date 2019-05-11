package riverntorch

//
// NewbaseApproach Initializes and returns a new object of baseApproach.
//
func NewBaseApproach(crossers RiverCrossers) *baseApproach {
	a := new(baseApproach)
	a.noOfcrossers = len(crossers)
	a.sideA = crossers

	newc := make([]*Person, len(crossers), len(crossers))
	a.sideB = RiverCrossers(newc)

	a.solution = &Solution{approachName: "Base Approach"}
	a.solution.Initialize()

	///no. of steps required for approach
	a.fwdSteps = a.noOfcrossers - 1
	a.backSteps = a.noOfcrossers - 2
	return a
}

//
// baseApproach datatype contains the configs and functionalities that
// are common to ClubSlowestApproach and FastestBearerApproach.
//
type baseApproach struct {
	//no. of people to cross.
	noOfcrossers int

	//count of forward and backward steps involved.
	fwdSteps  int
	backSteps int

	//meta and details of the solution
	solution *Solution

	//details of people on A and B.
	sideA RiverCrossers
	sideB RiverCrossers
}

//
// isBaseCase check if the input is one of the base cases.
// if so prepares the solution accordingly and notifies.
//
func (this *baseApproach) isBaseCase() bool {
	if this.noOfcrossers == 1 {
		this.solution.duration = this.sideA[0].Duration
		this.solution.set = append(this.solution.set, []Person{*this.sideA[0]})
		return true
	}
	if this.noOfcrossers == 2 {
		this.solution.duration = this.sideA[1].Duration
		this.solution.set = append(this.solution.set, []Person{*this.sideA[0], *this.sideA[1]})
		return true
	}
	return false
}
