package main

import (
	"fmt"
)

//  worker pool
// func worker(id int, jobs <-chan int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	for j := range jobs {
// 		fmt.Println("woeker ", id, " job ", j)
// 		time.Sleep(time.Second)

// 	}

// }
// func main() {

// 	var wg sync.WaitGroup
// 	job := make(chan int)

// 	//allocating worker
// 	for w := 0; w < 3; w++ {
// 		wg.Add(1)
// 		go worker(w, job, &wg)
// 	}

// 	//defining jobs
// 	for i := 0; i < 5; i++ {
// 		job <- i
// 	}
// 	close(job)
// 	wg.Wait()

// }

//pipeline

// func main() {
// 	ch := make(chan int)
// 	result := square(ch)
// 	go func() {
// 		for i := 0; i < 10; i++ {

// 			ch <- i
// 		}
// 		close(ch)
// 	}()

// 	for v := range result {
// 		fmt.Println(v)
// 	}

// }

// func square(ch chan int) chan int {
// 	result := make(chan int)
// 	go func() {
// 		for v := range ch {
// 			result <- v * v
// 		}
// 		close(result)
// 	}()
// 	return result

// }

//fan in pattern

// func worker(ch chan string, msg string, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	ch <- msg

// }

// func main() {
// 	var wg sync.WaitGroup
// 	ch := make(chan string)
// 	msg := []string{"this", "is", "msg"}
// 	wg.Add(len(msg))
// 	for _, v := range msg {
// 		go worker(ch, v, &wg)
// 	}

// 	go func() {
// 		wg.Wait()
// 		close(ch)
// 	}()

// 	// print output

// 	for v := range ch {
// 		fmt.Printf(v)
// 	}

// }

//string swap

// func main() {
// 	str := "my name is  "

// 	run := []rune(str)

// 	for i, j := 0, len(run)-1; i < j; i, j = i+1, j-1 {
// 		run[i], run[j] = run[j], run[i]
// 	}

// 	fmt.Println(string(run))
// }

//second largest number

// func secondLargest(num []int) int {

// 	first, sec, thi := math.MinInt, math.MinInt, math.MinInt

// 	for _, v := range num {

// 		if v > first {
// 			thi = sec
// 			sec = first
// 			first = v
// 		} else if v > sec && first != v {
// 			thi = sec
// 			sec = v
// 		} else if v > thi && sec != v {
// 			thi = v
// 		}

// 	}
// 	fmt.Println(first, sec, thi)
// 	return thi
// }

// func main() {
// 	arr := []int{1, 6, 5, 9, 4, 3, 894, 78, 79, 54, 1, 15}
// 	fmt.Print(secondLargest(arr))
// }

// reversing an array

// func main() {
// 	arr := []int{1, 5, 15, 48, 35, 9}
// 	for i := 0; i < len(arr)/2; i++ {
// 		temp := arr[i]
// 		arr[i] = arr[len(arr)-1-i]
// 		arr[len(arr)-1-i] = temp
// 	}
// 	fmt.Println(arr)

// }

// find contiguos subarray( kadane's algo)

// func maxSubArray(nums []int) (int, []int) {
// 	// subArr := make([]int, 0)
// 	var subArr []int
// 	start, end, s := 0, 0, 0
// 	maxSoFar, maxEndingHere := nums[0], nums[0]
// 	for i := 1; i < len(nums); i++ {
// 		if nums[i] > maxEndingHere+nums[i] {
// 			maxEndingHere = nums[i]
// 			s = i
// 		} else {
// 			maxEndingHere = maxEndingHere + nums[i]
// 		}
// 		if maxEndingHere > maxSoFar {
// 			maxSoFar = maxEndingHere
// 			start = s
// 			end = i
// 		}
// 	}
// 	subArr = nums[start : end+1]
// 	return maxSoFar, subArr
// }

// func main() {
// 	arr := []int{1, -2, 5, -3, 4, -1, 6}
// 	fmt.Println(maxSubArray(arr))
// }

//string is palindrome

// func ifPalin(str string) bool {
// 	run := []rune(str)

// 	for i, j := 0, len(run)-1; i < j; i, j = i+1, j-1 {
// 		run[i], run[j] = run[j], run[i]
// 	}
// 	rev := string(run)
// 	if rev == str {
// 		return true
// 	} else {
// 		return false
// 	}
// }
// func main() {
// 	str := "kanak"
// 	fmt.Println(ifPalin(str))
// }

//3 sum = target

//group anargams

// func groupAnagrams(strs []string) [][]string {
// 	anagramMap := make(map[string][]string)

// 	for _, str := range strs {
// 		// Convert the string to a slice of runes for sorting
// 		chars := strings.Split(str, "")
// 		sort.Strings(chars)
// 		sortedStr := strings.Join(chars, "")
// 		anagramMap[sortedStr] = append(anagramMap[sortedStr], str)
// 	}

// 	result := make([][]string, 0, len(anagramMap))
// 	for _, group := range anagramMap {
// 		result = append(result, group)
// 	}
// 	return result
// }

// func main() {
// 	input1 := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
// 	fmt.Println(groupAnagrams(input1)) // Output: [["bat"],["nat","tan"],["ate","eat","tea"]]

// 	input2 := []string{""}
// 	fmt.Println(groupAnagrams(input2)) // Output: [[""]]

// 	input3 := []string{"a"}
// 	fmt.Println(groupAnagrams(input3)) // Output: [["a"]]
// }

//product of array except self

// func productExceptSelf(nums []int) []int {
// 	n := len(nums)
// 	result := make([]int, n)

