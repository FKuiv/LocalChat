import { useEffect, useState } from "react";
import HomePage from "./pages/HomePage";
import LoginPage from "./pages/LoginPage";
import { api, baseUrl } from "./endpoints";

export default function App() {
  const [loggedIn, setLoggedIn] = useState(false);

  useEffect(() => {
    api
      .get(baseUrl)
      .then((res) => {
        if (res.status == 200) {
          setLoggedIn(true);
        }
      })
      .catch((err) => {
        if (err.response.status == 401 || err.response.status == 403) {
          if (
            localStorage.getItem("Session") != null ||
            localStorage.getItem("UserId") != null
          ) {
            localStorage.removeItem("UserId");
            localStorage.removeItem("Session");
            window.location.reload();
          }
        } else {
          console.log("err from /", err);
        }
      });
  }, []);

  return (
    <div className="maxHeight">
      {!loggedIn ? (
        <LoginPage />
      ) : (
        <>
          <HomePage />
        </>
      )}
    </div>
  );
}
