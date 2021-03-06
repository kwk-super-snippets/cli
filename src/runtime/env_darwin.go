package runtime

const defaultEnvString = `kwkenv: "1"
editors:
#  Specify one app for each file type to edit.
#  sh: [vim]
#  go: [gogland]
#  py: [vscode]
#  url: [textedit]
  default: ["textedit"]
apps:
  webstorm: ["open", "-a", "webstorm", "$DIR"]
  textedit: ["open", "-e", "$FULL_NAME"]
  vscode: ["open", "-a", "Visual Studio Code", "$DIR"]
  vim: ["vi", "$FULL_NAME" ]
  emacs: ["emacs", "$FULL_NAME" ]
  nano: ["nano", "$FULL_NAME" ]
  default: ["open", "-t", "$FULL_NAME"]
runners:
  sh: ["/bin/bash", "-c", "$SNIP"]
  url: ["open", "$SNIP"]
  url-covert: ["/Applications/Google Chrome.app/Contents/MacOS/Google Chrome", "--incognito", "$SNIP"]
  js: ["node", "-e", "$SNIP"] #nodejs
  py: ["python", "-c", "$SNIP"] #python
  php: ["php", "-r", "$SNIP"] #php
  scpt: ["osascript", "-e", "$SNIP"] #applescript
  applescript: ["osascript", "-e", "$SNIP"] #applescript
  rb: ["ruby", "-e", "$SNIP"] #ruby
  pl: ["perl", "-E", "$SNIP" ] #perl
  exs: ["elixir", "-e", "$SNIP"] # elixir
  java:
    compile: ["javac", "$FULL_NAME"]
    run: ["java", "$CLASS_NAME"]
  scala:
    compile: ["scalac", "-d", "$DIR", "$FULL_NAME"]
    run: ["scala", "$NAME"]
  cs: #c sharp (dotnet core) Under development
    compile: ["dotnet", "restore", "/Volumes/development/go/src/github.com/rjarmstrong/kwk/src/dotnet/project.json"]
    run: ["dotnet", "run", "--project", "/Volumes/development/go/src/github.com/rjarmstrong/kwk/src/dotnet/project.json", "$FULL_NAME",]
  go: #golang
    run: ["go", "run", "$FULL_NAME"]
  rs: #rust
    compile: ["rustc", "-o", "$NAME", "$FULL_NAME"]
    run: ["$NAME"]
  cpp: # c++
    compile: ["g++", "$FULL_NAME", "-o", "$FULL_NAME.out" ]
    run: ["$FULL_NAME.out"]
  path: ["echo", "$SNIP" ]
  xml: ["echo", "$SNIP"]
  json: ["echo", "$SNIP"]
  yml: ["echo", "$SNIP"]
  md:
    run: ["mdless", "$FULL_NAME"]
  default: ["echo", "$SNIP"]
security: #https://gist.github.com/pmarreck/5388643
  encrypt: []
  decrypt: []
  sign: []
  verify: []`
