const ver = "v1"

let path = {
	infoAll: "/" + ver + "/info/all",
	infoDisk: "/" + ver + "/info/disk",
	
	processAll: "/" + ver + "/process/all",
	processDetails: "/" + ver + "/process/details",
	
	sysInfoAll: "/" + ver + "/sysInfo/all",
	sysInfoHostname: "/" + ver + "/sysInfo/hostname",
}

var activeServer = localStorage.getItem("activeServer")
if(!activeServer) {
	activeServer = ""
}

const infoAll = activeServer + path.infoAll;
const infoDisk = activeServer + path.infoDisk;

const processAll = activeServer + path.processAll;
const processDetails = activeServer + path.processDetails;

const sysInfoAll = activeServer + path.sysInfoAll;
const sysInfoHostname = activeServer + path.sysInfoHostname;

exports.path = path;

exports.infoAll = infoAll;
exports.infoDisk = infoDisk;

exports.processAll = processAll;
exports.processDetails = processDetails;

exports.sysInfoAll = sysInfoAll;
exports.sysInfoHostname = sysInfoHostname;