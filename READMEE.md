# Golang Deltas - One Stop SDE2 Preparation
## Module 1: Golang

### Introduce Yourself
I have a background in B.Tech IT and passout of	 2020. After college, I joined Sensibol as a Golang Backend Developer. My primary role was developing microservices using Golang, AWS, MongoDB, and Redis.

I resigned in April to take care of my father's cancer treatment. Now he is well, I am looking for new opportunities.

### What are the projects you woked on Sensibol, describe them
I worked on PDL (Phonographic Digital Limited), a music distribution and royalty management platform. 

And, also worked on Singshala, which is similar to TikTok but with extra feature of analyzing the audio components and providing rankings based on the analysis. 

Both projects used Event Driven microservices architectures.

I last used PostgreSQL in college, but I have been using MongoDB extensively at work.

### Microservices vs Monolith
Microservices are better for large projects where scaling and almost zero downtime are required. Bug fixing and maintaining the codebase are easier. A disadvantage of microservices can be inter-service network calls.

### Authentication vs  Authorization
Authentication is about verifying identity. Users typically provide credentials such as a username and password, or biometric data, and The system checks these credentials against stored data to confirm their validity.

Authorization is about granting permissions based on that verified identity. The system checks the user’s permissions and roles to determine what they can access or modify. 

### Golang Garbage Collection
Golang uses automatic garbage collection to manage memory. Developers do not need to allocate or deallocate memory manually, which reduces memory-related errors. But one disadvantage is CPU and Memory spkies after short interval at Production Servers.

### Pointer
A pointer holds the memory address of a variable, struct, function, or any data type.

You can reference a variable, to get its address, using the & operator.

You can dereference a pointer, to access the value present at the memory address, by using the * operator.
```go
a:=9
p:=&a // p' is a pointer that stores the address of 'a' (referencing)
fmt.Println("Value at address p:", *p) 
// Dereferencing p to get the value (outputs: 9)
*p = 10 // Change the value of 'a' using the pointer
```
### Goroutine vs Thread
Goroutines are designed for concurrency, meaning multiple tasks can run using context switching. Threads are designed for parallelism, meaning multiple tasks can run simultaneously on multiple CPU cores.

Goroutines have a dynamic stack size and are managed by the Go runtime. Threads have a fixed-size stack and are managed by the OS kernel.

### Method Dispatching
golang use Receiver function for method dispatching and has 2 way to dispatch methods at runtime.
```go
func (obj *class_name)method_name(argument int) (returns_name bool){
    //body
}
```
Pointer receiver function :  As obj is refrence of the Struct so any modification inside the function will affect the original Struct. More memory-efficient and can result in faster execution, especially for large structs.

```go
func (obj class_name)method_name(argument int) (returns_name bool){
    //body
}
```
Value receiver function: As obj is copy of the Struct so any modification inside the function will not affect the original Struct. 


### Closure in golang
A closure is a special type of anonymous function that can use variables, that declared outside of the function. 

Closures treat functions as values(First class citigen), allowing us to assign functions to variables, pass functions as arguments, and return functions from other functions. 
```go
v:= func (f func (i int) int) func()int{
    c:=f(3)
    return func()int{
        c++
        return c
    }
}
f:=v()// v() returns a function address
f()// call the function which v() returns
```
### Interfaces in golang
Interfaces allow us to define contracts, which are abstract methods, have no body/implementations of the methods.

A Struct which wants to implements the Interface need to write the body of every abstract methods the interface holds.

We can compose interfaces together.

An empty interface can hold any type of values.

name.(type) give us the Type the interface will hold at runtime. or we can use reflect.TypeOf(name)

```go
func ty(i interface{}) {
	switch i.(type) {
	case int:
		fmt.Println("Integer")
	default:
		fmt.Println("No idea")
	}
	fmt.Println(reflect.TypeOf(i))
}
func main() {
	ty(67.89)
}
```

### Panic Defer Recover combo
- panic is use to cause a Runtime Error and Stop the execution.
- When a function return or panicking, then Defer blocks are called according to Last in First out manner, the last defer will execute first.
- Recover is use to regain the execution from a panicking situation and handle it properly then stop execution. Recover is useful for close any connection like db and websockets etc.
```go
func div(num int) int {
	if num == 0 {
		panic("Not divisible by 0")
	} else {
		return 27 / num
	}
}
func rec() {
	r := recover()
	if r != nil {
		fmt.Println("I am recovering from Panic")
		fmt.Println("I am Fine Now")
	}
}
func main() {
	defer rec()
	fmt.Println(div(0))
	fmt.Println("Main Regained") // Will not executed if divisble by 0
}
```
### Array vs Slice
Array can not Grow and Shrink dynamically at runtime, Slice can. Slice is just references to an existing array of a fixed length.

### Concurency Primitives
Concurency Primitives are tools that are provided by any programming languages to handle execution behaviors of Concurent tasks.

In golang we have Mutex, Semaphore, Channels as concurency primitives.

Mutex is used to protect shared resources from being accessed by multiple threads simultaneously. We Lock() the Shared resource before the Operation, after operation done we Unlock() the resource. RLock(),RUnlock() are used for allowing concurrent Reads.

