package main

import "fmt"

func main() {
   const LENGTH int = 10
   const WIDTH int = 5  
   var area int
   const a, b, c = 1, false, "str" //多重赋值

   area = LENGTH * WIDTH
   fmt.Printf("面积为 : %d", area)
   println(a, b, c)  
   println()
   testFunc()
}
func sumInt(){

	sum:=0
	for i:=1;i<=10;i++{
		sum+=i
	}
	fmt.Println(sum)

}

func testFunc(){
	strings := []string{"google", "runoob"}
	for i, s := range strings {
			fmt.Println(i, s)
	}

	numbers := [6]int{1, 2, 3, 5}
	for i,x:= range numbers {
			fmt.Printf("第 %d 位 x 的值 = %d\n", i,x)
	}  
	
}