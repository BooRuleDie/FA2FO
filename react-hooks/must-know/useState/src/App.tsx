import { useState } from "react";

function App() {
    const [count, setCount] = useState(0);

    function decreaseCount() {
        setCount((prevState) => prevState - 1)
    }

    function increaseCount() {
        setCount((prevState) => prevState + 1)
    }

    return (
        <>
            <h1 className="font-mono text-center my-3 font-bold">useState</h1>
            <div className="flex items-center justify-center gap-4 px-6">
                <button
                    onClick={decreaseCount}
                    className="bg-gray-200 hover:bg-gray-300 text-gray-800 font-bold py-2 px-4 rounded-l hover:cursor-pointer"
                >
                    -
                </button>
                <span className="text-2xl font-mono w-12 text-center">{count}</span>
                <button
                    onClick={increaseCount}
                    className="bg-gray-200 hover:bg-gray-300 text-gray-800 font-bold py-2 px-4 rounded-r hover:cursor-pointer"
                >
                    +
                </button>
            </div>
        </>
    );
}

export default App;
