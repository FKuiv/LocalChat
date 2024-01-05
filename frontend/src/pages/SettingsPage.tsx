import LogoutButton from "@/components/login/LogoutButton";
import { Button, Flex, Loader } from "@mantine/core";
import { useDisclosure } from "@mantine/hooks";
import React, { Suspense } from "react";

const CreateGroupModal = React.lazy(
  () => import("@/components/modals/CreateGroupModal")
);

const SettingsPage = () => {
  const [opened, { open, close }] = useDisclosure(false);

  return (
    <Flex direction="column" h="100%" gap={20}>
      <Suspense fallback={<Loader color="blue" />}>
        <CreateGroupModal opened={opened} onClose={close} />
      </Suspense>
      <Button onClick={open}>Create group chat</Button>
      <LogoutButton />
    </Flex>
  );
};

export default SettingsPage;
