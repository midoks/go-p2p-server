#!/bin/bash
PATH=/bin:/sbin:/usr/bin:/usr/sbin:/usr/local/bin:/usr/local/sbin:~/bin


TAGRT_DIR=/usr/local/gop2p_dev
mkdir -p $TAGRT_DIR
cd $TAGRT_DIR


if [ ! -d $TAGRT_DIR/go-p2p-server ]; then
	git clone https://github.com/midoks/go-p2p-server
	cd $TAGRT_DIR/go-p2p-server
else
	cd $TAGRT_DIR/go-p2p-server
	git pull https://github.com/midoks/go-p2p-server
fi

go mod tidy
go mod vendor


rm -rf imail
go build ./


cd $TAGRT_DIR/imail/scripts

sh make.sh

systemctl daemon-reload

service imail restart

cd $TAGRT_DIR/imail && ./imail -v


