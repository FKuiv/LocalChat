import { useEffect, useState } from "react";
import HomePage from "./pages/HomePage";
import LoginPage from "./pages/LoginPage";
import { ping } from "./api";

export default function App() {
  const [loggedIn, setLoggedIn] = useState(false);

  useEffect(() => {
    ping()
      .then((res) => {
        if (res.status == 200) {
          setLoggedIn(true);
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
