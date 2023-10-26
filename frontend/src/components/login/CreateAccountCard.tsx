import { Login } from "@/types/Login";
import UserLoginCard from "../ui/UserLoginCard";
import { FC } from "react";

const CreateAccountCard: FC = () => {
  const handleSubmit = (values: Login) => {
    console.log(values);
  };

  return <UserLoginCard onSubmit={handleSubmit} buttonLabel="Create" />;
};
export default CreateAccountCard;
