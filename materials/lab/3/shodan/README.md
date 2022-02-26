Lab 3b modifications
Ally Hays
2/25/22

I chose option 2: Extend host.go to build up more complex queries.
I modified host.go to accept page requests as well as query requests.
These are passed from the command line, so I added type conversion to ensure that only numbers are used for the page number, and an error handler to stop the code when the page given is not a number.
I modified main.go to accept 3 arguments from the command line and changed the usage error message accordingly. I also modified the function call to pass the correct command line arguments to HostSearch.