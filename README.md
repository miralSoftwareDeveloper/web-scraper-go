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
3. It OnHtml handler will trigger.
4. It store html data in type struct.
5. It create csv file in project folder.
6. It write header columns and data in csv.

