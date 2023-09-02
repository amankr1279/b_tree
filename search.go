package main

func search(n *node, val int) *node {
	if n == nil {
		return nil
	}
	numKeys := len(n.keys)
	numChildren := len(n.children)
	if numKeys == 0 {
		return nil
	}
	if n.keys[0] <= val && val <= n.keys[numKeys-1] {
		for i := 0; i < numKeys-1; i++ {
			if n.keys[i] == val {
				return n
			}
			if n.keys[i] < val && n.keys[i+1] > val && numChildren >= (i+1) {
				return search(n.children[i+1], val)
			}
		}
		if n.keys[numKeys-1] == val {
			return n
		}
	} else if n.keys[numKeys-1] < val {
		if numChildren == 1+numKeys {
			return search(n.children[numKeys], val)
		}
	} else {
		if !n.isLeaf {
			return search(n.children[0], val)
		}
	}
	return nil
}