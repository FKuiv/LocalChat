import { FC } from "react";
import UserLoginCard from "../ui/UserLoginCard";
import { Login } from "@/types/Login";
import { UserEndpoints, api } from "@/endpoints";

const UserLogin: FC = () => {
  const handleSubmit = (values: Login) => {
    api
      .post(UserEndpoints.login, values)
      .then((res) => {
        if (res.status == 200) {
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
