package main

func deleteNode(n **node, key int) {
	if *n == nil {
		return
	}

	// Find the index of the key to delete
	index := findIndex((*n).keys, key)

	if index < len((*n).keys) && (*n).keys[index] == key {
		// Key is found in this node
		if (*n).isLeaf {
			// Case 1: If the key is in a leaf node, simply remove it
			deleteFromLeaf(*n, index)
		} else {
			// Case 2: If the key is in an internal node
			if len((*n).children[index].keys) >= T {
				// Case 2a: If the child node has at least T keys, find the predecessor and replace the key
				predecessor := getMax(&(*n).children[index])
				(*n).keys[index] = predecessor
				deleteNode(&(*n).children[index], predecessor)
			} else if len((*n).children[index+1].keys) >= T {
				// Case 2b: If the child node has fewer than T keys but its right sibling has at least T keys,
				// find the successor and replace the key
				successor := getMin(&(*n).children[index+1])
				(*n).keys[index] = successor
				deleteNode(&(*n).children[index+1], successor)
			} else {
				// Case 2c: If both the child and its right sibling have less than T keys, merge the child and its sibling
				mergeChildren(n, index)
				deleteNode(&(*n).children[index], key)
			}
		}
	} else {
		// The key is not in this node, find the child where it could be
		if len((*n).children[index].keys) < T {
			// Case 3: If the child has less than T keys, balance it
			balanceChildren(n, index)
		}

		deleteNode(&(*n).children[index], key)
	}
}

func deleteFromLeaf(n *node, index int) {
	// Remove the key at the specified index from the leaf node
	copy(n.keys[index:], n.keys[index+1:])
	n.keys = n.keys[:len(n.keys)-1]
}

func getMax(n **node) int {
	// Find and return the maximum key in the subtree rooted at node n
	for !(*n).isLeaf {
		*n = (*n).children[len((*n).keys)]
	}
	return (*n).keys[len((*n).keys)-1]
}

func getMin(n **node) int {
	// Find and return the minimum key in the subtree rooted at node n
	for !(*n).isLeaf {
		*n = (*n).children[0]
	}
	return (*n).keys[0]
}

func mergeChildren(n **node, index int) {
	// Merge the child at index with its right sibling
	child := (*n).children[index]
	rightSibling := (*n).children[index+1]

	// Move the key from the current node to the child node
	child.keys = append(child.keys, (*n).keys[index])

	// Move keys from the right sibling to the child
	child.keys = append(child.keys, rightSibling.keys...)

	// Remove the key and the right sibling
	copy((*n).keys[index:], (*n).keys[index+1:])
	(*n).keys = (*n).keys[:len((*n).keys)-1]

	// If the current node is not the root and becomes empty, remove it and adjust the parent
	if len((*n).keys) == 0 && (*n) != (*n).children[0] {
		copy((*n).children[index:], (*n).children[index+1:])
		(*n).children = (*n).children[:len((*n).children)-1]
	}
}

func balanceChildren(n **node, index int) {
	// Balance the child at index by borrowing a key from its left or right sibling
	if index > 0 && len((*n).children[index-1].keys) >= T {
		// Borrow from the left sibling
		child := (*n).children[index]
		leftSibling := (*n).children[index-1]

		// Move the largest key from the left sibling to the current node
		borrowedKey := leftSibling.keys[len(leftSibling.keys)-1]
		leftSibling.keys = leftSibling.keys[:len(leftSibling.keys)-1]
		(*n).keys = append((*n).keys[:index-1], borrowedKey)
		// Adjust the child pointer
		child.children = append([]*node{child.children[0]}, child.children...)
		// Adjust the left sibling pointer
		leftSibling.children = leftSibling.children[:len(leftSibling.children)-1]
	} else if index < len((*n).children)-1 && len((*n).children[index+1].keys) >= T {
		// Borrow from the right sibling
		child := (*n).children[index]
		rightSibling := (*n).children[index+1]

		// Move the smallest key from the right sibling to the current node
		borrowedKey := rightSibling.keys[0]
		rightSibling.keys = rightSibling.keys[1:]
		(*n).keys = append((*n).keys[:index], borrowedKey)
		// Adjust the child pointer
		child.children = append(child.children, child.children[len(child.children)-1])
		// Adjust the right sibling pointer
		rightSibling.children = rightSibling.children[1:]
	} else {
		// Merge the child with either its left or right sibling
		if index > 0 {
			// Merge with the left sibling
			mergeChildren(n, index-1)
		} else {
			// Merge with the right sibling
			mergeChildren(n, index)
		}
	}

	// Recursively balance the child node only if it has fewer than T-1 keys
	if len((*n).children[index].keys) < T-1 {
		balanceChildren(&(*n).children[index], index)
	}
}

func findIndex(keys []int, key int) int {
	// Find the index of the key in the given slice of keys using binary search
	left, right := 0, len(keys)-1
	for left <= right {
		mid := left + (right-left)/2
		if keys[mid] == key {
			return mid
		} else if keys[mid] < key {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}
