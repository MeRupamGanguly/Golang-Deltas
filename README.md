# Golang-Deltas
## Module 1: GoLang
### Introduce Yourself: 
I have a background in B.Tech IT and passout of	 2020. After college, I joined Sensibol as a Golang Backend Developer. My primary role was developing microservices using Golang, AWS, MongoDB, and Redis. I resigned in April to take care of my father's cancer treatment. Now he is well, I am looking for new opportunities.

### What projects do you woked on Sensibol, describe: 
I worked on PDL (Phonographic Digital Limited), a music distribution and royalty management platform. 

Additionally, I worked on Singshala, which is similar to TikTok but with extra feature of analyzing the audio components of videos and providing rankings based on the analysis. 

Both projects used Golang, MongoDB, Redis, AWS S3, SQS, SNS, Lambda, etc. Both had microservices architectures. PDL transitioned from a domain-driven approach to an event-driven one, while Singshala was domain-driven and complete. 

Since Sensibol is very conservative regarding hiring, only three backend developers worked on both the projects, so my involvement was extensive.

### Microservices vs Monolith:
Microservices are better for large projects where scaling and almost zero downtime are required. Bug fixing and maintaining the codebase are easier. A disadvantage of microservices can be inter-service network calls.

### Authentication vs  Authorization:
Authentication is about verifying identity. Users typically provide credentials such as a username and password, biometric data (like fingerprints or facial recognition), or security tokens. The system checks these credentials against stored data to confirm their validity.

Authorization is about granting permissions based on that verified identity. Once authenticated, the system checks the user’s permissions and roles to determine what they can access or modify. Authorization policies might be based on roles, permissions, or other criteria.

### Golang Garbage Collection:
Golang uses automatic garbage collection to manage memory. Developers do not need to allocate or deallocate memory manually, which reduces memory-related errors.

### Pointer:
Pointer is a Memory address of the Variable or Struct or Function or any kind of Data type.
Pointer can Referenced and Dereferenced by * &
&a gets the memory address of the variable a.
*p accesses the value stored at the memory address pointed to by p.

#1256 Memory address of a
a:=9
p:=&a // P contains the Memory adress of a // Referencing
*p // Dereferencing

### Goroutine vs Thread:
Goroutines are designed for concurrency, meaning multiple tasks can run using context switching. Threads are designed for parallelism, meaning multiple tasks can run simultaneously on multiple CPU cores.

Goroutines have a dynamic stack size and are managed by the Go runtime. Threads have a fixed-size stack and are managed by the OS kernel.

### What is Closure in golang:
A closure is a special type of anonymous function that can use variables, that declared outside of the function. Closures treat functions as values, allowing us to assign functions to variables, pass functions as arguments, and return functions from other functions. 

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

### Interfaces in golang:
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
### Panic Defer Recover combo:
panic is use to cause a Runtime Error and Stop the execution.
When a function return or panicking then Defer blocks are called according to Last in First out manner, the last defer will execute first.
Recover is use to regain the execution from a panicking situation and handle it properly then stop execution. Recover is useful for close any connection like db and websockets etc.
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
### Array vs Slice: 
Array can not Grow and Shrink dynamically at runtime, Slice can. Slice is just references to an existing array of a fixed length.

### Method Dispatching:
golang use Receiver function for method dispatching and has 2 way to dispatch methods at runtime.

Pointer receiver function: As obj is refrence of the Struct so any modification inside the function will affect the original Struct. More memory-efficient and can result in faster execution, especially for large structs.
```go
func (obj *class_name)method_name(argument int) (returns_name bool){
    //body
}
```
Value receiver function: As obj is copy of the Struct so any modification inside the function will not affect the original Struct. 
```go
func (obj class_name)method_name(argument int) (returns_name bool){
    //body
}
```
### Concurency Primitives:
Concurency Primitives are tools that are provided by any programming languages to handle execution behaviors of Concurent tasks.

In golang we have Mutex, Semaphore, Channels as concurency primitives.

Mutex is used to protect shared resources from being accessed by multiple threads simultaneously.

Semaphore is used to protect shared pool of resources from being accessed by multiple threads simultaneously. Semaphore is a Counter which start from Number of Reosurces. When one thread using the reosurces Semaphore decremented by 1. If semaphore value is 0 then thread will wait untils its value greater than 0. When one thread done with the resources then Semaphore incremented by 1.

Channel is used to communicate via sending and receiving data and provide synchronisation between multiple gorountines. If channel have a value then execution blocked until reader reads from the channel.
Channel can be buffered, allowing goroutines to send multiple values without blocking until the buffer is full. 

Waitgroup is used when we want the function should wait until goroutines complete its task.
Waitgroup has Add() function which increments the wait-counter for each goroutine.
Wait() is used for wait until wait-counter became zero.
Done() decrement wait-counter and it called when goroutine complete its task.

### Map Synchronisation:
In golang if multiple goroutines try to acess map at same time, then the operations leads to Panic for RACE or DEADLOCK (fatal error: concurrent map read and map write).
So we need proper codes for handeling Map.
We use MUTEX for LOCK and UNLOCK the Map operations like Read and Write. 
```go
func producer(m *map[int]string, wg *sync.WaitGroup, mu *sync.RWMutex) {
	vm := *m
	for i := 0; i < 5; i++ {
		mu.Lock()
		vm[i] = fmt.Sprint("$", i)
		mu.Unlock()
	}
	m = &vm
	wg.Done()
}
func consumer(m *map[int]string, wg *sync.WaitGroup, mu *sync.RWMutex) {
	vm := *m
	for i := 0; i < 5; i++ {
		mu.RLock()
		fmt.Println(vm[i])
		mu.RUnlock()
	}
	wg.Done()
}
func main() {
	m := make(map[int]string)
	m[0] = "1234"
	m[3] = "2345"
	wg := sync.WaitGroup{}
	mu := sync.RWMutex{}
	for i := 0; i < 5; i++ {
		wg.Add(2)
		go producer(&m, &wg, &mu)
		go consumer(&m, &wg, &mu)
	}
	wg.Wait()
}
```
### Describe Channel comunication with task distributions:
Lets imagine we have a number n, we have to find 0 to n numbers are prime or not.
```go
func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
func primeHelper(a []int, ch chan<- map[int]bool, wg *sync.WaitGroup) {
	time.Sleep(time.Second)
	defer wg.Done()
	m := make(map[int]bool)
	for i := range a {
		m[a[i]] = isPrime(a[i])
	}
	ch <- m
}
func main() {
	startTime := time.Now()
	var wg sync.WaitGroup
	n := 12
	arr := []int{}
	for i := 0; i < n; i++ {
		arr = append(arr, i)
	}
	length := len(arr)
	goroutines := 4
	part := length / goroutines
	ch := make(chan map[int]bool, goroutines)
	ma := make(map[int]bool)
	for i := 0; i < goroutines; i++ {
		wg.Add(1)
		s := i * part
		e := s + part
		if e > length {
			e = length
		}
		go primeHelper(arr[s:e], ch, &wg)
	}
	wg.Wait()
	close(ch)
	for i := range ch {
		for k, v := range i {
			ma[k] = v
		}
	}
	fmt.Println(ma)
	fmt.Println("Time Taken: ", time.Since(startTime))
}
```

### Select Statement:
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

### SOLID Principles:
SOLID priciples are guidelines for designing Code base that are easy to understand maintain and extend over time.

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
	Count   int  `json:"Count"`
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
### URL DECODER Generic:
```go
package main

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"reflect"
	"strconv"

	"github.com/gorilla/mux"
)

// Domain Layer --------------------------
type Counter struct {
	Count int `json:"count"`
}

func NewService() *Counter {
	return &Counter{}
}

type service interface {
	Get(ctx context.Context, id string) (int, error)
	Add(ctx context.Context, c Counter) error
}

// Transport Layer  --------------------------

// queryDecoder extracts URL query parameters into the given struct
func queryDecoder(r *http.Request, inp interface{}) error {
	q := r.URL.Query()
	inpVal := reflect.ValueOf(inp)
	if inpVal.Kind() != reflect.Ptr || inpVal.IsNil() {
		return errors.New("inp must be a non-nil pointer to a struct")
	}

	inpVal = inpVal.Elem()
	if inpVal.Kind() != reflect.Struct {
		return errors.New("inp must be a pointer to a struct")
	}

	inpType := inpVal.Type()
	for i := 0; i < inpType.NumField(); i++ {
		field := inpType.Field(i)
		tag := field.Tag.Get("url")
		if tag == "" {
			continue
		}

		values, ok := q[tag]
		if ok && len(values) > 0 {
			value := values[0]
			switch field.Type.Kind() {
			case reflect.Int:
				d, err := strconv.Atoi(value)
				if err != nil {
					return err
				}
				inpVal.Field(i).SetInt(int64(d))
			case reflect.String:
				inpVal.Field(i).SetString(value)
			case reflect.Float32:
				d, err := strconv.ParseFloat(value, 32)
				if err != nil {
					return err
				}
				inpVal.Field(i).SetFloat(d)
			case reflect.Float64:
				d, err := strconv.ParseFloat(value, 64)
				if err != nil {
					return err
				}
				inpVal.Field(i).SetFloat(d)
			case reflect.Bool:
				d, err := strconv.ParseBool(value)
				if err != nil {
					return err
				}
				inpVal.Field(i).SetBool(d)
			default:
				return errors.New("unsupported field type: " + field.Type.Kind().String())
			}
		}
	}
	return nil
}

type transport struct {
	svc service
}

func NewTransport(svc service) *transport {
	return &transport{
		svc: svc,
	}
}

type getRequest struct {
	Id string `url:"id"`
}

type getResponse struct {
	Count   int  `json:"count"`
	Success bool `json:"success"`
}
type addRequest struct {
	Counter Counter `json:"counter"`
}

type addResponse struct {
	Success bool `json:"success"`
}

// Get handler retrieves data based on the ID
func (trans *transport) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	req := &getRequest{}
	if err := queryDecoder(r, req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id := req.Id
	data, err := trans.svc.Get(context.Background(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(getResponse{Success: true, Count: data})
}
func (trans *transport) Add(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	req := addRequest{}
	json.NewDecoder(r.Body).Decode(&req)
	err := trans.svc.Add(context.Background(), req.Counter)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(addResponse{Success: true})
}

// Service Layer  --------------------------

// Get retrieves the count value based on the ID (for simplicity, ID is not used)
func (svc *Counter) Get(ctx context.Context, id string) (int, error) {
	return svc.Count, nil
}

// Add increments the counter
func (svc *Counter) Add(ctx context.Context, c Counter) error {
	svc.Count += c.Count
	return nil
}
//  --------------------------
func main() {
	r := mux.NewRouter()
	svc := NewService()
	ts := NewTransport(svc)
	r.HandleFunc("/get", ts.Get).Methods("GET")
	r.HandleFunc("/add", ts.Add).Methods("POST")
	http.ListenAndServe(":5000", r)
}
```

## Module 2: Logic Building
### Arrays
- Find the maximum/minimum element in an array.
```go
func findMaxMinOfArray(arr []int) (max, min int) {
	max = 0
	min = arr[0]
	for i := 0; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
		if arr[i] < min {
			min = arr[i]
		}
	}
	return max, min
}
```
- Find the index of a specific element in an array.
```go
func findIndex(a []int, inp int) int {
	for i := range a {
		if a[i] == inp {
			return i
		}
	}
	return -1
}
```
- Reverse an array.
```go
func reverseArray(arr []int) []int {
	l := 0
	r := len(arr) - 1
	for l < r {
		temp := arr[l]
		arr[l] = arr[r]
		arr[r] = temp
		l++
		r--
	}
	return arr
}
```
- Check if an array is sorted.
```go
func arraySortCheck(a []int) bool {
	for i := 1; i < len(a); i++ {
		if a[i-1] > a[i] {
			return false
		}
	}
	return true
}
```
- Find the second largest element in an array.
```go
func secondLargest(a []int) int {
	max := 0
	secondMax := 0
	for i := range a {
		if a[i] > max {
			max = a[i]

		} else {
			if a[i] > secondMax {
				secondMax = a[i]
			}
		}
	}
	return secondMax
}
```
- Remove duplicates from an array.
```go
func removeDup(arr []int) []int {
	l := 0
	for i := 1; i < len(arr); i++ {
		if arr[i] != arr[l] {
			l++
			arr[l] = arr[i]
		}
	}
	return arr[:l+1]
}
func removeDuplicates(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		if arr[i] == arr[i+1] {
			temp := arr[i+1:]
			arr = arr[:i]
			arr = append(arr, temp...)
		}
	}
	return arr
}
```
- Find the average of elements in an array.
```go
func average(a []int) int {
	sum := 0
	for i := range a {
		sum += a[i]
	}
	return sum / len(a)
}
```
- Move all zeroes to the end of the array while maintaining the order of non-zero elements.
```go
func moveZeroes(arr []int) []int {
	outArr := make([]int, len(arr))
	outIndex := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] != 0 {
			outArr[outIndex] = arr[i]
			outIndex++
		}
	}
	return outArr
}
```
- Rotate an array to the right by k steps.
```go
func rotateKStepsRight(a []int, k int) []int {
	part := a[:k]
	a = a[k:]
	a = append(a, part...)
	return a
}
```
- Two Sum
```go
func twoSum(arr []int, target int) (int, int) {
	for i := 1; i < len(arr); i++ { // Adjacent Sum check
		if arr[i-1]+arr[i] == target {
			return i, i + 1
		}
	}
	for r := len(arr) - 1; r > 0; r-- { // Check every sequences
		l := 0
		for l < r {
			if arr[l]+arr[r] == target {
				return l, r
			}
			l++
		}
	}
	return 0, 0
}
func twoSumOpt(arr []int, target int) (int, int) {
	hmap := make(map[int]int)
	for i := range arr {
		complement := target - arr[i]
		k, v := hmap[complement]
		if v {
			return i, k
		} else {
			hmap[arr[i]] = i
		}
	}
	return -1, -1
}
```
- Product of Array Except Self
```go
func productOfArrayExceptSelf(a []int) []int {
	var outArr []int
	for i := 0; i < len(a); i++ {
		prod := 1
		for j := 0; j < len(a); j++ {
			if j == i {
				continue
			}
			prod = prod * a[j]
		}
		outArr = append(outArr, prod)
	}
	return outArr
}
```
- Find the Missing Number
```go
func missingNumber(a []int) int {
	length := len(a) + 1
	actualSum := length * (length + 1) / 2
	sum := 0
	for i := range a {
		sum += a[i]
	}
	missingNum := actualSum - sum
	return missingNum
}
```
- Maximum Subarray

