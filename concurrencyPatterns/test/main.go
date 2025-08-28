package main

import "fmt"

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

func ifPalin(str string) bool {
	run := []rune(str)

	for i, j := 0, len(run)-1; i < j; i, j = i+1, j-1 {
		run[i], run[j] = run[j], run[i]
	}
	rev := string(run)
	if rev == str {
		return true
	} else {
		return false
	}
}
func main() {
	str := "kanak"
	fmt.Println(ifPalin(str))
}
