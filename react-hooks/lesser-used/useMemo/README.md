# useMemo

![useMemo Demo](./images/demo.png)

`useMemo` in React acts like a cache for computed values. When the cache key (determined by the dependency array) stays the same on re-render, it just gives you the cached value instead of recalculating it. This is especially useful for expensive calculations, but it also solves a subtle React problem with objects and arrays.

In JavaScript, even if two objects or arrays have the same values, they're not considered identical because they're different references in memory. For example, `{}` !== `{}` and `[] !== []`, but `1 === 1`. So, if you use an object or array as a dependency in your `useEffect` dependency array, the effect will run on every render, since a new object/array reference is created every time—even if the content is the same.

`useMemo` solves this: when you create an object or array inside a `useMemo` and use that memoized value in your `useEffect` dependency array, your effect will only re-run when the actual underlying dependency changes (not just because of a new reference).

See the example: [@App.tsx](./src/App.tsx)