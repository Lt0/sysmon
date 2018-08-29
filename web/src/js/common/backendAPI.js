const ver = "v1"

var activeServer = localStorage.getItem("activeServer")
if(!activeServer) {
	activeServer = ""
}

const infoAll = activeServer + "/" + ver + "/info/all"
const infoDisk = activeServer + "/" + ver + "/info/disk"

const processAll = activeServer + "/" + ver + "/process/all"
const processDetails = activeServer + "/" + ver + "/process/details"

const sysInfoAll = activeServer + "/" + ver + "/sysInfo/all"


exports.infoAll = infoAll;
exports.infoDisk = infoDisk;

exports.processAll = processAll;
exports.processDetails = processDetails;

exports.sysInfoAll = sysInfoAll;