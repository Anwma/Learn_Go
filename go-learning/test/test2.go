package main
/*
#include <stdio.h>
// 声明Go中的函数
extern void GoFunction();
void CFunction() {
printf("CFunction!\n");
GoFunction();
}
*/
import "C"
func main() {
	C.CFunction()
}