
import BaseHTTPServer


class RequestHandler(BaseHTTPServer.BaseHTTPRequestHandler):
   ''' handle http request by returning a fixed 'page' '''

   #page template...
   def do_GET(selfself):
     page = self.create_page()
     self.send_page(page)

    def create_page(self):
      # ...fill in

    def send_page(self, page):
      #....

#------------------------------------------------

if __name__ == '__main__':
  serverAddress = ('', 8080)
  server = BaseHTTPServer.HTTPServer(serverAddress, RequestHandler)
  server.server_forever()










