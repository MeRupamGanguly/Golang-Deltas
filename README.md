# Golang-Deltas

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

## Databases

Horizontal Scaling : Distributing data and workload across multiple servers or nodes.

Sharding: Dividing the dataset into smaller, manageable parts called shards and distributing these shards across multiple servers. Each server (or shard) handles a subset of the data.

MongoDB is a NoSql Database which is used for Concurent operations(multi-document ACID transactions) and Horizontal scalability.

PostgreSQL is ideal for applications requiring strong consistency, complex joins, and analytical queries, such as financial systems, CRM applications, and data warehousing.

Redis is an in-memory data structure store known for its exceptional speed and simplicity, It excels in scenarios requiring low latency and high throughput for read and write operations.

MongoDB and Redis excel in horizontal scaling, while PostgreSQL supports vertical scaling better. The ability to scale horizontally can directly influence OPS(Operation Per Second) performance in distributed and large-scale applications.

### MongoDB

#### Scalability:
- Sharding enables horizontal scaling.
- Each shard contains a subset of the data, distributed based on a shard key.
- MongoDB uses config servers to store metadata like mapping between shards and shard keys.
- Config servers allow MongoDB routers (`mongos` instances) to send queries to the appropriate shards based on the shard key.
- `mongos` instances also manage load balancing across the shards. They distribute incoming queries evenly across shards.
- Adding more `mongos` instances can improve the throughput and scalability.
- MongoDB uses replica sets to provide redundancy and automatic failover.
- Each replica set consists of multiple nodes (typically three or more): one primary node for read and write operations and secondary nodes that replicate data from the primary. If the primary node fails, a new primary is elected from the remaining nodes in the replica set, ensuring continuous availability.
- MongoDB’s oplog is a capped collection that records all write operations (inserts, updates, deletes) in the order they occur.

#### Operations
`$eq`: Matches values that are equal to a specified value.
`$ne`: Matches all values that are not equal to a specified value.
`$gt`, $gte, $lt, $lte: Greater than, greater than or equal to, less than, less than or equal to, respectively.
`$in`: Matches any of the values specified in an array.
`$nin`: Matches none of the values specified in an array.
`$set`: Sets the value of a field in a document.
`$unset`: Removes the specified field from a document.
`$inc`: Increments the value of the field by a specified amount.
`$push`: Adds an element to an array.
`$addToSet`: Adds elements to an array only if they do not already exist.
`$pull`: Removes all instances of a value from an array.
`$rename`: Renames a field.
`$project`: Reshapes documents by including, excluding, or transforming fields.
`$elemMatch`: Matches documents that contain an array field with at least one element that matches all the specified query criteria.
`$all`: Matches arrays that contain all elements specified in the query.
`$size`: Matches arrays with a specific number of elements.
`$sort`: Orders documents.
`$limit`: Limits the number of documents.
`$skip`: Skips a specified number of documents.
`$match`: Filters documents.
`$unwind`: Imagine you have a collection where each document contains an array field. When you apply $unwind to an array field in a document, MongoDB will create separate documents where Each new document will have the same values for all other fields except the one array field, for each element of the array.

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


`$lookup` stage in MongoDB aggregation allows you to perform a left outer join to retrieve documents from another collection and include them in your result set.

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

`$text`: Performs text search. 
Text search is case insensitive by default. Text search is also diacritic insensitive (e.g., "café" would match "cafe").
MongoDB calculates a relevance score (score) for each document based on the frequency and proximity of the search terms in the indexed fields.

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

MongoDB supports various types of indexes to accommodate different query patterns and optimize performance:

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



### $elemMatch v/s . operator:

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
    "_id": "6671d6fa65fcd889d2dc2255",
    "name": "Company 5",
    "employee": [
      { "name": "Employee 8", "age": 38, "is_active": true, 
	  "roles": [] },
      {  "name": "Employee 11", "age": 26, "is_active": true, 
	  "roles": ["Law"] }
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
  }
]
```
The dot operator allows us to get a result where each of the two arrays in a document meets one condition individually, but neither array meets all conditions on its own.
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

## Docker and Kubernetis:
Docker is a platform that automate the deployment of applications inside lightweight, portable containers. 

Containers package up everything like code, runtime, system tools, libraries, and settings, that an application needs for runing, so it runs consistently across different computing environments.

Basic Concepts:

Image: A snapshot of a filesystem. An image contains everything needed to run an application: code, runtime, libraries, and environment variables. Images are read-only.

Container: A running instance of an image. Containers can write to their own filesystem (but not to the image) and have a writable layer on top of the image.

Dockerfile: A script with a set of instructions on how to build a Docker image. It specifies the base image, the application code, and any additional dependencies or configurations.

Docker Hub: A cloud-based registry where Docker images are stored and shared.

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
List all containers with their status:
```bash
docker ps -a --format "table {{.ID}}\t{{.Names}}\t{{.Status}}\t{{.Ports}}"
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
Stops a running container.
```bash
docker stop <container>
```
Start a stopped container:
```bash
docker start <container>
```
Restart a container:
```bash
docker restart <container>
```
Deletes a stopped container.
```bash
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
docker exec -it mongodb mongo
```
```bash
docker stop mongodb
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
#### Dockerfile
A Dockerfile is a script containing a series of instructions on how to build a Docker image. Each instruction in the Dockerfile creates a new layer in the Docker image, which is then used to build and run containers.

Dockerfile Commands
```bash
FROM:
Definition: Specifies the base image for the new image.
Example: FROM ubuntu:20.04
```
```bash
RUN:
Definition: Executes a command during the image build process.
Example: RUN apt-get update && apt-get install -y nginx
```
```bash
COPY:
Definition: Copies files from the host to the image.
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
version: '3.8'
services:
  mongodb:
    image: mongo:latest
    container_name: mongodb
    ports:
      - "27017:27017"
    volumes:
      - mongodb_data:/data/db
    networks:
      - mynetwork

volumes:
  mongodb_data:

networks:
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
    environment:
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
    command: ["redis-server", "--appendonly", "yes"]
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
Kubernetes is a powerful open-source platform designed to automate deploying, scaling, and managing containerized applications.  Kubernetes uses YAML files to define resources.

Basic Concepts
Cluster: A set of nodes (machines) that run containerized applications managed by Kubernetes.

Node: A single machine (virtual or physical) that is part of a Kubernetes cluster. Each node runs a container runtime (like Docker), and necessary services to manage containers.

Pod: The smallest deployable unit in Kubernetes, consisting of one or more containers that share storage, networking, and a specification for how to run the containers.
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
Service: An abstraction that defines a logical set of Pods and a policy to access them, typically using a stable IP address and DNS name.
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
Deployment: A controller that manages the deployment of Pods, ensuring the desired number of replicas are running at any time.
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
apiVersion: apps/v1
kind: Deployment
metadata:
  name: postgres
  namespace: mynamespace
spec:
  replicas: 1
  selector:
    matchLabels:
      app: postgres
  template:
    metadata:
      labels:
        app: postgres
    spec:
      containers:
      - name: postgres
        image: postgres:latest
        ports:
        - containerPort: 5432
        env:
        - name: POSTGRES_DB
          value: "testdb"
        - name: POSTGRES_USER
          value: "postgres"
        - name: POSTGRES_PASSWORD
          value: "password"
        volumeMounts:
        - name: postgres-storage
          mountPath: /var/lib/postgresql/data
      volumes:
      - name: postgres-storage
        emptyDir: {}
---
apiVersion: v1
kind: Service
metadata:
  name: postgres
  namespace: mynamespace
spec:
  ports:
  - port: 5432
  selector:
    app: postgres
```

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


