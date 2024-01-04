import { AppShell, Burger, useMantineTheme } from "@mantine/core";
import { useDisclosure, useMediaQuery } from "@mantine/hooks";
import Navbar from "@/components/navigation/Navbar";
import Logo from "@/components/ui/Logo";
import ChatGroups from "@/components/home/ChatGroups";
import SettingsPage from "./SettingsPage";
import UserCarousel from "@/components/home/UserCarousel";
import UserAvatar from "@/components/ui/UserAvatar";
import Cookie from "universal-cookie";
import UserSettingsModal from "@/components/modals/UserSettingsModal";
import { useDispatch } from "react-redux";
import { setModalOpen } from "@/redux/userSlice";

const HomePage = () => {
  const [opened, { toggle }] = useDisclosure();
  const theme = useMantineTheme();
  const isMobile = useMediaQuery(`(max-width: ${theme.breakpoints.sm})`);
  const cookies = new Cookie();
  const dispatch = useDispatch();

  return (
    <AppShell
      header={{ height: 60 }}
      navbar={{ width: 300, breakpoint: "sm", collapsed: { mobile: !opened } }}
    >
      <AppShell.Header className="header" px={10}>
        <Burger
          style={{ flexBasis: 1 }}
          opened={opened}
          onClick={toggle}
          hiddenFrom="sm"
          size="sm"
        />
        <Logo />
        <UserAvatar
          userId={cookies.get("UserId")}
          altName="Me"
          size="md"
          onClick={() => dispatch(setModalOpen(true))}
        />
      </AppShell.Header>
      <AppShell.Navbar p="md">
        {isMobile ? <SettingsPage /> : <Navbar />}
      </AppShell.Navbar>
      <AppShell.Main>
        <UserSettingsModal />
        <UserCarousel />
        <ChatGroups />
      </AppShell.Main>
    </AppShell>
  );
};

export default HomePage;
