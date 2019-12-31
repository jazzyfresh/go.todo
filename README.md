go.todo
-------

> Everybody deserves tools to live a meaningful, productive life

## Setup
1. Install the Go programming language
- [ ] Check if Go is already installed
```
#! which go
/usr/local/go/bin/go
```
- [ ] Verify your installation
```
#! go version
go version go1.12.4 linux/amd64
```
- [ ] Add Go binaries to your path
```
#! export PATH=$PATH:~/go/bin
```

2. Download the project codebase
- [ ] Change your working directory to the place where Go will be looking for code
(confusing, too much to explain why it is like this, Google had a monorepo,
resulting in problematic packaging patterns,
but anyways this is defined in Step 1 of setup.)
- [ ] (One time only) Create a directory for your github account
- [ ] Download your fork of the project. This creates a directory for your forked project,
and copies it from the origin repository (GitHub).
```
#! cd $YOUR_GO_HOME
#! mkdir $YOUR_GITHUB_USERNAME
#! git clone https://github.com/$YOUR_GITHUB_USERNAME/go.todo.git
```
- [ ] Change directories in to the project directory and look around
```
#! cd go.todo
#! ls -lG
```

3. Build & Run the web server
```
#! go install
#! go.todo
Starting go.todo server
Go to http://localhost:8080 to get to it
```


4. Test the web server
- [ ] Verify the web server is running & working correctly
by going to the link http://localhost:8080 in your browser
- [ ] Verify the web server is running & working correctly
by curling the link in the terminal
```
#! curl "localhost:8080"
BOOYAH
```
- [ ] Verify the task creation functionality with curl
```
#!  curl -X POST "localhost:8080/task" -H "Content-Type: application/json" -d '{"name":"post a task"}'             
* post a task
#!  curl -X POST "localhost:8080/task" -H "Content-Type: application/json" -d '{"name":"add completed functionality"}'                                                                                                                                
* post a task                                                                                                                      
* add completed functionality
#!  curl -X POST "localhost:8080/task" -H "Content-Type: application/json" -d '{"name":"add lists functionality"}'
* post a task
* add completed functionality
* add lists functionality
#!  curl -X POST "localhost:8080/task"       # Empty task bodies do not add a task
* post a task
* add completed functionality
* add lists functionality
```
- [ ] Meanwhile, the running server process should print out these logs and look like this so far
```
#!  go.todo
Starting go.todo server
Go to http://localhost:8080 to get to it
{post a task}
{add completed functionality}
{add lists functionality}
{}
```


