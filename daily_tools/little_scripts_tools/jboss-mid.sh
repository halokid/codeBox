#!/bin/bash

#!/bin/bash
DEST_DIR="/work/"
wget -c -O $DEST_DIR"jboss.tar.gz" --no-check-certificate http://xxxxxx/repos/download/jboss.tar.gz
tar -zxf $DEST_DIR"jboss.tar.gz" -C $DEST_DIR
rm -f $DEST_DIR"jboss.tar.gz"


