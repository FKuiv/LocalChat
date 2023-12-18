import { Flex, Image, Title } from "@mantine/core";
import { FC } from "react";
import logo from "../../media/logo.png";

const Logo: FC = () => {
  return (
    <Flex m="auto" align="center" justify="center">
      <Image h={40} w="auto" fit="contain" src={logo} />
      <Title>Localchat</Title>
    </Flex>
  );
};

export default Logo;
