from typing import *
def merge_sort(student_marks:List[Tuple[int,str]])->List[Tuple[int,str]]:
    if len(student_marks)<=1:
        return student_marks
    
    mid=len(student_marks)//2
    left_half=merge_sort(student_marks[:mid])
    right_half=merge_sort(student_marks[mid:])
    return merge(left_half,right_half) 

def merge(left: List[Tuple[int,str]], right:List[Tuple[int,str]]):
    result=[]
    i=j=0
    while i < len(left) and j < len(right):
        # Compare by score
        if left[i][0] < right[j][0]:
            result.append(left[i])
            i += 1
        elif left[i][0] > right[j][0]:
            result.append(right[j])
            j += 1
        else:
            if left[i][1] < right[j][1]:
                result.append(left[i])
                i += 1
            else:
                result.append(right[j])
                j += 1
    result.extend(left[i:])
    result.extend(right[j:])
    return result

data: List[Tuple[int, str]] = [
    (92, "Bob"),
    (85, "Charlie"),
    (85, "Alice"),
    (90, "Diana"),
    (88, "Ethan")
]

print(merge_sort(data))

