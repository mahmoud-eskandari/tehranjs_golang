package main

import "fmt"

type rect struct {
	w int
	h  int
}
func (a rect) getArea() int{
	return a.w*a.h
}
func (a *rect) setW(ww int){
	a.w = ww
}
func main(){
	rect_obj := rect{100,20}
	rect_obj.setW(10)
	fmt.Println(rect_obj.getArea())
}