// 	// Step 1: Populate result with products of elements to the left of i
// 	leftProduct := 1
// 	for i := 0; i < n; i++ {
// 		result[i] = leftProduct
// 		leftProduct *= nums[i]
// 	}

// 	// Step 2: Multiply with products of elements to the right of i
// 	rightProduct := 1
// 	for i := n - 1; i >= 0; i-- {
// 		result[i] *= rightProduct
// 		rightProduct *= nums[i]
// 	}

// 	return result
// }

// func main() {
// 	fmt.Println(productExceptSelf([]int{2, 4, 3, 6}))      // Output: [24, 12, 8, 6]
// 	fmt.Println(productExceptSelf([]int{-1, 1, 0, -3, 3})) // Output: [0, 0, 9, 0, 0]
// }

// func decodeString(s string) string {

// 	var numStack []int
// 	var strStack []string
// 	currentStr := ""
// 	num := 0

// 	for _, char := range s {
// 		if unicode.IsDigit(char) {
// 			// Build the current number (could be multiple digits)
// 			digit, _ := strconv.Atoi(string(char))
// 			num = num*10 + digit
// 		} else if char == '[' {
// 			// Push current number and string to stacks
// 			numStack = append(numStack, num)
// 			strStack = append(strStack, currentStr)
// 			// Reset current number and string
// 			num = 0
// 			currentStr = ""
// 		} else if char == ']' {
// 			// Pop number from numStack
// 			repeatCount := numStack[len(numStack)-1]
// 			numStack = numStack[:len(numStack)-1]
// 			// Pop string from strStack
// 			prevStr := strStack[len(strStack)-1]
// 			strStack = strStack[:len(strStack)-1]
// 			// Repeat currentStr repeatCount times
// 			repeated := ""
// 			for i := 0; i < repeatCount; i++ {
// 				repeated += currentStr
// 			}
// 			// Append repeated string to previous string
// 			currentStr = prevStr + repeated
// 		} else {
// 			// Regular character, append to current string
// 			currentStr += string(char)
// 		}
// 	}
// 	return currentStr
// }

// func main() {
// 	fmt.Println(decodeString("3[a]2[bc]"))
// }

//threeSumClosest
// func threeSumClosest(nums []int, target int) int {
// 	sort.Ints(nums)
// 	closest := nums[0] + nums[1] + nums[2]
// 	for i := 0; i < len(nums)-2; i++ {

// 		left, right := i+1, len(nums)-1
// 		for left < right {
// 			sum := nums[i] + nums[left] + nums[right]
// 			if abs(sum-target) < abs(closest-target) {
// 				closest = sum
// 			}
// 			if sum < target {
// 				left++ // Move left to increase sum
// 			} else {
// 				right-- // Move right to decrease sum
// 			}

// 		}
// 	}
// 	return closest
// }

// func abs(x int) int {
// 	if x < 0 {
// 		return -x
// 	}
// 	return x
// }
// func main() {
// 	nums := []int{-1, 2, 1, -4}
// 	target := 1
// 	fmt.Println(threeSumClosest(nums, target)) // Output: 2
// }

// valid parantheses

// func isValid(s string) bool {
// 	stack := []rune{}
// 	bracketMap := map[rune]rune{')': '(', '}': '{', ']': '['}

// 	for _, char := range s {
// 		switch char {
// 		case '(', '{', '[':
// 			stack = append(stack, char)
// 		case ')', '}', ']':
// 			if len(stack) == 0 || stack[len(stack)-1] != bracketMap[char] {
// 				return false
// 			}
// 			stack = stack[:len(stack)-1]
// 		}

// 	}
// 	return len(stack) == 0

// }

// func main() {
// 	fmt.Println(isValid("()"))
// }

//phone number combination

// func letterCombinations(digits string) []string {
// 	if digits == "" {
// 		return []string{}
// 	}
// 	mapping := map[byte]string{
// 		'2': "abc", '3': "def", '4': "ghi", '5': "jkl",
// 		'6': "mno", '7': "pqrs", '8': "tuv", '9': "wxyz",
// 	}
// 	var result []string

// 	var backtrack func(index int, curr string)
// 	backtrack = func(index int, curr string) {
// 		if index == len(digits) {
// 			result = append(result, curr)
// 			return
// 		}
// 		letters := mapping[digits[index]]
// 		for i := 0; i < len(letters); i++ {
// 			backtrack(index+1, curr+string(letters[i]))
// 		}
// 	}
// 	backtrack(0, "")
// 	return result

// }
// func main() {
// 	fmt.Println(letterCombinations("23"))
// }

// remove duplicate from sorted array

// func removeDuplicates(nums []int) int {
// 	if len(nums) == 0 {
// 		return 0
// 	}
// 	j := 0
// 	for i := 1; i < len(nums); i++ {
// 		if nums[i] != nums[j] {
// 			j++
// 			nums[j] = nums[i]
// 		}
// 	}
// 	return j + 1

// }

// func main() {
// 	fmt.Println(removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))
// }

// routine leaks

// func main() {
// 	go func() {
// 		for {
// 			time.Sleep(10 * time.Second)
// 		}
// 	}()

// 	time.Sleep(1 * time.Second)
// 	fmt.Println("Number of Goroutines:", runtime.NumGoroutine())
// }

// find out unique slice using generics
func uniqueSlice[t comparable](input []t) []t {
	var result []t
	seen := make(map[t]bool)
	for _, v := range input {
		if !seen[v] {
			seen[v] = true
			result = append(result, v)
		}

	}

	return result
}

func main() {
	fmt.Println(uniqueSlice([]int{1, 2, 4, 5, 3, 4, 3, 5, 1, 2}))
	fmt.Println(uniqueSlice([]string{"a", "a", "a", "f", "g", "d", "d", "g"}))
}
