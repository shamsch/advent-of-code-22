def main ( ) :
    max_val = []
    sum = 0
    while True:
        n = input()
        if n.isnumeric():
            sum += int(n) 
        elif n == "e":
            break
        else:
            # when max has less than 3 elements, append sum
            if len(max_val)<3:
                max_val.append(sum)
            # when max has 3 elements, check if sum is greater than any of them
            else:
                if sum > min(max_val):
                    max_val.remove(min(max_val))
                    max_val.append(sum)
            sum = 0 
    
    total = 0
    for i in max_val:
        total += i
    print(total)
   

if __name__ == "__main__" :
    main()