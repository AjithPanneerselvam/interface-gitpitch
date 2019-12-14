### A Deep Dive into Go Interfaces 

@snap[south-east] 
**Ajith Panneerselvam**
@snapend 

--- 
### Interfaces

An interface is a device or a system that unrelated entities use to interact 

+++
### Types in Go

<p>&nbsp;</p>
@ul 
- Concrete Type (int32, float64)
- Abstract Type (io.Reader, io.Writer)
@end

--- 
### Go Interface 
Collection of method signatures 

+++ 
@snap[north span-100]
### Downloader Interface  
@snapend 

@snap[west span-100]
```go 
type Downloader interface {
    Download(url string) (uuid.UUID, error) 
    Pause(downloadID uuid.UUID) error 
    Resume(downloadID uuid.UUID) error 
    Cancel(downloadID uuid.UUID) error 
}
```
@snapend

+++
@snap[north span-100]
#### HTTP Downloader implements Downloader 
@snapend

@snap[west text-06 span-100] 
```go 
type httpDownloader struct{
    protocol string 
}

func (h *httpDownloader) Download(url string) (uuid.UUID, error) {
    // download ...
    return uuid.New(), nil 
}

func (h *httpDownloader) Pause(downloadID uuid.UUID) error {
    // pause the download 
    return nil
}

func (h *httpDownloader) Resume(downloadID uuid.UUID) error {
    // resume the download ... 
    return nil
}

func (h *httpDownloader) Cancel(downloadID uuid.UUID) error {
    // cancel the download ...
    return nil
}
```
@snapend

Note: 
- Work with the code
- Type can implement superset of the methods declared in the interface to satisfy
- Point out that type can be anything. Need not only by struct. Can be int, string, ...
- Assignn Torrent downloader 

---
### Interface Type  

@ul 
- Static Type 
- Dynamic Type 
    - Dynamic value 
@ulend

+++  
@snap[north span-100] 
### Downloader Type  
@snapend 

@snap[span-100 text-06 text-center] 
```go 
func main() {
    var downloader Downloader = &httpDownloader{
        protocol: "http/1.1",
    }

    fmt.Printf("type of downloader is %T", downloader)
    fmt.Printf("value of downloader is %v", downloader)
}
```
@snapend

Note: 
- Show the dynamic type when downloader is not assigned

---
@snap[north span-100]
### Good to know beforehand  
@snapend 

@ul
- Can embed interfaces (Composition)
- Cannot have duplicate methods
- No cycles 
- No pointer for interface 
@ulend 

--- 
@snap[north span-100]
### Empty interface
@snapend

**interface{}**

- It says nothing 
- It is the largest set in Go 

Note: 
- Show the code 

+++
@snap[north span-100]
### Type Assertions
@snapend 

@snap[text-gray]
**i.(Type)**
@snapend

- To check if value held by interface type variable either implements desired interface or is of a concrete type

- i should be an interface 
@snapend

@snapend

Note: 
- Type can be a concrete type or interface 
- Show example with Type as interface and concrete type 

+++
@snap[span-100]
### Safer way to do type assertions
@snapend 

<p>&nbsp;</p>
**value, ok := i.(Type)**

Note: 

- Modify the existing code to show them how to perform safe type assertions 

---
### Readers and Writers <3 

+++ 
@snap[north span-100]
#### Readers and Writers method signature 
@snapend 

``` go 
type Reader interface {
    Read(p []byte) (n int, err error)
}
```

```go 
type Writer interface {
    Write(p []byte) (n int, err error)
}
```

Note: 
- Readers: Readers  have data and you read it out, and make use of that data  
- Writers: You have data and you want to write it into a writer where something happens to that data

- Reader Best Practices: 
    - Read reads up to len(p) bytes into p. It returns the number of bytes read (0 <= n <= len(p)) and any error encountered
    - A successfyk read may return err == EOF or err == nil, depends on source and implementation 
    - Even if Read returns n < len(p), it may use all of p as scratch space during the call 
    - Implmentations must not retain p 
    - If some data is available but not len(p) bytes. Read conventionally returns what is available instead of waiting for more
    - Implementations of Read are discouraged from returning a zero byte count with a nil error, except when len(p) == 0. 
    - [Continuation] Callers should not treat a return of 0 and nil as indicating that nothing happened; in particular it does not indicate EOF. 
    - Process n > 0 bytes returned before considering the error

--- 
#### Unit testing is made easy with interfaces 

---
@snap[north span-100]
#### Object-Oriented like features offered by Interface 
@snapend 

@ul
- Polymorphism 
- Inheritance (Composition) 
- Abstraction 
@ulend 

---
### Interface Segregation Principle 
@quote[Clients should not be forced to depend on methods they do not use.](Robert C. Martin)

---
@quote[A great rule of thumb for Go is accept interfaces, return structs.](Jack Lindamood)

---
### The bigger the interface the weaker the abstraction

---
### Maintain backward compatibility

---
@snap[text-black] 
### Thank you!
@snapend
