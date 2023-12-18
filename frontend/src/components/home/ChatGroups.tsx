import { getAllUserGroups } from "@/api/group";
import { Group } from "@/types/group";
import { Flex, Container } from "@mantine/core";
import { FC, useEffect, useState } from "react";

const ChatGroups: FC = () => {
  const [groups, setGroups] = useState<Group[]>();

  useEffect(() => {
    getAllUserGroups().then((res) => {
      setGroups(res.data);
    });
  }, []);

  return (
    <Flex w="100%" direction="column">
      {groups?.map((group: Group) => (
        <ChatGroup {...group} key={group.id} />
      ))}
    </Flex>
  );
};

const ChatGroup: FC<Group> = (group) => {
  const handleClick = () => {};

  return (
    <Container
      w="100%"
      h={60}
      bg="var(--mantine-color-grape-9)"
      style={{ borderBottom: "1px solid black" }}
      onClick={handleClick}
    >
      {group.name}
    </Container>
  );
};

export default ChatGroups;
