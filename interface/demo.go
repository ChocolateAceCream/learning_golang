package main

import "fmt"

/ Golang program to illustrate the
// concept of polymorphism
package main
 
import (
    "fmt"
)
 
// defining an interface
type Reading interface{
 
    // declaring interface method
    reading_time() int
}
 
// declaring a struct
type Book struct{
     
    // declaring struct variables
    name string
    page_count int
}
 
// declaring a struct
type Newspaper struct{
 
    // declaring struct variables
    name string
    page_count int
}
 
// declaring a struct
type Magazine struct{
 
    // declaring struct variables
    name string
    page_count int
}
 
// function to calculate reading
// time for book
func (book Book) reading_time() int {
     
    // taking average speed
    // of 10 mins per page
    read_time := 10 * book.page_count
    return read_time
}
 
// function to calculate reading
// time for newspaper
func (newspaper Newspaper) reading_time() int {
 
    // taking average speed
    // of 30 mins per page
    read_time := 30 * newspaper.page_count
    return read_time
}
 
// function to calculate reading
// time for magazine
func (magazine Magazine) reading_time() int {
     
    // taking average speed
    // of 5 mins per page
    read_time := 5 * magazine.page_count
    return read_time
}
 
// function to calculate reading time
func calcReadingTime(ReadingTime []Reading) int {
 
        totalTime := 0
     
    // looping through elements
    // of the Reading array   
        for _, t := range ReadingTime {
         
        // run time polymorphism, call to
        // method depends on object being
        // referred at run time
            totalTime += t.reading_time()
        }
     
        return totalTime
}
 
// main function
func main() {
     
    // declaring a book instance
    book1 := Book{
        name: "Goosebumps",
        page_count: 150,
    }
     
    // declaring a newspaper instance
    newspaper1 := Newspaper{
        name: "TOI",
        page_count: 12,
    }
     
    // declaring a magazine instance
    magazine1 := Magazine{
        name: "Forbes",
        page_count: 40,
    }
     
    // array of type Reading interface
    ReadingTime := []Reading{book1, newspaper1, magazine1}
     
    // total reading time calculated
        totalTime := calcReadingTime(ReadingTime)
     
    // Printing total time for reading
        fmt.Printf("Total Time is %d minutes.\n", totalTime)
}