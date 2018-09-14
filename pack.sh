#!/bin/bash

set -e

ARCH=$1
HOST_ARCH=""
SUPP_ARCH="amd64 386 arm64 arm mips64 mips"
VER=$(git tag | tail -n1)
WS=$PWD

get_host_arch() {
	ARCH_RAW=$(uname -m)
	if [ "$ARCH_RAW" = "x86_64" ]; then
		HOST_ARCH="amd64"
	elif [ "$ARCH_RAW" = "i386" ] || [ "$ARCH_RAW" = "i486" ] || [ "$ARCH_RAW" = "i586" ] || [ "$ARCH_RAW" = "i686" ]; then
		HOST_ARCH="386"
	elif [ "$ARCH_RAW" = "aarch64" ]; then
		HOST_ARCH="arm64"
	elif [ "$ARCH_RAW" = "armv6l" ] || [ "$ARCH_RAW" = "armv7l" ]; then
		HOST_ARCH="arm"
	else
		HOST_ARCH="unknown"
	fi
}


#cd web
#npm install
#npm run build
#cd $WS

build_binary_local() {
	echo "build binary with musl-gcc..."
	export CC="musl-gcc"
	go build -o sysmon -ldflags '-linkmode "external" -extldflags "-static"'
}

build_binary_cross() {
	echo "cross build binary for $arch ...."

	arch=$1
	export GOARCH="$arch"

	# only support building all ARCH on amd64 platform and installed docker
	if [ "$HOST_ARCH" = "amd64" ]; then
		if [ "$arch" = "arm" ] || [ "$arch" = "mips" ]; then
			docker run --rm -v $PWD:$PWD -e GOPATH="/vob/golang" -w $PWD -e GOARCH=$arch -it lightimehpq/golang-386 go build
		else
			go build
		fi
	fi
}

build_binary() {
	arch=$1
	if [ "$arch" = "$HOST_ARCH" ]; then
		build_binary_local
	else
		build_binary_cross $arch
	fi
}

pack() {
	echo "Packing..."
	arch=$1

	DST=$(mktemp -d)
	mkdir $DST/sysmon
	DST=$DST/sysmon
	cp -rf sysmon views conf web/dist install.sh etc $DST
	mkdir $DST/web
	mv $DST/dist $DST/web/
	sed -i 's/runmode = dev/runmode = prod/g' $DST/conf/app.conf
	cd $DST/../ && tar Jcf $WS/sysmon-server-${arch}-${VER}.tar.xz *
	cd $WS
}

build_pack_all() {
	for arch in $SUPP_ARCH; do
		echo "###### Building all - $arch ######"
		build_binary $arch
		pack $arch
	done
}

get_host_arch
if [ -z "$ARCH" ]; then
	build_binary $HOST_ARCH
	pack $HOST_ARCH
elif [ "$ARCH" = "all" ]; then
	build_pack_all
else
	build_binary $ARCH
	pack $ARCH
fi
