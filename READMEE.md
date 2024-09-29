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
```go
func odder(n int, ch1 chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	if n%2 != 0 {
		ch1 <- n
	}
}
func evener(n int, ch2 chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	if n%2 == 0 {
		ch2 <- n
	}
}
func main() {
	var intSlice = []int{91, 42, 23, 14, 15, 76, 87, 28, 19, 95}
	ch1 := make(chan int) // Channel for odd numbers
	ch2 := make(chan int) // Channel for even numbers
	var wg sync.WaitGroup

	for _, num := range intSlice {
		wg.Add(2) // Add two for odd and even checks
		go odder(num, ch1, &wg)
		go evener(num, ch2, &wg)
	}
	// Close channels after all goroutines are done
	go func() {
		wg.Wait()
		close(ch1)
		close(ch2)
	}()
	go func() {
		for odd := range ch1 {
			fmt.Println("ODD :", odd)
		}
	}()
	go func() {
		for even := range ch2 {
			fmt.Println("EVEN:", even)
		}
	}()
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
Using Shared Variable
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
	parts := (len(arr) + noGoRo - 1) / noGoRo
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
Using Channel
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
	for _, v := range a {
		m[v] = f(v)
	}
	time.Sleep(time.Second)
	ch <- m
}
func main() {
	m := make(map[int]bool)
	wg := sync.WaitGroup{}
	arr := arrGen()
	noGoRo := 6
	parts := (len(arr) + noGoRo - 1) / noGoRo
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
	for i := 0; i < len(arr); i++ {
		fmt.Println(i, " : ", m[i])
	}
}
func arrGen() []int {
	var a []int
	for i := 0; i < 100; i++ {
		a = append(a, i)
	}
	return a
}
```
### Select Statement
Assume a development scenerio where we have 3 s3 Buckets. We spawn 3 GO-Routines each one uploading a File on each S3 bucket at same time. We have to Return SignedUrl of the file so user can stream the File as soon as possible. Now we do not have to wait for 3 S3 Upload operation, when one s3 upload done we can send the SignedUrl of the File to the User so he can Stream. And Other two S3 Upload will continue at same time. This is the Scenerio when Select Statement can help.

Select statement is used for Concurency comunication between multiple goroutines. Select have multiple Case statement related to channel operations. Select block the execution unitl one of its case return. If multiple case returns at same time, then one random case is selected for returns. If no case is ready and there's a default case, it executes immediately. If there's no default case, select blocks until at least one case is ready.
```go
func work(ctx context.Context, ch chan<- string) {
	rand.NewSource(time.Now().Unix())
	r := rand.Intn(6)
	t1 := time.Duration(r) * time.Second
	ctx, cancel := context.WithTimeout(ctx, t1)
	defer cancel()
	select {
	case <-time.After(t1):
		ch <- "Connection established"
	case <-ctx.Done():
		ch <- "Context Expired"
	}
}

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	go work(context.Background(), ch1)
	go work(context.Background(), ch2)
	select {
	case res := <-ch1:
		fmt.Println("ch1 ", res)
	case res := <-ch2:
		fmt.Println("ch2 ", res)
	}
}
```

```go
func uploadToS3(bucket string, wg *sync.WaitGroup, ch chan<- string) {
	defer wg.Done()
	time.Sleep(time.Second)
	ch <- fmt.Sprintf("https://%s.s3.amazonaws.com/file", bucket)
}

func main() {
	buckets := []string{"bucket1", "bucket2", "bucket3"}
	var wg sync.WaitGroup
	ch := make(chan string, len(buckets))

	for _, bucket := range buckets {
		wg.Add(1)
		go uploadToS3(bucket, &wg, ch)
	}
	go func() {
		wg.Wait()
		close(ch)
	}()
	// Wait for the first completed upload
	//  As soon as one of the goroutines finishes uploading to an S3 bucket and sends a signed URL through the urls channel, this case will be selected, and the url variable will be assigned that value. 
	select {
	case url := <-ch:
		fmt.Println("First completed upload URL:", url)
	case <-time.After(5 * time.Second): // Optional timeout
		fmt.Println("No uploads completed in time.")
	}
}
```
### SOLID Principles
SOLID priciples are designe guidelines for Code base that are easy to understand maintain and extend over time.

