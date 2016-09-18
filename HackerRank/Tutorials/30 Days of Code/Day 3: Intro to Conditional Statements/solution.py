"""
Given an integer, n, perform the following conditional actions:
    If n is odd, print Weird
    If n is even and in the inclusive range of 2 to 5, print Not Weird
    If n is even and in the inclusive range of 6 to 20, print Weird
    If n is even and greater than 20, print Not Weird
"""

def weird(n):
    if n % 2 != 0:
        print("Weird")
    else:
        if 2 <= n and n <= 5:
            print("Not Weird")
        if 6 <= n and n <= 20:
            print("Weird")
        if 20 < n:
            print("Not Weird")
    return 0;


def main():
    N = int(input())
    weird(N);


if __name__ == '__main__':
    main()
