# go.todo

> Everybody deserves tools to live a meaningful, productive life
>                                                - Me

This is a proof-of-concept productivity project, which is a bunch of jargon for --
this is an app that aims to help you organize your thoughts into dependencies and next actions,
enabling you to live a productive life, whatever that means to you.

Right now, it's a dinky web server built in Go.
Long term, it will be a fully functional (double meaning, ask me later)
application for all clients.

## Setup
#### 1. Install the Go programming language
- [ ] Check if Go is already installed
```
#! which go
/usr/local/go/bin/go
```
- [ ] Download the Go language binary & setup Go directories
```
TODO
```
- [ ] Verify your installation
```
#! go version
go version go1.12.4 linux/amd64
```
- [ ] Add Go binaries to your path (so you can run executables later)
```
#! export PATH=$PATH:~/go/bin
```

#### 2. Download the project codebase
- [ ] Change your working directory to the place where Go will be looking for code
(confusing, too much to explain why it is like this, Google had a monorepo,
resulting in problematic packaging patterns,
but anyways this is defined in Step 1 of setup.)
```
#! cd $YOUR_GO_HOME
```
- [ ] (One time only) Create a directory for your github account
```
#! mkdir $YOUR_GITHUB_USERNAME
```
- [ ] Download your fork of the project. This creates a directory for your forked project,
and copies it from the origin repository (GitHub).
```
#! git clone https://github.com/$YOUR_GITHUB_USERNAME/go.todo.git
```
- [ ] Change directories in to the project directory and look around
```
#! cd go.todo
#! ls
main.go  README.md  Task.go
```

#### 3. Build & Run the web server
```
#! go install
#! go.todo
Starting go.todo server
Go to http://localhost:8080 to get to it
```


#### 4. Test the web server
- [ ] Verify the web server is running & working correctly
by going to the link http://localhost:8080 in your browser
- [ ] Verify the web server is running & working correctly
by curling the link in the terminal
```
#! curl "localhost:8080"
BOOYAH
```
- [ ] Verify the task functionality with curl
```
#!  curl -X POST "localhost:8080/task" -H "Content-Type: application/json" -d '{"name":"post a task","completed":tr
ue}'                                                                                                                               
- [x] post a task
#!  curl -X POST "localhost:8080/task" -H "Content-Type: application/json" -d '{"name":"add completed functionality
"}'                                                                                                                                
- [x] post a task                                                                                                                  
- [ ] add completed functionality                                                                                                  
#!  curl -X POST "localhost:8080/task" -H "Content-Type: application/json" -d '{"name":"add completed functi[1/104]
","completed":true}'
- [x] post a task
- [x] add completed functionality
#!  curl -X POST "localhost:8080/task" -H "Content-Type: application/json" -d '{"name":"add lists functionality"}'
- [x] post a task
- [x] add completed functionality
- [ ] add lists functionality
#!  curl -X POST "localhost:8080/task" -H "Content-Type: application/json" -d '{"name":"backup to file"}'
- [x] post a task
- [x] add completed functionality
- [ ] add lists functionality
- [ ] backup to file
```
- [ ] Meanwhile, the running server process should print out these logs and look like this so far
```
#!  go.todo
Starting go.todo server
Go to http://localhost:8080 to get to it
{post a task true}
{add completed functionality false}
{add completed functionality true}
{add lists functionality false}
{backup to file false}
```

## TODO
- [x] init project
- [x] init readme
- [x] post a task
- [x] add completed functionality
- [ ] add lists functionality
  - [ ] `parent<>child` task relationship
- [ ] backup to file

