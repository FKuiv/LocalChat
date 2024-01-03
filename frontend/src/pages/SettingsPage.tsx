import { logoutUser } from "@/api/user";
import { Button, Flex, Loader } from "@mantine/core";
import { useDisclosure } from "@mantine/hooks";
import React, { Suspense } from "react";

const CreateGroupModal = React.lazy(
  () => import("@/components/modals/CreateGroupModal")
);

const SettingsPage = () => {
  const [opened, { open, close }] = useDisclosure(false);
  const handleLogout = () => {
    logoutUser().then(() => {
      window.location.reload();
    });
  };
  return (
    <Flex direction="column" h="100%" gap={20}>
      <Suspense fallback={<Loader color="blue" />}>
        <CreateGroupModal opened={opened} onClose={close} />
      </Suspense>
      <Button onClick={open}>Create group chat</Button>
      <Button onClick={handleLogout}>Logout</Button>
    </Flex>
  );
};

export default SettingsPage;
