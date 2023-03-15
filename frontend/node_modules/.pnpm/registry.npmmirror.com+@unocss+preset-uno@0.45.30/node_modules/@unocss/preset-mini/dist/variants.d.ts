import { Variant, VariantObject } from '@unocss/core';
import { T as Theme } from './types-c22910b5.js';
import { PresetMiniOptions } from './index.js';
import './colors-67ae184f.js';
import './default-016bd463.js';
import './utilities-d1833377.js';

declare const calcMaxWidthBySize: (size: string) => string;
declare const variantBreakpoints: Variant<Theme>;

declare const variantCombinators: Variant[];

declare const variantPrint: Variant;
declare const variantCustomMedia: VariantObject;

declare const variantSupports: VariantObject;

declare const variantColorsMediaOrClass: (options?: PresetMiniOptions) => Variant[];

declare const variants: (options: PresetMiniOptions) => Variant<Theme>[];

declare const variantLanguageDirections: Variant[];

declare const variantSelector: Variant;
declare const variantCssLayer: Variant;
declare const variantInternalLayer: Variant;
declare const variantScope: Variant;
declare const variantVariables: Variant;

declare const variantPseudoClassesAndElements: VariantObject;
declare const variantPseudoClassFunctions: VariantObject;
declare const variantTaggedPseudoClasses: (options?: PresetMiniOptions) => VariantObject[];
declare const partClasses: VariantObject;

declare const variantImportant: Variant;

declare const variantNegative: Variant;

export { calcMaxWidthBySize, partClasses, variantBreakpoints, variantColorsMediaOrClass, variantCombinators, variantCssLayer, variantCustomMedia, variantImportant, variantInternalLayer, variantLanguageDirections, variantNegative, variantPrint, variantPseudoClassFunctions, variantPseudoClassesAndElements, variantScope, variantSelector, variantSupports, variantTaggedPseudoClasses, variantVariables, variants };
