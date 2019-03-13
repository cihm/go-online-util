package oop

//加法子類
type Add struct{ BaseNum }

func (a *Add) Opt() int { return a.num1 + a.num2 }
