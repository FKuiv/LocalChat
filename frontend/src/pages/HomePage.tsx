import { FC } from "react";
import { AppShell, Burger, Button, useMantineTheme } from "@mantine/core";
import { useDisclosure, useMediaQuery } from "@mantine/hooks";
import Navbar from "@/components/navigation/Navbar";
import Logo from "@/components/ui/Logo";
import { UserEndpoints, api } from "@/endpoints";
import SettingsPage from "./SettingsPage";
import Chats from "@/components/home/Chats";

const HomePage: FC = () => {
  const [opened, { toggle }] = useDisclosure();
  const theme = useMantineTheme();
  const isMobile = useMediaQuery(`(max-width: ${theme.breakpoints.sm})`);

  const handleLogout = () => {
    api
      .get(UserEndpoints.logout)
      .then((res) => {
        console.log("logout res", res);
        window.location.reload();
      })
      .catch((err) => {
        console.log("logout err", err);
      });
  };

  return (
    <AppShell
      header={{ height: 60 }}
      navbar={{ width: 300, breakpoint: "sm", collapsed: { mobile: !opened } }}
    >
      <AppShell.Header className="header">
        <Burger opened={opened} onClick={toggle} hiddenFrom="sm" size="sm" />
        <Logo />
      </AppShell.Header>
      <AppShell.Navbar p="md">
        {isMobile ? <SettingsPage /> : <Navbar />}
      </AppShell.Navbar>
      <AppShell.Main>
        {isMobile && <Chats />}
        <Button onClick={handleLogout}>Logout</Button>
      </AppShell.Main>
    </AppShell>
  );
};

export default HomePage;
