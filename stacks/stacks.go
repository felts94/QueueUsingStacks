package stacks

type Stack []interface{}

func (s *Stack) Push(elem interface{}) {
	temp := &Stack{elem}
	*s = append(*temp, *s...)
}

func (s *Stack) Pop() interface{} {
	if len(*s) > 0 {
		elem := (*s)[0]
		*s = (*s)[1:]
		return elem
	}
	return nil
}

func (s *Stack) ToString() string {
	sString := "STACK TOP"
	sCopy := &Stack{}
	*sCopy = *s

	elem := sCopy.Pop()

	for elem != nil {
		sString = sString + "\n" + elem.(string)
		elem = sCopy.Pop()
		//fmt.Println(elem)
	}

	return sString + "\nSTACK BOTTOM"
}
