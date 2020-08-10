


"""

Quick sort implemetnation. The idea of quick sort is to partition the whole array based on the pivot point. The pivot is selected such that
the left side is less than the large side of the arra.



a = [2, 8, 7,20, 12, 4]

"""


def quick_sort(a, l, h):

    if l < h:
        p  = partition(a, l, h)
        quick_sort(a, l, p -1)
        quick_sort(a, p+1, h)

def partition(a, l, h):

    i = l - 1

    for j in range(l, h):

        if a[j] < a[h] :
            i += 1
            a[i], a[j] = a[j], a[i]
    a[i+1], a[h] = a[h], a[i+1]
    return i + 1


a = [2, 9, 1,10, 3,9]
quick_sort(a, 0, len(a) - 1)


print(a)