### Strings
- Reverse a string.
```go
func reverseString(s string) string {
	arr := []rune(s)
	l := 0
	r := len(arr) - 1
	for l < r {
		t := arr[l]
		arr[l] = arr[r]
		arr[r] = t
		l++
		r--
	}
	return string(arr)
}
```
- Check if a string is a palindrome.
```go
func palindrome(s string) bool {
	arr := []rune(s)
	l := 0
	r := len(arr) - 1
	for l < r {
		if arr[l] != arr[r] {
			return false
		}
		l++
		r--
	}
	return true
}
```
- Count the number of vowels in a string.
```go
func noOfVowels(s string) int {
	arr := []rune(strings.ToLower(s))
	count := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] == rune('a') || arr[i] == rune('e') || arr[i] == rune('i') || arr[i] == rune('o') || arr[i] == rune('u') {
			count++
		}
	}
	return count
}
```
- Find the length of the longest substring without repeating characters.
```go
func longestSubstringWithoutRepeat(s string) int {
	arr := []rune(s)
	uni := make(map[rune]int)
	count := 0
	for i := 0; i < len(arr); i++ {
		_, ok := uni[arr[i]]
		if ok {
			return count
		}
		uni[arr[i]] = 0
		count++
	}
	return count
}
```
- Check if two strings are anagrams.
```go
func anagram(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	arr1 := []rune(s1)
	arr2 := []rune(s2)
	m1 := make(map[rune]int)
	m2 := make(map[rune]int)

	for i := 0; i < len(arr1); i++ {
		v, ok := m1[arr1[i]]
		if ok {

			m1[arr1[i]] = v + 1
		} else {
			m1[arr1[i]] = 0
		}
		v, ok = m2[arr2[i]]
		if ok {

			m2[arr2[i]] = v + 1
		} else {
			m2[arr2[i]] = 0
		}
	}
	for k, v1 := range m1 {
		v2, ok := m2[k]
		if ok && v1 != v2 {
			return false
		}
	}
	return true
}
```
- Find the first non-repeated character in a string.
```go
func nonRepeatingFirstChar(s string) string {
	arr := []rune(s)
	m := make(map[rune]int)
	for i := 0; i < len(arr); i++ {
		v, ok := m[arr[i]]
		if ok {
			m[arr[i]] = v + 1
		} else {
			m[arr[i]] = 0
		}
	}
	for _, v := range s {
		if m[v] == 0 {
			return string(v)
		}
	}
	return ""
}
```
- Find all permutations of a string.
- Check if a string contains all unique characters.
```go
func checkUnique(s string) bool {
	arr := []rune(s)
	m := make(map[rune]int)
	for i := range arr {
		v, ok := m[arr[i]]
		if ok {
			m[arr[i]] = v + 1
			return false
		} else {
			m[arr[i]] = 0
		}
	}
	return true
}
```
- Longest Common Prefix

### Linked Lists
- Print all elements in a linked list.
```go
type Node struct {
	data  int
	right *Node
}

type list struct {
	Nodes *Node
}

func NewList() *list {
	return &list{}
}
func (l *list) Add(d int) {
	node := Node{
		data: d,
	}
	if l.Nodes == nil {
		l.Nodes = &node
		return
	}
	c := l.Nodes
	for c.right != nil {
		c = c.right
	}
	c.right = &node
}
func (l *list) Remove(n int) {
	if l.Nodes == nil {
		return
	}
	c := l.Nodes
	for i := 0; i < n-1; i++ {
		if c.right != nil {
			c = c.right
		}
	}
	if c.right.right != nil {
		c.right = c.right.right
	}
}
func (l *list) print() {
	c := l.Nodes
	for c.right != nil {
		fmt.Println(c.data)
		c = c.right
	}
	fmt.Println(c.data)
	fmt.Println("------------------------")
}
func (l *list) NthAdd(d int, n int) {
	node := Node{
		data: d,
	}
	if l.Nodes == nil {
		l.Nodes = &node
	}
	if n == 0 {
		node.right = l.Nodes
		l.Nodes = &node
	}

	c := l.Nodes
	for i := 0; i < n-1; i++ {
		if c.right != nil {
			c = c.right
		}
	}
	node.right = c.right
	c.right = &node
}
func main() {
	l := NewList()
	l.Add(2)
	l.Add(4)
	l.Add(6)
	l.Add(9)
	l.Add(12)
	l.Add(16)
	l.print()
	l.NthAdd(17, 8)
	l.print()
	l.NthAdd(14, 3)
	l.print()
	l.Remove(3)
	l.print()
	l.Remove(2)
	l.print()
}
```
- Find a node by its value.
- Reverse a linked list.
- Check if a linked list is cyclic.
- Find the middle element of a linked list.
- Merge two sorted linked lists.
- Detect a Cycle in a Linked List
- Remove Nth Node From End of List
### Stacks Queues
- Implement a stack using an array.
```go
type stack struct {
	d []string
}

func NewStack(d []string) *stack {
	return &stack{
		d: d,
	}
}
func (s *stack) push(data string) {
	s.d = append(s.d, data)
}
func (s *stack) pop() string {
	if len(s.d) < 1 {
		return ""
	}
	t := s.d[len(s.d)-1]
	s.d = s.d[:len(s.d)-1]
	return t
}
```
- Valid Parentheses
```go
func validParenthesesCheck(arr []string) bool {
	o := NewStack([]string{})
	m := make(map[string]string)
	m["{"] = "}"
	m["["] = "]"
	m["("] = ")"
	m["}"] = "{"
	m["]"] = "["
	m[")"] = "("
	for i := 0; i < len(arr); i++ {
		char := arr[i]
		if len(o.d) > 0 {
			oelem := o.pop()
			if m[char] != oelem {
				o.push(oelem)
				o.push(char)
			}
		} else {
			o.push(char)
		}
	}
	if len(o.d) == 0 {
		return true
	} else {
		return false
	}
}
```
- Implement a stack using a linked list.
- Check for balanced parentheses in an expression.
- Sort a stack using another stack.
- Find the minimum element in a stack.
- Perform stack operations (push, pop, peek).
- Evaluate a postfix expression using a stack.
- Evaluate Reverse Polish Notation
- Implement a queue using an array.
- Implement a queue using a linked list.
- Implement a circular queue.
- Implement a priority queue.
### Sorting and Searching
- Implement Bubble Sort.
- Implement Selection Sort.
- Implement Insertion Sort.
- Find an element in a sorted array using Binary Search.
- Count the number of occurrences of an element in a sorted array.
- Perform Linear Search on an unsorted array.
- Sort an array using Merge Sort.
- Sort an array using Quick Sort.
### DP
- Subset Sum
- Climbing Stairs
- Longest Common Subsequence
- Coin Change
- Longest Common Prefix




## Module 3: Databases

Horizontal Scaling : Distributing data and workload across multiple servers.

Sharding: Dividing the dataset into smaller, manageable parts called shards and distributing these shards across multiple servers. Each server handles a subset of the data.

MongoDB is a NoSql Database which is used for Concurent operations(multi-document ACID transactions) and Horizontal scalability.

PostgreSQL is ideal for applications requiring strong consistency, complex joins, and analytical queries, such as financial systems, CRM applications, and data warehousing.

Redis is an in-memory data structure store known for its exceptional read write speed and simplicity.

MongoDB and Redis excel in horizontal scaling, while PostgreSQL supports vertical scaling better. The ability to scale horizontally can directly influence OPS(Operation Per Second) performance in distributed and large-scale applications.

### MongoDB

```bash
docker run -it --name ubuntu -p 3000:3000 ubuntu:20.04
```
```bash
rupam@opulence:~$ docker ps -a
CONTAINER ID   IMAGE          COMMAND       CREATED          STATUS          PORTS                                       NAMES
62c2825a182f   ubuntu:20.04   "/bin/bash"   26 seconds ago   Up 25 seconds   0.0.0.0:3000->3000/tcp, :::3000->3000/tcp   ubuntu
```
```bash
apt update
apt install mongodb
apt install wget
wget https://golang.org/dl/go1.23.1.linux-amd64.tar.gz
apt install nano
tar -xvf go1.23.1.linux-amd64.tar.gz
mv go /usr/local
nano .bashr
export GOROOT=/usr/local/go
export GOPATH=$HOME/go
export PATH=$GOPATH/bin:$GOROOT/bin:$PATH
go version
nano main.go
go mod init deltas
go mod tidy
go run main.go 
mongod
mongo
show dbs
use MyDatabase
db.Album.find({}).pretty()
exit
exit
docker start ubuntu
docker exec -it ubuntu /bin/bash
```

DB:-
Album
```json
{
	"_id" : "66e0502252c36ebe41835851",
	"name" : "A1",
	"label_id" : "66e0502252c36ebe41835850",
	"releases" : [
		{
			"platform_id" : "66e0502252c36ebe4183584b",
			"status" : "Pending",
			"artists" : [
				{
					"_id" : "66e0502252c36ebe4183584c",
					"name" : "Artist1"
				},
				{
					"_id" : "66e0502252c36ebe4183584d",
					"name" : "Artist3"
				}
			]
		},
		{
			"platform_id" : "66e0502252c36ebe4183584e",
			"status" : "Pending",
			"artists" : [
				{
					"_id" : "66e0502252c36ebe4183584f",
					"name" : "King"
				}
			]
		}
	]
}
{
	"_id" : "66e0502252c36ebe41835858",
	"name" : "A2",
	"label_id" : "66e0502252c36ebe41835857",
	"releases" : [
		{
			"platform_id" : "66e0502252c36ebe41835852",
			"status" : "Pending",
			"artists" : [
				{
					"_id" : "66e0502252c36ebe41835853",
					"name" : "Artist4"
				},
				{
					"_id" : "66e0502252c36ebe41835854",
					"name" : "Artist5"
				}
			]
		},
		{
			"platform_id" : "66e0502252c36ebe41835855",
			"status" : "Pending",
			"artists" : [
				{
					"_id" : "66e0502252c36ebe41835856",
					"name" : "Artist3"
				}
			]
		}
	]
}

```

Song
```json
{
	"_id" : "66e0502252c36ebe41835859",
	"name" : "A1 Song 1",
	"album_id" : "66e0502252c36ebe41835851"
}
{
	"_id" : "66e0502252c36ebe4183585a",
	"name" : "A1 Song 2",
	"album_id" : "66e0502252c36ebe41835851"
}
{
	"_id" : "66e0502252c36ebe4183585b",
	"name" : "A2 Song 1",
	"album_id" : "66e0502252c36ebe41835858"
}

```

#### Find by name:
db.Album.find({"name":"A2"}).pretty()

#### Find by PlatformId:
db.Album.find({"release.platform_id":"66e0502252c36ebe41835855"}).pretty()

To control which fields are included in the result, we can use projection.
```json
db.Album.find({"releases.platform_id":"66e0502252c36ebe41835852"},{"releases.platform_id":1,"releases.status":1})
```

#### Aggregate:
In aggregate we have multiple Stage operations like Match, Group, Sort, Project, Limit etc. We pass one Satge Filtered data to Another stage as Input. then current Stage performe its operation and pass the result to next satge as input.

$match: This stage filters documents in the collection. It’s similar to a find query but used in the aggregation pipeline.

$project: This stage reshapes each document. It allows you to include or exclude fields or create new ones.



$filter: This operator creates a new array containing only those elements of the input array that match a specified condition.

