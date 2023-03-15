export declare function rotatePoint({ x, y }: {
    x: number;
    y: number;
}): {
    x: number;
    y: number;
};
/**
* Returns a deep copy of selected node (copy of itself and it's children).
* If selected node or it's children have no '_key' attribute it will assign a new one.
**/
export declare function deepCopy(node: any): {
    _key: string;
};
