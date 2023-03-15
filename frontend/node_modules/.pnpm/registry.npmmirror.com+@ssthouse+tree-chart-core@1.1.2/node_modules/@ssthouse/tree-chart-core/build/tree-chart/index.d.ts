import { TreeDataset, Direction, TreeLinkStyle } from './tree-chart';
interface TreeConfig {
    nodeWidth: number;
    nodeHeight: number;
    levelHeight: number;
}
interface TreeChartCoreParams {
    treeConfig?: TreeConfig;
    linkStyle?: TreeLinkStyle;
    direction?: Direction;
    collapseEnabled: boolean;
    dataset: TreeDataset;
    svgElement: SVGElement;
    domElement: HTMLDivElement;
    treeContainer: HTMLDivElement;
}
export default class TreeChartCore {
    treeConfig: TreeConfig;
    linkStyle: TreeLinkStyle;
    direction: Direction;
    collapseEnabled: boolean;
    dataset: TreeDataset;
    svgElement: SVGElement;
    svgSelection: any;
    domElement: HTMLDivElement;
    treeContainer: HTMLDivElement;
    nodeDataList: any[];
    linkDataList: any[];
    initTransformX: number;
    initTransformY: number;
    currentScale: number;
    constructor(params: TreeChartCoreParams);
    init(): void;
    getNodeDataList(): any[];
    getInitialTransformStyle(): Record<string, string>;
    zoomIn(): void;
    zoomOut(): void;
    restoreScale(): void;
    setScale(scaleNum: any): void;
    getTranslate(): number[];
    isVertical(): boolean;
    /**
   * 根据link数据,生成svg path data
   */
    private generateLinkPath;
    private generateCurceLinkPath;
    private generateStraightLinkPath;
    updateDataList(): void;
    private draw;
    /**
   * Returns updated dataset by deep copying every nodes from the externalData and adding unique '_key' attributes.
   **/
    private updatedInternalData;
    private buildTree;
    private enableDrag;
    initTransform(): void;
    onClickNode(index: number): void;
    /**
     * call this function to update dataset
     * notice : you need to update the view rendered by `nodeDataList` too
     * @param dataset the new dataset to show in chart
     */
    updateDataset(dataset: TreeDataset): void;
    /**
     * release all dom reference
     */
    destroy(): void;
}
export {};
