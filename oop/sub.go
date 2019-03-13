package oop

//減法子類
type Sub struct{ BaseNum }

func (s *Sub) Opt() int { return s.num1 - s.num2 }
