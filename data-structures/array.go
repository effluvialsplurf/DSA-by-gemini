/*  an array is a contiguous memory allocation of a fixed size
a dynamic array is a contiguous memory allocation of a variable size
arrays allow for
	random access to elements
	constant time access to elements
	constant time insertion and deletion of elements at the end
	linear time insertion and deletion of elements at the beginning and in the middle
*/

package data_structures

// dynamic array, variable size and data type
type Array struct {
	capacity int
	size     int
	elements []interface{}
}

func InitializeArray(initialCapacity int) *Array {
	return &Array{
		capacity: initialCapacity,
		size:     0,
		elements: make([]interface{}, initialCapacity),
	}
}

// now we need functions to add, remove, and get elements
func (a *Array) Append(element interface{}) {
	// if the array is empty, we need to initialize it
	if a.size == 0 {
		a.elements[0] = element
		a.capacity = 1
		a.size = 1
		return
	}
	// if the array is full, we need to resize it
	if a.size == a.capacity {
		a.capacity *= 2
		a.elements = append(a.elements, element)
		a.size++
		return
	}
	// if the array is not full, we can add the element to the end
	a.elements[a.size] = element
	a.size++
}
