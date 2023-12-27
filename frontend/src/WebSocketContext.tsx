import { createContext, useState, useEffect } from "react";
import useWebSocket, { ReadyState } from "react-use-websocket";
import { WebsocketEndpoints } from "@/api/endpoints";
import { MessageRequest } from "./types/message";
import { SendJsonMessage } from "react-use-websocket/dist/lib/types";

type WebSocketContextValue = {
  sendJsonMessage: SendJsonMessage;
  messageHistory: MessageRequest[];
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
  const [messageHistory, setMessageHistory] = useState<MessageRequest[]>([]);

  const { sendJsonMessage, lastJsonMessage, readyState } = useWebSocket(
    WebsocketEndpoints.base,
    {
      onOpen: () => console.log("ws opened"),
      onClose: () => console.log("ws CLOSED"),
      shouldReconnect: (closeEvent) => {
        console.log("The ws close event:", closeEvent);
        return true;
      },
    }
  );

  const websocketMessage = lastJsonMessage as MessageRequest;

  useEffect(() => {
    if (websocketMessage !== null) {
      setMessageHistory((prev) => prev.concat(websocketMessage));
    }
  }, [websocketMessage]);

  return (
    <WebSocketContext.Provider
      value={{ sendJsonMessage, messageHistory, readyState }}
    >
      {children}
    </WebSocketContext.Provider>
  );
};
