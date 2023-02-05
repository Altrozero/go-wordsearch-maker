# Go Wordsearch

A go command line tool for making quick printable wordsearchs for kids. Create a PNG that can be printed, allows for you to set the grid size, capitalization and words.

## Example
Produce a 15x15 PNG to image.png called colors
```
go run *.go -w=15 -h=15 -cap=false -title="Colours" \
-words="black, silver, gray, white, maroon, red, purple, fuchsia, green, lime, olive, yellow, navy, blue, teal, aqua, aquamarine" \
-png=true -file="image.png"
```

# Flags
## -words (string)
Comma seperate value of the words you want to hide. If it fails to place a word it'll spit it out in console.
## -w (int)
The width of the grid
## -h (int)
The height of the grid
## -diag (bool)
Allow diagonal placements. For younger children turn off
## -backward (bool)
Allow backward placements. For younger children turn off
## -stopFill (bool)
A debug command for showing not filling in unused squares with garbage
## -cap (bool)
If letters should be capitalized or lowercase. For younger children use lowercase.

## -png (bool)
If to save to a png or not. If true you must specify a file with -file
## -file (string)
The file to save the image to, you should add .png to the string
## -title (string)
The title of the wordsearch, by default is "wordsearch"