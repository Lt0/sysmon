var axios = require('axios');
var baip = require("./backendAPI");


let serversKey = "servers";
let activeServerKey = "activeServer";
let hostnameKey = "hostnames"

// 获取 web app 的服务器 IP
function getHostServer(){
	return location.protocol + "//" + location.hostname + ":" + location.port;
}

// 获取当前保存的服务器
function getServers() {
	return JSON.parse(localStorage.getItem(serversKey));
}

// 获取当前连接的服务器
function getActiveServer() {
	return localStorage.getItem(activeServerKey);
}

// 保存所有的服务器
function setServers(servers){
	localStorage.setItem(serversKey, JSON.stringify(servers));
}

// 添加一个服务器
function addServer(server, hostname) {
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
    	
// 设置当前活跃 server
function setActiveServer(server) {
	localStorage.setItem(activeServerKey, server);
}

// 删除本地保存的一个 server
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

// 清空本地保存的所有 servers
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

// 获取当前一个 server 的 hostname，并保存到 store.hsotname
function getHostname(server, store) {
	let url = server + baip.path.sysInfoHostname;
	axios.get(url).then(function(res){
	    store.hostname = res.data;
	}).catch(function(err){
	    console.log("get SysInfo all failed: " + err);
	})
}

// 获取当前活跃 server 的 hostname，并保存到 store.hsotname
function getActiveHostname(store) {
	getHostname(getActiveServer(), store);
}

// 获取当前保存的所有 server 的 hostname，并保存到 hostnames，访问 server 的 hostname: hostnames[server].hostname
function getHostnames(hostnames) {
	let servers = getServers();
	for(let i in servers) {
		hostnames[servers[i]] = {"hostname": ""};
		getHostname(servers[i], hostnames[servers[i]]);
	}
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
	
	getHostname,
	getActiveHostname,
	getHostnames,
}