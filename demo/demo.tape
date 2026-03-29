# From repo root (PATH must include vhs, ffmpeg, ttyd):
#   vhs demo/demo.tape
Output assets/demo.gif

Require go

Set Shell zsh
Set FontSize 15
Set Width 1120
Set Height 640
Set Padding 36
Set Margin 28
Set MarginFill "#000000"
Set BorderRadius 10
Set WindowBar Colorful
# Same palette as "Aardvark Blue" but terminal background is black; outer canvas via MarginFill.
Set Theme {"name":"Aardvark Blue","black":"#191919","red":"#aa342e","green":"#4b8c0f","yellow":"#dbba00","blue":"#1370d3","magenta":"#c43ac3","cyan":"#008eb0","white":"#bebebe","brightBlack":"#454545","brightRed":"#f05b50","brightGreen":"#95dc55","brightYellow":"#ffe763","brightBlue":"#60a4ec","brightMagenta":"#e26be2","brightCyan":"#60b6cb","brightWhite":"#f7f7f7","background":"#000000","foreground":"#dddddd","cursor":"#007acc","selection":"#bfdbfe"}
Set TypingSpeed 75ms
Set PlaybackSpeed 1.15
Set CursorBlink false

Hide
Type "go build -o billbot ./cmd/billbot"  Sleep 200ms  Enter
Sleep 5s
Type "clear"  Sleep 200ms  Enter
Sleep 300ms
Show

Type "./billbot --help"  Sleep 300ms  Enter
Sleep 4s

Type "clear"  Sleep 200ms  Enter
Sleep 400ms

Type "./billbot --provider cc testdata/anthropic/min.jsonl"  Sleep 300ms  Enter
Sleep 5s

Sleep 2s
