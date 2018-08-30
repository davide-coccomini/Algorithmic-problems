# Given a tree of integers and return the number of non-empty unival trees
# A unival tree is a tree which contain all nodes with the same number.

# DUMB SOLUTION

# O(n)
def isUnival(root):
    if root == null:
        return True
    if root.left != null and root.left.value != root.value:
        return False
    if root.right != null and root.right.value != root.value:  
        return False
    if isUnival(root.left) and isUnival(root.right):
        return True
    return False

# O(n^2)
def countUnivals(root):
    if root == null:
        return 0

    count = countUnivals(root.left) + countUnivals(root.right)
    
    if isUnival(root):
        count += 1
    return count
    


# GOOD SOLUTION O(n)
def countUnivals(root):
    count, is_unival = explore(root)
    return count

def explore(root):
    if root == null: return (0, True)
    left_count, is_left_unival = explore(root.left)
    right_count, is_right_unival = explore(root.right)
    
    is_unival = True
    if not is_left_unival or not is_right_unival:
        is_unival = False

    if root.left != null and root.left.value != root.value:
        is_unival = False

    if root.right != null and root.right.value != root.value:
        is_unival = False

    if is_unival:
        return (left_count + right_count +1, True)
    else:
        return (left_count + right_count, False)



# Asked by Google