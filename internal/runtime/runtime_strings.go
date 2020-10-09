// Code generated by "stringer -linecomment -type=runEvent -output=runtime_strings.go"; DO NOT EDIT.

package runtime

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[startRunEvent-0]
	_ = x[endRunEvent-1]
	_ = x[nextStepEvent-2]
}

const _runEvent_name = "Start RunEnd RunNext Run Step"

var _runEvent_index = [...]uint8{0, 9, 16, 29}

func (i runEvent) String() string {
	if i < 0 || i >= runEvent(len(_runEvent_index)-1) {
		return "runEvent(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _runEvent_name[_runEvent_index[i]:_runEvent_index[i+1]]
}