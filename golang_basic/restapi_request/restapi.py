from SimpleHTTPServer import SimpleHTTPRequestHandler
import SocketServer
import simplejson

class S(SimpleHTTPRequestHandler):
    def _set_headers(self):
        self.send_response(200)
        self.send_header('Content-type', 'text/html')
        self.end_headers()

    def do_GET(self):
        print "got get request %s" % (self.path)
        if self.path == '/':
          self.path = '/index.html'
          return SimpleHTTPRequestHandler.do_GET(self)

    def do_POST(self):
        print "got post!!"
        content_len = int(self.headers.getheader('content-length', 0))
        # post_body = self.rfile.read(content_len)
        post_body = '{"servname": "test_serv", "ip": "1.1.1.1"}' 
        test_data = simplejson.loads(post_body)
        print "post_body(%s)" % (test_data)
        return SimpleHTTPRequestHandler.do_POST(self)

def run(handler_class=S, port=80):
    httpd = SocketServer.TCPServer(("", port), handler_class)
    print 'Starting httpd...'
    httpd.serve_forever()


if __name__ == "__main__":
    run()
