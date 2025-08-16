# React Hooks

### Must Know Hooks

1. [**`useState`**](./must-know/useState)
2. [**`useEffect`**](./must-know/useEffect)
3. [**`useContext`**](./must-know/useContext)

### Lesser Used Hooks

1. [**`useRef`**](./lesser-used/useRef)
2. [**`useMemo`**](./lesser-used/useMemo)
3. [**`useCallback`**](./lesser-used/useCallback)
4. [**`useReducer`**](./lesser-used/useReducer)
5. [**`useTransition`**](./lesser-used/useTransition)
6. [**`useDeferredValue`**](./lesser-used/useDeferredValue)

### Optional Hooks

1. [**`useLayoutEffect`**](./optional/useLayoutEffect)
2. [**`useDebugValue`**](./optional/useDebugValue)
3. [**`useImperativeHandle`**](./optional/useImperativeHandle)
4. [**`useId`**](./optional/useId)

### Custom Hooks

1. [**`useToggle`**](./custom/useToggle)

### Experimental Hooks

### Startup Project

```bash
# generate the project in current directory
npm create vite@latest . -- --template react-ts
npm install tailwindcss @tailwindcss/vite
```

# vite.config.ts
```ts
import { defineConfig } from 'vite'
import tailwindcss from '@tailwindcss/vite'
export default defineConfig({
  plugins: [
    tailwindcss(),
  ],
})
```

# index.css
```css
@import "tailwindcss";
```