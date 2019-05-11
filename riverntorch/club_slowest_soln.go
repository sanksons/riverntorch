package riverntorch

import (
	"math"
)

func NewClubSlowestApproach(crossers RiverCrossers) *clubSlowestApproach {
	c := &clubSlowestApproach{
		NewBaseApproach(crossers),
	}
	c.solution.approachName = c.GetName()
	return c
}

//
// clubSlowestApproach uses the idea of clubbing the Slowest members together,
// so that time of forward journey can be reduced.
// However, this incurs an extra cost on backward journey. So, instead of always
// using the fastest rider to return the torch, it uses the combination of fastest and
// second fastest riders.
//
type clubSlowestApproach struct {
	*baseApproach
}

// GetName implementation of GetName().
func (this *clubSlowestApproach) GetName() string {
	return "clubSlowestApproach"
}

// Cross implementationn of Cross().
func (this *clubSlowestApproach) Cross() Solution {
	//check for base case.
	if this.isBaseCase() {
		return *this.solution
	}

	maxSteps := this.fwdSteps + this.backSteps
	var stepsTaken int
	var needSlowest bool
	for {
		var chosen PeopleOnBridge
		//check if we need to choose fastest ones or slowest ones.
		if needSlowest {
			chosen = this.pick2Slowest()
			needSlowest = false
		} else {
			chosen = this.pick2Fastest()
			needSlowest = true
		}
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
		//return back torch with the help of fastest bearer on side B.
		this.returnTorch()
		stepsTaken += 1
	}
	return *this.solution
}

//
// pick the fastest and second fastest person from SideA.
//
func (this *clubSlowestApproach) pick2Fastest() PeopleOnBridge {
	return PeopleOnBridge{0, 1}
}

//
// pick the slowest and second slowest person from SideA.
//
func (this *clubSlowestApproach) pick2Slowest() PeopleOnBridge {

	var fastOne, slowOne int = -1, -1
	for i := (this.noOfcrossers - 1); i >= 0; i-- {
		if this.sideA[i] == nil {
			continue
		}
		if slowOne == -1 {
			slowOne = i
			continue
		}
		if fastOne == -1 {
			fastOne = i
			continue
		}
		break
	}
	return PeopleOnBridge{fastOne, slowOne}
}

//
// transfer2SideB transafers chosen people to sideB.
//
func (this *clubSlowestApproach) transfer2SideB(p PeopleOnBridge) {

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
// returnTorch using the fastest person on sideB.
//
func (this *clubSlowestApproach) returnTorch() {

	var torchBearerIndex int
	if this.sideB[0] == nil {
		torchBearerIndex = 1
	}
	//pick the fastest person to bear torch.
	this.sideA[torchBearerIndex] = this.sideB[torchBearerIndex]

	//clearoff bearer from sideB.
	this.sideB[torchBearerIndex] = nil

	//add the cost of travelling.
	this.solution.duration += this.sideA[torchBearerIndex].Duration
	this.solution.set = append(this.solution.set, []Person{*this.sideA[torchBearerIndex]})
}

//
// timeSpentInBckwrdJrny calculates the total time spent in backward journey.
//
func (this *clubSlowestApproach) timeSpentInBckwrdJrny() int {

	stepsBySecFastest := int(math.Floor(float64(this.backSteps / 2)))

	stepsByFastest := this.backSteps - stepsBySecFastest

	return (stepsByFastest * this.sideA[0].Duration) + (stepsBySecFastest * this.sideA[1].Duration)
}

//
// timeSavedByClubbing calculates the time saved by clubbing slowest members.
// = (sum of saved slower members) - (count of skipped members * duration of second fastest)
//
func (this *clubSlowestApproach) timeSavedByClubbing() int {

	var sumOfSavedMembers, countOfSavedMembers int
	for i := this.noOfcrossers - 2; i >= 2; i -= 2 {
		sumOfSavedMembers += this.sideA[i].Duration
		countOfSavedMembers += 1
	}
	timeSaved := sumOfSavedMembers - (countOfSavedMembers * this.sideA[1].Duration)
	return timeSaved
}
