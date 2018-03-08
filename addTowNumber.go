/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    carrier := 0
    prev := &ListNode{}
    result := prev;
    for{
            if l1 == nil && l2!=nil {
                data := l2.Val + carrier
                carrier,prev = getPrev(data,prev)
                l2 = l2.Next
            } else if l1 != nil && l2==nil {
                data := l1.Val + carrier
                carrier,prev = getPrev(data,prev)
                l1 = l1.Next
            } else if l1 == nil && l2 == nil{
                if carrier>0{
                    prev.Next = &ListNode{Val:carrier,Next:nil}
                    carrier = 0
                    prev = prev.Next
                }
    		    break
                    
            }else{
                data := l1.Val + l2.Val + carrier
                carrier,prev = getPrev(data,prev)
                l1 = l1.Next
                l2 = l2.Next
            }
    }
	return result.Next
}

func getPrev(data int,prev *ListNode) (carrier int,prevs *ListNode){
    carrier = data / 10
	val := data % 10
    prev.Next = &ListNode{Val:val,Next:nil}
    prevs = prev.Next
    return
}