var rsc = require('./resources')
var process = require('./processes')
var fs = require('./fs')
var fmtSize = require('./fmtSize')

export default {
    bus: null,

    rsc: rsc,
    process: process,
    fmtSize: fmtSize,
    fs: fs,
}