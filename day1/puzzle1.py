def main ( ) :
    max = 0
    sum = 0
    while True:
        n = input()
        if n.isnumeric():
            sum += int(n) 
        elif n == "e":
            break
        else:
            if sum > max:
                max = sum
            sum = 0 
    print(max)

if __name__ == "__main__" :
    main()
