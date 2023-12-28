import { getAllUserGroups } from "@/api/group";
import { Group } from "@/types/group";
import GetOtherUsername from "@/utils/GetOtherUsername";
import { Flex, Container } from "@mantine/core";
import { useEffect, useState } from "react";
import { useNavigate } from "react-router";
import Cookie from "universal-cookie";

const ChatGroups = () => {
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

const ChatGroup = (group: Group) => {
  const navigate = useNavigate();
  const cookies = new Cookie();

  const handleClick = () => {
    navigate(`/chat/${group.id}`);
  };
  console.log(GetOtherUsername(group.usernames, cookies.get("UserId")), group);
  return (
    <Container
      w="100%"
      h={60}
      bg="var(--mantine-color-grape-9)"
      style={{ borderBottom: "1px solid black" }}
      onClick={handleClick}
    >
      {group.is_dm
        ? GetOtherUsername(group.usernames, cookies.get("UserId"))
        : group.name}
    </Container>
  );
};

export default ChatGroups;