Single Responsibility:- A Struct/Class should have only a single reason to change. Fields of Author shoud not placed inside Book Struct.
```go
type Book struct{
  ISIN string
  Name String
  AuthorID string
}
type Author struct{
  ID string
  Name String
}
```
Assume One Author decided later, he does not want to Disclose its Real Name to Spread. So we can Serve Frontend by Alias instead of Real Name. Without Changing Book Class/Struct, we can add Alias in Author Struct. By that, Existing Authors present in DB will not be affected as Frontend will Change Name only when it Founds that - Alias field is not empty.
```go
type Book struct{
  ISIN string
  Name String
  AuthorID string
}
type Author struct{
  ID string
  Name String
  Alias String
}

```
Open Close:- Struct and Functions should be open for Extension but closed for modifications. New functionality to be added without changing existing Code.
```go
type Shape interface{
	Area() float64
}
type Rectangle struct{
	W float64
	H float64
}
type Circle struct{
	R float64
}
```
Now we want to Calculate Area of Rectangle and Circle, so Rectangle and Circle both can Implements Shape Interface by Write Body of the Area() Function.
```go
func (r Rectangle) Area()float64{
	return r.W * r.H
}
func (c Circle)Area()float64{
	return 3.14 * c.R * c.R
}
```
Now we can create a Function PrintArea() which take Shape as Arguments and Calculate Area of that Shape. So here Shape can be Rectangle, Circle. In Future we can add Triangle Struct which implements Shape interface by writing Body of Area. Now Traingle can be passed to PrintArea() with out modifing the PrintArea() Function.
```go
func PrintArea(shape Shape) {
	fmt.Printf("Area of the shape: %f\n", shape.Area())
}

// In Future
type Triangle struct{
	B float64
	H float54
}
func (t Triangle)Area()float64{
	return 1/2 * t.B * t.H
}

func main(){
	rect:= Rectangle{W:5,H:3}
	cir:=Circle{R:3}
	PrintArea(rect)
	PrintArea(cir)
	// In Future
	tri:=Triangle{B:4,H:8}
	PrintArea(tri)
}
```
Liskov Substitution:- Super class Object can be replaced by Child Class object without affecting the correctness of the program.
```go
type Bird interface{
	Fly() string
}
type Sparrow struct{
	Name string
}
type Penguin struct{
	Name string
}
```
Sparrow and Pengin both are Bird, But Sparrow can Fly, Penguin Not. ShowFly() function take argument of Bird type and call Fly() function. Now as Penguin and Sparrow both are types of Bird, they should be passed as Bird within ShowFly() function.
```go
func (s Sparrow) Fly() string{
	return "Sparrow is Flying"
}
func (p Penguin) Fly() string{
	return "Penguin Can Not Fly"
}

func ShowFly(b Bird){
	fmt.Println(b.Fly())
}
func main() {
	sparrow := Sparrow{Name: "Sparrow"}
	penguin := Penguin{Name: "Penguin"}
  // SuperClass is Bird,  Sparrow, Penguin are the SubClass
	ShowFly(sparrow)
	ShowFly(penguin)
}
```
Interface Segregation:- A class should not be forced to implements interfaces which are not required for the class. Do not couple multiple interfaces together if not necessary then. 
```go
// The Printer interface defines a contract for printers with a Print method.
type Printer interface {
	Print()
}
// The Scanner interface defines a contract for scanners with a Scan method.
type Scanner interface {
	Scan()
}
// The NewTypeOfDevice interface combines Printer and Scanner interfaces for
// New type of devices which can Print and Scan with it new invented Hardware.
type NewTypeOfDevice interface {
	Printer
	Scanner
}
```

Dependecy Inversion:- Class should depends on the Interfaces not the implementations of methods.

