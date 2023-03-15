import { defineStore } from 'pinia'
import { darkTheme } from 'naive-ui'

import { getColorPalette, addColorAlpha } from '@/utils/color'
type ColorType = 'primary' | 'info' | 'success' | 'warning' | 'error'
type ColorScene = '' | 'Suppl' | 'Hover' | 'Pressed' | 'Active'
type ColorKey = `${ColorType}Color${ColorScene}`
type ThemeColor = Partial<Record<ColorKey, string>>
interface ColorAction {
    scene: ColorScene
    handler: (color: string) => string
}

interface ThemeState {
    darkMode: boolean
    primaryColor: string
    infoColor: string
    successColor: string | undefined
    warningColor: string | undefined
    errorColor: string | undefined
}

export const useThemeStore = defineStore('theme', {
    state: (): ThemeState => {
        return {
            darkMode: false,
            primaryColor: "#5ac8fa",//"#2d8cf0",
            infoColor: "#909399",
            successColor: undefined,
            warningColor: undefined,
            errorColor: undefined
        }
    },
    getters: {
        themeColor: state => {
            return state.primaryColor
        },
        naiveTheme(state) {
            return state.darkMode ? darkTheme : undefined
        },
        naiveThemeColors: state => {
            const themeColors: ThemeColor = {}
            const colorActions: ColorAction[] = [
                { scene: '', handler: color => color },
                { scene: 'Suppl', handler: color => color },
                { scene: 'Hover', handler: color => getColorPalette(color, 5) },
                { scene: 'Pressed', handler: color => getColorPalette(color, 7) },
                { scene: 'Active', handler: color => addColorAlpha(color, 0.1) }
            ]

            const { primaryColor, infoColor, successColor, warningColor, errorColor } = state
            const colors: [ColorType, string | undefined][] = [
                ['primary', primaryColor],
                ['info', infoColor],
                ['success', successColor],
                ['warning', warningColor],
                ['error', errorColor]
            ]
            colors.forEach(color => {
                const [colorType, colorValue] = color
                if (colorValue === undefined) {
                    return
                }
                colorActions.forEach(action => {
                    const colorKey: ColorKey = `${colorType}Color${action.scene}`
                    themeColors[colorKey] = action.handler(colorValue)
                })
            })

            return themeColors
        }

    },
    actions: {
        setThemeColor(themeColor: string): void {
            this.primaryColor = themeColor
            document.documentElement.style.cssText = `--primary-color: ${themeColor}`
        },
    },
    persist: {
        enabled: true,
        strategies: [
            {
                storage: localStorage,
            },
        ],
    }
})