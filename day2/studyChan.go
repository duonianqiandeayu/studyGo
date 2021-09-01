package main

import (
	"os"
	"bufio"
	"fmt"
)

func sum(s []int,c chan int){
	sum:=0
	for _, n:=range s{
		sum+=n
	}
	c <- sum
	for true{
		n:=<-c
		sum+=n
		c <- sum
	}
}

func main(){
	s:=[]int{ 1,9,5,6,32,0,4}

	c:=make(chan int)
	go sum(s,c)

	for i := range c {
		fmt.Println(i)
	}

}