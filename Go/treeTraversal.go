
  /* A binary tree Node has data, pointer to left child 
   and a pointer to right child */
package main
type Node struct
{ 
    data int; 
    left *Node; 
    right *Node; 
}; 
  
/* Iterative function for inorder tree 
   traversal */
func inOrder(root *Node){
    s := []*Node{}
    curr := root
  
    for  curr != nil || len(s) == 0 { 
        /* Reach the left most Node of the 
           curr Node */
        for curr !=  nil { 
            /* place pointer to a tree node on 
               the stack before traversing 
              the node's left subtree */
            s = append(s,curr)
            curr = curr.left; 
        } 
  
        /* Current must be NULL at this point */
        curr = s[len(s)-1];
		s[len(s)-1] = nil   // Erase last element (write zero value)
		s = s[:len(s)-1]   // Truncate slice
  
        /* we have visited the node and its 
           left subtree.  Now, it's right 
           subtree's turn */
        curr = curr.right; 
  
    } /* end of while */
} 
  