input: "$releases": This specifies the array to filter, which is the releases field of each document.

as: "release": This defines a variable name (release) that represents each element of the releases array during filtering.

cond: This specifies the condition for filtering. Only elements that match this condition are included in the output array.
This condition checks if the platform_id field of the release object equals

$: This denotes an aggregation operator or field. For example, $sum is an aggregation operator, and $fieldName accesses the value of fieldName in the document.

$$: This denotes a variable that’s scoped within the aggregation pipeline. These variables are typically used in expressions to refer to values from specific stages, particularly when dealing with operators like $map, $filter, $reduce, etc.


```json
db.Album.aggregate([
    {
        $match: {
            "releases.platform_id": "66e0502252c36ebe41835855"
        }
    },
    {
        $project: {
            releases: { #This specifies that we’re working with the releases field in the documents.
                $filter: {
                    input: "$releases",
                    as: "release",
                    cond: {
                        $eq: ["$$release.platform_id", "66e0502252c36ebe41835855"]
                    }
                }
            }
        }
    }
]).pretty()
{
	"_id" : "66e0502252c36ebe41835858",
	"releases" : [
		{
			"platform_id" : "66e0502252c36ebe41835855",
			"status" : "Pending",
			"artists" : [
				{
					"_id" : "66e0502252c36ebe41835856",
					"name" : "Artist3"
				}
			]
		}
	]
}
```

#### Dot Operator:
First, we need to locate the documents where the releases array contains an item with a specific platform_id. It filters documents where at least one releases item has the platform_id equal to "66e0502252c36ebe41835855".
- `db.Album.find({"releases.platform_id":"66e0502252c36ebe41835855"}).pretty()`
```json
{
	"_id" : "66e0502252c36ebe41835858",
	"name" : "A2",
	"label_id" : "66e0502252c36ebe41835857",
	"releases" : [
		{
			"platform_id" : "66e0502252c36ebe41835852",
			"status" : "Pending",
			"artists" : [
				{
					"_id" : "66e0502252c36ebe41835853",
					"name" : "Artist4"
				},
				{
					"_id" : "66e0502252c36ebe41835854",
					"name" : "Artist5"
				}
			]
		},
		{
			"platform_id" : "66e0502252c36ebe41835855",
			"status" : "Pending",
			"artists" : [
				{
					"_id" : "66e0502252c36ebe41835856",
					"name" : "Artist3"
				}
			]
		}
	]
}
```
#### Projection:
To control which fields are included in the result, you use projection. In this case, you want to include only the releases that match the platform_id and exclude the rest.

- `db.Album.find({"releases.platform_id":"66e0502252c36ebe41835855"},{"releases.platform_id":1}).pretty()`
```json
{
	"_id" : "66e0502252c36ebe41835858",
	"releases" : [
		{
			"platform_id" : "66e0502252c36ebe41835852"
		},
		{
			"platform_id" : "66e0502252c36ebe41835855"
		}
	]
}
```
#### Aggregation
However, simple projection alone won’t filter out the unwanted releases items within the array. For that, we use the aggregation framework.

Aggregation Pipeline Steps:
- Match Stage: Filters documents to include only those documents which match the criteria. 
`{$match:{"releases.platform_id":"66e0502252c36ebe41835855"}}`
	This filters the documents to include only those where the releases array contains at least one item with the platform_id equal to "66e0502252c36ebe41835855".
- Project Stage: Reshapes each document by including only the relevant parts.
```json
{
    $project: {
        releases: {
            $filter: { // A special operator to filter elements from an array.
                input: "$releases",   // Access the `releases` array // specifies that we are working with the releases array.
                as: "release",   // Temporary name for each array element // gives each item in the releases array a temporary name (release).
                cond: { $eq: ["$$release.platform_id", "66e0502252c36ebe41835855"] } // Condition for filtering 
				// $$release to refer to the current element in the array. 
				// The $$ prefix denotes that release is a variable, 
				// and you use it to access fields within each element of the array.
            }
        }
    }
}	
```
- `db.Album.aggregate([{$match:{"releases.platform_id":"66e0502252c36ebe41835855"}},{$project:{releases:{$filter:{input:"$releases",as:"release", cond:{$eq:["$$release.platform_id","66e0502252c36ebe41835855"]}}}}}]).pretty()`
```json
{
	"_id" : "66e0502252c36ebe41835858",
	"releases" : [
		{
			"platform_id" : "66e0502252c36ebe41835855",
			"status" : "Pending",
			"artists" : [
				{
					"_id" : "66e0502252c36ebe41835856",
					"name" : "Artist3"
				}
			]
		}
	]
}
```
#### Update Status from Pending to Approved : Update specific elements within nested arrays.
```bash
db.Album.find({"releases.status":"Pending"}).pretty()
```
`$` used to refer to a specific element within an array that matches the query criteria. 

```bash
db.Album.updateMany(
{"releases.status":"Pending"},
{$set:{"releases.$[release].status":"Approved"}},
{arrayFilters:[{"release.status":"Pending"}]}
)
```
`arrayFilter` determine which element of the array should updated. Each object within the releases array is an element. The element(release object) where "status": "Pending" is the one targeted by the array filter.
`$[release]` in the update statement, refers to the elements that match the conditions written within the arrayFilters.

#### Update Artist Name from Artis2 to King : Update specific elements within nested arrays.
In this example, releases is an array that contains objects, and each object has an artists array.

In MongoDB, when using arrayFilters in an update operation that targets nested arrays, each filter within the arrayFilters is used to specify conditions for different levels of nested arrays. 

```bash
db.Album.updateMany(
{"releases.artists.name":"Artist2"},
{$set:{"releases.$[release].artists.$[artist].name":"King"}},
{arrayFilters:[{"release.artists.name":"Artist2"},{"artist.name":"Artist2"}]}
)
```

#### Elemmatch vs Dot operator:
- `$elemMatch` is typically used in queries to find documents where at least one element in an array meets all the criteria. So it returns documents where condition-1 and condition-2 must matched with same array-element of same document.

- `.` The dot operator is typically used in queries to find documents where each element of the array, meets one condition individually. So it returns documents where condition-1 matched with array-element-1 and condition-2 matched with array-element-2 of same document.

```json
[
  {
    "_id": "6671d6fa65fcd889d2dc222b",
    "name": "Company 3",
    "employee": [
      { "name": "Employee 0", "age": 26, "is_active": true, 
	  "roles": ["IT", "Finance", "Law"] },
      {  "name": "Employee 40", "age": 29, "is_active": true, 
	  "roles": ["Sell"] }
    ]
  },
  {
    "_id": "6671d6fa65fcd889d2dc222a",
    "name": "Company 4",
    "employee": [
      {  "name": "Employee 3", "age": 36, "is_active": true, 
	  "roles": ["IT", "Finance", "Law"] },
      {  "name": "Employee 80", "age": 41, "is_active": true, 
	  "roles": ["Finance"] },
      {  "name": "Employee 20", "age": 37, "is_active": false, 
	  "roles": ["Finance"] }
    ]
  },
    {
    "_id": "6671d6fa65fcd889d2dc2255",
    "name": "Company 5",
    "employee": [
      { "name": "Employee 8", "age": 38, "is_active": true, 
	  "roles": [] },
      {  "name": "Employee 11", "age": 26, "is_active": true, 
	  "roles": ["Law"] }
    ]
  }
]
```
```go
db.Company.find({"employee.age":{$gt:30},"employee.is_active":true, "employee.roles":{$in:["Law"]}});
// Result: Company 5,Company 4 
// Why: Employee 8 has age 38 and role nil. 
// but Employee 11 has age 26 and role Law. 
// This document satisfy two condition with two different array element.

db.Company.find({"employee":{$elemMatch:{"age":{$gt:30},"is_active":true,"roles":{$in:["Law"]}}}});
// Result: Company 4 
// Why: Employee 8 has age 38 and role nil. 
// but Employee 11 has age 26 and role Law. 
// This document does not satisfy two condition with same array element. 
// So Company 5 not satisfied the Criteria.
```

#### Scalability:
- Sharding enables horizontal scaling.
- Each shard contains a subset of the data, distributed based on a shard key.
- MongoDB uses config servers to store metadata like mapping between shards and shard keys. Using config servers, MongoDB routers (`mongos` instances) send queries to the appropriate shards based on the shard key.
- MongoDB routers - `mongos` instances also manage load balancing across the shards. They distribute incoming queries evenly across shards.
- Adding more `mongos` instances can improve the throughput and scalability.
- MongoDB uses replica sets to provide redundancy and automatic failover.
- Each replica set consists of multiple nodes (typically three or more): one primary node for read and write operations and secondary nodes that replicate data from the primary. If the primary node fails, a new primary is elected from the remaining nodes in the replica set, ensuring continuous availability.
- MongoDB’s oplog is a collection that records all write operations (inserts, updates, deletes) in the order they occur.

#### Operations
- `$eq`: Matches values that are equal to a specified value.
- `$ne`: Matches all values that are not equal to a specified value.
- `$gt`, $gte, $lt, $lte: Greater than, greater than or equal to, less than, less than or equal to, respectively.
- `$in`: Matches any of the values specified in an array.
- `$nin`: Matches none of the values specified in an array.
- `$set`: Sets the value of a field in a document.
- `$unset`: Removes the specified field from a document.
- `$inc`: Increments the value of the field by a specified amount.
- `$push`: Adds an element to an array.
- `$addToSet`: Adds elements to an array only if they do not already exist.
- `$pull`: Removes all instances of a value from an array.
- `$rename`: Renames a field.
- `$project`: Reshapes documents by including, excluding, or transforming fields.
- `$elemMatch`: Matches documents that contain an array field with `at least one element` that matches `all` the specified query criteria.
- `$all`: Matches arrays that contain all elements specified in the query.
- `$size`: Matches arrays with a specific number of elements.
- `$sort`: Orders documents.
- `$limit`: Limits the number of documents.
- `$skip`: Skips a specified number of documents.
- `$match`: Filters documents.
- `$unwind`: Imagine you have a collection where each document contains an array field. When you apply $unwind to an array field in a document, MongoDB will create separate documents for each element of the array, where Each new document will have the same values for all other fields except the one array field.][]

```json
{
  "_id": 1,
  "name": "Product ABC",
  "tags": ["electronics", "smartphone", "tech"]
}
```
```go
db.products.aggregate([
  { $unwind: "$tags" }
]);
```
```json
{ "_id": 1, "name": "Product ABC", "tags": "electronics" }
{ "_id": 1, "name": "Product ABC", "tags": "smartphone" }
{ "_id": 1, "name": "Product ABC", "tags": "tech" }
```

`$group`: Groups documents by a specified identifier.

```json
[
  { "_id": 1, "item": "book", "price": 10, "quantity": 2,"tags": ["school", "notebook"]  },
  { "_id": 2, "item": "pen", "price": 5, "quantity": 5, "tags": ["school", "stationery"] },
  { "_id": 3, "item": "book", "price": 10, "quantity": 1,"tags": ["school", "stationery"] },
  { "_id": 4, "item": "eraser", "price": 2, "quantity": 3, "tags": ["school", "stationery"] }
]
```
```go
db.sales.aggregate([
  {
    $group: {
      _id: "$item",  // Group by the "item" field
      totalSales: { $sum: { $multiply: ["$price", "$quantity"] } }  // Calculate total sales for each item
    }
  }
]);
```
```json
[
  { "_id": "book", "totalSales": 30 },
  { "_id": "pen", "totalSales": 25 },
  { "_id": "eraser", "totalSales": 6 }
]
```
```python
db.sales.aggregate([
  { $match: { tags: "school" } }, #  Filters documents to pass only those that match the specified condition.
  { $group: { _id: "$item", totalAmount: { $sum: { $multiply: ["$price", "$quantity"] } } } }, #  Groups documents by a specified identifier expression and applies accumulator expressions. Operators like $sum, $multiply, etc., perform specific computations.
  { $sort: { totalAmount: -1 } }, #  Orders the documents.
  { $limit: 3 } # Limits the number of documents passed to the next stage.
]);
```


- `$lookup` is a stage in MongoDB aggregation, that allows you to perform a left outer join to retrieve documents from another collection and include them in your result set.

```json
// Orders Collection:
{ "_id": 1, "order_id": "A001", "product_id": 101, "quantity": 2 }
{ "_id": 2, "order_id": "A002", "product_id": 102, "quantity": 1 }
```
```json
// Products Collection:
{ "_id": 101, "name": "Laptop", "price": 1200 }
{ "_id": 102, "name": "Mouse", "price": 30 }
```
```go
db.orders.aggregate([
  {
    $lookup: {
      from: "products",
      localField: "product_id",
      foreignField: "_id",
      as: "product_details"
    }
  }
]);
```
```json
[
  {
    "_id": 1,
    "order_id": "A001",
    "product_id": 101,
    "quantity": 2,
    "product_details": [
      { "_id": 101, "name": "Laptop", "price": 1200 }
    ]
  },
  {
    "_id": 2,
    "order_id": "A002",
    "product_id": 102,
    "quantity": 1,
    "product_details": [
      { "_id": 102, "name": "Mouse", "price": 30 }
    ]
  }
]
```

