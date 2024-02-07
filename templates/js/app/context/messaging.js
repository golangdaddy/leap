import { createContext, useContext, useState } from "react";

const MessagingContext = createContext()

export function MessagingProvider({ children }) {
  const [messagedata, setMessagingdata] = useState({
    "test": [],
    "latest": [],
    "feed": [],
  });
  return (
    <MessagingContext.Provider value={[messagedata, setMessagingdata]}>{children}</MessagingContext.Provider>
  );
}

export function useMessagingContext() {
    return useContext(MessagingContext)
}