package safari

/* Compare function for types using regular compare operator, for other types or sort 
   orders, supply you own compare function. */

type OpCompareable interface {
	int | int16 | int32 | int64 | float32 | float64  | string
}

func OpCompAsc[N OpCompareable](a, b N) bool {
	return a < b
}

func OpCompDesc[N OpCompareable](a, b N) bool {
	return a > b
}


/* Generic Quicksort */

func Partition[N any](input []N, low, high int, comp func(N,N)bool) (int) {
	pivot := input[high]
	i := low
	for j := low; j < high; j++ {
		if comp(input[j], pivot) {
			input[i], input[j] = input[j], input[i]
			i++
		}
	}
	input[i], input[high] = input[high], input[i]
	return  i
}

func Sort[N any](input []N, low, high int, comp func(N,N)bool ) []N {
	if low < high {
		partition := Partition(input, low, high, comp)
		input = Sort(input, low, partition-1, comp)
		input = Sort(input, partition+1, high, comp)
	}
	return input
}

func SortAll[N any](input []N, comp func(N,N)bool) []N {
	return Sort(input,0,len(input)-1,comp)
}