- `$text`: Performs text search. 
Text search is case insensitive by default. Text search is also diacritic insensitive (e.g., "café" would match "cafe").
MongoDB calculates a relevance score (score) for each document based on the frequency and proximity of the search terms in the indexed fields.
Text searches in MongoDB require a text index on the field(s) you want to search.
```json
{
  "_id": ObjectId("60d02e9c25c156ae22df2b73"),
  "title": "The Catcher in the Rye",
  "author": "J.D. Salinger",
  "genre": "Fiction",
  "summary": "The Catcher in the Rye is a novel by J. D. Salinger, 
  partially published in serial form in 1945–1946 and as a novel in 
  1951. It was originally intended for adults but is often read by 
  adolescents for its themes of angst and alienation, and as a critique on superficiality in society."
}
```
```go
// Text searches in MongoDB require a text index on the field(s) you want to search.
db.books.createIndex({ summary: "text" });
```
This query searches for documents in the books collection where the summary field contains the word "adolescents".

- The `$text `operator specifies that this is a text search.
- `$search` is used to specify the search string. 
- This query not only finds documents matching "adolescents" but also sorts them by the relevance score in descending order.
```go
db.books.find(
  { $text: { $search: "adolescents" } },
  { score: { $meta: "textScore" } }
).sort({ score: { $meta: "textScore" } });
```

MongoDB supports various types of indexes for different query patterns and optimize performance:

Single Field Index: Indexes on a single field of a document.

Compound Index: Indexes on multiple fields within a document.

Multikey Index: Indexes on arrays of values (each value in the array is indexed).

Partial Indexing : Partial indexing involves creating an index on documents that satisfy a filter expression.  Indexes only those documents that match the filter, ignoring documents that do not meet the criteria. 
By indexing only a subset of documents, you reduce the index size compared to indexing the entire collection.  Queries that match the indexed subset of documents can benefit from faster query execution because MongoDB only needs to scan the indexed subset.
```go
db.books.createIndex(
  { price: 1 }, 
  // specifies that we are creating an ascending index on the price field.
  { partialFilterExpression: { available: true } } 
  // specifies that the index should only include documents where the available field is true.
);
```
```go
// This query can utilize our partial index on price where 
// available is true to efficiently retrieve and sort books that are currently available.
db.books.find({ available: true }).sort({ price: 1 });
```

Text Index: Special index type for performing full-text searches on string content.

Geospatial Index: Indexes for geospatial data, supporting queries that calculate geometries based on proximity.

Hashed Index: Indexes where MongoDB hashes the indexed field's values, typically used for sharding.



Assume we are working with this Schema
```go
type Status string

const (
	Pending  Status = "Pending"
	Approved Status = "Approved"
	Rejected Status = "Rejected"
)

type Artist struct {
	Id   string
	Name string
}
type Release struct {
	PlatformId string
	Status     Status
	Artists    []Artist
}
type Album struct {
	Id       string
	Name     string
	LabelId  string
	Releases []Release
}
type Song struct {
	Id      string
	Name    string
	AlbumId string
}

type Label struct {
	Id   string
	Name string
}
```
- Find Albums by Name:
```go
db.album.find({"Name":"Requested Name"})
```
```go
filter:=bson.M{"Name":name}
collection.FindOne(ctx,filter).Decode(&album)
```
- Find Albums by Status enum:
```go
db.album.find({"Releases.Status":"Approved"})
```
```go
filter:=bson.M{"Releases.Status":status}
cursor,err:=collection.Find(ctx,filter)
defer cursor.Close(ctx)
err = cursor.All(ctx,&albums)
```
- Find all artists associated with a specific album:
```go
db.album.find({"Id":"album id"},{"Releases.Artists":1})
```
```go
filter:=bson.M{"Id":albumId}
projection:=bson.M{"Releases.Artists":1}
err = collection.FindOne(ctx, filter, options.FindOne().SetProjection(projection)).Decode(&result)
```
- Find albums where at least one release has the status "Approved" and includes a specific artist (e.g., an artist with the name "Artist Name"):
```go
db.albums.find({
	"Releases":{
		"$elemmatch":{
			"Status":"Approved",
			"Artists.Name":"requested artist name"
		}
	}
})
```
```go
filter:=bson.M{
	"Releases":bson.M{
		"$elemmatch":bson.M{
			"Status":status,
			"Artists.Name":artistName
		},
	},
}
cursor, err := collection.Find(ctx, filter)
defer cursor.Close(ctx)
err = cursor.All(ctx, &albums)
```
- Count the number of albums associated with each label
```go
db.album.aggregate([
	{
		"$group":{
			"_id":"$LabelId",
			"count":{"$sum":1}
		}
	}
])
```
```go
pipeline := mongo.Pipeline{
    {{Key: "$group", Value: bson.D{{Key: "_id", Value: "$LabelId"}, {Key: "count", Value: bson.D{{Key: "$sum", Value: 1}}}}}},
}
cursor, err := collection.Aggregate(ctx, pipeline)
defer cursor.Close(ctx)
err = cursor.All(ctx, &results)
```
- Find albums with more than 3 releases:
```go
db.album.find({
	"Releases":{
		"$exists": true,
		"$not":{"$size":3}
	}
})
```
```go
pipeline := mongo.Pipeline{
    {{Key: "$match", Value: bson.M{"Releases": bson.M{"$exists": true}}}},
    {{Key: "$project", Value: bson.M{"ReleasesCount": bson.M{"$size": "$Releases"}}}},
    {{Key: "$match", Value: bson.M{"ReleasesCount": bson.M{"$gt": minReleases}}}},
}
cursor, err := collection.Aggregate(ctx, pipeline)
defer cursor.Close(ctx)
err = cursor.All(ctx, &albums)
```
- Find albums where at least one release is on a specific platform (e.g., "Spotify") and has a certain status (e.g., "Approved"):
```go
db.album.find({
	"Releases": {
    "$elemMatch": {
      "PlatformId": "Spotify",
      "Status": "Approved"
    }
  }
})
```
```go
filter := bson.M{
        "Releases": bson.M{
            "$elemMatch": bson.M{
                "PlatformId": platformID,
                "Status":     status,
            },
        },
    }
cursor, err := collection.Find(ctx, filter)
defer cursor.Close(ctx)
err = cursor.All(ctx, &albums)
```
- Get a list of albums with their songs:
```go
pipeline := mongo.Pipeline{
    {{"$lookup", bson.D{
        {Key: "from", Value: "songs"},
        {Key: "localField", Value: "Releases.PlatformId"},
        {Key: "foreignField", Value: "AlbumId"},
        {Key: "as", Value: "songDetails"},
    }}},
}
cursor, err := albumsCollection.Aggregate(ctx, pipeline)
defer cursor.Close(ctx)
err = cursor.All(ctx, &results)
```
- Get a list of albums along with their labels and songs:
```go
pipeline := mongo.Pipeline{
    {{"$lookup", bson.D{
        {Key: "from", Value: "labels"},
        {Key: "localField", Value: "LabelId"},
        {Key: "foreignField", Value: "Id"},
        {Key: "as", Value: "labelDetails"},
    }}},
    {{"$unwind", "$labelDetails"}},
    {{"$lookup", bson.D{
        {Key: "from", Value: "songs"},
        {Key: "let", Value: bson.D{{Key: "albumId", Value: "$Id"}}},
        {Key: "pipeline", Value: mongo.Pipeline{
            {{"$match", bson.D{{Key: "$expr", Value: bson.D{{Key: "$eq", Value: bson.A{"$AlbumId", "$$albumId"}}}}}},
        }},
        {Key: "as", Value: "songDetails"},
    }}},
}
cursor, err := albumsCollection.Aggregate(ctx, pipeline)
defer cursor.Close(ctx)
err = cursor.All(ctx, &results)
```
- Get a list of songs along with their album and label information:
```go
 pipeline := mongo.Pipeline{
    {{"$lookup", bson.D{
        {Key: "from", Value: "albums"},
        {Key: "localField", Value: "AlbumId"},
        {Key: "foreignField", Value: "Id"},
        {Key: "as", Value: "albumDetails"},
    }}},
    {{"$unwind", "$albumDetails"}},
    {{"$lookup", bson.D{
        {Key: "from", Value: "labels"},
        {Key: "localField", Value: "albumDetails.LabelId"},
        {Key: "foreignField", Value: "Id"},
        {Key: "as", Value: "labelDetails"},
    }}},
    {{"$unwind", "$labelDetails"}},
}
cursor, err := songsCollection.Aggregate(ctx, pipeline)
defer cursor.Close(ctx)
err = cursor.All(ctx, &results)
```
- Update Release Status in an Album
```go
db.albums.updateOne(
    { "id": "album1", "releases.platformId": "platform1" },
    { $set: { "releases.$.status": "Approved" } }
);
```
```go
filter := bson.M{
	"id": albumId,
	"releases.platformId": platformId,
}
update := bson.M{
	"$set": bson.M{
	"releases.$.status": newStatus,
	},
}
result, err := collection.UpdateOne(context.TODO(), filter, update)
```
- Add a new Artist to the Artists array within a specific Release in an Album:
```go
db.albums.updateOne(
    { "id": "album1", "releases.platformId": "platform1" },
    { $push: { "releases.$.artists": { "id": "artist1", "name": "New Artist" } } }
);
```
```go
filter := bson.M{
	"id": albumId,
	"releases.platformId": platformId,
}
update := bson.M{
	"$push": bson.M{
	"releases.$.artists": artist,
	},
}
result, err := collection.UpdateOne(context.TODO(), filter, update)
```
- Remove a specific Artist from the Artists array within a specific Release in an Album:
```go
db.albums.updateOne(
    { "id": "album1", "releases.platformId": "platform1" },
    { $pull: { "releases.$.artists": { "id": "artist1" } } }
);
```
```go
filter := bson.M{
	"id": albumId,
	"releases.platformId": platformId,
}
update := bson.M{
	"$pull": bson.M{
	"releases.$.artists": bson.M{"id": artistId},
	},
}
result, err := collection.UpdateOne(context.TODO(), filter, update)
```
- Update multiple documents in a collection, for example, to set a status for all albums with a certain label
```go
db.albums.updateMany(
    { "labelId": "label123" },
    { $set: { "status": "Approved" } }
);
```
```go
filter := bson.M{
	"labelId": labelId,
}
update := bson.M{
	"$set": bson.M{
		"releases.$[].status": newStatus,
	},
}
result, err := collection.UpdateMany(context.TODO(), filter, update)
```
- Perform an upsert (update if exists, otherwise insert), use the Upsert option:
```go
db.albums.updateOne(
    { "id": "album2" },
    { $set: { "id": "album2", "name": "Another Album", "labelId": "label123", "releases": [] } },
    { upsert: true }
);
```
```go
filter := bson.M{
	"id": album.Id,
}
update := bson.M{
	"$set": album,
}
opts := options.Update().SetUpsert(true)
result, err := collection.UpdateOne(context.TODO(), filter, update, opts)
```
-  Incrementing a field within a document:
```go
db.albums.updateOne(
    { "id": "album1" },
    { $inc: { "releaseCount": 1 } }
);
```
```go
filter := bson.M{
	"id": albumId,
}
update := bson.M{
	"$inc": bson.M{
	"releaseCount": 1,
	},
}
result, err := collection.UpdateOne(context.TODO(), filter, update)
```
## Module 4: GRPC
## Module 5: RabitMQ

Setting Up RabbitMQ

```bash
docker run -d --name rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:3-management
```

RabitMq is a Message Broker which facilates communication between Distributed systems by allowing microservices to Sending and Receiving messages. It implements the Advance Message Queuing Protocol(AMQP). 

RabitMq helps in Decoupling microservices by allowing then to communicate via Queues instead of each other directly. That way we can increase the Consumers according to the loads. 

In RabitMQ Producer send messages. Consumer reveive messages.
Queue act as a Buffer that stores messages until Consumer reads it.
Exchange routes the messages to Queues based on routing rules and the type of exchange.
Routing Key is used by Exchnages to route messages to Queues.

Type of Exchanges:
When you want a message to be routed to a specific queue then use direct exchange, that routes messages with a specific routing key to queues that are bound to the exchange with that same routing key.

When you need flexible routing based on multiple criteria then use topic exchange, that routes messages to queues based on pattern matching of routing keys using wildcards.

When you want to broadcast messages to multiple queues then use fanout exchange, that routes messages to all queues bound to it, regardless of the routing key.

When routing decisions are based on message attributes rather than routing keys then use headers exchange that routes messages based on message header attributes. Routing is determined by matching headers to the binding criteria.

Message acknowledgements are a mechanism that informs RabbitMQ whether a message has been successfully processed or not. 

With automatic acknowledgements, messages are acknowledged by the broker as soon as they are delivered to the consumer. This is done by setting the autoAck parameter to true in the Consume method.  If a consumer crashes or fails to process the message, the message is lost and not requeued for processing.

