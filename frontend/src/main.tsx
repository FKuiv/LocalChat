import React from "react";
import ReactDOM from "react-dom/client";
import App from "./App.tsx";
import "./index.css";

import "@mantine/core/styles.css";

import { MantineProvider } from "@mantine/core";
import { Provider } from "react-redux";
import store from "./redux/store.ts";
import { BrowserRouter } from "react-router-dom";
import { WebSocketProvider } from "./WebSocketContext.tsx";

// const theme = createTheme({
// });

ReactDOM.createRoot(document.getElementById("root")!).render(
  <Provider store={store}>
    <MantineProvider defaultColorScheme="dark">
      <BrowserRouter>
        <React.StrictMode>
          <WebSocketProvider>
            <App />
          </WebSocketProvider>
        </React.StrictMode>
      </BrowserRouter>
    </MantineProvider>
  </Provider>
);