```go
// The MessageSender interface defines a contract for 
//sending messages with a SendMessage method.
type MessageSender interface {
	SendMessage(msg string) error
}
// EmailSender and SMSClient structs implement 
//the MessageSender interface with their respective SendMessage methods.
type EmailSender struct{}

func (es EmailSender) SendMessage(msg string) error {
	fmt.Println("Sending email:", msg)
	return nil
}
type SMSClient struct{}

func (sc SMSClient) SendMessage(msg string) error {
	fmt.Println("Sending SMS:", msg)
	return nil
}
type NotificationService struct {
	Sender MessageSender
}
```
The NotificationService struct depends on MessageSender interface, not on concrete implementations (EmailSender or SMSClient). This adheres to Dependency Inversion, because high-level modules (NotificationService) depend on abstractions (MessageSender) rather than details.
```go
func (ns NotificationService) SendNotification(msg string) error {
	return ns.Sender.SendMessage(msg)
}
func main() {
	emailSender := EmailSender{}

	emailNotification := NotificationService{Sender: emailSender}

	emailNotification.SendNotification("Hello, this is an email notification!")
}
```
### Create Post and Get APIs:
```go
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/gorilla/mux"
)
// ---------- Domain -----------
type Counter struct {
	Count int
}

// ---------- Service -----------
type Service interface {
	AddCounter()
	GetCounter(id string) int
}
type service struct {
	Counter Counter
}

func NewService(c Counter) *service {
	return &service{Counter: c}
}
func (s *service) AddCounter() {
	s.Counter.Count++
}
func (s *service) GetCounter(id string) int {
	fmt.Println(id)
	return s.Counter.Count
}


// ---------- Transport -----------
type httpTransport struct {
	Service // Interfaces use for Dependency Inversion
}

func NewHttpTransport(s Service) *httpTransport {
	return &httpTransport{
		Service: s,
	}
}

type AddReq struct {
}
type AddRes struct {
	Success bool `json:"success"`
}
type GetReq struct {
	Id string `json:"id"`
}
type GetUrlReq struct {
	Id string `url:"id"`
}
type GetRes struct {
	Count   int  `json:"count"`
	Success bool `json:"success"`
}

func (t *httpTransport) Add(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	t.AddCounter()
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(AddRes{Success: true})
}
func (t *httpTransport) Get(w http.ResponseWriter, r *http.Request) {
	var req GetReq
	json.NewDecoder(r.Body).Decode(&req)
	count := t.GetCounter(req.Id)
	json.NewEncoder(w).Encode(GetRes{Count: count, Success: true})
}
func (t *httpTransport) GetUrl(w http.ResponseWriter, r *http.Request) {
	u, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		json.NewEncoder(w).Encode(GetRes{Count: -1, Success: false})
	}
	id := u["id"][0]
	count := t.GetCounter(id)
	json.NewEncoder(w).Encode(GetRes{Count: count, Success: true})
}

// ---------- Main -----------
func main() {
	// Service
	s := NewService(Counter{})
	// Trnasport
	t := NewHttpTransport(s)
	// Routing
	r := mux.NewRouter()
	r.HandleFunc("/api/add", t.Add).Methods("POST")
	r.HandleFunc("/api/get", t.Get).Methods("GET")
	r.HandleFunc("/api/geturl", t.GetUrl).Methods("GET")
	// Server
	http.ListenAndServe(":3000", r)
}
```

### Golang Unit Testing
```go
func ComplexAlgo(a []int, t int) (int, int) {
	m := make(map[int]int)
	for i := range a {
		m[a[i]] = i
	}
	for i := range a {
		snd := t - a[i]
		v, ok := m[snd]
		if ok {
			return i, v
		}
	}
	return -1, -1
}
```
```go
import (
	"fmt"
	"testing"
	"time"

	"math/rand"
)
type inps struct {
	a []int
	t int
}
type exps struct {
	first  int
	second int
}
func TestComplexAlgo(t *testing.T) {
	a1 := []int{2, 4, 6, 8, 10, 12}
	t1 := 10
	a2 := []int{1, 2, 4, 6, 8, 10}
	t2 := 16
	a3 := []int{2, 3, 4, 6, 12}
	t3 := 15
	tests := []struct {
		input    inps
		expected exps
	}{
		{inps{a: a1, t: t1}, exps{first: 0, second: 3}},
		{inps{a: a2, t: t2}, exps{first: 3, second: 5}},
		{inps{a: a3, t: t3}, exps{first: 1, second: 4}},
	}

	for i := range tests {
		f, s := ComplexAlgo(tests[i].input.a, tests[i].input.t)
		if f != tests[i].expected.first && s != tests[i].expected.second {
			t.Errorf("ComplexAlgo(%d)=%d but got %d,%d", tests[i].input, tests[i].expected, f, s)
		}
	}
}
func BenchmarkComplexAlgo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rand.NewSource(time.Now().Unix())
		arraySize := 10
		randomArray := generateRandomArray(arraySize)
		targetSum := rand.Intn(20*2-2*1+1) + 2*1
		fmt.Println(randomArray, targetSum)
		ComplexAlgo(randomArray, targetSum)
	}
}
func generateRandomArray(size int) []int {
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = rand.Intn(20-1+1) + 1
	}
	return arr
}
```
## Module 2: AWS:
### Theory
- AWS IAM (Identity and Access Management) helps you manage who can access your resources in Amazon Web Services. You can create users, groups, and roles to control permissions. This means you can decide who can do what, like viewing or changing files. IAM also allows for secure sign-in using passwords and multi-factor authentication. Overall, it helps keep your AWS environment safe and organized.

- An AWS User Group is a collection of users that you can manage together. It helps you organize users who need similar access permissions. By assigning permissions to the group, you simplify managing user access. When you add a user to the group, they automatically get those permissions. This makes it easier to keep everything organized and secure in AWS.

