var rsc = require('./resources')
var process = require('./processes')
var fs = require('./fs')
var fmtSize = require('./fmtSize')
var sysInfo = require('./sysInfo')
var bapi = require('./backendAPI')
import sapi from './serverAPI'

export default {
    bus: null,

    rsc: rsc,
    process: process,
    fmtSize: fmtSize,
    fs: fs,
    sysInfo: sysInfo,
    bapi: bapi,
    sapi: sapi,
}