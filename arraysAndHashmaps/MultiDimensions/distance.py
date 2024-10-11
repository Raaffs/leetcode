import math

ls = [[1, 3], [-2, 2], [4, 5], [10, -1]] 
n = 2  
result = []

for _ in range(n):
    if not ls:
        break 
    min_val = math.inf
    curr_smallest = None
    
    for l in ls:
        distance = math.sqrt(l[0]**2 + l[1]**2)
        if distance < min_val:
            min_val = distance
            curr_smallest = l
    
    if curr_smallest:
        ls.remove(curr_smallest)
        result.append(curr_smallest)

print("result:", result)
