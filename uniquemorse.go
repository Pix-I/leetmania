// https://leetcode.com/problems/unique-morse-code-words/submissions/
func uniqueMorseRepresentations(words []string) int {
    var m = make(map[string]bool)
    morse := []string{".-","-...","-.-.","-..",".","..-.","--.","....","..",".---","-.-",".-..","--","-.","---",".--.","--.-",".-.","...","-","..-","...-",".--","-..-","-.--","--.."}
    for i:=0;i<len(words);i++{
        asciiBytes := []byte(words[i])
        str := ""
        for j:=0;j<len(asciiBytes);j++{
            str = str + morse[asciiBytes[j]-97]
        }
        m[str] = true
    }
    return len(m)
}