- AWS SG, or Security Group, acts like a virtual firewall for your Amazon Web Services resources. It controls what traffic can enter or leave your instances. You can set rules based on IP addresses, protocols, and ports. This helps protect your applications from unwanted access. Security Groups are easy to manage and can be adjusted anytime as your needs change.

- You use AWS Security Groups with services like EC2 (Elastic Compute Cloud) to control access to virtual servers. They also work with RDS (Relational Database Service) to protect your databases from unauthorized access. For AWS Lambda functions, Security Groups help manage network access when they connect to other services. Additionally, you can use them with Elastic Load Balancers to ensure traffic is properly filtered. Overall, Security Groups are important for securing various AWS resources and services.

- An AWS Availability Zone is a separate location within a region that has its own power, cooling, and networking. Each zone is designed to be isolated from failures in other zones. This means if one zone goes down, your applications can still run in another zone. Using multiple Availability Zones helps improve the reliability and availability of your services. It's a key part of building resilient cloud applications in AWS.

- An Application Load Balancer (ALB) distributes incoming application traffic across multiple targets, like EC2 instances or containers. It operates at the application layer, allowing it to route requests based on content, such as URL paths and HTTP headers. ALB performs health checks on registered targets to ensure that traffic is only sent to healthy instances. It can also handle SSL termination, reducing the load on backend servers. Overall, ALB enhances application performance, scalability, and reliability.

- An Auto Scaling Group (ASG) in AWS automatically adjusts the number of EC2 instances based on demand. It can scale out by adding instances when traffic increases and scale in by removing instances during low demand. ASGs perform health checks to replace any unhealthy instances, ensuring reliability. You can configure scaling policies based on metrics like CPU usage or set scheduled scaling for predictable traffic patterns. Overall, ASGs help maintain application performance while optimizing costs.

- Amazon S3 (Simple Storage Service)
Amazon S3 is a scalable object storage service that allows you to store and retrieve any amount of data from anywhere on the web. It ensures durability and security with features like versioning and lifecycle policies, organizing data in buckets. Common uses include backups, static website hosting, and big data analytics, all under a pay-as-you-go pricing model. S3 offers various storage classes to optimize cost and performance, including S3 Standard for frequently accessed data and S3 Intelligent-Tiering, which adjusts based on access patterns. S3 Glacier and S3 Glacier Deep Archive provide cost-effective options for long-term archival storage with longer retrieval times.

- Amazon SQS (Simple Queue Service)
Amazon SQS is a fully managed message queuing service that decouples and scales microservices and applications. It enables sending, storing, and receiving messages between components without loss, supporting standard queues for high throughput and FIFO queues for ordered processing. SQS ensures reliable message processing and handles workload spikes, making it ideal for asynchronous communication.

- Amazon SNS (Simple Notification Service)
Amazon SNS is a fully managed messaging service for sending notifications to multiple subscribers, including applications and devices. It supports various communication methods like SMS, email, and push notifications and allows the creation of topics for efficient message broadcasting. SNS can trigger workflows and alerts in real time, integrating well with other AWS services for event-driven architectures.

- AWS Lambda
AWS Lambda is a serverless computing service that runs your code without requiring server management. You simply upload your code, and Lambda handles execution, scaling, and availability. It supports various programming languages and can be triggered by events from other AWS services, making it ideal for event-driven applications and APIs. You pay only for the compute time used, making it cost-effective for variable workloads.

- Amazon Route 53 is a scalable DNS web service that translates domain names into IP addresses. It routes internet traffic to resources like websites by directing users to the nearest server based on latency or location. Route 53 includes health checks to monitor resource availability and reroutes traffic from unhealthy endpoints. It also supports domain registration, simplifying the management of domain names and DNS settings.

## Module 3: Docker Kubernetis:
- Docker is a platform that allows developers to easily package and deploy applications in lightweight, portable containers. These containers include everything needed to run the application, ensuring consistency across different environments. This makes it simpler to build, share, and manage applications efficiently. 

- Docker Images: Read-only templates used to create containers. They contain everything needed to run an application, including code, runtime, libraries, and environment variables.  Images are built in layers, allowing efficient storage and sharing. Changes in images only add new layers instead of rewriting the entire image.

- Docker provides several networking options to allow containers to communicate with each other and with external systems, including bridge networks, host networks, and overlay networks. Bridge networks are the default network type in Docker, allowing containers on the same host to communicate with each other while isolating them from other networks. Host networks remove network isolation between the container and the Docker host, allowing the container to use the host’s networking stack directly, which can improve performance but reduces isolation. Overlay networks enable containers running on different Docker hosts (in a multi-host setup) to communicate with each other, typically used in orchestrated environments like Docker Swarm or Kubernetes.

