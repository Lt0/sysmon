#!/bin/sh

SCRIPTS_ROOT=etc/init
APP_PATH=$PWD
INIT_TYPE=""


get_init_type() {
	comm=$(cat /proc/1/comm)
	if [ "$comm" = "systemd" ]; then
		INIT_TYPE="systemd"
		return
	elif [ "$comm" = "init"  ]; then
		ver=$(/sbin/init --version 2>/dev/null | grep upstart)
		if [ -n "$ver" ]; then
			INIT_TYPE="upstart"
			return
		fi
	fi

	INIT_TYPE="sysvinit"
}

install() {
	case $INIT_TYPE in
		"systemd" )
			echo install in systemd init
			if [ -d /usr/lib/systemd/system ]; then
				dst_path="/usr/lib/systemd/system"
			elif [ -d /lib/systemd/system ]; then
				dst_path="/lib/systemd/system"
			else
				dst_path="/etc/systemd/system"
			fi
			cp $SCRIPTS_ROOT/systemd/sysmon.service $dst_path/
			sed -i "s|APP_PATH|$APP_PATH|g" $dst_path/sysmon.service
			systemctl daemon-reload
			systemctl enable sysmon
			systemctl start sysmon
		;;
		"upstart" )
			echo install in upstart init
			cp $SCRIPTS_ROOT/upstart/sysmon.conf /etc/init/sysmon.conf
			sed -i "s|APP_PATH=.*|APP_PATH=\"$APP_PATH\"|g" /etc/init/sysmon.conf
			service sysmon start
		;;
		"sysvinit" )
			echo install in sysvinit init
		;;
		* )
			echo install in unknown init
		;;
	esac
}

uninstall() {
	case $INIT_TYPE in
		"systemd" )
			echo uninstall sysmon from systemd init
			systemctl stop sysmon
			rm -rvf /lib/systemd/system/sysmon.service
			find /etc/systemd/system -name "sysmon.*" | xargs rm -vf
			systemctl daemon-reload
		;;
		"upstart" )
			rm -rvf /etc/init/sysmon.conf
			echo uninstall sysmon from upstart init
		;;
		"sysvinit" )
			echo uninstall sysmon from sysvinit init
		;;
		* )
			echo uninstall sysmon from unknown init
		;;
	esac
}

get_init_type
case $1 in
	u | -u | --uninstall )
		uninstall
	;;
	h | -h | --help )
		show_help
	;;
	* )
		install
	;;
esac
