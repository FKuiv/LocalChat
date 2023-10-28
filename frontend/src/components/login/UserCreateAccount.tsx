import { Login } from "@/types/Login";
import UserLoginCard from "../ui/UserLoginCard";
import { FC } from "react";
import { UserEndpoints, api } from "@/endpoints";

const UserCreateAccount: FC = () => {
  const handleSubmit = (values: Login) => {
    api
      .post(UserEndpoints.base, values)
      .then((res) => {
        console.log("user creation res", res);
      })
      .catch((err) => console.log(err));
  };

  return <UserLoginCard onSubmit={handleSubmit} buttonLabel="Create" />;
};
export default UserCreateAccount;
