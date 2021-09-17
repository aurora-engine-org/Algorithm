package graph

type Element struct {
	value int
	next  *Element
}

type Queue struct {
	 q    *Element
	 end  *Element
}

func (qe *Queue) EnQueue(value int)  {
	if qe.q==nil {
		e:=&Element{value: value,next: nil}
		qe.end=e
		qe.q=e
		return
	}
	e:=&Element{value: value,next: nil}
	qe.end.next=e
	qe.end=e
}

func (qe *Queue) DeQueue() *Element {
	if qe.q==nil{
		return nil
	}
	 if qe.q.next!=nil{
	 	t:=qe.q
	 	qe.q=qe.q.next
	 	return t
	 }else {
	 	t:=qe.q
	 	qe.q=nil
	 	return t
	 }
}