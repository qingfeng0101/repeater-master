import { defineConfig, presetUno } from 'unocss';

export default defineConfig({
  exclude: ['node_modules', 'dist', '.git', '.vscode', 'public'],
  presets: [presetUno()],
  shortcuts: {
    'flex-x-center': 'flex justify-center',
    'flex-y-center': 'flex items-center',
    'flex-center': 'flex justify-center items-center',
    'wh-full': 'w-full h-full',
  },
  theme: {
    colors: {
      primary: 'var(--primary-color)',
    }
  }
});