import UserCreateAccount from "@/components/login/UserCreateAccount";
import UserLogin from "@/components/login/UserLogin";
import Logo from "@/components/ui/Logo";
import { Flex, Card, Center, Tabs } from "@mantine/core";

const LoginPage = () => {
  return (
    <Center className="maxHeight">
      <Flex
        direction="column"
        justify="center"
        align="center"
        w="100%"
        h="100%"
        gap="xl"
      >
        <Flex>
          <Logo />
        </Flex>
        <Card className="center" miw="30%" maw="70%" h="80%">
          <Tabs defaultValue="createAccount">
            <Tabs.List grow mx={"auto"} w="50%" mb={"xl"}>
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
        </Card>
      </Flex>
    </Center>
  );
};

export default LoginPage;