- Data Persistence: Volumes are used to store data generated by containers, allowing for data persistence beyond the lifecycle of a container. They enable easy sharing of data between containers and the host system.

- A Dockerfile is a script that contains a series of commands to assemble a Docker image. It defines the base image, dependencies, and instructions for building the image.
```dockerfile
# Step 1: Use the official Golang image to build the application
FROM golang:1.20 AS builder
```
FROM: This command specifies the base image for the container. Here, we’re using the official Go image version 1.20. The AS builder syntax allows us to name this stage "builder," enabling us to reference it later in a multi-stage build.

```dockerfile
# Step 2: Set the Current Working Directory inside the container
WORKDIR /app
```
WORKDIR: This command sets the working directory inside the container. If the directory does not exist, it will be created. Subsequent commands like COPY and RUN will operate relative to this directory.
```dockerfile
# Step 3: Copy go.mod and go.sum files to the working directory
COPY go.mod go.sum ./
```
COPY: This command copies files from the host machine into the container. In this case, we're copying go.mod and go.sum, which contain information about dependencies for the Go application. The ./ denotes the current working directory inside the container.
```dockerfile
# Step 4: Download dependencies
RUN go mod download
```
RUN: This command executes a command in a new layer on top of the current image and commits the results. Here, we're using it to download the Go module dependencies specified in go.mod. This layer caches the dependencies, improving build times if the dependencies don’t change.
```dockerfile
# Step 5: Copy the source code into the container
COPY . .
```
COPY: Again, this command is used to copy files. This time, it copies the entire application source code from the host (the current directory) to the working directory inside the container (/app).
```dockerfile
# Step 6: Build the Go application
RUN go build -o myapp .
```
RUN: This command compiles the Go application. The -o myapp option specifies the output file name. The . indicates that the build context is the current directory, which contains the Go source files.
```dockerfile
# Step 7: Use a smaller base image for the final image
FROM alpine:latest
```
FROM: This command starts a new build stage using a minimal base image, in this case, Alpine Linux. This is beneficial for reducing the final image size and minimizing security vulnerabilities.
```dockerfile
# Step 8: Set the working directory in the final image
WORKDIR /app
```
WORKDIR: Sets the working directory again for the final image. This ensures that the final container runs with the same directory structure.
```dockerfile
# Step 9: Copy the compiled binary from the builder stage
COPY --from=builder /app/myapp .
```
COPY --from=builder: This command copies the compiled Go binary from the previous build stage (named "builder"). It pulls the myapp binary from the /app directory in the builder stage and places it in the current working directory of the final image.
```dockerfile
# Step 10: Expose the port that the application will run on
EXPOSE 8080
```
EXPOSE: This command documents which ports the container listens on at runtime. It does not actually publish the port; this is done with the docker run command. Here, we're indicating that our application listens on port 8080.
```dockerfile
# Step 11: Command to run the application
CMD ["./myapp"]
```

CMD: This command specifies the default command to run when the container starts. It can be overridden by command-line arguments provided to docker run. Here, we’re running the compiled Go application.

`ARG APP_VERSION=1.0.0`
Defines a build argument that can be passed during the build process. This example sets a default value for APP_VERSION, which can be overridden.

`LABEL maintainer="yourname@example.com" ...`
Adds metadata to the image. Labels can provide useful information about the image, such as the maintainer's contact information, version, and description.

`ADD . /app`
Copies all files from the current directory on the host machine to the /app directory in the container. The ADD command also supports URL downloads and can automatically extract tar files.

`ENV APP_ENV=production APP_DEBUG=false`
Sets environment variables inside the container. These can be accessed by the 
application at runtime, allowing for configurable behavior.

`EXPOSE 5000`
Informs Docker that the application listens on port 5000. This command does not publish the port; it serves as documentation and is useful for when you run the container.

`VOLUME ["/app/data"]`
Creates a mount point for a volume at /app/data, which can be used to persist data generated by the application. This ensures that data is not lost when the container stops or is removed.

`ENTRYPOINT ["python", "app.py"]`
Specifies the command that should be executed when the container starts. This makes it easier to run the application with the specified entry point.