With manual acknowledgements, the consumer explicitly acknowledges the message after processing it. This is done by setting the autoAck parameter to false in the Consume method and using the msg.Acknowledge method to confirm that the message was processed successfully.

Negative acknowledgements are used to indicate that a message was not processed successfully. This can be done with the msg.Nack method, and you can choose whether to requeue the message or discard it.

Always use manual acknowledgements for critical message processing to ensure that messages are only removed from the queue after successful processing.

Use RabbitMQ management tools to monitor queues, message rates, and processing statuses. This helps in identifying issues with message processing and acknowledgements.

Message Persistence ensures that messages are stored on disk rather than kept only in memory. This way, even if RabbitMQ crashes, the messages are not lost and can be recovered upon restart. Writing messages to disk can impact performance, so there is a trade-off between reliability and throughput.

When a Queue is Declared Durable: The queue definition is stored to disk. If RabbitMQ crashes and restarts, the queue structure will be restored. Declaring a queue as durable does not mean that the messages will be persistent unless they are also marked as persistent.

To make a message persistent, set the DeliveryMode property of the amqp.Publishing structure to 2 when publishing the message. 

We can use message Time-To-Live (TTL) settings to automatically remove old messages from queues that are no longer needed.


When a message cannot be processed by its original queue, it is redirected to a dead-letter exchange based on predefined criteria. 
Once messages are routed to the DLX, you can handle them in several ways:  Use the DLX to examine failed messages and understand why they failed. This can be useful for debugging and improving message processing logic.  Implement logic to retry processing failed messages after some time or under certain conditions. Set up alerts or monitoring to be notified when messages are routed to the DLX, allowing you to respond to issues promptly. Periodically clean up old messages from the DLX to prevent excessive disk usage and maintain system performance.
Design retry mechanisms to handle temporary failures and reprocess messages before routing them to the DLX.

You need to declare a dead-letter exchange (DLX) where messages that cannot be delivered to the original queue will be routed.

```go
// Declare a dead-letter exchange
err = ch.ExchangeDeclare(
    "dlx_exchange", // Exchange name
    "direct",       // Exchange type
    true,           // Durable
    false,          // Auto-delete
    false,          // Internal
    nil,            // Arguments
)
if err != nil {
    log.Fatalf("Failed to declare a DLX: %v", err)
}
```
You need to declare a queue and set its x-dead-letter-exchange argument to specify the DLX.
```go
_, err = ch.QueueDeclare(
    "my_queue",    // Queue name
    true,          // Durable
    false,         // Auto-delete
    false,         // Exclusive
    false,         // No-wait
    amqp.Table{
        "x-dead-letter-exchange": "dlx_exchange", // Specify the DLX
    },
)
if err != nil {
    log.Fatalf("Failed to declare a queue: %v", err)
}
```
You also need to bind the queue to the DLX to ensure messages are routed correctly.
```go
// Bind the DLX to a queue
err = ch.QueueBind(
    "my_queue",    // Queue name
    "",            // Routing key
    "dlx_exchange", // Exchange name
    false,         // No-wait
    nil,           // Arguments
)
if err != nil {
    log.Fatalf("Failed to bind queue to DLX: %v", err)
}
```
Messages are typically routed to the DLX under the following conditions:

When a message in the queue exceeds its Time-To-Live (TTL) and is not consumed.
When a queue’s length exceeds its maximum limit (if configured), excess messages are moved to the DLX.
When a message is rejected by a consumer (e.g., using msg.Nack with requeue set to false).
When a message cannot be routed to any queue because of routing issues (e.g., when using exchanges).
```go
    amqp.Table{
        "x-message-ttl": 60000, // TTL in milliseconds
        "x-dead-letter-exchange": "dlx_exchange",
		// OR
		"x-max-length": 1000, // Max number of messages
        "x-dead-letter-exchange": "dlx_exchange",
	}
```

Example:
- Album Service: Publish messages about album creation and updates.
- Song Service: Publish messages about song creation and updates.
- Notification Service: Receive and handle specific types of messages based on routing keys.

A Fanout Exchange (broadcast_logs) for general announcements that should be broadcasted to all interested parties.

A Direct Exchange (direct_logs) for specific messages with routing keys.

Each service publishes messages to both the fanout exchange (broadcast_logs) and the direct exchange (direct_logs).

The fanout exchange distributes messages to all bound queues, useful for general notifications that all consumers should receive.
The direct exchange routes messages based on the routing key, allowing specific consumers to handle particular types of events.

album_service.go
```go
package main

import (
    "log"
    "github.com/streadway/amqp"
    "encoding/json"
)

type Album struct {
    ID     string `json:"id"`
    Title  string `json:"title"`
    Artist string `json:"artist"`
}

func publishMessage(ch *amqp.Channel, exchange, routingKey string, album Album) {
    body, err := json.Marshal(album)
    if err != nil {
        log.Fatalf("Failed to marshal album: %v", err)
    }

    err = ch.Publish(
        exchange,      // Exchange
        routingKey,    // Routing key (ignored by fanout exchange)
        false,         // Mandatory
        false,         // Immediate
        amqp.Publishing{
            ContentType: "application/json",
            Body:        body,
            DeliveryMode: amqp.Persistent, // Make message persistent
        },
    )
    if err != nil {
        log.Fatalf("Failed to publish a message: %v", err)
    }
    log.Printf("Published album event to %s: %s", exchange, body)
}

func main() {
    conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
    if err != nil {
        log.Fatalf("Failed to connect to RabbitMQ: %v", err)
    }
    defer conn.Close()

    ch, err := conn.Channel()
    if err != nil {
        log.Fatalf("Failed to open a channel: %v", err)
    }
    defer ch.Close()

    // Declare fanout exchange
    err = ch.ExchangeDeclare(
        "broadcast_logs", // Name
        "fanout",         // Type
        true,             // Durable
        false,            // Auto-deleted
        false,            // Internal
        nil,              // Arguments
    )
    if err != nil {
        log.Fatalf("Failed to declare fanout exchange: %v", err)
    }

    // Declare direct exchange
    err = ch.ExchangeDeclare(
        "direct_logs", // Name
        "direct",      // Type
        true,          // Durable
        false,         // Auto-deleted
        false,         // Internal
        nil,           // Arguments
    )
    if err != nil {
        log.Fatalf("Failed to declare direct exchange: %v", err)
    }

    album := Album{ID: "1", Title: "Greatest Hits", Artist: "Artist Name"}
    
    // Publish to direct exchange with routing key
    publishMessage(ch, "direct_logs", "album.created", album)

    // Publish to fanout exchange
    publishMessage(ch, "broadcast_logs", "", album)
}
```

song_service.go
```go
package main

import (
    "log"
    "github.com/streadway/amqp"
    "encoding/json"
)

type Song struct {
    ID     string `json:"id"`
    Title  string `json:"title"`
    AlbumID string `json:"album_id"`
}

func publishMessage(ch *amqp.Channel, exchange, routingKey string, song Song) {
    body, err := json.Marshal(song)
    if err != nil {
        log.Fatalf("Failed to marshal song: %v", err)
    }

    err = ch.Publish(
        exchange,      // Exchange
        routingKey,    // Routing key (ignored by fanout exchange)
        false,         // Mandatory
        false,         // Immediate
        amqp.Publishing{
            ContentType: "application/json",
            Body:        body,
            DeliveryMode: amqp.Persistent, // Make message persistent
        },
    )
    if err != nil {
        log.Fatalf("Failed to publish a message: %v", err)
    }
    log.Printf("Published song event to %s: %s", exchange, body)
}

func main() {
    conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
    if err != nil {
        log.Fatalf("Failed to connect to RabbitMQ: %v", err)
    }
    defer conn.Close()

    ch, err := conn.Channel()
    if err != nil {
        log.Fatalf("Failed to open a channel: %v", err)
    }
    defer ch.Close()

    // Declare fanout exchange
    err = ch.ExchangeDeclare(
        "broadcast_logs", // Name
        "fanout",         // Type
        true,             // Durable
        false,            // Auto-deleted
        false,            // Internal
        nil,              // Arguments
    )
    if err != nil {
        log.Fatalf("Failed to declare fanout exchange: %v", err)
    }

    // Declare direct exchange
    err = ch.ExchangeDeclare(
        "direct_logs", // Name
        "direct",      // Type
        true,          // Durable
        false,         // Auto-deleted
        false,         // Internal
        nil,           // Arguments
    )
    if err != nil {
        log.Fatalf("Failed to declare direct exchange: %v", err)
    }

    song := Song{ID: "1", Title: "Hit Song", AlbumID: "1"}
    
    // Publish to direct exchange with routing key
    publishMessage(ch, "direct_logs", "song.created", song)

    // Publish to fanout exchange
    publishMessage(ch, "broadcast_logs", "", song)
}
```
Notification Service: It listens to both exchanges.
It binds to the fanout exchange without specifying a routing key, so it receives all broadcast messages.
It binds to the direct exchange with specific routing keys to process messages related to particular events.

notification_service.go
```go
package main

import (
    "log"
    "github.com/streadway/amqp"
    "encoding/json"
)

func main() {
    conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
    if err != nil {
        log.Fatalf("Failed to connect to RabbitMQ: %v", err)
    }
    defer conn.Close()

    ch, err := conn.Channel()
    if err != nil {
        log.Fatalf("Failed to open a channel: %v", err)
    }
    defer ch.Close()

    // Declare fanout exchange
    err = ch.ExchangeDeclare(
        "broadcast_logs", // Name
        "fanout",         // Type
        true,             // Durable
        false,            // Auto-deleted
        false,            // Internal
        nil,              // Arguments
    )
    if err != nil {
        log.Fatalf("Failed to declare fanout exchange: %v", err)
    }

    // Declare direct exchange
    err = ch.ExchangeDeclare(
        "direct_logs", // Name
        "direct",      // Type
        true,          // Durable
        false,         // Auto-deleted
        false,         // Internal
        nil,           // Arguments
    )
    if err != nil {
        log.Fatalf("Failed to declare direct exchange: %v", err)
    }

    // Declare queues
    q1, err := ch.QueueDeclare(
        "",    // Name (let RabbitMQ generate a name)
        false, // Durable
        true,  // Delete when unused
        true,  // Exclusive
        false, // No-wait
        nil,   // Arguments
    )
    if err != nil {
        log.Fatalf("Failed to declare a queue: %v", err)
    }

    q2, err := ch.QueueDeclare(
        "",    // Name (let RabbitMQ generate a name)
        false, // Durable
        true,  // Delete when unused
        true,  // Exclusive
        false, // No-wait
        nil,   // Arguments
    )
    if err != nil {
        log.Fatalf("Failed to declare a queue: %v", err)
    }

    // Bind queues to fanout exchange
    err = ch.QueueBind(
        q1.Name,         // Queue name
        "",              // Routing key (ignored by fanout exchange)
        "broadcast_logs", // Exchange name
        false,
        nil,
    )
    if err != nil {
        log.Fatalf("Failed to bind queue to fanout exchange: %v", err)
    }

    // Bind queue to direct exchange with routing keys
    routingKeys := []string{"album.created", "song.created", "label.created"}
    for _, key := range routingKeys {
        err = ch.QueueBind(
            q2.Name,          // Queue name
            key,              // Routing key
            "direct_logs",    // Exchange name
            false,
            nil,
        )
        if err != nil {
            log.Fatalf("Failed to bind queue to direct exchange with routing key %s: %v", key, err)
        }
    }

    msgs1, err := ch.Consume(
        q1.Name, // Queue
        "",      // Consumer tag
        true,    // Auto-ack
        false,   // Exclusive
        false,   // No-local
        false,   // No-wait
        nil,     // Arguments
    )
    if err != nil {
        log.Fatalf("Failed to register a consumer for fanout queue: %v", err)
    }

    msgs2, err := ch.Consume(
        q2.Name, // Queue
        "",      // Consumer tag
        true,    // Auto-ack
        false,   // Exclusive
        false,   // No-local
        false,   // No-wait
        nil,     // Arguments
    )
    if err != nil {
        log.Fatalf("Failed to register a consumer for direct queue: %v", err)
    }

    go func() {
        for msg := range msgs1 {
            log.Printf("Received message from fanout exchange: %s", msg.Body)
        }
    }()

    go func() {
        for msg := range msgs2 {
            log.Printf("Received message from direct exchange: %s", msg.Body)
        }
    }()

    log.Println("Waiting for messages. To exit press CTRL+C")
    select {} // Block forever
}
```

## Module 6: AWS
## Module 7: Docker and Kubernetis
Docker is a platform that automate the deployment of applications inside lightweight, portable containers. 

Containers package up everything like code, runtime, system tools, libraries, and settings, that an application needs for runing, so it runs consistently across different computing environments.

Basic Concepts:

Image: Docker Image is a snapshot of a filesystem. An image contains everything needed to run an application: code, runtime, libraries, and environment variables. Images are read-only.

Container: Docker Container is a running instance of an image. Containers can write to their own filesystem (but not to the image).
Containers have a writable layer on top of the image.

