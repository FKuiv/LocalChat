import { FC } from "react";
import UserLoginCard from "../ui/UserLoginCard";
import { Login } from "@/types/Login";

const LoginCard: FC = () => {
  const handleSubmit = (values: Login) => {
    console.log(values);
  };

  return <UserLoginCard onSubmit={handleSubmit} buttonLabel="Login" />;
};

export default LoginCard;
