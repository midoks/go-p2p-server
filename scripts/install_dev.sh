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


#go generate internal/assets/conf/conf.go
#go generate internal/assets/public/public.go
#go generate internal/assets/templates/templates.go

#rm -rf go-p2p-server
go build ./
mv go-p2p-server gop2p

cd $TAGRT_DIR/go-p2p-server/scripts

sh make.sh

systemctl daemon-reload

service gop2p restart

cd $TAGRT_DIR/go-p2p-server && ./gop2p -v


