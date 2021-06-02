package linkedlist

func ReverseList(header *node) (reverseHeader *node) {
	if header == nil {
		return nil
	}
	var (
		pre  *node
		curr = header
	)
	for curr != nil {
		next := curr.next
		if next == nil {
			reverseHeader = curr
		}
		curr.next = pre
		pre = curr
		curr = next
	}
	return reverseHeader
}
