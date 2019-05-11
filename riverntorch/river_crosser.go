package riverntorch

var _ RiverCrosser = (*fastestBearerApproach)(nil)
var _ RiverCrosser = (*clubSlowestApproach)(nil)

type RiverCrosser interface {
	//get name of the approach been used
	GetName() string

	//cross all people from SideA to SideB.
	//gives back solution object, with details about the process.
	Cross() Solution
}

//
// GetRiverCrosser used to get instance of RiverCrosser.
//
func GetRiverCrosser(people RiverCrossers) RiverCrosser {

	fbApproach := NewFastestBearerApproach(people)
	csApproach := NewClubSlowestApproach(people)

	// No. of passers are 3 or less, both approach works same.
	// So, lets use fbApproach for such a scenario.
	if fbApproach.noOfcrossers <= 3 {
		return fbApproach
	}

	// Test which approach to use, beforing applying.
	// Use ClubSlowestApproach, if the time saved by clubbing is greater than
	// time lost by not using fastest torchbearer always.
	// Else, use the FastestBearerApproach

	//extra time spent by ClubSlowestApproach during Backward Journey.
	extraTime := csApproach.timeSpentInBckwrdJrny() - fbApproach.timeSpentInBckwrdJrny()
	//time saved by clubbing slower members together.
	timeSaved := csApproach.timeSavedByClubbing()

	if timeSaved > extraTime {
		return csApproach
	}
	return fbApproach
}
