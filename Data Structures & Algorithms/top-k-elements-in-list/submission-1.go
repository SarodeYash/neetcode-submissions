
func topKFrequent(nums []int, k int) []int {
type mapdata struct{
	key int 
	value int
} 
count:=make(map[int]int)
data:= []mapdata{}

for _,v :=range nums{
count[v]++
}
for k,v:=range count{
	data=append(data,mapdata{
		key:k,
		value:v,
	})
}

sort.Slice(data,func(i ,j int )bool {
	return data[i].value>data[j].value
 })

 result:=[]int{}

 for i:=0;i<k;i++{
	result =append(result, data[i].key)	
 }
 return result

}