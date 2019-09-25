package main

import "fmt"

/**
给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。

你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。

--------示列
给定 nums = [2, 7, 11, 15], target = 9

因为 nums[0] + nums[1] = 2 + 7 = 9
所以返回 [0, 1]


 */


func main() {

	nums :=[]int{2, 7, 11, 15}
	target:=9

	b:=addRangeSum(nums,target)

	fmt.Println(b[0])
	fmt.Println(b[1])

}



/**
暴力循环
 */
func addSum(nums []int,target int) []int {
	count:=len(nums)
	var arr []int
	for i:=0;i<count;i++ {
		for j:=i+1;j<count;j++ {
			if (nums[i]+nums[j]==target) {
				arr=[]int{i,j}
				break
			}
		}
	}

	return arr


}

/**
空间换时间
 */
func addRangeSum(nums []int,target int) []int{
	//定一个map集合，然后将值为索引，健为值
	var maps=make(map[int]int)
	var arr []int
	//将值赋给map集合
	for k,i:=range nums{
		maps[i]=k
	}
	//判断是否存在
	for k,i:=range nums{
		curnt:=target-i;
		x1,x2:=maps[curnt]
		if(x2 && x1!=k) {
			arr=[]int{x1,k}
		}
	}
	return arr
}
