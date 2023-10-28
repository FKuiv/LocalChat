import UserCreateAccount from "@/components/login/UserCreateAccount";
import UserLogin from "@/components/login/UserLogin";
import { Center, Tabs } from "@mantine/core";
import { FC } from "react";

const LoginPage: FC = () => {
  return (
    <Center className="maxHeight">
      <Tabs defaultValue="createAccount" w="50%">
        <Tabs.List grow mx="auto" w="50%" mb={"md"}>
          <Tabs.Tab value="createAccount">Create Account</Tabs.Tab>
          <Tabs.Tab value="login">Login</Tabs.Tab>
        </Tabs.List>
        <Tabs.Panel value="createAccount">
          <UserCreateAccount />
        </Tabs.Panel>

        <Tabs.Panel value="login">
          <UserLogin />
        </Tabs.Panel>
      </Tabs>
    </Center>
  );
};

export default LoginPage;
