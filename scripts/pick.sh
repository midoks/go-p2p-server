#!/bin/bash
PATH=/bin:/sbin:/usr/bin:/usr/sbin:/usr/local/bin:/usr/local/sbin:~/bin


VERSION=0.9.9
curPath=`pwd`
rootPath=$(dirname "$curPath")

PACK_NAME=gop2p

# go tool dist list
mkdir -p $rootPath/tmp/build
mkdir -p $rootPath/tmp/package

source ~/.bash_profile

cd $rootPath
LDFLAGS='-w -s'
echo $LDFLAGS
build_app(){

	if [ -f $rootPath/tmp/build/gop2p ]; then
		rm -rf $rootPath/tmp/build/gop2p
		rm -rf $rootPath/gop2p
	fi

	if [ -f $rootPath/tmp/build/gop2p.exe ]; then
		rm -rf $rootPath/tmp/build/gop2p.exe
		rm -rf $rootPath/gop2p.exe
	fi

	echo "build_app" $1 $2

	echo "export CGO_ENABLED=1 GOOS=$1 GOARCH=$2"
	echo "cd $rootPath && go build gop2p.go"

	export CGO_ENABLED=0 GOOS=$1 GOARCH=$2
	# export CGO_ENABLED=1 GOOS=linux GOARCH=amd64
	# export CGO_LDFLAGS="-static"

	if [ $1 == "windows" ];then
		
		if [ $2 == "amd64" ]; then
			export CC=x86_64-w64-mingw32-gcc
			export CXX=x86_64-w64-mingw32-g++
		else
			export CC=i686-w64-mingw32-gcc
			export CXX=i686-w64-mingw32-g++
		fi

		cd $rootPath && go build -o gop2p.exe -ldflags "${LDFLAGS}" gop2p.go

		# -ldflags="-s -w"
		# cd $rootPath && go build gop2p.go && /usr/local/bin/strip gop2p
	fi

	if [ $1 == "linux" ]; then
		export CC=x86_64-linux-musl-gcc
		if [ $2 == "amd64" ]; then
			export CC=x86_64-linux-musl-gcc

		fi

		if [ $2 == "386" ]; then
			export CC=i486-linux-musl-gcc
		fi

		if [ $2 == "arm64" ]; then
			export CC=aarch64-linux-musl-gcc
		fi

		if [ $2 == "arm" ]; then
			export CC=arm-linux-musleabi-gcc
		fi

		cd $rootPath && go build -o gop2p -ldflags "${LDFLAGS}" gop2p.go 
	fi

	if [ $1 == "darwin" ]; then
		echo "cd $rootPath && go build -v  -ldflags '${LDFLAGS}'"
		cd $rootPath && go build -o gop2p -v -ldflags "${LDFLAGS}" 
	fi
	

	cp -r $rootPath/scripts $rootPath/tmp/build
	cp -r $rootPath/LICENSE $rootPath/tmp/build
	cp -r $rootPath/README.md $rootPath/tmp/build
	cp -r $rootPath/data $rootPath/tmp/build

	cd $rootPath/tmp/build && xattr -c * && rm -rf ./*/.DS_Store && rm -rf ./*/*/.DS_Store


	if [ $1 == "windows" ];then
		cp $rootPath/gop2p.exe $rootPath/tmp/build
	else
		cp $rootPath/gop2p $rootPath/tmp/build
	fi

	# tar.gz
	cd $rootPath/tmp/build && tar -zcvf ${PACK_NAME}_${VERSION}_$1_$2.tar.gz ./ && mv ${PACK_NAME}_${VERSION}_$1_$2.tar.gz $rootPath/tmp/package
	# bz
	#cd $rootPath/tmp/build && tar -jcvf ${PACK_NAME}_${VERSION}_$1_$2.tar.bz2 ./ && mv ${PACK_NAME}_${VERSION}_$1_$2.tar.bz2 $rootPath/tmp/package
}

golist=`go tool dist list`
echo $golist

# build_app linux amd64
# build_app linux 386
# build_app linux arm64
# build_app linux arm
# build_app darwin amd64
build_app windows 386
build_app windows amd64

