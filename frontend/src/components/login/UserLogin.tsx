import { FC } from "react";
import UserLoginCard from "../ui/UserLoginCard";
import { Login } from "@/types/Login";

const UserLogin: FC = () => {
  const handleSubmit = (values: Login) => {
    console.log(values);
  };

  return <UserLoginCard onSubmit={handleSubmit} buttonLabel="Login" />;
};

export default UserLogin;
