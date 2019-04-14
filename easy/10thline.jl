#https://leetcode.com/problems/tenth-line/
#Technically a Bash challenge but whatever.
result = open("demo.txt") do f
    result = ""
    linecounter = 0
    for l in eachline(f)
        linecounter += 1
        if(linecounter==10)
            result = l
        end
    end
    (result)
end
print(result)
