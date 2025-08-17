import { useState, useMemo, useEffect } from "react";

export default function App() {
    const [number, setNumber] = useState(1);
    const [dark, setDark] = useState(false);

    // useMemo caches the result of slowFunction unless 'number' changes
    const doubleNumber = useMemo(() => {
        return slowFunction(number);
    }, [number]);

    // Memoize themeStyles so its reference only changes when 'dark' changes.
    // This prevents unnecessary effect invocations if used as a dependency.
    const themeStyles = useMemo(
        () => ({
            backgroundColor: dark ? "black" : "white",
            color: dark ? "white" : "black",
        }),
        [dark],
    );

    // Thanks to the memoized themeStyles object,
    // this useEffect runs only when themeStyles changes (i.e. when 'dark' toggles).
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

// Simulates an expensive calculation.
function slowFunction(num: number) {
    console.log("Calling Slow Function");
    for (let i = 0; i <= 1000000000; i++) continue;
    return num * 2;
}
