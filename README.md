# heic-to-jpg
Converts Apple .heic files to .jpg


Install this necessary tool: `sudo apt-get install libheif-examples`


Put all your `.heic` files in the `heics` folder. These files can end with `.heic ` or `.HEIC`.
When you run `main.go` with `go run main.go` it will convert all the `heic` files into `.jpg` files using the same prefix name and store them in the `jpgs` folder.