def quicksort(arr):
    if len(arr) <= 1:
        return arr
    pivot = arr[len(arr) // 2]
    left = [x for x in arr if x < pivot]
    middle = [x for x in arr if x == pivot]
    right = [x for x in arr if x > pivot]
    return quicksort(left) + middle + quicksort(right)
    
def main():
    arr = [5, 2, 9, 1, 5, 6]
    sorted_arr = quicksort(arr)
    print(sorted_arr)

if __name__ == '__main__':
    main()
