import UserLoginCard from "../ui/UserLoginCard";
import { Login } from "@/types/user";
import { loginUser } from "@/api/user";

const UserLogin = () => {
  const handleSubmit = (values: Login) => {
    loginUser(values).then((res) => {
      if (res.status == 200) {
        window.location.reload();
      }
    });
  };

  return <UserLoginCard onSubmit={handleSubmit} buttonLabel="Login" />;
};

export default UserLogin;
