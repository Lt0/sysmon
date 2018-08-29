let serversKey = "servers";
let activeServerKey = "activeServer";

function getHostServer(){
	return location.protocol + "//" + location.hostname + ":" + location.port;
}

function getServers() {
	return JSON.parse(localStorage.getItem(serversKey));
}

function getActiveServer() {
	return localStorage.getItem(activeServerKey);
}

function setServers(servers){
	localStorage.setItem(serversKey, JSON.stringify(servers));
}

function addServer(server) {
	let servers = getServers();
	if(!servers) {
		servers = [];
	}
	for(let i in servers) {
		if(server == servers[i]) {
			return "This server already exist!";
		}
	}
	servers.push(server);
	setServers(servers);
	return null;
}
    	
function setActiveServer(server) {
	localStorage.setItem(activeServerKey, server);
}

function delServer(server) {
	console.log("delServer ", server);
	let servers = getServers();
    for(let item in servers) {
    	console.log("checking ", servers[item]);
    	if(servers[item] == server) {
    		servers.splice(item, 1);
    		break;
    	}
    }
    setServers(servers);
}

function clearServers() {
	let activeServer = getActiveServer();
	let hostServer = getHostServer();
	let newServers = [];
	
	newServers.push(activeServer);
	
	let proto = location.protocol;
	if(proto == "http:" || "https:") {
		if(hostServer != activeServer) {
			newServers.push(hostServer)
		}
	}
	
	setServers(newServers);
}


export default {
	getHostServer,
	getServers,
	getActiveServer,
	setServers,
	addServer,
	setActiveServer,
	delServer,
	clearServers,
}