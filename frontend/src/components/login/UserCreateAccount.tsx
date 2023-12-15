import { Login } from "@/types/user";
import UserLoginCard from "../ui/UserLoginCard";
import { FC } from "react";
import { createUser, loginUser } from "@/api/user";

const UserCreateAccount: FC = () => {
  const handleSubmit = (values: Login) => {
    createUser(values).then((res) => {
      if (res.status == 200) {
        loginUser(values).then((loginRes) => {
          if (loginRes.status == 200) {
            window.location.reload();
          }
        });
      }
    });
  };

  return <UserLoginCard onSubmit={handleSubmit} buttonLabel="Create" />;
};
export default UserCreateAccount;
