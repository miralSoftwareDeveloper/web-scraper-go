# Web-Scraper in Golang
This project is very basic for scrapping webpages by using Golang package called Colly.

## Packages used
1. Colly - For Scrapping web pages.
2. OS - For creating csv file to store data.
3. log -For logging the errors.
4. fmt- For print to console.
5. encoding/csv - For writing to csv file.

## What Program does
1. It visit web url.
2. It target "li.Product" class.
3. OnHtml Callback function will trigger.
4. It store html data in type struct.
5. It create csv file in project folder.
6. It write header columns and data in csv.

## What I learnt
1. _ discard operator in Golang for ignoring the value
```
   Example :- Standard Forloop - for index, value := range collection {}
              If you want to ignore index just replace it with underscore "_"
               for _,value :=range collection{}
```
2. defer - In Go, the defer keyword is used to schedule a function call to be executed just before the surrounding function returns.
           This allows you to ensure that certain cleanup or finalization tasks are performed regardless of how the function exits—
           whether it returns normally, panics, or encounters a runtime error.
```
   Example :-
               func processFile(filename string) error {
               file, err := os.Open(filename)
                if err != nil {
                    return err
                }
                defer file.Close() // Ensure the file is closed when the function exits
            
                // ... read and process the file ...
            
                return nil // File will be closed automatically when this function returns
            }
```
3.Go Commands used
```
go mod init filename.go - for creating project.
go run filename.go - it will run the script.
go build filename.go - it will build the project and create .exe file in project folder.
```
## Challenge
Colly OnHTML callback function was not called because it was called after Visit callback function.
I have change the order I have define OnHTML callback along with other callback fuction prior to Visit.

## On going Challenge
I want to remove header anchor links which autogenerated at right side of each header in this readme file.

## Comments & Suggestions.
I welcome comments,suggestions & pull requests for mutual learning in any of my repositiory.
   
              

