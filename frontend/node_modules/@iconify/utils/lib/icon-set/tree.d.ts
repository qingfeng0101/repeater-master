import { IconifyJSON } from '@iconify/types';

declare type ParentIconsList = string[];
declare type ParentIconsTree = Record<string, ParentIconsList | null>;
/**
 * Resolve icon set icons
 *
 * Returns parent icon for each icon
 */
declare function getIconsTree(data: IconifyJSON, names?: string[]): ParentIconsTree;

export { ParentIconsList, ParentIconsTree, getIconsTree };
