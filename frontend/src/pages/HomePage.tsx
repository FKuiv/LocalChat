import { FC } from "react";
import { AppShell, Burger, Button } from "@mantine/core";
import { useDisclosure } from "@mantine/hooks";
import Navbar from "@/components/navigation/Navbar";
import Logo from "@/components/ui/Logo";
import { UserEndpoints, api } from "@/endpoints";
import UsersCarousel from "@/components/home/UsersCarousel";

const HomePage: FC = () => {
  const [opened, { toggle }] = useDisclosure();

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
        <UsersCarousel />
        <Button onClick={handleLogout}>Logout</Button>
      </AppShell.Header>

      <AppShell.Navbar p="md">
        <Navbar />
      </AppShell.Navbar>

      <AppShell.Main>Main</AppShell.Main>
    </AppShell>
  );
};

export default HomePage;
