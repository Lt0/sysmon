let unit = 1024;


// ks 是以 KB 为单位的大小
// acc 是保存精度，也就是小数点位数, 默认为 0
function fmtKBSize(ks, acc){
    if (!acc) {
        acc = 0;
    }
    if (ks < unit) {
        return ks.toFixed(acc) + "KB";
    }

    let ms = ks/unit;
    if (ms < unit) {
        return ms.toFixed(acc) + "MB";
    }

    let gs = ms/unit;
    if (gs < unit) {
        return gs.toFixed(acc) + "GB";
    }

    let ts = gs/unit;
    if (ts < unit) {
        return ts.toFixed(acc) + "TB";
    }

    let ps = ts/unit;
    return ps.toFixed(acc) + "PB";
}

exports.fmtKBSize = fmtKBSize;