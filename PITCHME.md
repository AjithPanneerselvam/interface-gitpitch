## A deep dive into Go Interfaces 

--- 
### Interfaces

An interface is a device or a system that unrelated entities use to interact 

+++
### Types in Go

<p>&nbsp;</p>
@ul 
- Concrete Type (int32, float64)
- Abstract Type (io.Reader, io.Writer)
@ulend

--- 
### Go Interface 
Collection of method signatures 

+++ 
@snap[north span-100]
#### Downloader Interface  
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

Note: 
- Work with the code
- Type can implement superset of the methods declared in the interface to satisfy
- Point out that type can be anything. Need not only by struct. Can be int, string, ...
- Assignn Torrent downloader 

@snapend

---
## Interface Type  

@ul 
- Static Type 
- Dynamic Type 
    - Dynamic value 
@ulend

+++  
@snap[north span-40 text-center] 
#### Interface Type  
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
#### Good to know beforehand  
@ul
- Embed interfaces
- Cannot have duplicate methods
- No cycles 
- No pointer for interface 
- Be wary of receivers 
@ulend 

--- 
#### Empty interface

@snap[midpoint]
```go 
var i interface{}
```
@snapend

- It says nothing 
- It is the largest set in Go 

Note: 

- Show the code 

+++
#### Type Assertions

- To check if value held by interface type variable either implements desired interface or is of a concrete type

@snap[midpoint]
```go 
i.(Type)
```
@snapend

- i should be an interface 


Note: 

- Type can be a concrete type or interface 
- Show example with Type as interface and concrete type 

+++
#### Safe way to do type assertions

```go 
value, ok := i.(Type)
```

Note: 

- Modify the existing code to show them how to perform safe type assertions 

+++ 
## Readers and Writers <3 