Semaphore is used to protect shared pool of resources from being accessed by multiple threads simultaneously. Semaphore is a Counter which start from Number of Reosurces. When one thread using the reosurces Semaphore decremented by 1. If semaphore value is 0 then thread will wait untils its value greater than 0. When one thread done with the resources then Semaphore incremented by 1.

Channel is used to communicate via sending and receiving data and provide synchronisation between multiple gorountines. If channel have a value then execution blocked until reader reads from the channel.
Channel can be buffered, allowing goroutines to send multiple values without blocking until the buffer is full. 

Waitgroup is used when we want the function should wait until goroutines complete its task.
Waitgroup has Add() function which increments the wait-counter for each goroutine.
Wait() is used for wait until wait-counter became zero.
Done() decrement wait-counter and it called when goroutine complete its task.

### Wait For Goroutine Completion
```go
func rou(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Hi from GORoutine")
}
func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go rou(&wg)
	wg.Wait()
}
```
### Describe Channel comunication
```go
func rou(ch chan<- int) {
	for i := 0; i < 5; i++ {
		time.Sleep(time.Second)
		ch <- i
	}
	close(ch)
}

func main() {
	ch := make(chan int)
	go rou(ch)
	for i := range ch {
		fmt.Println(i)
	}
}
```
```go
func rouW(ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		time.Sleep(time.Second)
		ch <- i
	}
	close(ch)
}
func rouR(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		fmt.Println(<-ch)
	}
}
func main() {
	ch := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(2)
	go rouW(ch, &wg)
	go rouR(ch, &wg)
	wg.Wait()
}
```

### Map Synchronisation
In golang if multiple goroutines try to acess map at same time, then the operations leads to Panic for RACE or DEADLOCK (fatal error: concurrent map read and map write).

We use MUTEX for LOCK and UNLOCK the Map operations like Read and Write. 

```go
func producer(m *map[int]string, mu *sync.Mutex, wg *sync.WaitGroup, ch chan<- bool) {
	defer wg.Done()
	vm := *m
	for i := 0; i < 5; i++ {
		mu.Lock()
		vm[i] = fmt.Sprint("$", i)
		mu.Unlock()
		if i == 0 {
			ch <- true
			close(ch)
		}
	}
	m = &vm
}
func consumer(m *map[int]string, mu *sync.Mutex, wg *sync.WaitGroup, ch <-chan bool) {
	defer wg.Done()
	vm := *m
	<-ch
	for i := 0; i < 5; i++ {
		mu.Lock()
		fmt.Println(vm[i])
		mu.Unlock()
	}
}
func main() {
	m := make(map[int]string)
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}
	ch := make(chan bool)
	wg.Add(2)
	go producer(&m, &mu, &wg, ch)
	go consumer(&m, &mu, &wg, ch)
	wg.Wait()
}
```
### Asume we have 20 Goroutines, we want to run only 3 Goroutines at a time, when one finished then another spawn

We can use Buffered channel which will be act as a semaphore to control the maximum number of running goroutines.
```go
func rou(i int, wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()
	ch <- i
	time.Sleep(time.Second)
	fmt.Println("$", i)
	<-ch
}
func main() {
	wg := sync.WaitGroup{}
	ch := make(chan int, 3)
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go rou(i, &wg, ch)
	}
	wg.Wait()
}
```
### Task distributions and combine
Shared Variable
```go
func prime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
func PrimeHelper(a []int, f func(n int) bool, m *map[int]bool, wg *sync.WaitGroup, mu *sync.Mutex) {
	defer wg.Done()
	vm := *m
	for i := range a {
		mu.Lock()
		vm[a[i]] = f(a[i])
		mu.Unlock()
	}
	time.Sleep(time.Second)
	m = &vm
}
func main() {
	m := make(map[int]bool)
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}
	arr := arrGen()
	noGoRo := 6
	parts := len(arr) / noGoRo
	for i := 0; i < noGoRo; i++ {
		wg.Add(1)
		s := parts * i
		e := s + parts
		if e > len(arr) {
			e = len(arr)
		}
		go PrimeHelper(arr[s:e], prime, &m, &wg, &mu)
	}
	wg.Wait()
	fmt.Println(m)
}
func arrGen() []int {
	var a []int
	for i := 0; i < 100; i++ {
		a = append(a, i)
	}
	return a
}
```
Channel
```go
package main

import (
	"fmt"
	"sync"
	"time"
)

func prime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
func PrimeHelper(a []int, f func(n int) bool, ch chan map[int]bool, wg *sync.WaitGroup) {
	defer wg.Done()
	m := make(map[int]bool)
	for i := range a {
		m[i] = f(a[i])
	}
	time.Sleep(time.Second)
	ch <- m
}
func main() {
	m := make(map[int]bool)
	wg := sync.WaitGroup{}
	arr := arrGen()
	noGoRo := 6
	parts := len(arr) / noGoRo
	ch := make(chan map[int]bool, noGoRo)
	for i := 0; i < noGoRo; i++ {
		wg.Add(1)
		s := parts * i
		e := s + parts
		if e > len(arr) {
			e = len(arr)
		}
		go PrimeHelper(arr[s:e], prime, ch, &wg)
	}
	wg.Wait()
	close(ch)
	for i := range ch {
		for k, b := range i {
			m[k] = b
		}
	}
	fmt.Println(m)
}
func arrGen() []int {
	var a []int
	for i := 0; i < 100; i++ {
		a = append(a, i)
	}
	return a
}

```