Docker Hub: A cloud-based registry where Docker images are stored and shared.

Docker volumes are used for persisting data of Containers, when we remove the container the volume associated with it, not removed until we remove the volume specifically, one volume can be shared among multiple containers.

#### Run Mongodb:
```bash
docker network create my-network
```
```bash
docker volume create mongodb-data
```
```bash
docker run -d \ 
# -d runs the container in detached mode.
  --name mongodb \ 
  # --name mongodb gives the container a name.
  --network my-network \ 
  # --network my-network connects the container to the network (if you created one).
  -v mongodb-data:/data/db \ 
  # -v mongodb-data:/data/db mounts the volume mongodb-data to /data/db in the container, which is where MongoDB stores its data.
  -p 27017:27017 \ 
  # -p 27017:27017 maps port 27017 on the host to port 27017 in the container.
  mongo:latest 
  # mongo:latest specifies the MongoDB image to use
```

```bash
docker exec -it mongodb /bin/bash
```
```bash
docker stop mongodb
```
Dockerfile: Dockerfile is a script for building a Docker image.  Each instruction in the Dockerfile creates a new layer in the Docker image, which is then used to build and run containers.

Dockerfile Commands
```bash
FROM:
Definition: Specifies the base image for the new image.
Example: FROM ubuntu:20.04
```
```bash
RUN:
Definition: Executes a command in Container during the image build process.
Example: RUN apt-get update && apt-get install -y nginx
```
```bash
COPY:
Definition: Copies files or directories from the host to the container.
Example: COPY index.html /usr/share/nginx/html/
```
```bash
CMD:
Definition: Provides the default command to run when a container starts.
Example: CMD ["nginx", "-g", "daemon off;"]
```
```bash
EXPOSE:
Definition: Documents the port on which the container will listen.
Example: EXPOSE 80
```
```bash
ENTRYPOINT:
Definition: Configures a container to run as an executable.
Example: ENTRYPOINT ["nginx", "-g", "daemon off;"]
```

The docker build command is used to create the image from the Dockerfile.
```bash
docker build -t my-image:latest .
```
. represents the current directory where dockerfile present.

Docker compose is a command is use for Start/Up multiple seervices, networks and volumes with single command `docker-compose up`
Docker compose  uses a YAML file (docker-compose.yml) to configure these Services, Networks and Volumes.

docker-compose.override.yml allow us to modify or extend the base docker-compose.yml configuration without changing the original file. They are useful where environment-specific settings or adjustments for local development, staging or prod are required.
When we run docker-compose up, Docker Compose automatically includes settings from docker-compose.override.yml in addition to docker-compose.yml

`docker-compose -f docker-compose.yml up`
This command ignores the docker-compose.override.yml file and only uses the settings from docker-compose.yml.

Docker networking enables containers to communicate with each other and with the outside world. Common network modes include:
- Bridge: Default network mode where containers are connected to a virtual bridge and can communicate with each other.

- Host: Containers share the host’s networks and IP address.

- Overlay: Used to connect containers across multiple Docker hosts.

- Macvlan: Each container have its own MAC address, making containers appear as physical network interfaces.


Docker Performance and Optimization:

- By combining multiple commands into a single RUN instruction, we can reduce the numer of image layers and ensure that temporary files are removed within the same layer where they are created.

- After building the Go application, you should remove source code, module files, and other temporary build artifacts that are not needed in the final image. This helps to keep the final image minimal and secure.

This Dockerfile uses a multi-stage build to first compile the Go application and then copy the resulting binary to a minimal final image.


#### Run Golang:
```bash
FROM golang:1.20 AS builder # provides the necessary environment for building the Go application.

WORKDIR /app # Set the working directory

COPY go.mod go.sum ./  # Copy the Go module files

RUN go mod download && \ # fetches dependencies,
    go mod verify # ensures their integrity.

COPY . .

RUN go build -o myapp . && \ # compiles the application.
    rm -rf /app/*.go /app/*.mod /app/*.sum /app/pkg /app/src #  removes the source files, module files, and directories that are not needed in the final image

FROM scratch # is an empty base image

COPY --from=builder /app/myapp /myapp # binary (myapp) is copied from the build stage

CMD ["/myapp"] # specifies the command to run the application.
```
- Use smaller base images like slim or alpine variants.
- Use multi-stage builds to keep the final image minimal and efficient.
- Minimize the number of layers: Combine commands in Dockerfile to reduce layers.
- Use .dockerignore: Exclude unnecessary files from the build context.
- Clean up temporary files: Remove unnecessary files in the RUN commands. A smaller image means faster build times, quicker uploads/downloads, and better performance in production.


Bind mounts allow you to mount a specific file or directory from the host filesystem into a container. Unlike volumes, bind mounts are not managed by Docker. Any changes to the file or directory on the host are reflected in the container, and vice versa.
```bash
docker run -it --rm -p 8080:8080 -v /path/to/your/project:/app my-go-app
```

The docker exec command allows you to run commands inside a running Docker container.

Docker logs can be accessed using `docker logs <container_id>`

Health checks are defined in the Dockerfile using the HEALTHCHECK instruction. This allows Docker to periodically run a command inside the container to determine if it is healthy. 


Install docker:
```bash
sudo apt install docker.io
```
Check Docker Version:
```bash
docker --version
```
Stop Docker:
```bash
sudo systemctl stop docker
```
Start Docker (if not already running):
```bash
sudo systemctl start docker
```
Lists all Docker images available locally.
```bash
docker images
```
Search for images on Docker Hub:
```bash
docker search <term>
```
Downloads an image from a Docker registry (e.g., Docker Hub).
```bash
docker pull <image>:<tag>
docker pull nginx:latest
```
Build or Creates a Docker image from a Dockerfile.
```bash
docker build -t <image>:<tag> <path>
docker build -t my-image:latest .
```
Deletes a Docker image from the local system.
```bash
docker rmi <image>:<tag>
docker rmi nginx:latest
```
Starts a new container from an image and optionally executes a command.
```bash
docker run [OPTIONS] <image>:<tag>
docker run -d -p 80:80 --name webserver nginx
```
Shows a list of all currently running containers.
```bash
docker ps
```
Lists all containers, including those that are stopped.
```bash
docker ps -a
```
Action of a running container.
```bash
docker stop <container>
docker start <container>
docker restart <container>
docker rm <container>
```
View logs from a container:
```bash
docker logs <container>
```
Execute a command in a running container: opens a Bash shell in the container
```bash
docker exec -it <container> <command>
docker exec -it webserver /bin/bash
```
Attach to a running container:
```bash
docker attach <container>
```
Inspect a Container or Image :  Provides detailed information about a container or image in JSON format.
```bash
docker inspect <container-id or image-id>
docker inspect d9b100f2f636
```
Tag an Image: Creates a new tag for an existing image.
```bash
docker tag <source-image> <target-image>
docker tag my-app:latest my-app:v1.0
```
Removes images that are not used by any containers.
```bash
docker image prune
```
Create a volume:
```bash
docker volume create <volume>
```
List volumes:
```bash
docker volume ls
```
Inspect a volume:
```bash
docker volume inspect <volume>
```
Remove a volume:
```bash
docker volume rm <volume>
```
List networks:
```bash
docker network ls
```
Create a network:
```bash
docker network create <network>
```
Inspect a network:
```bash
docker network inspect <network>
```
Remove a network:
```bash
docker network rm <network>
```
Check Docker system information:
```bash
docker info
```
Clean up unused Docker objects:
```bash
docker system prune
```
List all Docker resources:
```bash
docker system df
```
Start services defined in a docker-compose.yml file:
```bash
docker-compose up
```
Start services in detached mode:
```bash
docker-compose up -d
```
Stop services:
```bash
docker-compose down
```
List services:
```bash
docker-compose ps
```
Build or rebuild services:
```bash
docker-compose build
```
View logs from services:
```bash
docker-compose logs
```

#### Run Postgres:
```bash
docker volume create postgres-data
```
```bash
docker run -d \
  --name postgres \
  -v postgres-data:/var/lib/postgresql/data \
  -e POSTGRES_USER=myuser \
  -e POSTGRES_PASSWORD=mypassword \
  -e POSTGRES_DB=mydatabase \
  -p 5432:5432 \
  postgres:latest
```

```bash
docker exec -it postgres psql -U myuser -d mydatabase
```
```bash
docker stop postgres
```

#### Dockerfile for a Go Application:
```bash
# Stage 1: Build the Go binary
FROM golang:1.20 AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the Go Modules manifests
COPY go.mod go.sum ./

# Download the Go Modules dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go binary
RUN go build -o myapp .

# Stage 2: Create a smaller image with the Go binary
FROM debian:bullseye-slim

# Install necessary libraries (e.g., for SSL)
RUN apt-get update && apt-get install -y ca-certificates

# Set the Current Working Directory inside the container
WORKDIR /root/

# Copy the Go binary from the previous stage
COPY --from=builder /app/myapp .

# Expose port 8080 (or any port your application uses)
EXPOSE 8080

# Command to run the binary
CMD ["./myapp"]
```

#### Dockerfile for Redis
```bash
# Use the official Redis image from Docker Hub
FROM redis:latest

# Expose the default Redis port
EXPOSE 6379

# The Redis image already has a default entrypoint, so no CMD needed
```
#### Dockerfile for RabbitMQ
```bash
# Use the official RabbitMQ image with the management plugin
FROM rabbitmq:3-management

# Expose the default RabbitMQ ports
EXPOSE 5672 15672

# The RabbitMQ image already has a default entrypoint, so no CMD needed
```

#### Dockerfile for MongoDB:
```bash
# Use the official MongoDB image from Docker Hub
FROM mongo:latest

# Expose port 27017, the default port for MongoDB
EXPOSE 27017

# Optionally, you can add a default database or collection setup script
# COPY init-mongo.js /docker-entrypoint-initdb.d/

# The MongoDB image already has a default entrypoint, so no CMD needed

```
#### docker-compose.yml:
```yaml
version: '3.8' #Specifies the version of the Docker Compose file format to use.
services: # Each service is essentially a container, and you can configure its settings such as image, build context, and environment variables.
  mongodb:
    image: mongo:latest #  Specifies the Docker image to use for the service.
    container_name: mongodb
    ports: #  Maps container ports to host ports.
      - "27017:27017"
    volumes: # Mounts host directories or volumes into the container.
      - mongodb_data:/data/db
    networks:
      - mynetwork

volumes:
  mongodb_data:

networks: # Defines custom networks for services.
  mynetwork:
```
#### docker-compose.yml:
```yaml
version: '3.8'
services:
  golang:
    build: .
    container_name: golang_app
    ports:
      - "8080:8080"
    networks:
      - mynetwork

networks:
  mynetwork:
```
#### docker-compose.yml:
```yaml
version: '3.8'
services:
  postgres:
    image: postgres:latest
    container_name: postgres
    ports:
      - "5432:5432"
    environment: # Sets environment variables in the container.
      POSTGRES_DB: mydatabase
      POSTGRES_USER: user
      POSTGRES_PASSWORD: password
    volumes:
      - postgres_data:/var/lib/postgresql/data
    networks:
      - mynetwork

volumes:
  postgres_data:

networks:
  mynetwork:
```
#### docker-compose.yml:
```yaml
version: '3.8'
services:
  rabbitmq:
    image: rabbitmq:management
    container_name: rabbitmq
    ports:
      - "5672:5672"
      - "15672:15672"
    volumes:
      - rabbitmq_data:/var/lib/rabbitmq
    networks:
      - mynetwork

volumes:
  rabbitmq_data:

networks:
  mynetwork:
```
#### docker-compose.yml:
```yaml
version: '3.8'
services:
  redis:
    image: redis:latest
    container_name: redis
    ports:
      - "6379:6379"
    volumes:
      - redis_data:/data
    command: ["redis-server", "--appendonly", "yes"] # This specifies the command to run when the container starts.
    networks:
      - mynetwork

volumes:
  redis_data:

networks:
  mynetwork:
```

Create a directory for each service and place the respective docker-compose.yml and configuration files in that directory.

Run Docker Compose in each directory: This command will start the containers based on the configuration files.
```bash
docker-compose up
```
Access Services:

`MongoDB`: mongodb://localhost:27017

`PostgreSQL`: postgresql://user:password@localhost:5432/mydatabase

`RabbitMQ`: Management UI at http://localhost:15672 (default username/password: guest/guest)

`Redis`: redis://localhost:6379


#### Kubernetes 
Kubernetes is designed for automate the process of deploying, scaling, and managing containerized applications.  Kubernetes uses YAML files to define resources.

A Pod can contain one or more containers that share the same network, namespace and storage.
Containers in the same Pod share a common IP address and port space. They can communicate with each other using localhost and the same port numbers. Pods can communicate with each other using their IP addresses and defined ports. Pods are often managed by higher-level controllers like Deployments, ReplicaSets, which handle scaling, updates, and ensuring that the desired number of Pod replicas are running.

