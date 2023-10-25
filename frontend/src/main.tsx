import React from "react";
import ReactDOM from "react-dom/client";
import App from "./App.tsx";
import "./index.css";

import "@mantine/core/styles.css";

import { MantineProvider } from "@mantine/core";

// const theme = createTheme({
// });

ReactDOM.createRoot(document.getElementById("root")!).render(
  <MantineProvider defaultColorScheme="auto">
    <React.StrictMode>
      <App />
    </React.StrictMode>
  </MantineProvider>
);
