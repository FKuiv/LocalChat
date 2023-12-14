import { Image, Title } from "@mantine/core";
import { FC } from "react";
import logo from "../../media/logo.png";

const Logo: FC = () => {
  return (
    <>
      <Image maw={50} src={logo} />
      <Title>Localchat</Title>
    </>
  );
};

export default Logo;
