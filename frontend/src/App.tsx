import { useState, useEffect } from "react";
import HomePage from "./pages/HomePage";
import LoginPage from "./pages/LoginPage";
import { ping } from "./api";
import { useDispatch, useSelector } from "react-redux";
import { RootState } from "./redux/store";
import { setLoggedIn } from "./redux/userSlice";
import { Navigate, Route, Routes } from "react-router-dom";
import Chat from "./components/home/Chat";
import useWebSocket from "react-use-websocket";
import { WebsocketEndpoints } from "@/api/endpoints";
import { MessageRequest } from "./types/message";

export default function App() {
  const loggedIn = useSelector((state: RootState) => state.user.isLoggedIn);
  const dispatch = useDispatch();
  const [messageHistory, setMessageHistory] = useState<MessageRequest[]>([]);
  const { sendJsonMessage, lastJsonMessage, readyState } = useWebSocket(
    WebsocketEndpoints.base,
    {
      onOpen: () => console.log("ws opened"),
      onClose: () => console.log("ws CLOSED"),
      // Reconnect
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

    ping()
      .then((res) => {
        if (res.status == 200) {
          dispatch(setLoggedIn(true));
        }
      })
      .catch((err) => {
        if (err.response == undefined) {
          console.log("SERVER is DOWN");
        } else if (err.response.status == 401 || err.response.status == 400) {
          console.log("Not authenticated");
        } else {
          console.log("err from /", err);
        }
      });
  }, [dispatch, websocketMessage]);

  return (
    <Routes>
      <Route
        path="/login"
        element={!loggedIn ? <LoginPage /> : <Navigate to="/" replace />}
      />
      <Route
        path="/"
        element={loggedIn ? <HomePage /> : <Navigate to="/login" replace />}
      />
      <Route
        path="/chat/:groupId"
        element={
          <Chat
            readyState={readyState}
            websocketMessageHistory={messageHistory}
            sendJsonMessage={sendJsonMessage}
          />
        }
      />
      <Route path="*" element={<div>404</div>} />
    </Routes>
  );
}
