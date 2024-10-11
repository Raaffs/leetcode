def heapify(arr:list[int],size:int,i:int)->None:
    largest=i
    left=2*i+1
    right=2*i+2
    if left<size and arr[largest]<arr[left]:
        largest=left
    if right<size and arr[largest]<arr[right]:
        largest=right
    if largest!=i:
        arr[i], arr[largest] = arr[largest], arr[i] 
        heapify(arr,size,largest)

def sort(arr:list[int])->None:
    size=len(arr)
    for i in range(size//2-1,-1,-1):
        heapify(arr,size,i)
    for i in range(size-1,0,-1):
        arr[0],arr[i]=arr[i],arr[0]
        heapify(arr,i,0)

l=[2,4,1,5,20,10,1.3,100,2,21,232,12]
sort(l)
print(l)  