import { uuid } from "../base/uuid";
export function rotatePoint(_a) {
    var x = _a.x, y = _a.y;
    return {
        x: y,
        y: x,
    };
}
/**
* Returns a deep copy of selected node (copy of itself and it's children).
* If selected node or it's children have no '_key' attribute it will assign a new one.
**/
export function deepCopy(node) {
    var obj = { _key: uuid() };
    for (var key in node) {
        if (node[key] === null) {
            obj[key] = null;
        }
        else if (Array.isArray(node[key])) {
            obj[key] = node[key].map(function (x) { return deepCopy(x); });
        }
        else if (typeof node[key] === "object") {
            obj[key] = deepCopy(node[key]);
        }
        else {
            obj[key] = node[key];
        }
    }
    return obj;
}