```dockerfile
# Step 1: Use the official Python image as the base
FROM python:3.10-slim

# Step 2: Set build arguments
ARG APP_VERSION=1.0.0

# Step 3: Add metadata to the image
LABEL maintainer="yourname@example.com" \
      version="$APP_VERSION" \
      description="A simple Python web application"

# Step 4: Set the working directory
WORKDIR /app

# Step 5: Add application files
ADD . /app

# Step 6: Install dependencies
RUN pip install --no-cache-dir -r requirements.txt

# Step 7: Set environment variables
ENV APP_ENV=production \
    APP_DEBUG=false

# Step 8: Expose the port the app runs on
EXPOSE 5000

# Step 9: Create a volume for persistent data
VOLUME ["/app/data"]

# Step 10: Define the command to run the application
ENTRYPOINT ["python", "app.py"]
```

```bash
docker build -t my-python-app .
docker run -p 5000:5000 -v mydata:/app/data my-python-app
```

- Kubernetes is an open-source container orchestration platform designed to automate the deployment, scaling, and management of containerized applications.

- Users define the desired state of the system using YAML or JSON files, and Kubernetes works to maintain that state.

- Pods: The smallest deployable units, which can contain one or more containers.

- Kubernetes can automatically scale applications based on demand, using metrics like CPU or memory usage.
Load balancing distributes network traffic to ensure no single pod is overwhelmed.

- Role-Based Access Control (RBAC): RBAC is a method used to restrict system access to authorized users. In Kubernetes, RBAC allows administrators to define who can perform what actions on which resources within the cluster.

- Kubernetes provides a flat network structure, allowing all pods to communicate with each other by default.

- Network Policies are used to control the traffic flow between pods in a Kubernetes cluster. They allow you to define rules that govern how pods can communicate with each other and with external resources, enhancing the security of your applications.
	- Ingress Rules: Specify what inbound traffic is allowed to reach a pod.
	- Egress Rules: Define what outbound traffic is allowed from a pod.
	- Selectors: Used to select which pods the policies apply to.

- Services: Abstractions that define a logical set of pods and a policy to access them, enabling load balancing. Services provide stable endpoints for pods, even as they scale or restart.

- Deployments: Manage the lifecycle of pods, allowing for rolling updates and rollbacks.

- A Kubernetes cluster consists of a master node (control plane) and multiple worker nodes.
The master node manages the cluster, while worker nodes run the application containers.

- The master node is the control plane of the Kubernetes cluster, responsible for managing the cluster's state and making global decisions. It runs the API server, scheduler, and controller manager, coordinating the worker nodes and managing workload deployments.

- Worker nodes run the applications and services in the cluster. Each worker node contains components such as the kubelet, kube-proxy, and a container runtime, enabling it to communicate with the master node and manage containers.

- The API server acts as the front-end for the Kubernetes control plane, handling all requests for the cluster's resources. It provides a RESTful interface for clients to interact with the Kubernetes cluster, managing the state of the system and facilitating communication between components.

- The scheduler is responsible for assigning pods to specific worker nodes based on resource availability and constraints. It evaluates the current state of the cluster and makes placement decisions to optimize resource usage and maintain workload performance.

- The controller manager runs various controllers that regulate the state of the cluster, ensuring that the desired state matches the current state. It handles tasks such as replication, node management, and monitoring the health of pods and nodes.

- The kubelet is an agent running on each worker node, responsible for managing the lifecycle of containers. It ensures that the containers defined in pods are running as expected, communicating with the API server to report the node's status.

- The kube proxy manages network communication for services within the cluster, facilitating load balancing and service discovery. It maintains network rules that allow pods to communicate with each other and with external clients, ensuring reliable traffic routing.

- etcd is a distributed key-value store used for storing configuration data and the state of the Kubernetes cluster. It provides a reliable and consistent way to store and retrieve cluster metadata, enabling high availability and fault tolerance.

- The container runtime is responsible for running the containers on worker nodes. Kubernetes supports several runtimes, such as Docker and containerd, which handle the creation, management, and execution of containers.

#### To deploy a microservices architecture using Kubernetes, we can create services for User, Order, and an API Gateway, along with a MongoDB database. 

1. MongoDB Deployment and Service
First, create a Deployment for MongoDB that specifies the desired number of replicas and the container image to use. This deployment ensures that the MongoDB instance is always running and can be automatically rescheduled if it fails. You also define a Persistent Volume Claim (PVC) to provide persistent storage, ensuring that data is not lost when the pod is restarted. A Service is created to expose the MongoDB instance, allowing other services to communicate with it via a stable endpoint. The MongoDB service uses the default port 27017, which is specified in the service configuration.

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: mongodb
spec:
  replicas: 1
  selector:
    matchLabels:
      app: mongodb
  template:
    metadata:
      labels:
        app: mongodb
    spec:
      containers:
      - name: mongodb
        image: mongo:latest
        ports:
        - containerPort: 27017
        volumeMounts:
          - name: mongo-data
            mountPath: /data/db
      volumes:
        - name: mongo-data
          persistentVolumeClaim:
            claimName: mongo-pvc

