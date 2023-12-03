import { FC } from "react";
import UserLoginCard from "../ui/UserLoginCard";
import { Login } from "@/types/Login";
import { UserEndpoints, api } from "@/endpoints";

const UserLogin: FC = () => {
  const handleSubmit = (values: Login) => {
    api
      .post(UserEndpoints.login, values)
      .then((res) => {
        console.log("logni res", res);
        if (res.status == 200) {
          localStorage.setItem("UserId", res.data["user_id"]);
          localStorage.setItem("Session", res.data["id"]);
          api.defaults.headers.common["Session"] = res.data["id"];
          api.defaults.headers.common["UserId"] = res.data["user_id"];
          window.location.reload();
        }
      })
      .catch((err) => {
        console.log("login err", err);
      });
  };

  return <UserLoginCard onSubmit={handleSubmit} buttonLabel="Login" />;
};

export default UserLogin;
