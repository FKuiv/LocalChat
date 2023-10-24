import CreateAccountCard from "@/components/login/CreateAccountCard";
import LoginCard from "@/components/login/LoginCard";
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs";

export default function LoginPage() {
  return (
    <div className="flex flex-col justify-center items-center h-full">
      <Tabs
        defaultValue="createAccount"
        className="lg:w-1/4 lg:h-1/2 w-4/5 h-3/5"
      >
        <TabsList className="grid w-full grid-cols-2">
          <TabsTrigger value="createAccount">Create account</TabsTrigger>
          <TabsTrigger value="login">Login</TabsTrigger>
        </TabsList>
        <TabsContent value="createAccount">
          <CreateAccountCard />
        </TabsContent>
        <TabsContent value="login">
          <LoginCard />
        </TabsContent>
      </Tabs>
    </div>
  );
}
