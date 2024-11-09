**1) Introducing the net/http package (an interlude)**

**Basic Server Setup:**

-Uses net/http package to create a simple web server.
-http.HandleFunc("/", handler): Maps the root path ("/") to the handler function.
-http.ListenAndServe(":8080", nil): Starts the server on port 8080, logging any errors.

**Handler Function:**

-handler(w, r): Defines a basic HTTP handler function.
-fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:]): Sends a response using the URL path (drops the leading /).
-Example: Visiting http://localhost:8080/monkeys responds with Hi there, I love monkeys!

**Serving Wiki Pages:**

-Page Data: Saved in <title>.txt files.

**viewHandler:**
-Maps paths like "/view/test" to load data from test.txt.
-Uses loadPage(title) to read the file content.
-Sends the page title and content as HTML in the response.
-Improved main: Maps /view/ URLs to viewHandler.

**Steps to Run:**

-Save content in test.txt as "Hello world".
-Run with $ go build wiki.go and $ ./wiki.
-Visit http://localhost:8080/view/test to see "Hello world" displayed on the page titled "test".
