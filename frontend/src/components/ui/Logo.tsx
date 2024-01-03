import { Flex, Image, Title } from "@mantine/core";
import logo from "../../media/logo.png";

const Logo = () => {
  return (
    <Flex align="center" justify="center" style={{ flexGrow: 1 }}>
      <Image h={40} w="auto" fit="contain" src={logo} />
      <Title>Localchat</Title>
    </Flex>
  );
};

export default Logo;
