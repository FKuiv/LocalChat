import CreateAccountCard from "@/components/login/CreateAccountCard";
import LoginCard from "@/components/login/LoginCard";
import { Center, Tabs } from "@mantine/core";
import { FC } from "react";

const LoginPage: FC = () => {
  return (
    <Center className="maxHeight">
      <Tabs defaultValue="createAccount" w="50%">
        <Tabs.List grow mx="auto" w="50%" mb={"md"} className="border">
          <Tabs.Tab value="createAccount">Create Account</Tabs.Tab>
          <Tabs.Tab value="login">Login</Tabs.Tab>
        </Tabs.List>
        <Tabs.Panel value="createAccount">
          <CreateAccountCard />
        </Tabs.Panel>

        <Tabs.Panel value="login">
          <LoginCard />
        </Tabs.Panel>
      </Tabs>
    </Center>
  );
};

export default LoginPage;
