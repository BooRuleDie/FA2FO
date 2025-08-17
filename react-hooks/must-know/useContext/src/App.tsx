import { useTheme, useUpdateTheme } from "./ThemeContext";

function App() {
    const darkTheme = useTheme();
    const toggleTheme = useUpdateTheme();

    return (
        <div
            className={
                `min-h-screen flex flex-col items-center justify-center transition-colors duration-300 ` +
                (darkTheme
                    ? "bg-gray-900 text-gray-100"
                    : "bg-gray-100 text-gray-900")
            }
        >
            <h1 className="font-mono text-center mb-6 font-bold text-3xl">
                Theme Context
            </h1>
            <div
                className="flex flex-col items-center gap-4 px-6 py-8 rounded-lg shadow-lg bg-opacity-80"
                style={{
                    backgroundColor: darkTheme
                        ? "rgba(31,41,55,0.6)"
                        : "rgba(255,255,255,0.8)",
                }}
            >
                <span className="text-xl font-mono mb-2">
                    Current theme:{" "}
                    <span
                        className={
                            darkTheme ? "text-yellow-400" : "text-gray-800"
                        }
                    >
                        {darkTheme ? "Dark" : "Light"}
                    </span>
                </span>
                <button
                    onClick={toggleTheme}
                    className={
                        "bg-gray-200 hover:bg-gray-300 text-gray-800 cursor-pointer font-bold py-2 px-4 rounded transition-colors duration-200 " +
                        (darkTheme
                            ? "border border-gray-700"
                            : "border border-gray-300")
                    }
                >
                    Toggle Theme
                </button>
            </div>
        </div>
    );
}

export default App;