ReplicaSet is a controller that ensures a specified number of identical Pods are running at any given time.   If a Pod fails or is deleted, the ReplicaSet creates a new Pod to replace it. A ReplicaSet uses a label selector to identify which Pods it should manage. The selector matches Pods based on their labels, and the ReplicaSet ensures that the Pods with these labels are created or deleted as needed.

Services provide a stable endpoint(DNS name and IP address) for accessing a set of Pods. This stable endpoint is crucial because Pods can come and go due to scaling, updates, or failures, but Services provide a consistent way to access these Pods.

Kubernetes provides built-in mechanisms for service discovery, so applications can find and communicate with other services without hardcoding IP addresses.

Types of Services:

ClusterIP : The Service is only accessible within the Kubernetes cluster. It’s used for communication between different services within the cluster.

NodePort : This allows the Service to be accessed from outside the cluster by requesting <NodeIP>:<NodePort>. This type is often used for simple testing and debugging.

LoadBalancer : Creates an external load balancer (e.g., from cloud providers like AWS, GCP, Azure) and assigns a public IP to the Service. This allows the Service to be accessed from outside the cluster using that public IP. It’s often used in production environments.

Headless Service : When you don’t need load balancing or a cluster IP, you can create a headless Service by setting the clusterIP field to None. This allows direct access to the Pods without load balancing, useful for stateful applications.

Services use selectors to determine which Pods belong to the Service. 
selector : Specifies that the Service will route traffic to Pods with the label app: my-app


Deployment is a higher-level abstraction that manages ReplicaSets and Pods. Deployments provide version control and history for your application, making it easy to roll back to previous versions if needed.
Deployments support rolling updates, allowing you to update your application with zero downtime. Kubernetes gradually replaces old Pods with new ones according to the update strategy you specify.
If a new update causes issues, you can roll back to a previous version of your application. Kubernetes maintains a history of revisions, making it easy to revert to a stable state.
Deployments allow you to scale the number of replicas (Pods) up or down easily.


Kubernetes uses a flat network model where every Pod gets its own IP address. Pods communicate directly with each other using their IP addresses, 

Cluster is a set of machines (nodes) that work together to run containerized applications. The cluster consists of a control plane(Master Node) and a set of worker nodes, and it provides the environment needed to deploy, manage, and scale applications.
What are the main components of Kubernetes architecture?
- Master Node: The node that manages the cluster. It includes:

	- API Server: The API server exposes the Kubernetes API which is used by users and other components to interact with the cluster. When we run commands using kubectl or other tools, they communicate with the API Server.  When we deploy a new application, the API Server stores the desired state in etcd and updates the cluster to match this state

	- Controller Manager: Each controller watches the API Server for changes to resources and makes adjustments to ensure that the actual state matches the desired state. For example, if a ReplicaSet specifies that three replicas of a Pod should be running, the controller will create or delete Pods as necessary to meet this requirement.

	- Scheduler: Assigns Pods to Nodes based on resource availability and constraints. if a Pod requires 2 CPU cores and 4 GB of memory, the Scheduler will find a Node with sufficient available resources and assign the Pod to that Node.
	
	- etcd: A distributed key-value store that holds all the configuration data and state of the Kubernetes cluster.  When we make changes to the cluster (e.g., deploying an application or scaling a service), the API Server updates the state in etcd.  This allows Kubernetes to recover from failures by restoring the cluster state from etcd.

- Worker Nodes: worker node refers to a machine (virtual or physical) that runs the applications and workloads in your Kubernetes cluster. On AWS, worker nodes are typically Amazon EC2 instances that are part of your Kubernetes cluster.
Kubernetes Components on Worker Nodes: 
	- Kubelet: An agent that ensures containers are running in Pods.  The Kubelet communicates with the API Server to get the desired state of Pods. It then makes sure that the containers in those Pods are running as expected. If a container fails or crashes, the Kubelet will restart it to maintain the desired state. It also collects and reports metrics about the Node and the Pods running on it.

	- Kube-Proxy: Maintains network rules and load-balances traffic to Pods. Kube-Proxy manages network traffic routing by maintaining iptables or IPVS rules on Nodes. This ensures that network traffic is properly directed to the Pods. For example, if a Service exposes multiple Pods, Kube-Proxy ensures that traffic is distributed evenly across these Pods.

	- Container Runtime: Software that runs and manages containers (e.g., Docker, containerd).

We can scale a cluster by adding or removing worker nodes as needed to handle varying workloads.

AWS provides a managed Kubernetes service called EKS. When you create an EKS cluster, you can specify the EC2 instance types and configurations for your worker nodes.

Elastic Load Balancing (ELB): Kubernetes services that are exposed as LoadBalancer types, can use AWS ELB to provide external access and distribute traffic to the worker nodes.

Amazon EBS: For persistent storage, you can use Amazon Elastic Block Store (EBS) volumes that are attached to worker nodes.

Helm is a package manager for Kubernetes that simplifies the deployment, management, and versioning of applications and services on a Kubernetes cluster.

A Helm chart is a collection of Kubernetes YAML files organized into a directory structure that defines a Kubernetes application. It includes everything needed to run an application, such as deployments, services, ingress configurations, and more.

Helm is installed on your local machine or CI/CD pipeline and communicates with your Kubernetes cluster to deploy and manage applications.

A typical Helm chart directory includes
Chart.yaml: Contains metadata about the chart, such as its name, version, and description.
values.yaml: A file containing default configuration values that can be overridden by the user.
templates/: A directory with Kubernetes manifest templates that Helm will use to generate Kubernetes resources.
charts/: A directory where dependent charts can be stored.
README.md: Documentation about the chart.

Helm can install, upgrade, rollback, and delete releases. Each release can be upgraded or rolled back independently of others.


#### To deploy a Go (Golang) application on Amazon EKS (Elastic Kubernetes Service) with load balancing:

Install AWS CLI
Install kubectl
Install eksctl
Install Docker
Create the EKS Cluster Using eksctl
```bash
eksctl create cluster \
  --name my-cluster \
  --region us-west-2 \
  --nodes 3 \
  --node-type t3.medium \
  --with-oidc \
  --ssh-access \
  --ssh-public-key ~/.ssh/id_rsa.pub
```
Update kubectl Context
aws eks --region us-west-2 update-kubeconfig --name my-cluster


You can use MongoDB as a managed service (e.g., MongoDB Atlas) or deploy it directly on EKS.
Create a MongoDB Deployment YAML File
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
        - name: mongodb-data
          mountPath: /data/db
      volumes:
      - name: mongodb-data
        emptyDir: {}
```
Create a MongoDB Service YAML File
```yaml
apiVersion: v1
kind: Service
metadata:
  name: mongodb-service
spec:
  selector:
    app: mongodb
  ports:
    - protocol: TCP
      port: 27017
  type: ClusterIP
```
Apply MongoDB Resources
kubectl apply -f mongodb-deployment.yaml
kubectl apply -f mongodb-service.yaml


Create Your Go Application: Example main.go
```go
package main

import (
    "context"
    "fmt"
    "log"
    "net/http"
    "github.com/gin-gonic/gin"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func init() {
    var err error
    clientOptions := options.Client().ApplyURI("mongodb://mongodb-service:27017")
    client, err = mongo.Connect(context.TODO(), clientOptions)
    if err != nil {
        log.Fatal(err)
    }
}

func handler(c *gin.Context) {
    collection := client.Database("testdb").Collection("testcollection")
    result := collection.FindOne(context.TODO(), map[string]string{"hello": "world"})
    var res map[string]string
    err := result.Decode(&res)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, res)
}

func main() {
    r := gin.Default()
    r.GET("/", handler)
    r.Run(":8080")
}
```
Create a Dockerfile
```
# Use the official Golang image to build the app
FROM golang:1.20 AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the source code into the container
COPY . .

# Build the Go app
RUN go build -o myapp

# Start a new stage from scratch
FROM alpine:latest

# Set the Current Working Directory inside the container
WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/myapp .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./myapp"]
```
Build the Docker Image
docker build -t my-go-app:latest .


Push the Docker Image to ECR
Create an ECR Repository
aws ecr create-repository --repository-name my-go-app --region us-west-2
Authenticate Docker to ECR
aws ecr get-login-password --region us-west-2 | docker login --username AWS --password-stdin <account-id>.dkr.ecr.us-west-2.amazonaws.com


Tag and Push the Docker Image
docker tag my-go-app:latest <account-id>.dkr.ecr.us-west-2.amazonaws.com/my-go-app:latest
docker push <account-id>.dkr.ecr.us-west-2.amazonaws.com/my-go-app:latest

Deploy Your Application to EKS
Create a Deployment YAML File
```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: my-go-app
spec:
  replicas: 3
  selector:
    matchLabels:
      app: my-go-app
  template:
    metadata:
      labels:
        app: my-go-app
    spec:
      containers:
      - name: my-go-app
        image: <account-id>.dkr.ecr.us-west-2.amazonaws.com/my-go-app:latest
        ports:
        - containerPort: 8080
```
Apply the Deployment
kubectl apply -f deployment.yaml

Create a Service YAML File with Load Balancing
```yaml
apiVersion: v1
kind: Service
metadata:
  name: my-go-app-service
spec:
  selector:
    app: my-go-app
  ports:
    - protocol: TCP
      port: 80
      targetPort: 8080
  type: LoadBalancer
```
Apply the Service
kubectl apply -f service.yaml

Get the External IP Address:
kubectl get services
Look for the EXTERNAL-IP of the my-go-app-service
------------------------------------------------------
POD yaml
```yaml
apiVersion: v1
kind: Pod
metadata:
  name: mypod
spec:
  containers:
  - name: mycontainer
    image: nginx:latest
    ports:
    - containerPort: 80
```
Service: An abstraction that defines a logical set of Pods and a policy to access them, typically using a stable IP address and DNS name.  Services provide stable network identities and load balancing for Pods. They enable communication between Pods within the cluster and expose applications running in Pods to external traffic. Services can be of different types: ClusterIP (default, internal access), NodePort (exposes a port on each node), LoadBalancer (provisions an external load balancer), and ExternalName (maps to an external service).
```yaml
apiVersion: v1
kind: Service
metadata:
  name: myservice
spec:
  type: LoadBalancer
  selector:
    app: myapp
  ports:
    - protocol: TCP
      port: 80
      targetPort: 80
```
ReplicaSet Ensures a specified number of pod replicas are running at any given time. It monitors Pods and maintains the desired number of replicas, replacing any Pods that fail or are deleted.


Deployment: A higher-level abstraction that manages ReplicaSets and provides additional features like rolling updates, rollbacks, and declarative updates. A Deployment creates and manages ReplicaSets for maintaining application states and updates.
```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: mydeployment
spec:
  replicas: 3
  selector:
    matchLabels:
      app: myapp
  template:
    metadata:
      labels:
        app: myapp
    spec:
      containers:
      - name: mycontainer
        image: nginx:latest
        ports:
        - containerPort: 80
