import { useState, useEffect, useRef } from "react";

function App() {
    // State for current name input value
    const [name, setName] = useState("Mikasa");
    // Ref to keep track of render count without triggering re-renders
    const renderCount = useRef(1);
    // Ref to access the input element directly
    const inputRef = useRef<HTMLInputElement>(null);
    // Ref to store the previous name without causing a re-render
    const prevName = useRef("");

    useEffect(() => {
        // Increment render count whenever 'name' changes
        renderCount.current += 1;
    }, [name]);

    const handleFocus = () => {
        // Focus the input element if it exists
        if (inputRef.current) {
            inputRef.current.focus();
        }
    };

    const updateName = (e: React.ChangeEvent<HTMLInputElement>) => {
        // Store current name as previous before updating, avoids re-render
        prevName.current = name;
        setName(e.target.value);
    };

    return (
        <>
            <h1 className="font-mono text-center my-6 font-bold text-3xl">
                useRef
            </h1>
            <div className="max-w-md mx-auto p-6 space-y-4">
                <input
                    type="text"
                    value={name}
                    onChange={updateName}
                    ref={inputRef}
                    className="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
                    placeholder="Enter your name"
                />
                <p className="text-gray-700">
                    The name you entered is{" "}
                    <span className="font-semibold">{name}</span>
                </p>
                <p className="text-gray-700">
                    Previous name was{" "}
                    <span className="font-semibold">{prevName.current}</span>
                </p>
                <p className="text-gray-600 text-sm">
                    The component has been rendered {renderCount.current} times
                </p>
                <button
                    onClick={handleFocus}
                    className="px-4 py-2 bg-blue-500 text-white rounded-md hover:bg-blue-600 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2 cursor-pointer"
                >
                    Focus Input
                </button>
            </div>
        </>
    );
}

export default App;
