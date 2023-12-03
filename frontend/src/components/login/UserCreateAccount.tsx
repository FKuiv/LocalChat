import { Login } from "@/types/Login";
import UserLoginCard from "../ui/UserLoginCard";
import { FC } from "react";
import { UserEndpoints, api } from "@/endpoints";

const UserCreateAccount: FC = () => {
  const handleSubmit = (values: Login) => {
    api
      .post(UserEndpoints.base, values)
      .then((res) => {
        if (res.status == 200) {
          api
            .post(UserEndpoints.login, values)
            .then((loginRes) => {
              if (loginRes.status == 200) {
                localStorage.setItem("UserId", loginRes.data["user_id"]);
                localStorage.setItem("Session", loginRes.data["id"]);
                api.defaults.headers.common["Session"] = loginRes.data["id"];
                api.defaults.headers.common["UserId"] =
                  loginRes.data["user_id"];
                window.location.reload();
              }
            })
            .catch((loginErr) => {
              console.log("error logging in:", loginErr);
            });
        }
      })
      .catch((err) => console.log("creating user went wrong:", err));
  };

  return <UserLoginCard onSubmit={handleSubmit} buttonLabel="Create" />;
};
export default UserCreateAccount;
