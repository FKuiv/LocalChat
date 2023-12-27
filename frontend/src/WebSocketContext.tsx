import { createContext, useState, useEffect } from "react";
import useWebSocket, { ReadyState } from "react-use-websocket";
import { WebsocketEndpoints } from "@/api/endpoints";
import { SendJsonMessage } from "react-use-websocket/dist/lib/types";
import { Message } from "./types/message";

type WebSocketContextValue = {
  sendJsonMessage: SendJsonMessage;
  messageHistory: Message[];
  readyState: ReadyState;
};

export const WebSocketContext = createContext<WebSocketContextValue>({
  sendJsonMessage: () => {},
  messageHistory: [],
  readyState: -1,
});

export const WebSocketProvider = ({
  children,
}: {
  children?: React.ReactNode;
}) => {
  const [messageHistory, setMessageHistory] = useState<Message[]>([]);

  const { sendJsonMessage, lastJsonMessage, readyState } = useWebSocket(
    WebsocketEndpoints.base,
    {
      onOpen: () => console.log("ws OPEN"),
      onClose: () => console.log("ws CLOSED"),
      onError: (event) => console.error("ws ERROR", event),
      shouldReconnect: (closeEvent) => {
        console.log("The ws close event:", closeEvent);
        return true;
      },
    }
  );

  const websocketMessage = lastJsonMessage as Message;

  useEffect(() => {
    if (websocketMessage !== null) {
      setMessageHistory((prev) => prev.concat(websocketMessage));
    }
    console.log("THE READY state:", ReadyState[readyState]);
  }, [websocketMessage, readyState]);

  return (
    <WebSocketContext.Provider
      value={{ sendJsonMessage, messageHistory, readyState }}
    >
      {children}
    </WebSocketContext.Provider>
  );
};