```
Namespace: A way to divide cluster resources between multiple users or applications, providing a scope for names.

How does Kubernetes handle scaling of applications?
Kubernetes supports both manual and automatic scaling:

Adjust the number of replicas for a Deployment or ReplicaSet using the kubectl scale command.

Automatically scales the number of Pods in a Deployment or ReplicaSet based on metrics like CPU utilization or custom metrics. The Horizontal Pod Autoscaler monitors the current load and adjusts the number of Pods to maintain performance and resource utilization.

Helm is a package manager for Kubernetes that simplifies deploying and managing applications using Helm charts. Charts are pre-configured templates that describe Kubernetes resources and their configurations. 


Ingress in Kubernetes like a traffic manager for your web applications. Imagine you have a few different web services running in your Kubernetes cluster, like a blog, a shopping cart, and a user profile page. Each of these services needs to be accessed from outside your cluster (like from a web browser on your laptop).
Ingress is like a smart router that decides where incoming web traffic should go based on rules you define. For example, it can route traffic to the blog service if someone visits www.example.com/blog and to the shopping cart service if they visit www.example.com/cart. To make these rules work, you need an Ingress controller. This is like the actual router that follows the rules and directs the traffic accordingly. Popular controllers include Nginx and Traefik. Ingress controllers can also provide additional features like securing traffic with SSL (making sure your website is served over HTTPS) and balancing the load between different instances of your service to handle lots of traffic.
Traefik sits in front of your services and routes incoming requests to the correct service based on rules you define.
One of Traefik’s standout features is its ability to automatically discover services and apply configuration changes without needing manual updates. As new services are added or removed, Traefik automatically adjusts its routing rules.
Traefik can distribute incoming requests across multiple instances of a service, which helps in balancing the load
Traefik can handle SSL/TLS encryption for secure connections,
In Kubernetes, Traefik acts as an Ingress controller, meaning it manages the rules that route external HTTP/HTTPS traffic into the cluster, directing it to the appropriate services based on the Ingress rules you define.

Scaling in Amazon EKS (Elastic Kubernetes Service) involves both scaling the Kubernetes control plane and the worker nodes (EC2 instances) that run your applications.
AWS scales the control plane resources based on the workload and demand. This ensures that the Kubernetes API server remains responsive and capable of handling the requests made by the worker nodes and other components.
Worker nodes in EKS are EC2 instances that run the Kubernetes worker components (kubelet, container runtime). You can scale these nodes up or down based on the workload.
The Cluster Autoscaler monitors the Kubernetes scheduler. If it detects that there are unscheduled pods because there aren’t enough resources, it triggers scaling up of the cluster. It also monitors node utilization. If it finds underutilized nodes that can be removed without affecting the running pods (i.e., after scheduling those pods elsewhere), it triggers scaling down of the cluster.

EKS integrates with ELB (Application Load Balancer or Network Load Balancer) to distribute traffic across multiple pods. Load balancers scale automatically based on the number of targets (pods).

AWS Fargate provides serverless compute options that automatically handle scaling based on container resource requirements. 
With Fargate, you don’t need to provision, configure, or manage servers. You simply define your container requirements, and Fargate handles the rest.
Fargate automatically provisions the required compute resources (CPU and memory) based on your container specifications and workload. You pay only for the compute resources that you use. There’s no need to pay for or manage idle instances. Ideal for running microservices where each service can be scaled independently based on its resource requirements.


Kubernetes Federation is a mechanism for managing multiple Kubernetes clusters across different regions or cloud providers.


View Cluster Information:
```bash
kubectl cluster-info
```
Get Cluster Nodes:
```bash
kubectl get nodes
```
Get Cluster Info:
```bash
kubectl cluster-info
```
List Pods:
```bash
kubectl get pods
```
Get Pod Details:
```bash
kubectl describe pod pod_name
```
View Pod Logs:
```bash
kubectl logs pod_name
```
Execute Command in a Pod:
```bash
kubectl exec -it pod_name -- /bin/bash
```
Delete a Pod:
```bash
kubectl delete pod pod_name
```
List Services:
```bash
kubectl get services
```
Describe Service:
```bash
kubectl describe service service_name
```
Expose a Pod as a Service:
```bash
kubectl expose pod pod_name --type=LoadBalancer --port=80 --target-port=8080
```
Create a Deployment:
```bash
kubectl create deployment deployment_name --image=image_name:tag
```
List Deployments:
```bash
kubectl get deployments
```
Scale a Deployment:
```bash
kubectl scale deployment deployment_name --replicas=3
```
Update a Deployment:
```bash
kubectl set image deployment/deployment_name container_name=image_name:tag
```
Roll Back a Deployment:
```bash
kubectl rollout undo deployment/deployment_name
```
Get Deployment Details:
```bash
kubectl describe deployment deployment_name
```
Delete a Deployment:
```bash
kubectl delete deployment deployment_name
```
List Namespaces:
```bash
kubectl get namespaces
```
Create a Namespace:
```bash
kubectl create namespace namespace_name
```
Delete a Namespace:
```bash
kubectl delete namespace namespace_name
```
Apply Configuration from a File:
```bash
kubectl apply -f file_name.yaml
```
Delete Resources Defined in a File:
```bash
kubectl delete -f file_name.yaml
```
Get Resource Usage (CPU/Memory):
```bash
kubectl top nodes
kubectl top pods
```
Inspect Resource Details:
```bash
kubectl get all
kubectl describe pod pod_name
```
Port Forwarding:
```bash
kubectl port-forward pod_name local_port:container_port
```
#### Let's go through a detailed example of orchestrating a microservices architecture with Kubernetes, where each microservice is written in Go and uses different databases or messaging systems.

We'll create three Go microservices:

Service A: Uses MongoDB.
Service B: Uses PostgreSQL.
Service C: Uses Redis and RabbitMQ.

Each service will be containerized and deployed on Kubernetes.

Service A: MongoDB

```go
package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	clientOptions := options.Client().ApplyURI("mongodb://mongodb:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		collection := client.Database("testdb").Collection("testcol")
		result := collection.FindOne(context.Background(), map[string]interface{}{"name": "test"})
		var document map[string]interface{}
		if err := result.Decode(&document); err != nil {
			http.Error(w, "Error fetching data", http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, "Document: %v", document)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

```

Service B: PostgreSQL

```go
package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	connStr := "user=postgres password=password dbname=testdb host=postgres port=5432 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		row := db.QueryRow("SELECT data FROM test_table WHERE id = $1", 1)
		var data string
		if err := row.Scan(&data); err != nil {
			http.Error(w, "Error fetching data", http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, "Data: %s", data)
	})

	log.Fatal(http.ListenAndServe(":8081", nil))
}
```

Service C: Redis and RabbitMQ

```go
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-redis/redis/v8"
	"github.com/streadway/amqp"
	"golang.org/x/net/context"
)

func main() {
	// Redis
	rdb := redis.NewClient(&redis.Options{
		Addr: "redis:6379",
	})

	// RabbitMQ
	conn, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672/")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	ch, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}
	defer ch.Close()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Fetch data from Redis
		val, err := rdb.Get(context.Background(), "key").Result()
		if err != nil {
			http.Error(w, "Error fetching data from Redis", http.StatusInternalServerError)
			return
		}

		// Send message to RabbitMQ
		err = ch.Publish(
			"",      // exchange
			"queue", // routing key
			false,   // mandatory
			false,   // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte("Message from Go service"),
			})
		if err != nil {
			http.Error(w, "Error sending message to RabbitMQ", http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, "Redis Value: %s", val)
	})

	log.Fatal(http.ListenAndServe(":8082", nil))
}
```

Service A: Dockerfile:
```bash
FROM golang:1.20-alpine
WORKDIR /app
COPY main.go .
RUN go mod init myapp && go build -o myapp
CMD ["./myapp"]
```
Service B: Dockerfile:
```bash
FROM golang:1.20-alpine
WORKDIR /app
COPY main.go .
RUN go mod init myapp && go build -o myapp
CMD ["./myapp"]
```
Service C: Dockerfile:
```bash
FROM golang:1.20-alpine
WORKDIR /app
COPY main.go .
RUN go mod init myapp && go build -o myapp
CMD ["./myapp"]
```
To keep things organized, you might want to use a namespace.
namespace.yaml:
```yaml
apiVersion: v1
kind: Namespace
metadata:
  name: mynamespace
```
mongodb-deployment.yaml:
```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: mongodb
  namespace: mynamespace
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
        - name: mongodb-storage
          mountPath: /data/db
      volumes:
      - name: mongodb-storage
        emptyDir: {}
---
apiVersion: v1
kind: Service
metadata:
  name: mongodb
  namespace: mynamespace
spec:
  ports:
  - port: 27017
  selector:
    app: mongodb
```

postgres-deployment.yaml:

```yaml
apiVersion: apps/v1 # Specifies the API version for the Deployment resource
kind: Deployment # Indicates the type of Kubernetes resource being defined. In this case, it is a Deployment.
metadata: # Contains metadata about the Deployment resource.
  name: postgres # The name of the Deployment. This must be unique within the namespace.
  namespace: mynamespace
spec:
  replicas: 1 # he number of replicas (Pods) that should be running. Here, it’s set to 1, meaning only one Pod will be created.
  selector: # Defines how to select the Pods that belong to this Deployment
    matchLabels: # Specifies the labels used to identify the Pods. Here, it selects Pods with the label app: postgres
      app: postgres
  template: # Describes the Pods that will be created by this Deployment
    metadata: #  Metadata for the Pods created by this Deployment
      labels:
        app: postgres
    spec: # Specifies the desired state and configuration of the Deployment.
      containers: # Specifies the containers to run inside the Pods
      - name: postgres
        image: postgres:latest
        ports: # Defines the ports that the container exposes.
        - containerPort: 5432 # The port on which PostgreSQL listens inside the container.
        env: # Environment variables to set inside the container.
        - name: POSTGRES_DB
          value: "testdb"
        - name: POSTGRES_USER
          value: "postgres"
        - name: POSTGRES_PASSWORD
          value: "password"
        volumeMounts: # Specifies how volumes should be mounted into the container.
        - name: postgres-storage
          mountPath: /var/lib/postgresql/data
      volumes: #  Defines the volumes that can be mounted into the Pods.
      - name: postgres-storage
        emptyDir: {}
---
apiVersion: v1 # Specifies the API version for the Service resource
kind: Service # Indicates that this configuration defines a Service, which exposes a set of Pods as a network service.
metadata:
  name: postgres
  namespace: mynamespace
spec:
  ports:
  - port: 5432
  selector: # Defines how to select the Pods that this Service will route traffic to. 
    app: postgres # The label selector used to identify the Pods managed by the Service. It matches Pods with the label app: postgres.
```
Deployment: Manages the creation and scaling of Pods running PostgreSQL. It defines how many replicas to run, the container image, and environment variables for configuration.

Service: Exposes the PostgreSQL Pods to other services or external traffic within the cluster. It routes traffic to the Pods based on the label selector.

In this setup, the PostgreSQL Deployment creates and manages a single Pod with a persistent storage volume, while the Service provides a stable endpoint to access the PostgreSQL instance. The Deployment and Service are both created in the mynamespace namespace to ensure they operate within the same scope.


redis-deployment.yaml:
```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: redis
  namespace: mynamespace
spec:
  replicas: 1
  selector:
    matchLabels:
      app: redis
  template:
    metadata:
      labels:
        app: redis
    spec:
      containers:
      - name: redis
        image: redis:latest
        ports:
        - containerPort: 6379
---
apiVersion: v1
kind: Service
metadata:
  name: redis
  namespace: mynamespace
spec:
  ports:
  - port: 6379
  selector:
    app: redis
```

rabbitmq-deployment.yaml:
```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: rabbitmq
  namespace: mynamespace
spec:
  replicas: 1
  selector:
    matchLabels:
      app: rabbitmq
  template:
    metadata:
      labels:
        app: rabbitmq
    spec:
      containers:
      - name: rabbitmq
        image: rabbitmq:management
        ports:
        - containerPort: 5672
        - containerPort: 15672
---
apiVersion: v1
kind: Service
metadata:
  name: rabbitmq
  namespace: mynamespace
spec:
  ports:
  - port: 5672
    name: amqp
  - port: 15672
    name: management
  selector:
    app: rabbitmq
```
service-a-deployment.yaml:
```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: service-a
  namespace: mynamespace
spec:
  replicas: 1
  selector:
    matchLabels:
      app: service-a
  template:
    metadata:
      labels:
        app: service-a
    spec:
      containers:
      - name: service-a
        image: your-repo/service-a:latest
        ports:
        - containerPort: 8080
        env:
        - name: MONGO_URI
          value: "mongodb://mongodb:27017"
---
apiVersion: v1
kind: Service
metadata:
  name: service-a
  namespace: mynamespace
spec:
  ports:
  - port: 8080
  selector:
    app: service-a
```
service-b-deployment.yaml:

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: service-b
  namespace: mynamespace
spec:
  replicas: 1
  selector:
    matchLabels:
      app: service-b
  template:
    metadata:
      labels:
        app: service-b
    spec:
      containers:
      - name: service-b
        image: your-repo/service-b:latest
        ports:
        - containerPort: 8081
        env:
        - name: POSTGRES_URI
          value: "postgres://postgres:password@postgres:5432/testdb"
---
apiVersion: v1
kind: Service
metadata:
  name: service-b
  namespace: mynamespace
spec:
  ports:
  - port: 8081
  selector:
    app: service-b
```
service-c-deployment.yaml:
```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: service-c
  namespace: mynamespace
spec:
  replicas: 1
  selector:
    matchLabels:
      app: service-c
  template:
    metadata:
      labels:
        app: service-c
    spec:
      containers:
      - name: service-c
        image: your-repo/service-c:latest
        ports:
        - containerPort: 8082
        env:
        - name: REDIS_URI
          value: "redis:6379"
        - name: RABBITMQ_URI
          value: "amqp://guest:guest@rabbitmq:5672/"
---
apiVersion: v1
kind: Service
metadata:
  name: service-c
  namespace: mynamespace
spec:
  ports:
  - port: 8082
  selector:
    app: service-c
```

Deploying to Kubernetes

Apply Namespace:
```bash
kubectl apply -f namespace.yaml
```

Deploy Databases and Messaging Systems:
```bash
kubectl apply -f mongodb-deployment.yaml
kubectl apply -f postgres-deployment.yaml
kubectl apply -f redis-deployment.yaml
kubectl apply -f rabbitmq-deployment.yaml
```
Deploy Microservices:
```bash
kubectl apply -f service-a-deployment.yaml
kubectl apply -f service-b-deployment.yaml
kubectl apply -f service-c-deployment.yaml
```
Verify Deployments:
```bash
kubectl get pods -n mynamespace
kubectl get services -n mynamespace
```
Accessing Services
To access the services externally, you may need to use a LoadBalancer or NodePort service type for your microservices, or use port forwarding for testing.

Example: Port Forwarding:
```bash
kubectl port-forward svc/service-a 8080:8080 -n mynamespace
```

You can then access Service A at http://localhost:8080.


## Module 8: System Design