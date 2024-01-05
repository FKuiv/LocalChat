import { Avatar, Flex, Title } from "@mantine/core";
import logo from "../../media/logo.png";

const Logo = () => {
  return (
    <Flex align="center" justify="center" gap={8} style={{ flexGrow: 1 }}>
      <Avatar src={logo} alt="Localchat" radius="0" size="md" />
      <Title order={1} pt={10}>
        Localchat
      </Title>
    </Flex>
  );
};

export default Logo;
