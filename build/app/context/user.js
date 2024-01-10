import { createContext, useContext, useState } from "react";

const UserContext = createContext()

export function UserProvider({ children }) {
  const [userdata, setUserdata] = useState(null);
  return (
    <UserContext.Provider value={[userdata, setUserdata]}>{children}</UserContext.Provider>
  );
}

export function useUserContext() {
    return useContext(UserContext)
}