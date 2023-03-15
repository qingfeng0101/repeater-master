export function renderSize(kb:number) {
    var k = 1000, // or 1024
        sizes = ['KB', 'MB', 'GB', 'TB', 'PB', 'EB', 'ZB', 'YB'],
        i = Math.floor(Math.log(kb) / Math.log(k))

    return (kb / Math.pow(k, i)).toPrecision(3) + ' ' + sizes[i]
}