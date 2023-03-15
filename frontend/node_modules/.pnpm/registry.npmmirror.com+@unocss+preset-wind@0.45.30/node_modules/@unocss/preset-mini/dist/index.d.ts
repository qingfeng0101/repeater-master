import { Preflight, PresetOptions, Preset } from '@unocss/core';
export { c as colors } from './colors-67ae184f.js';
export { t as theme } from './default-016bd463.js';
import { T as Theme } from './types-c22910b5.js';
export { T as Theme, a as ThemeAnimation } from './types-c22910b5.js';
export { p as parseColor } from './utilities-d1833377.js';

declare const preflights: Preflight[];

interface DarkModeSelectors {
    /**
     * Selector for light variant.
     *
     * @default '.light'
     */
    light?: string;
    /**
     * Selector for dark variant.
     *
     * @default '.dark'
     */
    dark?: string;
}
interface PresetMiniOptions extends PresetOptions {
    /**
     * Dark mode options
     *
     * @default 'class'
     */
    dark?: 'class' | 'media' | DarkModeSelectors;
    /**
     * Generate pesudo selector as `[group=""]` instead of `.group`
     *
     * @default false
     */
    attributifyPseudo?: Boolean;
    /**
     * Prefix for CSS variables.
     *
     * @default 'un-'
     */
    variablePrefix?: string;
    /**
     * Utils prefix
     *
     * @default undefined
     */
    prefix?: string;
    /**
     * Generate preflight
     *
     * @default true
     */
    preflight?: boolean;
}
declare const presetMini: (options?: PresetMiniOptions) => Preset<Theme>;

export { DarkModeSelectors, PresetMiniOptions, presetMini as default, preflights, presetMini };
