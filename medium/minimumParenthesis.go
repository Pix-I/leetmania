func minAddToMakeValid(S string) int {
    if S == ""{
        return 0
    }
    var arr = make([]byte,0,1000)
    n := len(arr) - 1 // Top element
    arr = append(arr,S[0])
    for i:=1;i<len(S);i++{
        n = len(arr) -1
        if n>=0 && arr[n]=='('&& S[i]==')'{
            arr = arr[:n] // Pop
        } else{
            arr = append(arr,S[i])
        }
    }
    return len(arr)
}
