import { WebSocketContext } from "@/WebSocketContext";
import { getAllUserGroups } from "@/api/group";
import { Group } from "@/types/group";
import { Flex, Container } from "@mantine/core";
import { useEffect, useState, useContext } from "react";
import { useNavigate } from "react-router";

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
  const websocket = useContext(WebSocketContext);

  const handleClick = () => {
    websocket?.sendJsonMessage({
      content: "Hello from the frontend",
      group: group.id,
    });
    navigate(`/chat/${group.id}`);
  };

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
