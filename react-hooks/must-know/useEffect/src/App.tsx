import { useState, useEffect } from "react";

export default function App() {
    const [resourceType, setResourceType] = useState("posts");
    const [apiResponse, setApiResponse] = useState(null);
    console.log("render");
    useEffect(() => {
        fetch(`https://jsonplaceholder.typicode.com/${resourceType}`)
            .then((response) => response.json())
            .then((json) => {
                setApiResponse(json);
                console.log(json);
            });

        // Cleanup function example.
        // This is useful for things like removing event listeners, cancelling timers,
        // aborting fetch requests, etc., to prevent memory leaks.
        return () => {
            console.log(`Clean up for ${resourceType}`);
            // For example:
            // window.removeEventListener("resize", someHandler);
        };
    }, [resourceType]);
    return (
        <>
            <h1 className="text-center font-mono font-bold text-3xl my-12">useEffect</h1>
            <div className="flex flex-col items-center min-h-screen pb-10">
                <div className="mb-8 flex space-x-4">
                    <button
                        className={`px-4 py-2 rounded transition-colors duration-200 font-mono hover:cursor-pointer ${
                            resourceType === "posts"
                                ? "bg-gray-200 text-gray-800"
                                : "bg-white border border-gray-300 text-gray-800 hover:bg-gray-100"
                        }`}
                        onClick={() => setResourceType("posts")}
                    >
                        Posts
                    </button>
                    <button
                        className={`px-4 py-2 rounded transition-colors duration-200 font-mono hover:cursor-pointer ${
                            resourceType === "users"
                                ? "bg-gray-200 text-gray-800"
                                : "bg-white border border-gray-300 text-gray-800 hover:bg-gray-100"
                        }`}
                        onClick={() => setResourceType("users")}
                    >
                        Users
                    </button>
                    <button
                        className={`px-4 py-2 rounded transition-colors duration-200 font-mono hover:cursor-pointer ${
                            resourceType === "comments"
                                ? "bg-gray-200 text-gray-800"
                                : "bg-white border border-gray-300 text-gray-800 hover:bg-gray-100"
                        }`}
                        onClick={() => setResourceType("comments")}
                    >
                        Comments
                    </button>
                </div>
                <div className="w-full max-w-3/4">
                    <pre className="font-mono text-xs bg-gray-900 text-gray-100 p-4 rounded overflow-x-auto">
                        {apiResponse ? JSON.stringify(apiResponse, null, 2) : "Loading..."}
                    </pre>
                </div>
            </div>
        </>

    );
}
