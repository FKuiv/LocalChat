import React from "react";
import ReactDOM from "react-dom/client";
import App from "./App.tsx";
import "./index.css";

import "@mantine/core/styles.css";

import { MantineProvider } from "@mantine/core";
import { Provider } from "react-redux";
import store from "./redux/store.ts";

// const theme = createTheme({
// });

ReactDOM.createRoot(document.getElementById("root")!).render(
  <Provider store={store}>
    <MantineProvider defaultColorScheme="dark">
      <React.StrictMode>
        <App />
      </React.StrictMode>
    </MantineProvider>
  </Provider>
);
