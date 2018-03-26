import "sort"
func combinationSum2(candidates []int, target int) [][]int {
    sort.Ints(candidates)
    return dfs(candidates,0,target,[]int{},[][]int{})
}
func dfs(candidates []int,start int,target int,solution []int,result [][]int)(rets [][]int){
	rets = result
	if target < 0{
		return 
	}else if target == 0{
		rets = append(rets, solution)
		return
	}
	for i:=start;i<len(candidates);i++{
		if(i>start&&candidates[i]==candidates[i-1]){
			continue
		}
		repliaSolution := make([]int,len(solution))
		copy(repliaSolution, solution)
		repliaSolution = append(repliaSolution,candidates[i])
		rets = dfs(candidates,i+1,target-candidates[i],repliaSolution,rets)
		repliaSolution = repliaSolution[:len(repliaSolution)-1]
	}
}