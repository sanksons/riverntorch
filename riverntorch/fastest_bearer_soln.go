package riverntorch

//
// NewFastestBearerApproach returns the object of fastestBearerApproach
//
func NewFastestBearerApproach(crossers RiverCrossers) *fastestBearerApproach {
	f := &fastestBearerApproach{
		NewBaseApproach(crossers),
	}
	f.solution.approachName = f.GetName()
	return f
}

//
// fastestBearerApproach uses the idea of fastest person to always act as
// the torchbearer. This minimizes the torch returning time.
//
type fastestBearerApproach struct {
	*baseApproach
}

// GetName implementation of GetName().
func (this *fastestBearerApproach) GetName() string {
	return "fastestBearerApproach"
}

// Cross implementationn of Cross().
func (this *fastestBearerApproach) Cross() Solution {
	//check if its a base case.
	if this.isBaseCase() {
		return *this.solution
	}

	maxSteps := this.fwdSteps + this.backSteps
	var stepsTaken int
	for {
		//pick a fastest and a slowest person that exists on sideA.
		chosen := this.pick()
		if chosen.slowOne == -1 {
			//all have moved to sideB. STop here
			break
		}
		//transfer chosen people to sideB
		this.transfer2SideB(chosen)
		stepsTaken += 1

		if stepsTaken >= maxSteps {
			break
		}
		//return back torch with the help of fastest bearer.
		this.returnTorch()
		stepsTaken += 1
	}
	return *this.solution
}

//
// pick the fastest person and a slowest person from side A.
// to be moved to side B.
//
func (this *fastestBearerApproach) pick() PeopleOnBridge {

	var slowest int = -1
	for i := (this.noOfcrossers - 1); i > 0; i-- {
		if this.sideA[i] == nil {
			continue
		}
		slowest = i
		break
	}
	return PeopleOnBridge{0, slowest}
}

//
// transfer2SideB transfers chosen people to SideB.
//
func (this *fastestBearerApproach) transfer2SideB(p PeopleOnBridge) {

	a, b := p.fastOne, p.slowOne

	//move to sideB
	this.sideB[a] = this.sideA[a]
	this.sideB[b] = this.sideA[b]

	//clearOff from sideA.
	this.sideA[a], this.sideA[b] = nil, nil

	//add the cost of travelling.

	this.solution.duration += this.sideB[b].Duration
	this.solution.set = append(this.solution.set, []Person{*this.sideB[b], *this.sideB[a]})
}

//
// returnTorch always uses the fastest person to return torch.
//
func (this *fastestBearerApproach) returnTorch() {

	//pick the fastest person to bear torch.
	this.sideA[0] = this.sideB[0]

	//clearoff bearer from sideB.
	this.sideB[0] = nil

	//add the cost of travelling.
	this.solution.duration += this.sideA[0].Duration
	this.solution.set = append(this.solution.set, []Person{*this.sideA[0]})
}

// timeSpentInBckwrdJrny calculates the total time spent in return journey.
func (this *fastestBearerApproach) timeSpentInBckwrdJrny() int {
	return this.backSteps * this.sideA[0].Duration
}
