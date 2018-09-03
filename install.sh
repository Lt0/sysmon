#!/bin/sh

APP_PATH=$PWD

install() {
	cp onboot/upstart/sysmon.conf /etc/init/sysmon.conf
	sed -i "s|APP_PATH=.*|APP_PATH=\"$APP_PATH\"|g" /etc/init/sysmon.conf
}

install
