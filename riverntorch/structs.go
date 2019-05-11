package riverntorch

import (
	"bytes"
	"fmt"
	"sort"
	"strings"
)

type Person struct {
	//specifies the name of the person
	Name string `yaml:"name"`

	//specifies the time it takes for the person to cross.
	Duration int `yaml:"minutes"`
}

//
// people that needs to cross river.
//
type RiverCrossers []*Person

//
// sort the crossers as per the ascending time they take to cross
//
func (this RiverCrossers) Sort() RiverCrossers {
	sort.Sort(this)
	return this
}

//
// Override the default String() method. for custom printing.
//
func (this RiverCrossers) String() string {
	str := make([]string, len(this))
	for _, v := range this {
		str = append(str, fmt.Sprintf("%s %d", v.Name, v.Duration))
	}
	return strings.Join(str, "\n")
}

//
// Sorting methods.
//
func (this RiverCrossers) Len() int {
	return len(this)
}
func (this RiverCrossers) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}
func (this RiverCrossers) Less(i, j int) bool {
	return this[i].Duration < this[j].Duration
}

//
// PeopleOnBridge stores the index of people crossing bridge.
//
type PeopleOnBridge struct {
	fastOne, slowOne int
}

//
// Solution stores the meta and details about the approach used.
//
type Solution struct {
	//name of approach used to figure out solution.
	approachName string
	// time it took for crossers to cross.
	duration int
	// procedure/ steps followed for crossing.
	set [][]Person
}

// Initialize Solution construct.
func (this *Solution) Initialize() {
	this.set = make([][]Person, 0)
}

//
// Override the default String() method. for custom printing.
//
func (this Solution) String() string {
	buf := bytes.NewBuffer(nil)
	var frwd, bck int
	for _, v := range this.set {
		if len(v) == 2 {
			frwd++
			buf.WriteString(fmt.Sprintf("side A -> sideB : %s + %s\n", v[1].Name, v[0].Name))
		}
		if len(v) == 1 {
			bck++
			buf.WriteString(fmt.Sprintf("side A <- sideB : %s\n", v[0].Name))
		}
	}
	if len(this.set) > 1 {
		buf.WriteString(fmt.Sprintf("\nForward Steps: %d, Backward Steps: %d\n", frwd, bck))
	}
	buf.WriteString(fmt.Sprintf("Using %s\n\n", this.approachName))
	buf.WriteString("########################################\n")
	buf.WriteString(fmt.Sprintf("Total time Taken: %d minutes\n", this.duration))
	buf.WriteString("########################################\n")
	return buf.String()
}
