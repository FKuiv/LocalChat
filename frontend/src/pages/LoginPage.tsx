import CreateAccountCard from "@/components/login/CreateAccountCard";
import LoginCard from "@/components/login/LoginCard";

export default function LoginPage() {
  return (
    <div className="flex flex-col justify-center items-center h-full">
      <CreateAccountCard />
      <LoginCard />
    </div>
  );
}
