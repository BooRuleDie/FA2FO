import { useState, useMemo, useEffect } from "react";
export default function App() {
    const [number, setNumber] = useState(1);
    const [dark, setDark] = useState(false);

    // useMemo returns a cached value unless 'number' changes
    const doubleNumber = useMemo(() => {
        return slowFunction(number);
    }, [number]);

    // It's important to memoize themeStyles so it has the same reference unless 'dark' changes,
    // otherwise its reference would change on every render causing potential problems if used as a dependency.
    const themeStyles = useMemo(
        () => ({
            backgroundColor: dark ? "black" : "white",
            color: dark ? "white" : "black",
        }),
        [dark],
    );

    // prooves that now thanks to the memoized object
    // this useEffect won't get triggered on re-render caused
    // by doubleNumber chanage but only the themeStyles
    useEffect(() => {
        console.log("themeStyles changed");
    }, [themeStyles]);

    return (
        <>
            <h1 className="font-mono text-center my-6 font-bold text-3xl">useMemo</h1>
            <div className="flex items-center justify-center gap-6 px-6">
                <input
                    type="number"
                    value={number}
                    onChange={(e) => setNumber(parseInt(e.target.value, 10) || 0)}
                    className="border border-gray-300 rounded py-4 text-lg w-16 font-mono text-center focus:outline-none focus:ring-2 focus:ring-blue-500"
                />
                <button
                    onClick={() => setDark((prevDark) => !prevDark)}
                    className="cursor-pointer bg-gray-200 hover:bg-gray-300 text-gray-800 font-bold py-2 px-4 rounded hover:cursor-pointer transition"
                >
                    Change Theme
                </button>
                <div
                    style={themeStyles}
                    className="w-16 border-gray-300 text-2xl font-mono text-center flex items-center justify-center rounded py-4 px-8 border"
                    
                >
                    {doubleNumber}
                </div>
            </div>
        </>
    );
}

function slowFunction(num: number) {
    console.log("Calling Slow Function");
    for (let i = 0; i <= 1000000000; i++) continue;
    return num * 2;
}
