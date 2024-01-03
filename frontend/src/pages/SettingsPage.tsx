import { logoutUser } from "@/api/user";
import CreateGroupModal from "@/components/modals/CreateGroupModal";
import { Button, Flex } from "@mantine/core";
import { useDisclosure } from "@mantine/hooks";

const SettingsPage = () => {
  const [opened, { open, close }] = useDisclosure(false);
  const handleLogout = () => {
    logoutUser().then(() => {
      window.location.reload();
    });
  };
  return (
    <Flex direction="column" h="100%" gap={20}>
      <CreateGroupModal opened={opened} onClose={close} />
      <Button onClick={open}>Create group chat</Button>
      <Button onClick={handleLogout}>Logout</Button>
    </Flex>
  );
};

export default SettingsPage;
