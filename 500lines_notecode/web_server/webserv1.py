
import BaseHTTPServer


class RequestHandler(BaseHTTPServer.BaseHTTPRequestHandler):
   ''' handle http request by returning a fixed 'page' '''

   #page to send back
   Page = '''\
         <html>
          <body>
          <p>Hello web</p>
          </body>
          </html>
   '''

   # handle a get request
   def do_GET(self):
     self.send_response(200)
     self.send_header("Content-Type", "text/html")
     self.send_header("Content-Length", str(len(self.Page)))
     self.end_headers()
     self.wfile.writeZ(self.Page)

#------------------------------------------------

if __name__ == '__main__':
  serverAddress = ('', 8080)
  server = BaseHTTPServer.HTTPServer(serverAddress, RequestHandler)
  server.server_forever()










