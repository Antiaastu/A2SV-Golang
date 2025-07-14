package main
import (
	"bufio"
	"fmt"
	"os"
	"unicode"
	"strings"
)
func WordFrequency(text string) map[string]int{
	clean := ""
	for _, ch := range text{
		if unicode.IsLetter(ch) || unicode.IsSpace(ch){
			clean += string(unicode.ToLower(ch))
		}
	}
	freq := make(map[string]int)
	temp := ""
	for _, ch := range clean{
		if unicode.IsSpace(ch){
			if temp != ""{
				freq[temp]++
				temp = ""
			}
		} else{
			temp += string(ch)
		}
	}
	if temp != ""{
		freq[temp]++
	}
	return freq
}
func IsPalindrome(word string) bool{
	clean := ""
	for _, ch := range word{
		if unicode.IsLetter(ch) || unicode.IsDigit(ch) {
			clean += string(unicode.ToLower(ch))
		}
	}
	n := len(clean)
	for i := 0 ; i < n/2 ; i++{
		if clean[i] != clean[n-i-1]{
			return false
		}
	}
	return true
}
func main(){
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter a string for word frequency:")
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	fmt.Println("Word Frequency:", WordFrequency(text))
	fmt.Println("Enter a string to check palindrome:")
	strs, _ := reader.ReadString('\n')
	strs = strings.TrimSpace(strs)
	if IsPalindrome(strs) {
		fmt.Println("The string is a Palindrome")
	} else {
		fmt.Println("The string is not a Palindrome")
	}
}