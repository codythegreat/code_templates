import http.server
import socketserver

# set port to use for traffic onto server (use localhost:PORT)
PORT = 8080

# handler that serves files from current directory and below, mapped to HTTP requests.
Handler = http.server.SimpleHTTPRequestHandler

# uses internt TCP protocol, inputs are server address and request handler class.
with socketserver.TCPServer(("", PORT), Handler) as httpd:
    print("serving at port", PORT)

    # Call this to allow the server to process requests
    httpd.serve_forever()

    # You can call [servername].socket.close() to end the above process
