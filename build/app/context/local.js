import { createContext, useContext, useState } from "react";

const LocalContext = createContext()

export function LocalProvider({ children }) {
  const [userdata, setLocaldata] = useState(null);
  return (
    <LocalContext.Provider value={[userdata, setLocaldata]}>{children}</LocalContext.Provider>
  );
}

export function useLocalContext() {
    return useContext(LocalContext)
}