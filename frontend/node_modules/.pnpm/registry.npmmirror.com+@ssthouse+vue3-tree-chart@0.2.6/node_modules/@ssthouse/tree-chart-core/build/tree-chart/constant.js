export var DEFAULT_NODE_WIDTH = 100;
export var DEFAULT_NODE_HEIGHT = 100;
export var DEFAULT_LEVEL_HEIGHT = 200;
/**
 * Used to decrement the height of the 'initTransformY' to center diagrams.
 * This is only a hotfix caused by the addition of '__invisible_root' node
 * for multi root purposes.
 */
export var DEFAULT_HEIGHT_DECREMENT = 200;
export var ANIMATION_DURATION = 800;
export var MATCH_TRANSLATE_REGEX = /translate\((-?\d+)px, ?(-?\d+)px\)/i;
export var MATCH_SCALE_REGEX = /scale\((\S*)\)/i;
