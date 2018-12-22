
package main

import(
	"fmt"
	"time"
	)

func increment(pos int, k []int, size int)(int){
	if(pos+1 == size) {
		pos = 0;
	} else {
		pos++
	}
	for (k[pos] == 0) {
		pos = increment(pos,k,size)
	}
	return pos

}
func foo(c chan int, size int) {
	var army = make([]int, size)
	for i := 0; i < size; i++ {
		army[i] = 1
	}
	var left int = 1
	var pos int = 0

	for left!=pos {
		army[left] = 0
		pos = increment(pos,army,size)
		left = pos
		left = increment(left,army,size)
	}
	fmt.Print("Manual computation finished")
	c<-pos+1
}
func divisible(size int) (bool) {
	i:=0
	for(i!=1) {
		if(size==1) {
			return true
		} else{
			i = size%2;
			size= size / 2;
		}
	}
	return false
}

func foo2(c chan int, size int) {
	count := 0
	if(size==1) {
		c<-1
	}
	for(!(divisible(size))) {
		size--
		count++
	}
	fmt.Print("Finding closest power of two function finished")
	c<-(count*2)+1
}

func main() {
	c := make(chan int)
	//number of soldiers can be changed. be warned, going too high (around 8385000) causes memory overflow,
	//going too low can lead to inconsistencies in response time measurement.
	size := 41
	fmt.Println("\n\t\tNumber of soldiers: ", size, "\n-------------------------------------------------------------------------------------------")
	start := time.Now()
	go foo(c, size)
	go foo2(c, size)
	x := <-c
	elapsed1 := time.Since(start).Seconds()
	fmt.Print(": winning position ", x, "\ntook ", elapsed1, " seconds\n")
	k := <-c
	elapsed2 := time.Since(start).Seconds()
	fmt.Print(": winning position ", k, "\ntook ", elapsed2, " seconds\n")
	fmt.Println("\t\tAlgorithm 1 was ", elapsed2/elapsed1, " times faster")
	fmt.Println("-------------------------------------------------------------------------------------------")
}
