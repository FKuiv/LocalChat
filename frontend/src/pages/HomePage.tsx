import { AppShell, Burger, useMantineTheme } from "@mantine/core";
import { useDisclosure, useMediaQuery } from "@mantine/hooks";
import Navbar from "@/components/navigation/Navbar";
import Logo from "@/components/ui/Logo";
import ChatGroups from "@/components/home/ChatGroups";
import SettingsPage from "./SettingsPage";
import UserCarousel from "@/components/home/UserCarousel";

const HomePage = () => {
  const [opened, { toggle }] = useDisclosure();
  const theme = useMantineTheme();
  const isMobile = useMediaQuery(`(max-width: ${theme.breakpoints.sm})`);

  return (
    <AppShell
      header={{ height: 60 }}
      navbar={{ width: 300, breakpoint: "sm", collapsed: { mobile: !opened } }}
    >
      <AppShell.Header className="header">
        <Burger
          style={{ flexBasis: 1, marginLeft: 15 }}
          opened={opened}
          onClick={toggle}
          hiddenFrom="sm"
          size="sm"
        />
        <Logo />
      </AppShell.Header>
      <AppShell.Navbar p="md">
        {isMobile ? <SettingsPage /> : <Navbar />}
      </AppShell.Navbar>
      <AppShell.Main>
        <UserCarousel />
        <ChatGroups />
      </AppShell.Main>
    </AppShell>
  );
};

export default HomePage;
