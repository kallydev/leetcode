package solution

type DataStream struct {
	targetValue   int
	targetLength  int
	currentLength int
}

func Constructor(value int, k int) DataStream {
	return DataStream{
		targetValue:  value,
		targetLength: k,
	}
}

func (d *DataStream) Consec(num int) bool {
	if num == d.targetValue {
		d.currentLength += 1
	} else {
		d.currentLength = 0
	}

	return d.currentLength >= d.targetLength
}
