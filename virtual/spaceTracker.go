package virtual

import "container/list"

//SpaceTracker tracks and updates free frames in memory to speed up insertion and removal
type SpaceTracker struct {
	FreeFrames *list.List            //list with frame indexes that are free
	Location   map[int]*list.Element //hash with the position of the frames inside the list
}

//NewSpaceTracker creates new tracker
func NewSpaceTracker() *SpaceTracker {
	tracker := &SpaceTracker{
		FreeFrames: list.New(),
		Location:   make(map[int]*list.Element, 0),
	}
	return tracker
}

//Size returns amount of free spaces
func (t *SpaceTracker) Size() int {
	return len(t.Location)
}

//Remove Removes a specific frame from the list
func (t *SpaceTracker) Remove(frameIndex int) {

	node := t.Location[frameIndex]

	if node == nil {
		return
	}
	t.FreeFrames.Remove(node)
	delete(t.Location, frameIndex)
}

//Empty returns true if data structure empty
func (t *SpaceTracker) Empty() bool {
	return t.FreeFrames.Len() == 0
}

//Add adds a new frame location to the list
func (t *SpaceTracker) Add(frameIndex int) {
	t.FreeFrames.PushBack(frameIndex)
	t.Location[frameIndex] = t.FreeFrames.Back()
}

//Peek returns the next free avaliable frame index in memory
func (t *SpaceTracker) Peek() int {
	if t.Empty() {
		return -1
	}
	//type assertion
	nextIndex, _ := t.FreeFrames.Front().Value.(int)
	//return next index in list
	return nextIndex
}

//Pop returns and removes the next free avaliable frame index in memory
func (t *SpaceTracker) Pop() int {

	//return -1 if empty
	if t.Empty() {
		return -1
	}
	//type assertion
	nextIndex, _ := t.FreeFrames.Front().Value.(int)

	//remove index from data structure (hash and list)
	delete(t.Location, nextIndex)
	t.FreeFrames.Remove(t.FreeFrames.Front())
	return nextIndex
}