---
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: mongo-pvc
spec:
  accessModes:
    - ReadWriteOnce
  resources:
    requests:
      storage: 1Gi

---
apiVersion: v1
kind: Service
metadata:
  name: mongodb
spec:
  ports:
    - port: 27017
  selector:
    app: mongodb
```

2. User Service Deployment and Service
Next, set up a Deployment for the User Service, defining the container image and the number of replicas required for handling user-related operations. Environment variables are configured to specify the MongoDB connection string, ensuring that the service can access the database. A Service is also created for the User Service to allow internal communication with other components in the cluster. This service exposes the User Service on port 8080, enabling other services to route requests to it. The User Service will handle API calls related to user management and leverage the MongoDB database for data storage.

3. Order Service Deployment and Service
Then, deploy the Order Service using a similar approach to the User Service, specifying the container image and number of replicas. The Order Service also requires a MongoDB connection string set as an environment variable to interact with its dedicated database. A corresponding Service is created to expose the Order Service on port 8080, facilitating communication from the API Gateway and other services. This service will manage order-related operations, processing requests and storing data in MongoDB. By isolating the Order Service, you ensure better organization and scalability of your application.

4. API Gateway Deployment and Service
After setting up the microservices, deploy the API Gateway, which acts as the entry point for external clients. The API Gateway's Deployment specifies the container image and the number of replicas needed to handle incoming requests effectively. A Service is created for the API Gateway, typically configured as a LoadBalancer to expose it externally, allowing users to access the services. This service listens on port 8080 and routes requests to the appropriate backend services based on the API paths. The API Gateway centralizes access control, load balancing, and request routing, simplifying client interactions with the microservices.

5. Deploying the Application
Finally, save all the defined configurations in separate YAML files for easy management. Use the kubectl apply command to deploy each of the components to the Kubernetes cluster sequentially. This command instructs Kubernetes to create the resources defined in your manifests, such as Deployments and Services. Monitor the deployment process using kubectl get pods and kubectl get services to ensure that all components are running correctly. Once deployed, the services can communicate with one another, and the API Gateway can accept external requests, forming a complete microservices architecture.

#### Docker and Kubernetis commands:
`docker run -d -p 80:80 nginx`
This command creates and starts a container from a specified image. You can add options to customize the container's behavior, such as mapping ports and volumes.

`docker ps -a`
Lists all currently running containers. Use -a to see all containers, including stopped ones.

`docker exec -it my_container bash`
Runs a command in a running container. Useful for troubleshooting or interacting with the application inside the container.

`docker build -t my_image .`
Builds a Docker image from a Dockerfile. This command takes the context (the files and directories) and processes it according to the instructions in the Dockerfile.

`docker stop my_container`
Stops a running container. This sends a SIGTERM signal to the container's main process, allowing it to gracefully shut down.


`docker rm my_container`
Removes one or more stopped containers. This helps free up resources and keep your Docker environment clean.

`docker logs my_container`
Fetches the logs of a container. Useful for debugging or monitoring the application’s output.

`docker-compose up -d`
Starts up a multi-container application defined in a docker-compose.yml file. It handles networking and volume management for you.

`kubectl get pods`
Lists all pods in the current namespace, showing their status and other key information.

`kubectl apply -f deployment.yaml`
Applies the configuration defined in a file (like a YAML file) to create or update resources in the cluster.

`kubectl describe pod my-pod`
Provides detailed information about a specific pod, including its state, conditions, and events.	

`kubectl delete pod my-pod`
Deletes a specific pod from the cluster. If it's part of a deployment, Kubernetes will create a new one to replace it.

`kubectl exec -it my-pod -- /bin/bash`
Opens a shell in a running container within a pod, allowing for interactive troubleshooting or command execution.

`kubectl scale deployment my-deployment --replicas=3`
Changes the number of replicas of a deployment, allowing for scaling applications up or down easily.

`kubectl get services`
Lists all services in the current namespace, showing how pods are exposed within the cluster.

`kubectl rollout status deployment/my-deployment`
Checks the status of a rollout for a given deployment, helping to monitor updates or changes.

`kubectl logs my-pod`
Retrieves logs from a specific pod, useful for debugging and monitoring the application’s behavior.

`kubectl port-forward my-pod 8080:80`
Forwards a port from your local machine to a port on a pod, allowing for easy access to applications running in Kubernetes.

## Module 4: HR Round and Managerial Round:
### Tell me about yourself
I am a Golang developer with 3.3 years of experience in backend development. I enjoy building scalable apps and have a strong background in microservices. I’m passionate about clean code and passionated about always learning to stay updated with industry trends.

### What do you know about our company
Calsoft is a software development company that focuses on cloud computing, storage, networking, and virtualization technologies.

### Why did you choose Golang as your primary programming language
I chose Golang because it's fast, simple, and great at handling multiple tasks at once. It’s perfect for building microservices and managing large systems.

### What are your strengths as a developer
My strengths are problem-solving and understanding requirements, which helps me write efficient code.

### What is your biggest weakness
One of my weaknesses is being too detail-oriented. While it helps me produce high-quality work, it can sometimes make me spend more time on tasks than needed.

### How do you handle stress and tight deadlines
I prioritize my tasks and break down into manageable parts.I personally use Notion to Schedule my tasks. I maintain open communication with my team to ensure we’re aligned and stay organized.

### Describe a challenging project you worked on
The biggest challenge was ensuring smooth communication between SQS, Lambda, and SNS, and handling errors well. We set up the architecture to manage retries and keep data safe. In the end, the project successfully provided real-time insights and notifications while being scalable.

### How do you stay updated with new technologies and trends
I regularly follow industry blogs, participate in online courses, and engage in developer communities on platforms like GitHub and Stack Overflow.

### What motivates you in your work
I’m motivated by solving problems and the satisfaction of creating strong solutions.

### Where do you see yourself in five years
In five years, I see myself in a leadership role, possibly as a technical lead or architect.

### How do you approach code reviews
I see code reviews as chances to learn. I focus on giving and receiving helpful feedback to identify areas for improvement.

### Can you describe your experience with Agile methodologies
I have worked in Agile environments for most of my career. I take part in daily stand-ups, sprint planning, and testing.

### What is your experience with cloud services
I have hands-on experience with AWS and using services like EC2, Lambda, s3, sns, sqs. I appreciate the flexibility and scalability cloud platforms provide.

### How do you handle team conflicts
I believe in dealing with conflicts early and encouraging open discussions. I listen to everyone’s views and promote teamwork to find solutions that support our goals.

### Describe your experience with databases in Golang
I have worked with both SQL and NoSQL databases. I’m comfortable using ORMs(Object-Relational Mapping) like GORM and managing database migrations, ensuring optimal performance and data integrity.

GORM makes working with databases easier by letting developers focus on their data models instead of writing complex SQL queries. This means you can define your data structures in simple classes and GORM handles the database tasks for you, saving time and reducing errors.

### How do you ensure code quality
I use automated testing, continuous integration, and code reviews in my development process. I also promote writing clean, easy-to-maintain code and follow SOLID principles for the codebase.

### What techniques do you use for debugging
I employ a combination of logging, using tools like Delve, and unit tests to identify and resolve issues.

### How do you approach mentoring junior developers
I actively mentor by sharing best practices and encouraging others to be independent. I'm available for questions and give helpful feedback on their work.

### What is your experience with RESTful APIs
I have designed and implemented several RESTful APIs in Golang, focusing on clean architecture. I prioritize clear documentation and versioning to facilitate API usability.

### What role do you think Golang plays in the future of backend development
Golang’s emphasis on simplicity and performance makes it a strong candidate for cloud-native applications and microservices. Its growing ecosystem and community support suggest it will remain relevant in the future with Rust language.

### How do you handle technical debt in your projects
technical debt is quick fixes taken during development that may save time in the short term but can lead to problems later on.

I work on fixing technical debt during sprint planning, balancing new features with code improvements. By checking and writing down the debt regularly, I help stop it from piling up.

### Can you describe your experience with CI/CD pipelines
I have experience setting up CI/CD pipelines with Git and Docker. I use Git for version control to manage changes and Docker to create consistent applications. This helps with automated testing, easier builds, and smooth deployments, keeping code quality high and efficient.

### How do you ensure system scalability
I design systems with scalability in mind, utilizing microservices, load balancers, and horizontal scaling. I also monitor performance metrics to identify bottlenecks early.

### What strategies do you use for error handling in Golang
I use Go's error handling to manage errors in an organized way and log them with helpful context. For serious issues, I use panic and recover to keep the system stable.

### Describe your experience with testing in Golang
I write unit and integration tests using Go's testing package. I believe in test-driven development (TDD) to ensure code reliability and maintainability.

### How do you approach API documentation
I use tools like Swagger to generate and maintain API documentation, ensuring it is always up-to-date.

### How do you see the role of a backend developer evolving
I notice that backend developers are starting to handle more DevOps and cloud tasks because development and operations are becoming more connected. Focusing on automation and continuous delivery will be important.