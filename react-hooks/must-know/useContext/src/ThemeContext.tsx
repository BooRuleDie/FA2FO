import { useContext, createContext, useState } from "react";
import type { ReactNode } from "react";

/*
  We use two separate contexts here:
    - ThemeContext: holds the boolean value representing the theme (dark or light).
    - UpdateThemeContext: holds a function to update (toggle) the theme.

  This pattern (also called the Context + custom hooks split) makes it easier and safer for consumers.
  By using separate contexts, we avoid unnecessary re-renders: 
  components that only need to read the theme value won't re-render when the update function changes and vice versa.
*/

// The ThemeContext holds the current theme (true for dark, false for light by default)
const ThemeContext = createContext<boolean>(true);

// The UpdateThemeContext holds a function to toggle the theme
const UpdateThemeContext = createContext<() => void>(() => {});

/*
  Custom hook to access the current theme value.
  Using a custom hook avoids having to import both useContext and the context object everywhere.
  This makes the consuming code cleaner: just call `useTheme()` wherever needed.
*/
export function useTheme() {
    return useContext(ThemeContext);
}

/*
  Custom hook to access the theme updater (toggleTheme).
  Same motivation as above – keep consuming code simple and decoupled.
*/
export function useUpdateTheme() {
    return useContext(UpdateThemeContext);
}

/*
  The ThemeProvider manages the state (darkTheme) and provides 
  both the current theme and the toggle function down the component tree.

  This template is widely adopted in the React ecosystem for managing global state
  with context, since it allows for fine-grained control and minimal re-renders.
  The double provider pattern here keeps concerns separated for readability and performance.
*/
export function ThemeProvider({ children }: { children: ReactNode }) {
    // Manage the theme state locally in the provider.
    const [darkTheme, setDarkTheme] = useState<boolean>(true);

    // Toggle between dark and light theme values.
    const toggleTheme = () => {
        setDarkTheme((prevState) => !prevState);
    };

    // Provide both the value and the updater using nested Context Providers.
    return (
        <ThemeContext.Provider value={darkTheme}>
            <UpdateThemeContext.Provider value={toggleTheme}>
                {children}
            </UpdateThemeContext.Provider>
        </ThemeContext.Provider>
    );
}
