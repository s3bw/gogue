/*Package feed uses a queue*/
package feed

// Queue implementation of a queue which queues arrays of runes
type Queue interface {
	Enqueue(obj []rune)
	Dequeue() []rune
	IsEmpty() bool
	Size() int
	Iterate() <-chan []rune
}

// LinkedList implements the queue
type queue struct {
	First  *node
	Last   *node
	Length int
}

type node struct {
	Next  *node
	Value []rune
}

// CreateQueue returns a new queue
func CreateQueue() Queue {
	return &queue{}
}

func (q *queue) Enqueue(value []rune) {
	oldLast := q.Last
	q.Last = &node{}
	q.Last.Value = value

	if q.IsEmpty() {
		q.First = q.Last
	} else {
		oldLast.Next = q.Last
	}
	q.Length++
}

func (q *queue) Dequeue() []rune {
	if !q.IsEmpty() {
		item := q.First.Value
		q.Length--
		q.First = q.First.Next
		if q.Length == 0 {
			q.Last = q.First
		}
		return item
	}
	return []rune("")
}

func (q *queue) IsEmpty() bool {
	return q.Size() == 0
}

func (q *queue) Size() int {
	return q.Length
}

// Iterate goes through the list and passes the next item to the channel
// we have not called this function on the pointer (*) since we want to
// keep the items in the list, instead of 'Dequeuing' them. An alternative
// approach would be to return the 'Next' item instead of pop'ing each
// item
func (q queue) Iterate() <-chan []rune {
	ch := make(chan []rune)
	go func() {
		for {
			if q.IsEmpty() {
				break
			}
			ch <- q.Dequeue()
		}
		close(ch)
	}()
	return ch
}
